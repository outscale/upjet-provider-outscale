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

type PublicIpLinkObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	LinkPublicIPID *string `json:"linkPublicIpId,omitempty" tf:"link_public_ip_id,omitempty"`

	NicAccountID *string `json:"nicAccountId,omitempty" tf:"nic_account_id,omitempty"`

	RequestID *string `json:"requestId,omitempty" tf:"request_id,omitempty"`

	Tags []PublicIpLinkTagsObservation `json:"tags,omitempty" tf:"tags,omitempty"`
}

type PublicIpLinkParameters struct {

	// +kubebuilder:validation:Optional
	AllowRelink *bool `json:"allowRelink,omitempty" tf:"allow_relink,omitempty"`

	// +kubebuilder:validation:Optional
	NicID *string `json:"nicId,omitempty" tf:"nic_id,omitempty"`

	// +kubebuilder:validation:Optional
	PrivateIP *string `json:"privateIp,omitempty" tf:"private_ip,omitempty"`

	// +kubebuilder:validation:Optional
	PublicIP *string `json:"publicIp,omitempty" tf:"public_ip,omitempty"`

	// +crossplane:generate:reference:type=PublicIp
	// +kubebuilder:validation:Optional
	PublicIPID *string `json:"publicIpId,omitempty" tf:"public_ip_id,omitempty"`

	// Reference to a PublicIp to populate publicIpId.
	// +kubebuilder:validation:Optional
	PublicIPIDRef *v1.Reference `json:"publicIpIdRef,omitempty" tf:"-"`

	// Selector for a PublicIp to populate publicIpId.
	// +kubebuilder:validation:Optional
	PublicIPIDSelector *v1.Selector `json:"publicIpIdSelector,omitempty" tf:"-"`

	// +crossplane:generate:reference:type=github.com/outscale/upjet-provider-outscale/apis/vm/v1alpha1.Vm
	// +kubebuilder:validation:Optional
	VMID *string `json:"vmId,omitempty" tf:"vm_id,omitempty"`

	// Reference to a Vm in vm to populate vmId.
	// +kubebuilder:validation:Optional
	VMIDRef *v1.Reference `json:"vmIdRef,omitempty" tf:"-"`

	// Selector for a Vm in vm to populate vmId.
	// +kubebuilder:validation:Optional
	VMIDSelector *v1.Selector `json:"vmIdSelector,omitempty" tf:"-"`
}

type PublicIpLinkTagsObservation struct {
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type PublicIpLinkTagsParameters struct {
}

// PublicIpLinkSpec defines the desired state of PublicIpLink
type PublicIpLinkSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PublicIpLinkParameters `json:"forProvider"`
}

// PublicIpLinkStatus defines the observed state of PublicIpLink.
type PublicIpLinkStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PublicIpLinkObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// PublicIpLink is the Schema for the PublicIpLinks API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,upjet-provider-outscale}
type PublicIpLink struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PublicIpLinkSpec   `json:"spec"`
	Status            PublicIpLinkStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PublicIpLinkList contains a list of PublicIpLinks
type PublicIpLinkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PublicIpLink `json:"items"`
}

// Repository type metadata.
var (
	PublicIpLink_Kind             = "PublicIpLink"
	PublicIpLink_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PublicIpLink_Kind}.String()
	PublicIpLink_KindAPIVersion   = PublicIpLink_Kind + "." + CRDGroupVersion.String()
	PublicIpLink_GroupVersionKind = CRDGroupVersion.WithKind(PublicIpLink_Kind)
)

func init() {
	SchemeBuilder.Register(&PublicIpLink{}, &PublicIpLinkList{})
}
