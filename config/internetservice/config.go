package internetservice

import "github.com/upbound/upjet/pkg/config"

const internetServiceShortGroup = "internetService"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("outscale_internet_service", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = internetServiceShortGroup
		r.Kind = "InternetService"
	})
	p.AddResourceConfigurator("outscale_internet_service_link", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = internetServiceShortGroup
		r.Kind = "InternetServiceLink"
		r.References["internet_service_id"] = config.Reference{
			Type: "InternetService",
		}
		r.References["net_id"] = config.Reference{
			Type: "github.com/outscale/upjet-provider-outscale/apis/net/v1alpha1.Net",
		}
	})
}
