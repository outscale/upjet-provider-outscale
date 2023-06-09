/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	v1alpha1 "github.com/outscale/upjet-provider-outscale/apis/volume/v1alpha1"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Snapshot.
func (mg *Snapshot) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SourceSnapshotID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.SourceSnapshotIDRef,
		Selector:     mg.Spec.ForProvider.SourceSnapshotIDSelector,
		To: reference.To{
			List:    &SnapshotList{},
			Managed: &Snapshot{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SourceSnapshotID")
	}
	mg.Spec.ForProvider.SourceSnapshotID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SourceSnapshotIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VolumeID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.VolumeIDRef,
		Selector:     mg.Spec.ForProvider.VolumeIDSelector,
		To: reference.To{
			List:    &v1alpha1.VolumeList{},
			Managed: &v1alpha1.Volume{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VolumeID")
	}
	mg.Spec.ForProvider.VolumeID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VolumeIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this SnapshotAttributes.
func (mg *SnapshotAttributes) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SnapshotID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.SnapshotIDRef,
		Selector:     mg.Spec.ForProvider.SnapshotIDSelector,
		To: reference.To{
			List:    &SnapshotList{},
			Managed: &Snapshot{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SnapshotID")
	}
	mg.Spec.ForProvider.SnapshotID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SnapshotIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this SnapshotExportTask.
func (mg *SnapshotExportTask) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SnapshotID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.SnapshotIDRef,
		Selector:     mg.Spec.ForProvider.SnapshotIDSelector,
		To: reference.To{
			List:    &SnapshotList{},
			Managed: &Snapshot{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SnapshotID")
	}
	mg.Spec.ForProvider.SnapshotID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SnapshotIDRef = rsp.ResolvedReference

	return nil
}
