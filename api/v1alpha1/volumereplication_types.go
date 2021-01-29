/*
Copyright 2021.

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

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// Action the action operator need to perform
type Action string

const (
	// Promote indicates that the volume need to be Promoted
	Promote Action = "Promote"
)

// TypedLocalObjectReference contains enough information to let you locate the
// typed referenced object.
type TypedLocalObjectReference struct {
	// APIGroup is the group for the resource being referenced.
	// If APIGroup is not specified, the specified Kind must be in the core API group.
	// For any other third-party types, APIGroup is required.
	// +optional
	APIGroup *string `json:"apiGroup"`
	// Kind is the type of resource being referenced
	Kind string `json:"kind"`
	// Name is the name of resource being referenced
	Name string `json:"name"`
	// Namespace is the name of namespace where resource is located
	Namespace string `json:"namespace"`
}

// VolumeReplicationSpec defines the desired state of VolumeReplication
type VolumeReplicationSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Provisioner is the name of the driver
	Provisioner string `json:"provisioner,omitempty"`

	// Map of string keys and values that need to be sent to the CSI driver Request.
	// +optional
	Parameters map[string]string `json:"parameters,omitempty"`

	// Enabled indicates the Replication need to be enabled or not.
	Enabled *bool `json:"enabled,omitempty"`

	// Defines the operation the controller as to perform.
	// +optional
	Action *Action `json:"action,omitempty"`

	// This field can be used to specify either:
	// * An existing VolumeSnapshot object (snapshot.storage.k8s.io/VolumeSnapshot)
	// * An existing PVC (PersistentVolumeClaim)
	DataSource TypedLocalObjectReference `json:"dataSource"`
}

// VolumeReplicationStatus defines the observed state of VolumeReplication
type VolumeReplicationStatus struct {
	// Completed indicates that the operation is  completed or not
	Completed bool `json:"completed"`

	// StartTime is a timestamp representing the time at which the server
	// started operating on this object.
	// It is represented in RFC3339 form and is in UTC.
	StartTime metav1.Time `json:"startTime"`

	// CompletedTime is a timestamp representing the time at which the server
	// as completed the operation on this object.
	// It is represented in RFC3339 form and is in UTC.
	CompletedTime metav1.Time `json:"completedTime"`

	// Message indicates the additional information need to provided on this
	// object.
	Message string `json:"message,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// VolumeReplication is the Schema for the volumereplications API
type VolumeReplication struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VolumeReplicationSpec   `json:"spec,omitempty"`
	Status VolumeReplicationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VolumeReplicationList contains a list of VolumeReplication
type VolumeReplicationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VolumeReplication `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VolumeReplication{}, &VolumeReplicationList{})
}
