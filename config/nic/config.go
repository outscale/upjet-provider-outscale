package nic

import "github.com/upbound/upjet/pkg/config"

const nicShortGroup = "nic"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("outscale_nic", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = nicShortGroup
		r.Kind = "Nic"
		r.References["security_group_ids"] = config.Reference{
			Type: "github.com/outscale-vbr/upjet-provider-outscale/apis/securitygroup/v1alpha1.SecurityGroup",
		}
		r.References["subnet_id"] = config.Reference{
			Type: "github.com/outscale-vbr/upjet-provider-outscale/apis/subnet/v1alpha1.Subnet",
		}
	})
	p.AddResourceConfigurator("outscale_nic_private_ip", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = nicShortGroup
		r.Kind = "NicPrivateIp"
		r.References["nic_id"] = config.Reference{
			Type: "Nic",
		}
	})
}
