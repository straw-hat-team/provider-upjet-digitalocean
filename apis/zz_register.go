// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

// Package apis contains Kubernetes API for the provider.
package apis

import (
	"k8s.io/apimachinery/pkg/runtime"

	v1alpha1 "github.com/straw-hat-team/provider-digitalocean/apis/compute/v1alpha1"
	v1alpha1database "github.com/straw-hat-team/provider-digitalocean/apis/database/v1alpha1"
	v1alpha1digitalocean "github.com/straw-hat-team/provider-digitalocean/apis/digitalocean/v1alpha1"
	v1alpha1droplet "github.com/straw-hat-team/provider-digitalocean/apis/droplet/v1alpha1"
	v1alpha1kubernetes "github.com/straw-hat-team/provider-digitalocean/apis/kubernetes/v1alpha1"
	v1alpha1monitor "github.com/straw-hat-team/provider-digitalocean/apis/monitor/v1alpha1"
	v1alpha1networking "github.com/straw-hat-team/provider-digitalocean/apis/networking/v1alpha1"
	v1alpha1project "github.com/straw-hat-team/provider-digitalocean/apis/project/v1alpha1"
	v1alpha1storage "github.com/straw-hat-team/provider-digitalocean/apis/storage/v1alpha1"
	v1alpha1uptime "github.com/straw-hat-team/provider-digitalocean/apis/uptime/v1alpha1"
	v1alpha1apis "github.com/straw-hat-team/provider-digitalocean/apis/v1alpha1"
	v1beta1 "github.com/straw-hat-team/provider-digitalocean/apis/v1beta1"
	v1alpha1volume "github.com/straw-hat-team/provider-digitalocean/apis/volume/v1alpha1"
)

func init() {
	// Register the types with the Scheme so the components can map objects to GroupVersionKinds and back
	AddToSchemes = append(AddToSchemes,
		v1alpha1.SchemeBuilder.AddToScheme,
		v1alpha1database.SchemeBuilder.AddToScheme,
		v1alpha1digitalocean.SchemeBuilder.AddToScheme,
		v1alpha1droplet.SchemeBuilder.AddToScheme,
		v1alpha1kubernetes.SchemeBuilder.AddToScheme,
		v1alpha1monitor.SchemeBuilder.AddToScheme,
		v1alpha1networking.SchemeBuilder.AddToScheme,
		v1alpha1project.SchemeBuilder.AddToScheme,
		v1alpha1storage.SchemeBuilder.AddToScheme,
		v1alpha1uptime.SchemeBuilder.AddToScheme,
		v1alpha1apis.SchemeBuilder.AddToScheme,
		v1beta1.SchemeBuilder.AddToScheme,
		v1alpha1volume.SchemeBuilder.AddToScheme,
	)
}

// AddToSchemes may be used to add all resources defined in the project to a Scheme
var AddToSchemes runtime.SchemeBuilder

// AddToScheme adds all Resources to the Scheme
func AddToScheme(s *runtime.Scheme) error {
	return AddToSchemes.AddToScheme(s)
}
