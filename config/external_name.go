/*
Copyright 2022 Upbound Inc.
*/

package config

import "github.com/upbound/upjet/pkg/config"

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{
	// Import requires using a randomly generated ID from provider: nl-2e21sda
	"outscale_net": config.IdentifierFromProvider,
	"outscale_net_attributes": config.IdentifierFromProvider,
	"outscale_net_peering": config.IdentifierFromProvider,
	"outscale_net_peering_acceptation": config.IdentifierFromProvider,
	"outscale_internet_service": config.IdentifierFromProvider,
	"outscale_internet_service_link": config.IdentifierFromProvider,
	"outscale_dhcp_option": config.IdentifierFromProvider,
	"outscale_keypair": config.IdentifierFromProvider,
	"outscale_route": config.IdentifierFromProvider, 
	"outscale_route_table": config.IdentifierFromProvider, 
	"outscale_route_table_link": config.IdentifierFromProvider, 
	"outscale_subnet": config.IdentifierFromProvider, 
	"outscale_security_group_rule": config.IdentifierFromProvider, 
	"outscale_security_group": config.IdentifierFromProvider, 
	"outscale_vm": config.IdentifierFromProvider, 
	"outscale_public_ip": config.IdentifierFromProvider, 
	"outscale_public_ip_link": config.IdentifierFromProvider, 
	"outscale_load_balancer": config.IdentifierFromProvider, 
	"outscale_load_balancer_attributes": config.IdentifierFromProvider, 
	"outscale_load_balancer_listener_rule": config.IdentifierFromProvider, 
	"outscale_load_balancer_policy": config.IdentifierFromProvider, 
	"outscale_load_balancer_vms": config.IdentifierFromProvider, 
	"outscale_nat_service": config.NameAsIdentifier,
	"outscale_volume": config.IdentifierFromProvider,
	"outscale_volumes_link": config.NameAsIdentifier,
	"outscale_snapshot": config.IdentifierFromProvider,
	"outscale_snapshot_export_task": config.IdentifierFromProvider,
	"outscale_snapshot_attributes": config.IdentifierFromProvider,
	"outscale_image": config.IdentifierFromProvider,
	"outscale_image_export_task": config.IdentifierFromProvider,
	"outscale_image_launch_permission": config.IdentifierFromProvider,
	"outscale_nic": config.IdentifierFromProvider,
	"outscale_nic_link": config.IdentifierFromProvider,
	"outscale_nic_private_ip": config.IdentifierFromProvider,
	"outscale_vpn_connection": config.IdentifierFromProvider,
	"outscale_vpn_connection_route": config.IdentifierFromProvider,
	"outscale_client_gateway": config.IdentifierFromProvider,
	"outscale_virtual_gateway": config.IdentifierFromProvider,

}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		}
	}
}

// ExternalNameConfigured returns the list of all resources whose external name
// is configured manually.
func ExternalNameConfigured() []string {
	l := make([]string, len(ExternalNameConfigs))
	i := 0
	for name := range ExternalNameConfigs {
		// $ is added to match the exact string since the format is regex.
		l[i] = name + "$"
		i++
	}
	return l
}
