package virtualgateway

import "github.com/upbound/upjet/pkg/config"

const virtualGatewayShortGroup = "virtualgateway"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("outscale_virtual_gateway", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = virtualGatewayShortGroup
		r.Kind = "VirtualGateway"
	})
}