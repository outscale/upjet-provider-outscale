package routetable

import "github.com/upbound/upjet/pkg/config"

const routeTableShortGroup = "routetable"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("outscale_route_table", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = routeTableShortGroup
		r.Kind = "RouteTable"
		r.References["net_id"] = config.Reference{
			Type: "github.com/outscale-vbr/upjet-provider-outscale/apis/net/v1alpha1.Net",
		}
	})
	p.AddResourceConfigurator("outscale_route_table_link", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = routeTableShortGroup
		r.Kind = "RouteTableLink"
		r.References["route_table_id"] = config.Reference{
			Type: "RouteTable",
		}
		r.References["subnet_id"] = config.Reference{
			Type: "github.com/outscale-vbr/upjet-provider-outscale/apis/subnet/v1alpha1.Subnet",
		}
	})
}
