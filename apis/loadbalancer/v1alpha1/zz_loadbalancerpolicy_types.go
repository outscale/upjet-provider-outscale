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

type LoadBalancerPolicyApplicationStickyCookiePoliciesObservation struct {
	CookieName *string `json:"cookieName,omitempty" tf:"cookie_name,omitempty"`

	PolicyName *string `json:"policyName,omitempty" tf:"policy_name,omitempty"`
}

type LoadBalancerPolicyApplicationStickyCookiePoliciesParameters struct {
}

type LoadBalancerPolicyListenersObservation struct {
	BackendPort *float64 `json:"backendPort,omitempty" tf:"backend_port,omitempty"`

	BackendProtocol *string `json:"backendProtocol,omitempty" tf:"backend_protocol,omitempty"`

	LoadBalancerPort *float64 `json:"loadBalancerPort,omitempty" tf:"load_balancer_port,omitempty"`

	LoadBalancerProtocol *string `json:"loadBalancerProtocol,omitempty" tf:"load_balancer_protocol,omitempty"`

	PolicyNames []*string `json:"policyNames,omitempty" tf:"policy_names,omitempty"`

	ServerCertificateID *string `json:"serverCertificateId,omitempty" tf:"server_certificate_id,omitempty"`
}

type LoadBalancerPolicyListenersParameters struct {
}

type LoadBalancerPolicyLoadBalancerStickyCookiePoliciesObservation struct {
	PolicyName *string `json:"policyName,omitempty" tf:"policy_name,omitempty"`
}

type LoadBalancerPolicyLoadBalancerStickyCookiePoliciesParameters struct {
}

type LoadBalancerPolicyObservation struct {
	AccessLog map[string]*string `json:"accessLog,omitempty" tf:"access_log,omitempty"`

	ApplicationStickyCookiePolicies []LoadBalancerPolicyApplicationStickyCookiePoliciesObservation `json:"applicationStickyCookiePolicies,omitempty" tf:"application_sticky_cookie_policies,omitempty"`

	DNSName *string `json:"dnsName,omitempty" tf:"dns_name,omitempty"`

	HealthCheck map[string]*string `json:"healthCheck,omitempty" tf:"health_check,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Listeners []LoadBalancerPolicyListenersObservation `json:"listeners,omitempty" tf:"listeners,omitempty"`

	LoadBalancerStickyCookiePolicies []LoadBalancerPolicyLoadBalancerStickyCookiePoliciesObservation `json:"loadBalancerStickyCookiePolicies,omitempty" tf:"load_balancer_sticky_cookie_policies,omitempty"`

	NetID *string `json:"netId,omitempty" tf:"net_id,omitempty"`

	PublicIP *string `json:"publicIp,omitempty" tf:"public_ip,omitempty"`

	RequestID *string `json:"requestId,omitempty" tf:"request_id,omitempty"`

	SecuredCookies *bool `json:"securedCookies,omitempty" tf:"secured_cookies,omitempty"`

	SourceSecurityGroup map[string]*string `json:"sourceSecurityGroup,omitempty" tf:"source_security_group,omitempty"`

	// +kubebuilder:validation:Optional
	Tags []LoadBalancerPolicyTagsObservation `json:"tags,omitempty" tf:"tags,omitempty"`
}

type LoadBalancerPolicyParameters struct {

	// +kubebuilder:validation:Optional
	BackendVMIds []*string `json:"backendVmIds,omitempty" tf:"backend_vm_ids,omitempty"`

	// +kubebuilder:validation:Optional
	CookieExpirationPeriod *float64 `json:"cookieExpirationPeriod,omitempty" tf:"cookie_expiration_period,omitempty"`

	// +kubebuilder:validation:Optional
	CookieName *string `json:"cookieName,omitempty" tf:"cookie_name,omitempty"`

	// +kubebuilder:validation:Required
	LoadBalancerName *string `json:"loadBalancerName" tf:"load_balancer_name,omitempty"`

	// +kubebuilder:validation:Optional
	LoadBalancerType *string `json:"loadBalancerType,omitempty" tf:"load_balancer_type,omitempty"`

	// +kubebuilder:validation:Required
	PolicyName *string `json:"policyName" tf:"policy_name,omitempty"`

	// +kubebuilder:validation:Required
	PolicyType *string `json:"policyType" tf:"policy_type,omitempty"`

	// +kubebuilder:validation:Optional
	SecurityGroups []*string `json:"securityGroups,omitempty" tf:"security_groups,omitempty"`

	// +kubebuilder:validation:Optional
	Subnets []*string `json:"subnets,omitempty" tf:"subnets,omitempty"`

	// +kubebuilder:validation:Optional
	SubregionNames []*string `json:"subregionNames,omitempty" tf:"subregion_names,omitempty"`

	// +kubebuilder:validation:Optional
	Tags []LoadBalancerPolicyTagsParameters `json:"tags,omitempty" tf:"tags,omitempty"`
}

type LoadBalancerPolicyTagsObservation struct {
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type LoadBalancerPolicyTagsParameters struct {
}

// LoadBalancerPolicySpec defines the desired state of LoadBalancerPolicy
type LoadBalancerPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LoadBalancerPolicyParameters `json:"forProvider"`
}

// LoadBalancerPolicyStatus defines the observed state of LoadBalancerPolicy.
type LoadBalancerPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LoadBalancerPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LoadBalancerPolicy is the Schema for the LoadBalancerPolicys API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,upjet-provider-outscale}
type LoadBalancerPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LoadBalancerPolicySpec   `json:"spec"`
	Status            LoadBalancerPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LoadBalancerPolicyList contains a list of LoadBalancerPolicys
type LoadBalancerPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LoadBalancerPolicy `json:"items"`
}

// Repository type metadata.
var (
	LoadBalancerPolicy_Kind             = "LoadBalancerPolicy"
	LoadBalancerPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LoadBalancerPolicy_Kind}.String()
	LoadBalancerPolicy_KindAPIVersion   = LoadBalancerPolicy_Kind + "." + CRDGroupVersion.String()
	LoadBalancerPolicy_GroupVersionKind = CRDGroupVersion.WithKind(LoadBalancerPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&LoadBalancerPolicy{}, &LoadBalancerPolicyList{})
}
