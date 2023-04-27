//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LinkRouteTablesObservation) DeepCopyInto(out *LinkRouteTablesObservation) {
	*out = *in
	if in.LinkRouteTableID != nil {
		in, out := &in.LinkRouteTableID, &out.LinkRouteTableID
		*out = new(string)
		**out = **in
	}
	if in.Main != nil {
		in, out := &in.Main, &out.Main
		*out = new(bool)
		**out = **in
	}
	if in.RouteTableID != nil {
		in, out := &in.RouteTableID, &out.RouteTableID
		*out = new(string)
		**out = **in
	}
	if in.RouteTableToSubnetLinkID != nil {
		in, out := &in.RouteTableToSubnetLinkID, &out.RouteTableToSubnetLinkID
		*out = new(string)
		**out = **in
	}
	if in.SubnetID != nil {
		in, out := &in.SubnetID, &out.SubnetID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LinkRouteTablesObservation.
func (in *LinkRouteTablesObservation) DeepCopy() *LinkRouteTablesObservation {
	if in == nil {
		return nil
	}
	out := new(LinkRouteTablesObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LinkRouteTablesParameters) DeepCopyInto(out *LinkRouteTablesParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LinkRouteTablesParameters.
func (in *LinkRouteTablesParameters) DeepCopy() *LinkRouteTablesParameters {
	if in == nil {
		return nil
	}
	out := new(LinkRouteTablesParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoutePropagatingVirtualGatewaysObservation) DeepCopyInto(out *RoutePropagatingVirtualGatewaysObservation) {
	*out = *in
	if in.VirtualGatewayID != nil {
		in, out := &in.VirtualGatewayID, &out.VirtualGatewayID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoutePropagatingVirtualGatewaysObservation.
func (in *RoutePropagatingVirtualGatewaysObservation) DeepCopy() *RoutePropagatingVirtualGatewaysObservation {
	if in == nil {
		return nil
	}
	out := new(RoutePropagatingVirtualGatewaysObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoutePropagatingVirtualGatewaysParameters) DeepCopyInto(out *RoutePropagatingVirtualGatewaysParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoutePropagatingVirtualGatewaysParameters.
func (in *RoutePropagatingVirtualGatewaysParameters) DeepCopy() *RoutePropagatingVirtualGatewaysParameters {
	if in == nil {
		return nil
	}
	out := new(RoutePropagatingVirtualGatewaysParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RouteTable) DeepCopyInto(out *RouteTable) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouteTable.
func (in *RouteTable) DeepCopy() *RouteTable {
	if in == nil {
		return nil
	}
	out := new(RouteTable)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RouteTable) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RouteTableLink) DeepCopyInto(out *RouteTableLink) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouteTableLink.
func (in *RouteTableLink) DeepCopy() *RouteTableLink {
	if in == nil {
		return nil
	}
	out := new(RouteTableLink)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RouteTableLink) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RouteTableLinkList) DeepCopyInto(out *RouteTableLinkList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RouteTableLink, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouteTableLinkList.
func (in *RouteTableLinkList) DeepCopy() *RouteTableLinkList {
	if in == nil {
		return nil
	}
	out := new(RouteTableLinkList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RouteTableLinkList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RouteTableLinkObservation) DeepCopyInto(out *RouteTableLinkObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.LinkRouteTableID != nil {
		in, out := &in.LinkRouteTableID, &out.LinkRouteTableID
		*out = new(string)
		**out = **in
	}
	if in.Main != nil {
		in, out := &in.Main, &out.Main
		*out = new(bool)
		**out = **in
	}
	if in.RequestID != nil {
		in, out := &in.RequestID, &out.RequestID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouteTableLinkObservation.
func (in *RouteTableLinkObservation) DeepCopy() *RouteTableLinkObservation {
	if in == nil {
		return nil
	}
	out := new(RouteTableLinkObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RouteTableLinkParameters) DeepCopyInto(out *RouteTableLinkParameters) {
	*out = *in
	if in.RouteTableID != nil {
		in, out := &in.RouteTableID, &out.RouteTableID
		*out = new(string)
		**out = **in
	}
	if in.RouteTableIDRef != nil {
		in, out := &in.RouteTableIDRef, &out.RouteTableIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.RouteTableIDSelector != nil {
		in, out := &in.RouteTableIDSelector, &out.RouteTableIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.SubnetID != nil {
		in, out := &in.SubnetID, &out.SubnetID
		*out = new(string)
		**out = **in
	}
	if in.SubnetIDRef != nil {
		in, out := &in.SubnetIDRef, &out.SubnetIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.SubnetIDSelector != nil {
		in, out := &in.SubnetIDSelector, &out.SubnetIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouteTableLinkParameters.
func (in *RouteTableLinkParameters) DeepCopy() *RouteTableLinkParameters {
	if in == nil {
		return nil
	}
	out := new(RouteTableLinkParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RouteTableLinkSpec) DeepCopyInto(out *RouteTableLinkSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouteTableLinkSpec.
func (in *RouteTableLinkSpec) DeepCopy() *RouteTableLinkSpec {
	if in == nil {
		return nil
	}
	out := new(RouteTableLinkSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RouteTableLinkStatus) DeepCopyInto(out *RouteTableLinkStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouteTableLinkStatus.
func (in *RouteTableLinkStatus) DeepCopy() *RouteTableLinkStatus {
	if in == nil {
		return nil
	}
	out := new(RouteTableLinkStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RouteTableList) DeepCopyInto(out *RouteTableList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RouteTable, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouteTableList.
func (in *RouteTableList) DeepCopy() *RouteTableList {
	if in == nil {
		return nil
	}
	out := new(RouteTableList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RouteTableList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RouteTableObservation) DeepCopyInto(out *RouteTableObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.LinkRouteTables != nil {
		in, out := &in.LinkRouteTables, &out.LinkRouteTables
		*out = make([]LinkRouteTablesObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RequestID != nil {
		in, out := &in.RequestID, &out.RequestID
		*out = new(string)
		**out = **in
	}
	if in.RoutePropagatingVirtualGateways != nil {
		in, out := &in.RoutePropagatingVirtualGateways, &out.RoutePropagatingVirtualGateways
		*out = make([]RoutePropagatingVirtualGatewaysObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RouteTableID != nil {
		in, out := &in.RouteTableID, &out.RouteTableID
		*out = new(string)
		**out = **in
	}
	if in.Routes != nil {
		in, out := &in.Routes, &out.Routes
		*out = make([]RoutesObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouteTableObservation.
func (in *RouteTableObservation) DeepCopy() *RouteTableObservation {
	if in == nil {
		return nil
	}
	out := new(RouteTableObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RouteTableParameters) DeepCopyInto(out *RouteTableParameters) {
	*out = *in
	if in.NetID != nil {
		in, out := &in.NetID, &out.NetID
		*out = new(string)
		**out = **in
	}
	if in.NetIDRef != nil {
		in, out := &in.NetIDRef, &out.NetIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.NetIDSelector != nil {
		in, out := &in.NetIDSelector, &out.NetIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]TagsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouteTableParameters.
func (in *RouteTableParameters) DeepCopy() *RouteTableParameters {
	if in == nil {
		return nil
	}
	out := new(RouteTableParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RouteTableSpec) DeepCopyInto(out *RouteTableSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouteTableSpec.
func (in *RouteTableSpec) DeepCopy() *RouteTableSpec {
	if in == nil {
		return nil
	}
	out := new(RouteTableSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RouteTableStatus) DeepCopyInto(out *RouteTableStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouteTableStatus.
func (in *RouteTableStatus) DeepCopy() *RouteTableStatus {
	if in == nil {
		return nil
	}
	out := new(RouteTableStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoutesObservation) DeepCopyInto(out *RoutesObservation) {
	*out = *in
	if in.CreationMethod != nil {
		in, out := &in.CreationMethod, &out.CreationMethod
		*out = new(string)
		**out = **in
	}
	if in.DestinationIPRange != nil {
		in, out := &in.DestinationIPRange, &out.DestinationIPRange
		*out = new(string)
		**out = **in
	}
	if in.DestinationServiceID != nil {
		in, out := &in.DestinationServiceID, &out.DestinationServiceID
		*out = new(string)
		**out = **in
	}
	if in.GatewayID != nil {
		in, out := &in.GatewayID, &out.GatewayID
		*out = new(string)
		**out = **in
	}
	if in.NATServiceID != nil {
		in, out := &in.NATServiceID, &out.NATServiceID
		*out = new(string)
		**out = **in
	}
	if in.NetAccessPointID != nil {
		in, out := &in.NetAccessPointID, &out.NetAccessPointID
		*out = new(string)
		**out = **in
	}
	if in.NetPeeringID != nil {
		in, out := &in.NetPeeringID, &out.NetPeeringID
		*out = new(string)
		**out = **in
	}
	if in.NicID != nil {
		in, out := &in.NicID, &out.NicID
		*out = new(string)
		**out = **in
	}
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(string)
		**out = **in
	}
	if in.VMAccountID != nil {
		in, out := &in.VMAccountID, &out.VMAccountID
		*out = new(string)
		**out = **in
	}
	if in.VMID != nil {
		in, out := &in.VMID, &out.VMID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoutesObservation.
func (in *RoutesObservation) DeepCopy() *RoutesObservation {
	if in == nil {
		return nil
	}
	out := new(RoutesObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoutesParameters) DeepCopyInto(out *RoutesParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoutesParameters.
func (in *RoutesParameters) DeepCopy() *RoutesParameters {
	if in == nil {
		return nil
	}
	out := new(RoutesParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TagsObservation) DeepCopyInto(out *TagsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TagsObservation.
func (in *TagsObservation) DeepCopy() *TagsObservation {
	if in == nil {
		return nil
	}
	out := new(TagsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TagsParameters) DeepCopyInto(out *TagsParameters) {
	*out = *in
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TagsParameters.
func (in *TagsParameters) DeepCopy() *TagsParameters {
	if in == nil {
		return nil
	}
	out := new(TagsParameters)
	in.DeepCopyInto(out)
	return out
}