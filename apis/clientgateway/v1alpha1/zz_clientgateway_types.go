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

type ClientGatewayObservation struct {
	ClientGatewayID *string `json:"clientGatewayId,omitempty" tf:"client_gateway_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	RequestID *string `json:"requestId,omitempty" tf:"request_id,omitempty"`

	State *string `json:"state,omitempty" tf:"state,omitempty"`
}

type ClientGatewayParameters struct {

	// +kubebuilder:validation:Required
	BGPAsn *float64 `json:"bgpAsn" tf:"bgp_asn,omitempty"`

	// +kubebuilder:validation:Required
	ConnectionType *string `json:"connectionType" tf:"connection_type,omitempty"`

	// +kubebuilder:validation:Required
	PublicIP *string `json:"publicIp" tf:"public_ip,omitempty"`

	// +kubebuilder:validation:Optional
	Tags []TagsParameters `json:"tags,omitempty" tf:"tags,omitempty"`
}

type TagsObservation struct {
}

type TagsParameters struct {

	// +kubebuilder:validation:Optional
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

// ClientGatewaySpec defines the desired state of ClientGateway
type ClientGatewaySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ClientGatewayParameters `json:"forProvider"`
}

// ClientGatewayStatus defines the observed state of ClientGateway.
type ClientGatewayStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ClientGatewayObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ClientGateway is the Schema for the ClientGateways API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,upjet-provider-outscale}
type ClientGateway struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ClientGatewaySpec   `json:"spec"`
	Status            ClientGatewayStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ClientGatewayList contains a list of ClientGateways
type ClientGatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClientGateway `json:"items"`
}

// Repository type metadata.
var (
	ClientGateway_Kind             = "ClientGateway"
	ClientGateway_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ClientGateway_Kind}.String()
	ClientGateway_KindAPIVersion   = ClientGateway_Kind + "." + CRDGroupVersion.String()
	ClientGateway_GroupVersionKind = CRDGroupVersion.WithKind(ClientGateway_Kind)
)

func init() {
	SchemeBuilder.Register(&ClientGateway{}, &ClientGatewayList{})
}