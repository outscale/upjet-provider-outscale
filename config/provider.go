/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/upbound/upjet/pkg/config"

	oscclientgateway "github.com/outscale-vbr/upjet-provider-outscale/config/clientgateway"
	oscdhcoption "github.com/outscale-vbr/upjet-provider-outscale/config/dhcpoption"
	oscimage "github.com/outscale-vbr/upjet-provider-outscale/config/image"
	oscinternetservice "github.com/outscale-vbr/upjet-provider-outscale/config/internetservice"
	osckeypair "github.com/outscale-vbr/upjet-provider-outscale/config/keypair"
	oscloadbalancer "github.com/outscale-vbr/upjet-provider-outscale/config/loadbalancer"
	oscnat "github.com/outscale-vbr/upjet-provider-outscale/config/nat"
	oscnet "github.com/outscale-vbr/upjet-provider-outscale/config/net"
	oscnic "github.com/outscale-vbr/upjet-provider-outscale/config/nic"
	oscpublicip "github.com/outscale-vbr/upjet-provider-outscale/config/publicip"
	oscroute "github.com/outscale-vbr/upjet-provider-outscale/config/route"
	oscroutetable "github.com/outscale-vbr/upjet-provider-outscale/config/routetable"
	oscsecuritygroup "github.com/outscale-vbr/upjet-provider-outscale/config/securitygroup"
	oscsnapshot "github.com/outscale-vbr/upjet-provider-outscale/config/snapshot"
	oscsubnet "github.com/outscale-vbr/upjet-provider-outscale/config/subnet"
	oscvirtualgateway "github.com/outscale-vbr/upjet-provider-outscale/config/virtualgateway"
	oscvm "github.com/outscale-vbr/upjet-provider-outscale/config/vm"
	oscvolume "github.com/outscale-vbr/upjet-provider-outscale/config/volume"
	oscvpn "github.com/outscale-vbr/upjet-provider-outscale/config/vpn"
)

const (
	resourcePrefix = "upjet-provider-outscale"
	modulePath     = "github.com/outscale-vbr/upjet-provider-outscale"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))
	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		oscnet.Configure,
		oscinternetservice.Configure,
		oscdhcoption.Configure,
		oscsubnet.Configure,
		oscvm.Configure,
		oscroutetable.Configure,
		oscsecuritygroup.Configure,
		oscroute.Configure,
		oscpublicip.Configure,
		osckeypair.Configure,
		oscloadbalancer.Configure,
		oscnat.Configure,
		oscvolume.Configure,
		oscsnapshot.Configure,
		oscimage.Configure,
		oscnic.Configure,
		oscvpn.Configure,
		oscclientgateway.Configure,
		oscvirtualgateway.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
