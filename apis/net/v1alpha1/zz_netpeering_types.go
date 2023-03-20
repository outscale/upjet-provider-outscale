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

type NetPeeringObservation struct {
	AccepterNet map[string]*string `json:"accepterNet,omitempty" tf:"accepter_net,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	NetPeeringID *string `json:"netPeeringId,omitempty" tf:"net_peering_id,omitempty"`

	RequestID *string `json:"requestId,omitempty" tf:"request_id,omitempty"`

	SourceNet map[string]*string `json:"sourceNet,omitempty" tf:"source_net,omitempty"`

	State map[string]*string `json:"state,omitempty" tf:"state,omitempty"`
}

type NetPeeringParameters struct {

	// +crossplane:generate:reference:type=Net
	// +kubebuilder:validation:Optional
	AccepterNetID *string `json:"accepterNetId,omitempty" tf:"accepter_net_id,omitempty"`

	// Reference to a Net to populate accepterNetId.
	// +kubebuilder:validation:Optional
	AccepterNetIDRef *v1.Reference `json:"accepterNetIdRef,omitempty" tf:"-"`

	// Selector for a Net to populate accepterNetId.
	// +kubebuilder:validation:Optional
	AccepterNetIDSelector *v1.Selector `json:"accepterNetIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	SourceNetAccountID *string `json:"sourceNetAccountId,omitempty" tf:"source_net_account_id,omitempty"`

	// +crossplane:generate:reference:type=Net
	// +kubebuilder:validation:Optional
	SourceNetID *string `json:"sourceNetId,omitempty" tf:"source_net_id,omitempty"`

	// Reference to a Net to populate sourceNetId.
	// +kubebuilder:validation:Optional
	SourceNetIDRef *v1.Reference `json:"sourceNetIdRef,omitempty" tf:"-"`

	// Selector for a Net to populate sourceNetId.
	// +kubebuilder:validation:Optional
	SourceNetIDSelector *v1.Selector `json:"sourceNetIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Tags []NetPeeringTagsParameters `json:"tags,omitempty" tf:"tags,omitempty"`
}

type NetPeeringTagsObservation struct {
}

type NetPeeringTagsParameters struct {

	// +kubebuilder:validation:Optional
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

// NetPeeringSpec defines the desired state of NetPeering
type NetPeeringSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NetPeeringParameters `json:"forProvider"`
}

// NetPeeringStatus defines the observed state of NetPeering.
type NetPeeringStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NetPeeringObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// NetPeering is the Schema for the NetPeerings API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,upjet-provider-outscale}
type NetPeering struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NetPeeringSpec   `json:"spec"`
	Status            NetPeeringStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NetPeeringList contains a list of NetPeerings
type NetPeeringList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetPeering `json:"items"`
}

// Repository type metadata.
var (
	NetPeering_Kind             = "NetPeering"
	NetPeering_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: NetPeering_Kind}.String()
	NetPeering_KindAPIVersion   = NetPeering_Kind + "." + CRDGroupVersion.String()
	NetPeering_GroupVersionKind = CRDGroupVersion.WithKind(NetPeering_Kind)
)

func init() {
	SchemeBuilder.Register(&NetPeering{}, &NetPeeringList{})
}
