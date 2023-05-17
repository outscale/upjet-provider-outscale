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

type NicObservation struct {
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	IsSourceDestChecked *bool `json:"isSourceDestChecked,omitempty" tf:"is_source_dest_checked,omitempty"`

	LinkNic map[string]*string `json:"linkNic,omitempty" tf:"link_nic,omitempty"`

	LinkPublicIP map[string]*string `json:"linkPublicIp,omitempty" tf:"link_public_ip,omitempty"`

	MacAddress *string `json:"macAddress,omitempty" tf:"mac_address,omitempty"`

	NetID *string `json:"netId,omitempty" tf:"net_id,omitempty"`

	NicID *string `json:"nicId,omitempty" tf:"nic_id,omitempty"`

	PrivateDNSName *string `json:"privateDnsName,omitempty" tf:"private_dns_name,omitempty"`

	// +kubebuilder:validation:Optional
	PrivateIps []PrivateIpsObservation `json:"privateIps,omitempty" tf:"private_ips,omitempty"`

	RequestID *string `json:"requestId,omitempty" tf:"request_id,omitempty"`

	RequesterManaged *bool `json:"requesterManaged,omitempty" tf:"requester_managed,omitempty"`

	SecurityGroups []SecurityGroupsObservation `json:"securityGroups,omitempty" tf:"security_groups,omitempty"`

	State *string `json:"state,omitempty" tf:"state,omitempty"`

	SubregionName *string `json:"subregionName,omitempty" tf:"subregion_name,omitempty"`
}

type NicParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	PrivateIP *string `json:"privateIp,omitempty" tf:"private_ip,omitempty"`

	// +kubebuilder:validation:Optional
	PrivateIps []PrivateIpsParameters `json:"privateIps,omitempty" tf:"private_ips,omitempty"`

	// +crossplane:generate:reference:type=github.com/outscale-vbr/upjet-provider-outscale/apis/securitygroup/v1alpha1.SecurityGroup
	// +kubebuilder:validation:Optional
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" tf:"security_group_ids,omitempty"`

	// References to SecurityGroup in securitygroup to populate securityGroupIds.
	// +kubebuilder:validation:Optional
	SecurityGroupIdsRefs []v1.Reference `json:"securityGroupIdsRefs,omitempty" tf:"-"`

	// Selector for a list of SecurityGroup in securitygroup to populate securityGroupIds.
	// +kubebuilder:validation:Optional
	SecurityGroupIdsSelector *v1.Selector `json:"securityGroupIdsSelector,omitempty" tf:"-"`

	// +crossplane:generate:reference:type=github.com/outscale-vbr/upjet-provider-outscale/apis/subnet/v1alpha1.Subnet
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Reference to a Subnet in subnet to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// Selector for a Subnet in subnet to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Tags []TagsParameters `json:"tags,omitempty" tf:"tags,omitempty"`
}

type PrivateIpsObservation struct {
	LinkPublicIP map[string]*string `json:"linkPublicIp,omitempty" tf:"link_public_ip,omitempty"`

	PrivateDNSName *string `json:"privateDnsName,omitempty" tf:"private_dns_name,omitempty"`
}

type PrivateIpsParameters struct {

	// +kubebuilder:validation:Optional
	IsPrimary *bool `json:"isPrimary,omitempty" tf:"is_primary,omitempty"`

	// +kubebuilder:validation:Optional
	PrivateIP *string `json:"privateIp,omitempty" tf:"private_ip,omitempty"`
}

type SecurityGroupsObservation struct {
	SecurityGroupID *string `json:"securityGroupId,omitempty" tf:"security_group_id,omitempty"`

	SecurityGroupName *string `json:"securityGroupName,omitempty" tf:"security_group_name,omitempty"`
}

type SecurityGroupsParameters struct {
}

type TagsObservation struct {
}

type TagsParameters struct {

	// +kubebuilder:validation:Optional
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

// NicSpec defines the desired state of Nic
type NicSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NicParameters `json:"forProvider"`
}

// NicStatus defines the observed state of Nic.
type NicStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NicObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Nic is the Schema for the Nics API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,upjet-provider-outscale}
type Nic struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NicSpec   `json:"spec"`
	Status            NicStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NicList contains a list of Nics
type NicList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Nic `json:"items"`
}

// Repository type metadata.
var (
	Nic_Kind             = "Nic"
	Nic_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Nic_Kind}.String()
	Nic_KindAPIVersion   = Nic_Kind + "." + CRDGroupVersion.String()
	Nic_GroupVersionKind = CRDGroupVersion.WithKind(Nic_Kind)
)

func init() {
	SchemeBuilder.Register(&Nic{}, &NicList{})
}
