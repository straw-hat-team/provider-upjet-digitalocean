// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ResourcesInitParameters struct {

	// a list of uniform resource names (URNs) for the resources associated with the project
	// the resources associated with the project
	Resources []*string `json:"resources,omitempty" tf:"resources,omitempty"`
}

type ResourcesObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// the ID of the project
	// project ID
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// a list of uniform resource names (URNs) for the resources associated with the project
	// the resources associated with the project
	Resources []*string `json:"resources,omitempty" tf:"resources,omitempty"`
}

type ResourcesParameters struct {

	// the ID of the project
	// project ID
	// +crossplane:generate:reference:type=Project
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Reference to a Project to populate project.
	// +kubebuilder:validation:Optional
	ProjectRef *v1.Reference `json:"projectRef,omitempty" tf:"-"`

	// Selector for a Project to populate project.
	// +kubebuilder:validation:Optional
	ProjectSelector *v1.Selector `json:"projectSelector,omitempty" tf:"-"`

	// a list of uniform resource names (URNs) for the resources associated with the project
	// the resources associated with the project
	// +kubebuilder:validation:Optional
	Resources []*string `json:"resources,omitempty" tf:"resources,omitempty"`
}

// ResourcesSpec defines the desired state of Resources
type ResourcesSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ResourcesParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider ResourcesInitParameters `json:"initProvider,omitempty"`
}

// ResourcesStatus defines the observed state of Resources.
type ResourcesStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ResourcesObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Resources is the Schema for the Resourcess API.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,do}
type Resources struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.resources) || (has(self.initProvider) && has(self.initProvider.resources))",message="spec.forProvider.resources is a required parameter"
	Spec   ResourcesSpec   `json:"spec"`
	Status ResourcesStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ResourcesList contains a list of Resourcess
type ResourcesList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Resources `json:"items"`
}

// Repository type metadata.
var (
	Resources_Kind             = "Resources"
	Resources_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Resources_Kind}.String()
	Resources_KindAPIVersion   = Resources_Kind + "." + CRDGroupVersion.String()
	Resources_GroupVersionKind = CRDGroupVersion.WithKind(Resources_Kind)
)

func init() {
	SchemeBuilder.Register(&Resources{}, &ResourcesList{})
}