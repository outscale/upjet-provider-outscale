/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	v1alpha1 "github.com/outscale-vbr/upjet-provider-outscale/apis/publicip/v1alpha1"
	v1alpha11 "github.com/outscale-vbr/upjet-provider-outscale/apis/subnet/v1alpha1"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Nat.
func (mg *Nat) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.PublicIPID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.PublicIPIDRef,
		Selector:     mg.Spec.ForProvider.PublicIPIDSelector,
		To: reference.To{
			List:    &v1alpha1.PublicIpList{},
			Managed: &v1alpha1.PublicIp{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.PublicIPID")
	}
	mg.Spec.ForProvider.PublicIPID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.PublicIPIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SubnetID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.SubnetIDRef,
		Selector:     mg.Spec.ForProvider.SubnetIDSelector,
		To: reference.To{
			List:    &v1alpha11.SubnetList{},
			Managed: &v1alpha11.Subnet{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SubnetID")
	}
	mg.Spec.ForProvider.SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SubnetIDRef = rsp.ResolvedReference

	return nil
}
