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

type InternetServiceObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	InternetServiceID *string `json:"internetServiceId,omitempty" tf:"internet_service_id,omitempty"`

	NetID *string `json:"netId,omitempty" tf:"net_id,omitempty"`

	RequestID *string `json:"requestId,omitempty" tf:"request_id,omitempty"`

	State *string `json:"state,omitempty" tf:"state,omitempty"`
}

type InternetServiceParameters struct {

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

// InternetServiceSpec defines the desired state of InternetService
type InternetServiceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     InternetServiceParameters `json:"forProvider"`
}

// InternetServiceStatus defines the observed state of InternetService.
type InternetServiceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        InternetServiceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// InternetService is the Schema for the InternetServices API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,upjet-provider-outscale}
type InternetService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              InternetServiceSpec   `json:"spec"`
	Status            InternetServiceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// InternetServiceList contains a list of InternetServices
type InternetServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []InternetService `json:"items"`
}

// Repository type metadata.
var (
	InternetService_Kind             = "InternetService"
	InternetService_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: InternetService_Kind}.String()
	InternetService_KindAPIVersion   = InternetService_Kind + "." + CRDGroupVersion.String()
	InternetService_GroupVersionKind = CRDGroupVersion.WithKind(InternetService_Kind)
)

func init() {
	SchemeBuilder.Register(&InternetService{}, &InternetServiceList{})
}
