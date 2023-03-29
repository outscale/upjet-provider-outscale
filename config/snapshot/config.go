package snapshot

import "github.com/upbound/upjet/pkg/config"

const snapshotShortGroup = "snapshot"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("outscale_snapshot", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = snapshotShortGroup
		r.Kind = "Snapshot"
		r.References["source_snapshot_id"] = config.Reference{
			Type: "Snapshot",
		}
		r.References["volume_id"] = config.Reference{
			Type: "github.com/outscale-vbr/upjet-provider-outscale/apis/volume/v1alpha1.Volume",
		}
	})
	p.AddResourceConfigurator("outscale_snapshot_export_task", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = snapshotShortGroup
		r.Kind = "SnapshotExportTask"
		r.References["snapshot_id"] = config.Reference{
			Type: "Snapshot",
		}
	})
	p.AddResourceConfigurator("outscale_snapshot_attributes", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = snapshotShortGroup
		r.Kind = "SnapshotAttributes"
		r.References["snapshot_id"] = config.Reference{
			Type: "Snapshot",
		}
	})
}
