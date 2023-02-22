/*
Copyright 2023.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"
	"fmt"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	"github.com/mgrote/personal-iot/api/v1alpha1"
	"github.com/mgrote/personal-iot/internal"
)

// PoweroutletReconciler reconciles a Poweroutlet object
type PoweroutletReconciler struct {
	client.Client
	Scheme         *runtime.Scheme
	MQTTClientOpts *mqtt.ClientOptions
}

//+kubebuilder:rbac:groups=personal-iot.mgrote,resources=poweroutlets,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=personal-iot.mgrote,resources=poweroutlets/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=personal-iot.mgrote,resources=poweroutlets/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Poweroutlet object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.13.1/pkg/reconcile
func (r *PoweroutletReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx).WithValues("PoweroutletReconciler", req.NamespacedName)

	powerOutlet := &v1alpha1.Poweroutlet{}
	if err := r.Get(ctx, req.NamespacedName, powerOutlet); err != nil {
		logger.Error(err, "unable to fetch power outlet")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}
	logger.WithValues("switch", powerOutlet.Spec.Switch).Info("found power outlet")
	// nothing to do, leave
	if powerOutlet.Spec.Switch == powerOutlet.Status.Switch {
		logger.Info("desired switch state reached, nothing else to do", "", powerOutlet.Spec.Switch, "state", powerOutlet.Spec.Switch)
		return ctrl.Result{}, nil
	}
	currentState, err := r.reconcilePowerOutletState(ctx, powerOutlet)
	if err != nil {
		return ctrl.Result{}, err
	}
	logger.Info("reached state", "current state", currentState)
	powerOutlet.Status.Switch = *currentState

	if err := r.Status().Update(ctx, powerOutlet); err != nil {
		logger.Error(err, "update PowerOutlet status")
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

func (r *PoweroutletReconciler) reconcilePowerOutletState(ctx context.Context, powerOutlet *v1alpha1.Poweroutlet) (*string, error) {

	// TODO change timing
	// first subscribe and check current state
	// second other client send message
	// third check subscription again and check against msg.Duplicate and msg.MessageID

	mqttClient := mqtt.NewClient(r.MQTTClientOpts)
	if token := mqttClient.Connect(); token.Wait() && token.Error() != nil {
		return nil, fmt.Errorf("client could not connect MQTT broker %s, %w", r.MQTTClientOpts.Servers, token.Error())
	}

	// publish new state
	token := mqttClient.Publish(powerOutlet.Spec.MQTTCommandTopik, 1, true, powerOutlet.Spec.Switch)
	if !token.WaitTimeout(time.Second * 5) {
		return nil, fmt.Errorf("client could not publish to MQTT topik %s", powerOutlet.Spec.MQTTCommandTopik)
	}

	// check current state after publishing
	messageChannel := make(chan [2]string)
	messageHandler := func(client mqtt.Client, msg mqtt.Message) {
		messageChannel <- [2]string{msg.Topic(), string(msg.Payload())}
	}
	// TODO lecture ---> debug timing problems
	if token = mqttClient.Subscribe(powerOutlet.Spec.MQTTStatusTopik, 1, messageHandler); token.WaitTimeout(time.Second*5) && token.Error() != nil {
		return nil, fmt.Errorf(
			"client could not subscribe MQTT topik %s, %w", powerOutlet.Spec.MQTTStatusTopik, token.Error())
	}

	var currentState string
	for i := 0; i < 2; i++ {
		incoming := <-messageChannel
		currentState = incoming[1]
	}

	//mqttClient.Unsubscribe(powerOutlet.Spec.MQTTStatusTopik)
	mqttClient.Disconnect(500)

	// check for valid message format
	if currentState == internal.PowerOnSignal || currentState == internal.PowerOffSignal {
		return &currentState, nil
	}

	return nil, fmt.Errorf("unexpected state %s  found, expected where %s or %s", currentState, internal.PowerOnSignal, internal.PowerOffSignal)
}

// SetupWithManager sets up the controller with the Manager.
func (r *PoweroutletReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&v1alpha1.Poweroutlet{}).
		Complete(r)
}
