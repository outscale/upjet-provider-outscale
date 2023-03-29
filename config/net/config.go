package net

import "github.com/upbound/upjet/pkg/config"

const netShortGroup = "net"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("outscale_net", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = netShortGroup
		r.Kind = "Net"
	})
	p.AddResourceConfigurator("outscale_net_attributes", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = netShortGroup
		r.Kind = "NetAttributes"

	})
	p.AddResourceConfigurator("outscale_net_peering", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = netShortGroup
		r.Kind = "NetPeering"
		r.References["accepter_net_id"] = config.Reference{
			Type: "Net",
		}
		r.References["source_net_id"] = config.Reference{
			Type: "Net",
		}
	})
	p.AddResourceConfigurator("outscale_net_peering_acceptation", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = netShortGroup
		r.Kind = "NetPeeringAcceptation"
		r.References["net_peering_id"] = config.Reference{
			Type: "NetPeering",
		}
	})
}
