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

type VpnConnectionRouteObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	RequestID *string `json:"requestId,omitempty" tf:"request_id,omitempty"`
}

type VpnConnectionRouteParameters struct {

	// +kubebuilder:validation:Required
	DestinationIPRange *string `json:"destinationIpRange" tf:"destination_ip_range,omitempty"`

	// +crossplane:generate:reference:type=VpnConnection
	// +kubebuilder:validation:Optional
	VPNConnectionID *string `json:"vpnConnectionId,omitempty" tf:"vpn_connection_id,omitempty"`

	// Reference to a VpnConnection to populate vpnConnectionId.
	// +kubebuilder:validation:Optional
	VPNConnectionIDRef *v1.Reference `json:"vpnConnectionIdRef,omitempty" tf:"-"`

	// Selector for a VpnConnection to populate vpnConnectionId.
	// +kubebuilder:validation:Optional
	VPNConnectionIDSelector *v1.Selector `json:"vpnConnectionIdSelector,omitempty" tf:"-"`
}

// VpnConnectionRouteSpec defines the desired state of VpnConnectionRoute
type VpnConnectionRouteSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VpnConnectionRouteParameters `json:"forProvider"`
}

// VpnConnectionRouteStatus defines the observed state of VpnConnectionRoute.
type VpnConnectionRouteStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VpnConnectionRouteObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VpnConnectionRoute is the Schema for the VpnConnectionRoutes API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,upjet-provider-outscale}
type VpnConnectionRoute struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VpnConnectionRouteSpec   `json:"spec"`
	Status            VpnConnectionRouteStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VpnConnectionRouteList contains a list of VpnConnectionRoutes
type VpnConnectionRouteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VpnConnectionRoute `json:"items"`
}

// Repository type metadata.
var (
	VpnConnectionRoute_Kind             = "VpnConnectionRoute"
	VpnConnectionRoute_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VpnConnectionRoute_Kind}.String()
	VpnConnectionRoute_KindAPIVersion   = VpnConnectionRoute_Kind + "." + CRDGroupVersion.String()
	VpnConnectionRoute_GroupVersionKind = CRDGroupVersion.WithKind(VpnConnectionRoute_Kind)
)

func init() {
	SchemeBuilder.Register(&VpnConnectionRoute{}, &VpnConnectionRouteList{})
}