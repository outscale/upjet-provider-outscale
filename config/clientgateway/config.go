package clientgateway

import "github.com/upbound/upjet/pkg/config"

const clientGatewayShortGroup = "clientgateway"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("outscale_client_gateway", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = clientGatewayShortGroup
		r.Kind = "ClientGateway"
	})
}
