/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	v1alpha1 "github.com/outscale/upjet-provider-outscale/apis/net/v1alpha1"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this SecurityGroup.
func (mg *SecurityGroup) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.NetID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.NetIDRef,
		Selector:     mg.Spec.ForProvider.NetIDSelector,
		To: reference.To{
			List:    &v1alpha1.NetList{},
			Managed: &v1alpha1.Net{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.NetID")
	}
	mg.Spec.ForProvider.NetID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.NetIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this SecurityGroupRule.
func (mg *SecurityGroupRule) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.Rules); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.Rules[i3].SecurityGroupsMembers); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Rules[i3].SecurityGroupsMembers[i4].SecurityGroupID),
				Extract:      reference.ExternalName(),
				Reference:    mg.Spec.ForProvider.Rules[i3].SecurityGroupsMembers[i4].SecurityGroupIDRef,
				Selector:     mg.Spec.ForProvider.Rules[i3].SecurityGroupsMembers[i4].SecurityGroupIDSelector,
				To: reference.To{
					List:    &SecurityGroupList{},
					Managed: &SecurityGroup{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.Rules[i3].SecurityGroupsMembers[i4].SecurityGroupID")
			}
			mg.Spec.ForProvider.Rules[i3].SecurityGroupsMembers[i4].SecurityGroupID = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.Rules[i3].SecurityGroupsMembers[i4].SecurityGroupIDRef = rsp.ResolvedReference

		}
	}
	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SecurityGroupID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.SecurityGroupIDRef,
		Selector:     mg.Spec.ForProvider.SecurityGroupIDSelector,
		To: reference.To{
			List:    &SecurityGroupList{},
			Managed: &SecurityGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SecurityGroupID")
	}
	mg.Spec.ForProvider.SecurityGroupID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SecurityGroupIDRef = rsp.ResolvedReference

	return nil
}
