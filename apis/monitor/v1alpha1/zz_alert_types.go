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

type AlertInitParameters struct {

	// How to send notifications about the alerts. This is a list with one element, .
	// Note that for Slack, the DigitalOcean app needs to have permissions for your workspace. You can
	// read more in Slack's documentation
	// List with details how to notify about the alert. Support for Slack or email.
	Alerts []AlertsInitParameters `json:"alerts,omitempty" tf:"alerts,omitempty"`

	// The comparison for value.
	// This may be either GreaterThan or LessThan.
	// The comparison operator to use for value
	Compare *string `json:"compare,omitempty" tf:"compare,omitempty"`

	// The description of the alert.
	// Description of the alert policy
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The status of the alert.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// A list of IDs for the resources to which the alert policy applies.
	// The droplets to apply the alert policy to
	Entities []*string `json:"entities,omitempty" tf:"entities,omitempty"`

	// A list of tags. When an included tag is added to a resource, the alert policy will apply to it.
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The type of the alert.
	// This may be one of v1/insights/droplet/load_1, v1/insights/droplet/load_5, v1/insights/droplet/load_15,
	// v1/insights/droplet/memory_utilization_percent, v1/insights/droplet/disk_utilization_percent,
	// v1/insights/droplet/cpu, v1/insights/droplet/disk_read, v1/insights/droplet/disk_write,
	// v1/insights/droplet/public_outbound_bandwidth, v1/insights/droplet/public_inbound_bandwidth,
	// v1/insights/droplet/private_outbound_bandwidth, v1/insights/droplet/private_inbound_bandwidth,
	// v1/insights/lbaas/avg_cpu_utilization_percent, v1/insights/lbaas/connection_utilization_percent,
	// v1/insights/lbaas/droplet_health, v1/insights/lbaas/tls_connections_per_second_utilization_percent,
	// v1/insights/lbaas/increase_in_http_error_rate_percentage_5xx, v1/insights/lbaas/increase_in_http_error_rate_percentage_4xx,
	// v1/insights/lbaas/increase_in_http_error_rate_count_5xx, v1/insights/lbaas/increase_in_http_error_rate_count_4xx,
	// v1/insights/lbaas/high_http_request_response_time, v1/insights/lbaas/high_http_request_response_time_50p,
	// v1/insights/lbaas/high_http_request_response_time_95p, v1/insights/lbaas/high_http_request_response_time_99p,
	// v1/dbaas/alerts/load_15_alerts, v1/dbaas/alerts/cpu_alerts, v1/dbaas/alerts/memory_utilization_alerts, or
	// v1/dbaas/alerts/disk_utilization_alerts.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// The value to start alerting at, e.g., 90% or 85Mbps. This is a floating-point number.
	// DigitalOcean will show the correct unit in the web panel.
	Value *float64 `json:"value,omitempty" tf:"value,omitempty"`

	// The time frame of the alert. Either 5m, 10m, 30m, or 1h.
	Window *string `json:"window,omitempty" tf:"window,omitempty"`
}

type AlertObservation struct {

	// How to send notifications about the alerts. This is a list with one element, .
	// Note that for Slack, the DigitalOcean app needs to have permissions for your workspace. You can
	// read more in Slack's documentation
	// List with details how to notify about the alert. Support for Slack or email.
	Alerts []AlertsObservation `json:"alerts,omitempty" tf:"alerts,omitempty"`

	// The comparison for value.
	// This may be either GreaterThan or LessThan.
	// The comparison operator to use for value
	Compare *string `json:"compare,omitempty" tf:"compare,omitempty"`

	// The description of the alert.
	// Description of the alert policy
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The status of the alert.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// A list of IDs for the resources to which the alert policy applies.
	// The droplets to apply the alert policy to
	Entities []*string `json:"entities,omitempty" tf:"entities,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A list of tags. When an included tag is added to a resource, the alert policy will apply to it.
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The type of the alert.
	// This may be one of v1/insights/droplet/load_1, v1/insights/droplet/load_5, v1/insights/droplet/load_15,
	// v1/insights/droplet/memory_utilization_percent, v1/insights/droplet/disk_utilization_percent,
	// v1/insights/droplet/cpu, v1/insights/droplet/disk_read, v1/insights/droplet/disk_write,
	// v1/insights/droplet/public_outbound_bandwidth, v1/insights/droplet/public_inbound_bandwidth,
	// v1/insights/droplet/private_outbound_bandwidth, v1/insights/droplet/private_inbound_bandwidth,
	// v1/insights/lbaas/avg_cpu_utilization_percent, v1/insights/lbaas/connection_utilization_percent,
	// v1/insights/lbaas/droplet_health, v1/insights/lbaas/tls_connections_per_second_utilization_percent,
	// v1/insights/lbaas/increase_in_http_error_rate_percentage_5xx, v1/insights/lbaas/increase_in_http_error_rate_percentage_4xx,
	// v1/insights/lbaas/increase_in_http_error_rate_count_5xx, v1/insights/lbaas/increase_in_http_error_rate_count_4xx,
	// v1/insights/lbaas/high_http_request_response_time, v1/insights/lbaas/high_http_request_response_time_50p,
	// v1/insights/lbaas/high_http_request_response_time_95p, v1/insights/lbaas/high_http_request_response_time_99p,
	// v1/dbaas/alerts/load_15_alerts, v1/dbaas/alerts/cpu_alerts, v1/dbaas/alerts/memory_utilization_alerts, or
	// v1/dbaas/alerts/disk_utilization_alerts.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// The uuid of the alert.
	UUID *string `json:"uuid,omitempty" tf:"uuid,omitempty"`

	// The value to start alerting at, e.g., 90% or 85Mbps. This is a floating-point number.
	// DigitalOcean will show the correct unit in the web panel.
	Value *float64 `json:"value,omitempty" tf:"value,omitempty"`

	// The time frame of the alert. Either 5m, 10m, 30m, or 1h.
	Window *string `json:"window,omitempty" tf:"window,omitempty"`
}

type AlertParameters struct {

	// How to send notifications about the alerts. This is a list with one element, .
	// Note that for Slack, the DigitalOcean app needs to have permissions for your workspace. You can
	// read more in Slack's documentation
	// List with details how to notify about the alert. Support for Slack or email.
	// +kubebuilder:validation:Optional
	Alerts []AlertsParameters `json:"alerts,omitempty" tf:"alerts,omitempty"`

	// The comparison for value.
	// This may be either GreaterThan or LessThan.
	// The comparison operator to use for value
	// +kubebuilder:validation:Optional
	Compare *string `json:"compare,omitempty" tf:"compare,omitempty"`

	// The description of the alert.
	// Description of the alert policy
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The status of the alert.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// A list of IDs for the resources to which the alert policy applies.
	// The droplets to apply the alert policy to
	// +kubebuilder:validation:Optional
	Entities []*string `json:"entities,omitempty" tf:"entities,omitempty"`

	// A list of tags. When an included tag is added to a resource, the alert policy will apply to it.
	// +kubebuilder:validation:Optional
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The type of the alert.
	// This may be one of v1/insights/droplet/load_1, v1/insights/droplet/load_5, v1/insights/droplet/load_15,
	// v1/insights/droplet/memory_utilization_percent, v1/insights/droplet/disk_utilization_percent,
	// v1/insights/droplet/cpu, v1/insights/droplet/disk_read, v1/insights/droplet/disk_write,
	// v1/insights/droplet/public_outbound_bandwidth, v1/insights/droplet/public_inbound_bandwidth,
	// v1/insights/droplet/private_outbound_bandwidth, v1/insights/droplet/private_inbound_bandwidth,
	// v1/insights/lbaas/avg_cpu_utilization_percent, v1/insights/lbaas/connection_utilization_percent,
	// v1/insights/lbaas/droplet_health, v1/insights/lbaas/tls_connections_per_second_utilization_percent,
	// v1/insights/lbaas/increase_in_http_error_rate_percentage_5xx, v1/insights/lbaas/increase_in_http_error_rate_percentage_4xx,
	// v1/insights/lbaas/increase_in_http_error_rate_count_5xx, v1/insights/lbaas/increase_in_http_error_rate_count_4xx,
	// v1/insights/lbaas/high_http_request_response_time, v1/insights/lbaas/high_http_request_response_time_50p,
	// v1/insights/lbaas/high_http_request_response_time_95p, v1/insights/lbaas/high_http_request_response_time_99p,
	// v1/dbaas/alerts/load_15_alerts, v1/dbaas/alerts/cpu_alerts, v1/dbaas/alerts/memory_utilization_alerts, or
	// v1/dbaas/alerts/disk_utilization_alerts.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// The value to start alerting at, e.g., 90% or 85Mbps. This is a floating-point number.
	// DigitalOcean will show the correct unit in the web panel.
	// +kubebuilder:validation:Optional
	Value *float64 `json:"value,omitempty" tf:"value,omitempty"`

	// The time frame of the alert. Either 5m, 10m, 30m, or 1h.
	// +kubebuilder:validation:Optional
	Window *string `json:"window,omitempty" tf:"window,omitempty"`
}

type AlertsInitParameters struct {

	// List of email addresses to sent notifications to
	Email []*string `json:"email,omitempty" tf:"email,omitempty"`

	Slack []SlackInitParameters `json:"slack,omitempty" tf:"slack,omitempty"`
}

type AlertsObservation struct {

	// List of email addresses to sent notifications to
	Email []*string `json:"email,omitempty" tf:"email,omitempty"`

	Slack []SlackObservation `json:"slack,omitempty" tf:"slack,omitempty"`
}

type AlertsParameters struct {

	// List of email addresses to sent notifications to
	// +kubebuilder:validation:Optional
	Email []*string `json:"email,omitempty" tf:"email,omitempty"`

	// +kubebuilder:validation:Optional
	Slack []SlackParameters `json:"slack,omitempty" tf:"slack,omitempty"`
}

type SlackInitParameters struct {

	// The Slack channel to send alerts to
	Channel *string `json:"channel,omitempty" tf:"channel,omitempty"`

	// The webhook URL for Slack
	URL *string `json:"url,omitempty" tf:"url,omitempty"`
}

type SlackObservation struct {

	// The Slack channel to send alerts to
	Channel *string `json:"channel,omitempty" tf:"channel,omitempty"`

	// The webhook URL for Slack
	URL *string `json:"url,omitempty" tf:"url,omitempty"`
}

type SlackParameters struct {

	// The Slack channel to send alerts to
	// +kubebuilder:validation:Optional
	Channel *string `json:"channel" tf:"channel,omitempty"`

	// The webhook URL for Slack
	// +kubebuilder:validation:Optional
	URL *string `json:"url" tf:"url,omitempty"`
}

// AlertSpec defines the desired state of Alert
type AlertSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AlertParameters `json:"forProvider"`
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
	InitProvider AlertInitParameters `json:"initProvider,omitempty"`
}

// AlertStatus defines the observed state of Alert.
type AlertStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AlertObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Alert is the Schema for the Alerts API.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,do}
type Alert struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.alerts) || (has(self.initProvider) && has(self.initProvider.alerts))",message="spec.forProvider.alerts is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.compare) || (has(self.initProvider) && has(self.initProvider.compare))",message="spec.forProvider.compare is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.description) || (has(self.initProvider) && has(self.initProvider.description))",message="spec.forProvider.description is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.type) || (has(self.initProvider) && has(self.initProvider.type))",message="spec.forProvider.type is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.value) || (has(self.initProvider) && has(self.initProvider.value))",message="spec.forProvider.value is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.window) || (has(self.initProvider) && has(self.initProvider.window))",message="spec.forProvider.window is a required parameter"
	Spec   AlertSpec   `json:"spec"`
	Status AlertStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AlertList contains a list of Alerts
type AlertList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Alert `json:"items"`
}

// Repository type metadata.
var (
	Alert_Kind             = "Alert"
	Alert_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Alert_Kind}.String()
	Alert_KindAPIVersion   = Alert_Kind + "." + CRDGroupVersion.String()
	Alert_GroupVersionKind = CRDGroupVersion.WithKind(Alert_Kind)
)

func init() {
	SchemeBuilder.Register(&Alert{}, &AlertList{})
}