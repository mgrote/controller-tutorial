package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	cfg "sigs.k8s.io/controller-runtime/pkg/config/v1alpha1"
)

type MQTTConfig struct {
	Broker   *string `json:"broker,omitempty"`
	ClientID *string `json:"clientID,omitempty"`
	// TODO should be a secret
	UserName *string `json:"userName,omitempty"`
	Password *string `json:"password,omitempty"`
}

//+kubebuilder:object:root=true

type PersonalIOTConfig struct {
	metav1.TypeMeta                        `json:",inline"`
	cfg.ControllerManagerConfigurationSpec `json:",inline"`

	MQTTConfig MQTTConfig `json:"mqttConfig,omitempty"`
}

func init() {
	SchemeBuilder.Register(&PersonalIOTConfig{})
}
