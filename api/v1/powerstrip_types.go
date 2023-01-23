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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:resource:path=locations,scope=Namespaced,categories=caas,shortName=loc

// Location is the Schema for the location API
type Location struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	// TODO implement LocationSpec:Mood -> define a outlet/powerstrip-scenario

	Spec        LocationSpec    `json:"spec,omitempty"`
	Status      LocationStatus  `json:"status,omitempty"`
	Powerstrips PowerstripList  `json:"powerstrips,omitempty"`
	Outlets     PowerOutletList `json:"outlets,omitempty"`
}

//+kubebuilder:object:root=true

// LocationList contains a list of Location
type LocationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []Location `json:"items"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

type PowerOutlet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PowerOutletSpec   `json:"spec,omitempty"`
	Status PowerOutletStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

type PowerOutletList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Outlets []PowerOutlet `json:"outlets,omitempty"`
}

type PowerOutletSpec struct {
	SetTo string `json:"setto,omitempty"`
}

type PowerOutletStatus struct {
	// +kubebuilder:default:=off
	// +kubebuilder:validation:Enum:=on;off
	Power           string `json:"on,omitempty"`
	Consumption     int32  `json:"consumption,omitempty"`
	ConsumptionUnit string `json:"consumptionunit,omitempty"`
}

// PowerstripSpec defines the desired state of Powerstrip
type PowerstripSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of Powerstrip. Edit powerstrip_types.go to remove/update
	Foo     string        `json:"foo,omitempty"`
	Outlets []PowerOutlet `json:"poweroutlets,omitempty"`
}

// PowerstripStatus defines the observed state of Powerstrip
type PowerstripStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Consumption     int32  `json:"consumption,omitempty"`
	ConsumptionUnit string `json:"consumptionunit,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Powerstrip is the Schema for the powerstrips API
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

	Items []Powerstrip `json:"items"`
}

// LocationSpec contains the desired state of Location
type LocationSpec struct {
	Mood string `json:"mood"`
}

// LocationStatus contains the observed state of Location
type LocationStatus struct {
	Mood string `json:"mood"`
}

func init() {
	SchemeBuilder.Register(&Powerstrip{}, &PowerstripList{}, &PowerOutlet{}, &PowerstripList{}, &Location{}, &LocationList{})
}
