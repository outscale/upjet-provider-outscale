package nat

import "github.com/upbound/upjet/pkg/config"

const natShortGroup = "nat"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("outscale_nat_service", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = natShortGroup
		r.Kind = "Nat"
		r.References["public_ip_id"] = config.Reference{
			Type: "github.com/outscale-vbr/upjet-provider-outscale/apis/publicip/v1alpha1.PublicIp",
		}
		r.References["subnet_id"] = config.Reference{
			Type: "github.com/outscale-vbr/upjet-provider-outscale/apis/subnet/v1alpha1.Subnet",
		}
	})
}
