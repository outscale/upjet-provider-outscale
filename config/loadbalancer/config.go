package loadbalancer

import "github.com/upbound/upjet/pkg/config"

const loadBalancerShortGroup = "loadbalancer"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("outscale_load_balancer", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = loadBalancerShortGroup
		r.Kind = "LoadBalancer"
		r.References["security_groups"] = config.Reference {
			Type: "github.com/outscale-vbr/upjet-provider-outscale/apis/securitygroup/v1alpha1.SecurityGroup",
		}
		r.References["subnets"] = config.Reference {
			Type: "github.com/outscale-vbr/upjet-provider-outscale/apis/subnet/v1alpha1.Subnet",
		}
	})
	p.AddResourceConfigurator("outscale_load_balancer_attributes", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = loadBalancerShortGroup
		r.Kind = "LoadBalancerAttributes"
	})
	p.AddResourceConfigurator("outscale_load_balancer_listener_rule", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = loadBalancerShortGroup
		r.Kind = "LoadBalancerListernerRule"
	})
	p.AddResourceConfigurator("outscale_load_balancer_policy", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = loadBalancerShortGroup
		r.Kind = "LoadBalancerPolicy"
	})
	p.AddResourceConfigurator("outscale_load_balancer_vms", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = loadBalancerShortGroup
		r.Kind = "LoadBalancerVms"
		r.References["backend_vm_ids"] = config.Reference{
			Type: "github.com/outscale-vbr/upjet-provider-outscale/apis/vm/v1alpha1.Vm",
		}
	})

}