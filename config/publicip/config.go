package publicip

import "github.com/upbound/upjet/pkg/config"

const publicIpShortGroup = "publicip"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("outscale_public_ip", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = publicIpShortGroup
		r.Kind = "PublicIp"
	})
	p.AddResourceConfigurator("outscale_public_ip_link", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = publicIpShortGroup
		r.Kind = "PublicIpLink"
		r.References["public_ip_id"] = config.Reference{
			Type: "PublicIp",
		}
		r.References["vm_id"] = config.Reference{
			Type: "github.com/outscale/upjet-provider-outscale/apis/vm/v1alpha1.Vm",
		}
	})
}
