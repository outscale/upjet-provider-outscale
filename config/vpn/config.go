package vpn

import "github.com/upbound/upjet/pkg/config"

const vpnShortGroup = "vpn"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("outscale_vpn_connection", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = vpnShortGroup
		r.Kind = "VpnConnection"
		r.References["client_gateway_id"] = config.Reference{
			Type: "github.com/outscale-vbr/upjet-provider-outscale/apis/clientgateway/v1alpha1.ClientGateway",
		}
		r.References["virtual_gateway_id"] = config.Reference{
			Type: "github.com/outscale-vbr/upjet-provider-outscale/apis/virtualgateway/v1alpha1.VirtualGateway"}
	})
	p.AddResourceConfigurator("outscale_vpn_connection_route", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = vpnShortGroup
		r.Kind = "VpnConnectionRoute"
		r.References["vpn_connection_id"] = config.Reference{
			Type: "VpnConnection",
		}
	})
}
