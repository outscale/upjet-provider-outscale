package securitygroup

import "github.com/upbound/upjet/pkg/config"

const securityGroupShortGroup = "securitygroup"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("outscale_security_group", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = securityGroupShortGroup
		r.Kind = "SecurityGroup"
		r.References["net_id"] = config.Reference{
			Type: "github.com/outscale-vbr/upjet-provider-outscale/apis/net/v1alpha1.Net",
		}
	})
	p.AddResourceConfigurator("outscale_security_group_rule", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = securityGroupShortGroup
		r.Kind = "SecurityGroupRule"
		r.References["security_group_id"] = config.Reference {
			Type: "SecurityGroup",
		}
		r.References["rules.security_groups_members.security_group_id"] = config.Reference {
			Type: "SecurityGroup",
		}
	})
}
