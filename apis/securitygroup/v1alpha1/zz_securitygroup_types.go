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

type InboundRulesObservation struct {
	FromPortRange *float64 `json:"fromPortRange,omitempty" tf:"from_port_range,omitempty"`

	IPProtocol *string `json:"ipProtocol,omitempty" tf:"ip_protocol,omitempty"`

	IPRanges []*string `json:"ipRanges,omitempty" tf:"ip_ranges,omitempty"`

	SecurityGroupsMembers []map[string]*string `json:"securityGroupsMembers,omitempty" tf:"security_groups_members,omitempty"`

	ToPortRange *float64 `json:"toPortRange,omitempty" tf:"to_port_range,omitempty"`
}

type InboundRulesParameters struct {
}

type OutboundRulesObservation struct {
	FromPortRange *float64 `json:"fromPortRange,omitempty" tf:"from_port_range,omitempty"`

	IPProtocol *string `json:"ipProtocol,omitempty" tf:"ip_protocol,omitempty"`

	IPRanges []*string `json:"ipRanges,omitempty" tf:"ip_ranges,omitempty"`

	SecurityGroupsMembers []map[string]*string `json:"securityGroupsMembers,omitempty" tf:"security_groups_members,omitempty"`

	ToPortRange *float64 `json:"toPortRange,omitempty" tf:"to_port_range,omitempty"`
}

type OutboundRulesParameters struct {
}

type SecurityGroupObservation struct {
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	InboundRules []InboundRulesObservation `json:"inboundRules,omitempty" tf:"inbound_rules,omitempty"`

	OutboundRules []OutboundRulesObservation `json:"outboundRules,omitempty" tf:"outbound_rules,omitempty"`

	RequestID *string `json:"requestId,omitempty" tf:"request_id,omitempty"`

	SecurityGroupID *string `json:"securityGroupId,omitempty" tf:"security_group_id,omitempty"`
}

type SecurityGroupParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +crossplane:generate:reference:type=github.com/outscale-vbr/upjet-provider-outscale/apis/net/v1alpha1.Net
	// +kubebuilder:validation:Optional
	NetID *string `json:"netId,omitempty" tf:"net_id,omitempty"`

	// Reference to a Net in net to populate netId.
	// +kubebuilder:validation:Optional
	NetIDRef *v1.Reference `json:"netIdRef,omitempty" tf:"-"`

	// Selector for a Net in net to populate netId.
	// +kubebuilder:validation:Optional
	NetIDSelector *v1.Selector `json:"netIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	RemoveDefaultOutboundRule *bool `json:"removeDefaultOutboundRule,omitempty" tf:"remove_default_outbound_rule,omitempty"`

	// +kubebuilder:validation:Optional
	SecurityGroupName *string `json:"securityGroupName,omitempty" tf:"security_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	Tag map[string]*string `json:"tag,omitempty" tf:"tag,omitempty"`

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

// SecurityGroupSpec defines the desired state of SecurityGroup
type SecurityGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SecurityGroupParameters `json:"forProvider"`
}

// SecurityGroupStatus defines the observed state of SecurityGroup.
type SecurityGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SecurityGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SecurityGroup is the Schema for the SecurityGroups API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,upjet-provider-outscale}
type SecurityGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SecurityGroupSpec   `json:"spec"`
	Status            SecurityGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SecurityGroupList contains a list of SecurityGroups
type SecurityGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecurityGroup `json:"items"`
}

// Repository type metadata.
var (
	SecurityGroup_Kind             = "SecurityGroup"
	SecurityGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SecurityGroup_Kind}.String()
	SecurityGroup_KindAPIVersion   = SecurityGroup_Kind + "." + CRDGroupVersion.String()
	SecurityGroup_GroupVersionKind = CRDGroupVersion.WithKind(SecurityGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&SecurityGroup{}, &SecurityGroupList{})
}