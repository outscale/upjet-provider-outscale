package subnet

import "github.com/upbound/upjet/pkg/config"

const subnetShortGroup = "subnet"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider){
	p.AddResourceConfigurator("outscale_subnet", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = subnetShortGroup
		r.Kind = "Subnet"
		r.References["net_id"] = config.Reference {
			Type: "github.com/outscale-vbr/upjet-provider-outscale/apis/net/v1alpha1.Net",
		}
	})
}
