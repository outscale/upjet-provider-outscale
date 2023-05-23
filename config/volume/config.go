package volume

import "github.com/upbound/upjet/pkg/config"

const volumeShortGroup = "volume"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("outscale_volume", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = volumeShortGroup
		r.Kind = "Volume"
	})

	p.AddResourceConfigurator("outscale_volumes_link", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = volumeShortGroup
		r.Kind = "VolumesLink"
		r.References["vm_id"] = config.Reference{
			Type: "github.com/outscale/upjet-provider-outscale/apis/vm/v1alpha1.Vm",
		}
		r.References["volume_id"] = config.Reference{
			Type: "Volume",
		}
	})
}
