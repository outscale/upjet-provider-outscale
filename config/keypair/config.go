package keypair

import "github.com/upbound/upjet/pkg/config"

const keypairShortGroup = "keypair"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("outscale_keypair", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = keypairShortGroup
		r.Kind = "Keypair"
	})
}
