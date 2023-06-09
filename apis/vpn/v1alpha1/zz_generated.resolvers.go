/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	v1alpha1 "github.com/outscale/upjet-provider-outscale/apis/clientgateway/v1alpha1"
	v1alpha11 "github.com/outscale/upjet-provider-outscale/apis/virtualgateway/v1alpha1"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this VpnConnection.
func (mg *VpnConnection) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ClientGatewayID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ClientGatewayIDRef,
		Selector:     mg.Spec.ForProvider.ClientGatewayIDSelector,
		To: reference.To{
			List:    &v1alpha1.ClientGatewayList{},
			Managed: &v1alpha1.ClientGateway{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ClientGatewayID")
	}
	mg.Spec.ForProvider.ClientGatewayID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ClientGatewayIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VirtualGatewayID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.VirtualGatewayIDRef,
		Selector:     mg.Spec.ForProvider.VirtualGatewayIDSelector,
		To: reference.To{
			List:    &v1alpha11.VirtualGatewayList{},
			Managed: &v1alpha11.VirtualGateway{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VirtualGatewayID")
	}
	mg.Spec.ForProvider.VirtualGatewayID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VirtualGatewayIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this VpnConnectionRoute.
func (mg *VpnConnectionRoute) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VPNConnectionID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.VPNConnectionIDRef,
		Selector:     mg.Spec.ForProvider.VPNConnectionIDSelector,
		To: reference.To{
			List:    &VpnConnectionList{},
			Managed: &VpnConnection{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VPNConnectionID")
	}
	mg.Spec.ForProvider.VPNConnectionID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VPNConnectionIDRef = rsp.ResolvedReference

	return nil
}
