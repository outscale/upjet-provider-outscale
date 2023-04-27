//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetToVirtualGatewayLinksObservation) DeepCopyInto(out *NetToVirtualGatewayLinksObservation) {
	*out = *in
	if in.NetID != nil {
		in, out := &in.NetID, &out.NetID
		*out = new(string)
		**out = **in
	}
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetToVirtualGatewayLinksObservation.
func (in *NetToVirtualGatewayLinksObservation) DeepCopy() *NetToVirtualGatewayLinksObservation {
	if in == nil {
		return nil
	}
	out := new(NetToVirtualGatewayLinksObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetToVirtualGatewayLinksParameters) DeepCopyInto(out *NetToVirtualGatewayLinksParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetToVirtualGatewayLinksParameters.
func (in *NetToVirtualGatewayLinksParameters) DeepCopy() *NetToVirtualGatewayLinksParameters {
	if in == nil {
		return nil
	}
	out := new(NetToVirtualGatewayLinksParameters)
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

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualGateway) DeepCopyInto(out *VirtualGateway) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualGateway.
func (in *VirtualGateway) DeepCopy() *VirtualGateway {
	if in == nil {
		return nil
	}
	out := new(VirtualGateway)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VirtualGateway) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualGatewayList) DeepCopyInto(out *VirtualGatewayList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VirtualGateway, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualGatewayList.
func (in *VirtualGatewayList) DeepCopy() *VirtualGatewayList {
	if in == nil {
		return nil
	}
	out := new(VirtualGatewayList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VirtualGatewayList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualGatewayObservation) DeepCopyInto(out *VirtualGatewayObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.NetToVirtualGatewayLinks != nil {
		in, out := &in.NetToVirtualGatewayLinks, &out.NetToVirtualGatewayLinks
		*out = make([]NetToVirtualGatewayLinksObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualGatewayObservation.
func (in *VirtualGatewayObservation) DeepCopy() *VirtualGatewayObservation {
	if in == nil {
		return nil
	}
	out := new(VirtualGatewayObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualGatewayParameters) DeepCopyInto(out *VirtualGatewayParameters) {
	*out = *in
	if in.ConnectionType != nil {
		in, out := &in.ConnectionType, &out.ConnectionType
		*out = new(string)
		**out = **in
	}
	if in.NetToVirtualGatewayLinks != nil {
		in, out := &in.NetToVirtualGatewayLinks, &out.NetToVirtualGatewayLinks
		*out = make([]NetToVirtualGatewayLinksParameters, len(*in))
		copy(*out, *in)
	}
	if in.RequestID != nil {
		in, out := &in.RequestID, &out.RequestID
		*out = new(string)
		**out = **in
	}
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]TagsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.VirtualGatewayID != nil {
		in, out := &in.VirtualGatewayID, &out.VirtualGatewayID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualGatewayParameters.
func (in *VirtualGatewayParameters) DeepCopy() *VirtualGatewayParameters {
	if in == nil {
		return nil
	}
	out := new(VirtualGatewayParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualGatewaySpec) DeepCopyInto(out *VirtualGatewaySpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualGatewaySpec.
func (in *VirtualGatewaySpec) DeepCopy() *VirtualGatewaySpec {
	if in == nil {
		return nil
	}
	out := new(VirtualGatewaySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualGatewayStatus) DeepCopyInto(out *VirtualGatewayStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualGatewayStatus.
func (in *VirtualGatewayStatus) DeepCopy() *VirtualGatewayStatus {
	if in == nil {
		return nil
	}
	out := new(VirtualGatewayStatus)
	in.DeepCopyInto(out)
	return out
}