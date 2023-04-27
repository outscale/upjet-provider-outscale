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

type VolumesLinkObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	RequestID *string `json:"requestId,omitempty" tf:"request_id,omitempty"`

	State *string `json:"state,omitempty" tf:"state,omitempty"`
}

type VolumesLinkParameters struct {

	// +kubebuilder:validation:Optional
	DeleteOnVMTermination *bool `json:"deleteOnVmTermination,omitempty" tf:"delete_on_vm_termination,omitempty"`

	// +kubebuilder:validation:Required
	DeviceName *string `json:"deviceName" tf:"device_name,omitempty"`

	// +kubebuilder:validation:Optional
	ForceUnlink *bool `json:"forceUnlink,omitempty" tf:"force_unlink,omitempty"`

	// +crossplane:generate:reference:type=github.com/outscale-vbr/upjet-provider-outscale/apis/vm/v1alpha1.Vm
	// +kubebuilder:validation:Optional
	VMID *string `json:"vmId,omitempty" tf:"vm_id,omitempty"`

	// Reference to a Vm in vm to populate vmId.
	// +kubebuilder:validation:Optional
	VMIDRef *v1.Reference `json:"vmIdRef,omitempty" tf:"-"`

	// Selector for a Vm in vm to populate vmId.
	// +kubebuilder:validation:Optional
	VMIDSelector *v1.Selector `json:"vmIdSelector,omitempty" tf:"-"`

	// +crossplane:generate:reference:type=Volume
	// +kubebuilder:validation:Optional
	VolumeID *string `json:"volumeId,omitempty" tf:"volume_id,omitempty"`

	// Reference to a Volume to populate volumeId.
	// +kubebuilder:validation:Optional
	VolumeIDRef *v1.Reference `json:"volumeIdRef,omitempty" tf:"-"`

	// Selector for a Volume to populate volumeId.
	// +kubebuilder:validation:Optional
	VolumeIDSelector *v1.Selector `json:"volumeIdSelector,omitempty" tf:"-"`
}

// VolumesLinkSpec defines the desired state of VolumesLink
type VolumesLinkSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VolumesLinkParameters `json:"forProvider"`
}

// VolumesLinkStatus defines the observed state of VolumesLink.
type VolumesLinkStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VolumesLinkObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VolumesLink is the Schema for the VolumesLinks API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,upjet-provider-outscale}
type VolumesLink struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VolumesLinkSpec   `json:"spec"`
	Status            VolumesLinkStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VolumesLinkList contains a list of VolumesLinks
type VolumesLinkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VolumesLink `json:"items"`
}

// Repository type metadata.
var (
	VolumesLink_Kind             = "VolumesLink"
	VolumesLink_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VolumesLink_Kind}.String()
	VolumesLink_KindAPIVersion   = VolumesLink_Kind + "." + CRDGroupVersion.String()
	VolumesLink_GroupVersionKind = CRDGroupVersion.WithKind(VolumesLink_Kind)
)

func init() {
	SchemeBuilder.Register(&VolumesLink{}, &VolumesLinkList{})
}