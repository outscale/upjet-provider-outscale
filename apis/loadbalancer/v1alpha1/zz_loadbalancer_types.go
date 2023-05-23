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

type ApplicationStickyCookiePoliciesObservation struct {
	CookieName *string `json:"cookieName,omitempty" tf:"cookie_name,omitempty"`

	PolicyName *string `json:"policyName,omitempty" tf:"policy_name,omitempty"`
}

type ApplicationStickyCookiePoliciesParameters struct {
}

type ListenersObservation struct {
	PolicyNames []*string `json:"policyNames,omitempty" tf:"policy_names,omitempty"`
}

type ListenersParameters struct {

	// +kubebuilder:validation:Required
	BackendPort *float64 `json:"backendPort" tf:"backend_port,omitempty"`

	// +kubebuilder:validation:Required
	BackendProtocol *string `json:"backendProtocol" tf:"backend_protocol,omitempty"`

	// +kubebuilder:validation:Required
	LoadBalancerPort *float64 `json:"loadBalancerPort" tf:"load_balancer_port,omitempty"`

	// +kubebuilder:validation:Required
	LoadBalancerProtocol *string `json:"loadBalancerProtocol" tf:"load_balancer_protocol,omitempty"`

	// +kubebuilder:validation:Optional
	ServerCertificateID *string `json:"serverCertificateId,omitempty" tf:"server_certificate_id,omitempty"`
}

type LoadBalancerObservation struct {
	ApplicationStickyCookiePolicies []ApplicationStickyCookiePoliciesObservation `json:"applicationStickyCookiePolicies,omitempty" tf:"application_sticky_cookie_policies,omitempty"`

	DNSName *string `json:"dnsName,omitempty" tf:"dns_name,omitempty"`

	HealthCheck map[string]*string `json:"healthCheck,omitempty" tf:"health_check,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// +kubebuilder:validation:Required
	Listeners []ListenersObservation `json:"listeners,omitempty" tf:"listeners,omitempty"`

	LoadBalancerStickyCookiePolicies []LoadBalancerStickyCookiePoliciesObservation `json:"loadBalancerStickyCookiePolicies,omitempty" tf:"load_balancer_sticky_cookie_policies,omitempty"`

	NetID *string `json:"netId,omitempty" tf:"net_id,omitempty"`

	RequestID *string `json:"requestId,omitempty" tf:"request_id,omitempty"`

	SourceSecurityGroup map[string]*string `json:"sourceSecurityGroup,omitempty" tf:"source_security_group,omitempty"`
}

type LoadBalancerParameters struct {

	// +kubebuilder:validation:Optional
	AccessLog map[string]*string `json:"accessLog,omitempty" tf:"access_log,omitempty"`

	// +kubebuilder:validation:Optional
	BackendVMIds []*string `json:"backendVmIds,omitempty" tf:"backend_vm_ids,omitempty"`

	// +kubebuilder:validation:Required
	Listeners []ListenersParameters `json:"listeners" tf:"listeners,omitempty"`

	// +kubebuilder:validation:Required
	LoadBalancerName *string `json:"loadBalancerName" tf:"load_balancer_name,omitempty"`

	// +kubebuilder:validation:Optional
	LoadBalancerType *string `json:"loadBalancerType,omitempty" tf:"load_balancer_type,omitempty"`

	// +kubebuilder:validation:Optional
	PublicIP *string `json:"publicIp,omitempty" tf:"public_ip,omitempty"`

	// +kubebuilder:validation:Optional
	SecuredCookies *bool `json:"securedCookies,omitempty" tf:"secured_cookies,omitempty"`

	// +crossplane:generate:reference:type=github.com/outscale/upjet-provider-outscale/apis/securitygroup/v1alpha1.SecurityGroup
	// +kubebuilder:validation:Optional
	SecurityGroups []*string `json:"securityGroups,omitempty" tf:"security_groups,omitempty"`

	// References to SecurityGroup in securitygroup to populate securityGroups.
	// +kubebuilder:validation:Optional
	SecurityGroupsRefs []v1.Reference `json:"securityGroupsRefs,omitempty" tf:"-"`

	// Selector for a list of SecurityGroup in securitygroup to populate securityGroups.
	// +kubebuilder:validation:Optional
	SecurityGroupsSelector *v1.Selector `json:"securityGroupsSelector,omitempty" tf:"-"`

	// +crossplane:generate:reference:type=github.com/outscale/upjet-provider-outscale/apis/subnet/v1alpha1.Subnet
	// +kubebuilder:validation:Optional
	Subnets []*string `json:"subnets,omitempty" tf:"subnets,omitempty"`

	// References to Subnet in subnet to populate subnets.
	// +kubebuilder:validation:Optional
	SubnetsRefs []v1.Reference `json:"subnetsRefs,omitempty" tf:"-"`

	// Selector for a list of Subnet in subnet to populate subnets.
	// +kubebuilder:validation:Optional
	SubnetsSelector *v1.Selector `json:"subnetsSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	SubregionNames []*string `json:"subregionNames,omitempty" tf:"subregion_names,omitempty"`

	// +kubebuilder:validation:Optional
	Tags []TagsParameters `json:"tags,omitempty" tf:"tags,omitempty"`
}

type LoadBalancerStickyCookiePoliciesObservation struct {
	PolicyName *string `json:"policyName,omitempty" tf:"policy_name,omitempty"`
}

type LoadBalancerStickyCookiePoliciesParameters struct {
}

type TagsObservation struct {
}

type TagsParameters struct {

	// +kubebuilder:validation:Optional
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

// LoadBalancerSpec defines the desired state of LoadBalancer
type LoadBalancerSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LoadBalancerParameters `json:"forProvider"`
}

// LoadBalancerStatus defines the observed state of LoadBalancer.
type LoadBalancerStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LoadBalancerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LoadBalancer is the Schema for the LoadBalancers API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,upjet-provider-outscale}
type LoadBalancer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LoadBalancerSpec   `json:"spec"`
	Status            LoadBalancerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LoadBalancerList contains a list of LoadBalancers
type LoadBalancerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LoadBalancer `json:"items"`
}

// Repository type metadata.
var (
	LoadBalancer_Kind             = "LoadBalancer"
	LoadBalancer_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LoadBalancer_Kind}.String()
	LoadBalancer_KindAPIVersion   = LoadBalancer_Kind + "." + CRDGroupVersion.String()
	LoadBalancer_GroupVersionKind = CRDGroupVersion.WithKind(LoadBalancer_Kind)
)

func init() {
	SchemeBuilder.Register(&LoadBalancer{}, &LoadBalancerList{})
}
