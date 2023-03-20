package vm

import "github.com/upbound/upjet/pkg/config"

const vmShortGroup = "vm"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("outscale_vm", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = vmShortGroup
		r.Kind = "Vm"
		r.References["security_group_ids"] = config.Reference {
			Type: "github.com/outscale-vbr/upjet-provider-outscale/apis/securitygroup/v1alpha1.SecurityGroup",
		}
		r.References["subnet_id"] = config.Reference {
			Type: "github.com/outscale-vbr/upjet-provider-outscale/apis/subnet/v1alpha1.Subnet",
		}
		r.References["nics.subnet_id"] = config.Reference {
			Type: "github.com/outscale-vbr/upjet-provider-outscale/apis/subnet/v1alpha1.Subnet",
		}
		r.References["nics.nic_id"] = config.Reference {
			Type: "github.com/outscale-vbr/upjet-provider-outscale/apis/nic/v1alpha1.Nic",
		}
	})
}