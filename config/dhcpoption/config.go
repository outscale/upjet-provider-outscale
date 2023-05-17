package dhcpoption

import "github.com/upbound/upjet/pkg/config"

const dhcpOptionShortGroup = "dhcpoption"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("outscale_dhcp_option", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = dhcpOptionShortGroup
		r.Kind = "DhcpOption"
	})
}
