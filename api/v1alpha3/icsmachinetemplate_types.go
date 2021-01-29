/*


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

package v1alpha3

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ICSMachineTemplateSpec defines the desired state of ICSMachineTemplate
type ICSMachineTemplateSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of ICSMachineTemplate. Edit ICSMachineTemplate_types.go to remove/update
	Foo string `json:"foo,omitempty"`
}

// ICSMachineTemplateStatus defines the observed state of ICSMachineTemplate
type ICSMachineTemplateStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true

// ICSMachineTemplate is the Schema for the icsmachinetemplates API
type ICSMachineTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ICSMachineTemplateSpec   `json:"spec,omitempty"`
	Status ICSMachineTemplateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ICSMachineTemplateList contains a list of ICSMachineTemplate
type ICSMachineTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ICSMachineTemplate `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ICSMachineTemplate{}, &ICSMachineTemplateList{})
}
