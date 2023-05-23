/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

// Package apis contains Kubernetes API for the provider.
package apis

import (
	"k8s.io/apimachinery/pkg/runtime"

	v1alpha1 "github.com/outscale/upjet-provider-outscale/apis/clientgateway/v1alpha1"
	v1alpha1dhcpoption "github.com/outscale/upjet-provider-outscale/apis/dhcpoption/v1alpha1"
	v1alpha1image "github.com/outscale/upjet-provider-outscale/apis/image/v1alpha1"
	v1alpha1internetservice "github.com/outscale/upjet-provider-outscale/apis/internetservice/v1alpha1"
	v1alpha1keypair "github.com/outscale/upjet-provider-outscale/apis/keypair/v1alpha1"
	v1alpha1loadbalancer "github.com/outscale/upjet-provider-outscale/apis/loadbalancer/v1alpha1"
	v1alpha1nat "github.com/outscale/upjet-provider-outscale/apis/nat/v1alpha1"
	v1alpha1net "github.com/outscale/upjet-provider-outscale/apis/net/v1alpha1"
	v1alpha1nic "github.com/outscale/upjet-provider-outscale/apis/nic/v1alpha1"
	v1alpha1publicip "github.com/outscale/upjet-provider-outscale/apis/publicip/v1alpha1"
	v1alpha1route "github.com/outscale/upjet-provider-outscale/apis/route/v1alpha1"
	v1alpha1routetable "github.com/outscale/upjet-provider-outscale/apis/routetable/v1alpha1"
	v1alpha1securitygroup "github.com/outscale/upjet-provider-outscale/apis/securitygroup/v1alpha1"
	v1alpha1snapshot "github.com/outscale/upjet-provider-outscale/apis/snapshot/v1alpha1"
	v1alpha1subnet "github.com/outscale/upjet-provider-outscale/apis/subnet/v1alpha1"
	v1alpha1apis "github.com/outscale/upjet-provider-outscale/apis/v1alpha1"
	v1beta1 "github.com/outscale/upjet-provider-outscale/apis/v1beta1"
	v1alpha1virtualgateway "github.com/outscale/upjet-provider-outscale/apis/virtualgateway/v1alpha1"
	v1alpha1vm "github.com/outscale/upjet-provider-outscale/apis/vm/v1alpha1"
	v1alpha1volume "github.com/outscale/upjet-provider-outscale/apis/volume/v1alpha1"
	v1alpha1vpn "github.com/outscale/upjet-provider-outscale/apis/vpn/v1alpha1"
)

func init() {
	// Register the types with the Scheme so the components can map objects to GroupVersionKinds and back
	AddToSchemes = append(AddToSchemes,
		v1alpha1.SchemeBuilder.AddToScheme,
		v1alpha1dhcpoption.SchemeBuilder.AddToScheme,
		v1alpha1image.SchemeBuilder.AddToScheme,
		v1alpha1internetservice.SchemeBuilder.AddToScheme,
		v1alpha1keypair.SchemeBuilder.AddToScheme,
		v1alpha1loadbalancer.SchemeBuilder.AddToScheme,
		v1alpha1nat.SchemeBuilder.AddToScheme,
		v1alpha1net.SchemeBuilder.AddToScheme,
		v1alpha1nic.SchemeBuilder.AddToScheme,
		v1alpha1publicip.SchemeBuilder.AddToScheme,
		v1alpha1route.SchemeBuilder.AddToScheme,
		v1alpha1routetable.SchemeBuilder.AddToScheme,
		v1alpha1securitygroup.SchemeBuilder.AddToScheme,
		v1alpha1snapshot.SchemeBuilder.AddToScheme,
		v1alpha1subnet.SchemeBuilder.AddToScheme,
		v1alpha1apis.SchemeBuilder.AddToScheme,
		v1beta1.SchemeBuilder.AddToScheme,
		v1alpha1virtualgateway.SchemeBuilder.AddToScheme,
		v1alpha1vm.SchemeBuilder.AddToScheme,
		v1alpha1volume.SchemeBuilder.AddToScheme,
		v1alpha1vpn.SchemeBuilder.AddToScheme,
	)
}

// AddToSchemes may be used to add all resources defined in the project to a Scheme
var AddToSchemes runtime.SchemeBuilder

// AddToScheme adds all Resources to the Scheme
func AddToScheme(s *runtime.Scheme) error {
	return AddToSchemes.AddToScheme(s)
}
