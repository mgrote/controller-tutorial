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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// PowerstripSpec defines the desired state of Powerstrip
type PowerstripSpec struct {
	// Poweroutlets to be part of this powerstrip
	Outlets            []*Poweroutlet `json:"poweroutlets,omitempty"`
	MQTTStateTopik     string         `json:"mqttstatetopik,omitempty"`
	MQTTTelemetryTopik string         `json:"mqtttelemetrytopik,omitempty"`
	LocationName       string         `json:"location"`
}

// PowerstripStatus defines the observed state of Powerstrip
type PowerstripStatus struct {
	// Poweroutlets that are currently part of this powerstrip
	Outlets         []string `json:"poweroutlets,omitempty"`
	Location        string   `json:"location,omitempty"`
	Consumption     int32    `json:"consumption,omitempty"`
	ConsumptionUnit string   `json:"consumptionunit,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:resource:path=powerstrips,scope=Namespaced,categories=all;power,shortName=strip

// Powerstrip is the Schema for the powerstrips API
// A power strip hold one or more power outlets and provides a location.
type Powerstrip struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PowerstripSpec   `json:"spec,omitempty"`
	Status PowerstripStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// PowerstripList contains a list of Powerstrip
type PowerstripList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Powerstrip `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Powerstrip{}, &PowerstripList{})
}
