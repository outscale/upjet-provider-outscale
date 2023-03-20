package route

import "github.com/upbound/upjet/pkg/config"

const routeShortGroup = "route"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("outscale_route", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = routeShortGroup
		r.Kind = "Route"
		r.References["gateway_id"] = config.Reference {
			Type: "github.com/outscale-vbr/upjet-provider-outscale/apis/internetservice/v1alpha1.InternetService",
		}
		r.References["route_table_id"] = config.Reference {
			Type: "github.com/outscale-vbr/upjet-provider-outscale/apis/routetable/v1alpha1.RouteTable",
		}
		r.References["nat_service_id"] = config.Reference {
			Type: "github.com/outscale-vbr/upjet-provider-outscale/apis/nat/v1alpha1.Nat",
		}
	})
}