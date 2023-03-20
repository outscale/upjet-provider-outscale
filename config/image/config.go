package image

import "github.com/upbound/upjet/pkg/config"

const imageShortGroup = "image"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("outscale_image", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = imageShortGroup
		r.Kind = "Image"
		r.References["source_image_id"] = config.Reference {
			Type: "Image",
		}
		r.References["vm_id"] = config.Reference{
			Type: "github.com/outscale-vbr/upjet-provider-outscale/apis/vm/v1alpha1.Vm",
		}

	})
	p.AddResourceConfigurator("outscale_image_export_task", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = imageShortGroup
		r.References["image_id"] = config.Reference {
			Type: "Image",
		}
	})
	p.AddResourceConfigurator("outscale_image_launch_permission", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = imageShortGroup
		r.References["image_id"] = config.Reference {
			Type: "Image",
		}
	})
}