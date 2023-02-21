package mqttiot

import (
	"fmt"
	"os"

	mqtt "github.com/eclipse/paho.mqtt.golang"

	personaliotv1alpha1 "github.com/mgrote/personal-iot/api/v1alpha1"
)

const (
	mqttBroker   = "MQTT_BROKER"
	mqttClientID = "MQTT_CLIENT_ID"
	mqttUserName = "MQTT_USER"
	mqttPassWord = "MQTT_PASS"
)

func ClientOpts(mqttConfig personaliotv1alpha1.MQTTConfig) *mqtt.ClientOptions {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(*mqttConfig.Broker)
	opts.SetClientID(*mqttConfig.ClientID)
	opts.SetUsername(*mqttConfig.UserName)
	opts.SetPassword(*mqttConfig.Password)
	opts.SetCleanSession(true)
	return opts
}

func ClientOptsFromEnv() (*mqtt.ClientOptions, error) {
	opts := mqtt.NewClientOptions()
	broker, found := os.LookupEnv(mqttBroker)
	if !found {
		return nil, fmt.Errorf("unable to find environment var %s", mqttBroker)
	}
	clientID, found := os.LookupEnv(mqttClientID)
	if !found {
		return nil, fmt.Errorf("unable to find environment var %s", mqttClientID)
	}
	user, found := os.LookupEnv(mqttUserName)
	if !found {
		return nil, fmt.Errorf("unable to find environment var %s", mqttUserName)
	}
	pass, found := os.LookupEnv(mqttPassWord)
	if !found {
		return nil, fmt.Errorf("unable to find environment var %s", mqttPassWord)
	}
	opts.AddBroker(broker)
	opts.SetClientID(clientID)
	opts.SetUsername(user)
	opts.SetPassword(pass)
	opts.SetCleanSession(true)
	return opts, nil
}
