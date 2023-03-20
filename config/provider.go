/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/upbound/upjet/pkg/config"

	oscnet "github.com/outscale-vbr/upjet-provider-outscale/config/net"
	oscinternetservice "github.com/outscale-vbr/upjet-provider-outscale/config/internetservice"
	oscdhcoption "github.com/outscale-vbr/upjet-provider-outscale/config/dhcpoption"
	oscsubnet "github.com/outscale-vbr/upjet-provider-outscale/config/subnet"
	oscvm "github.com/outscale-vbr/upjet-provider-outscale/config/vm"
	oscroutetable "github.com/outscale-vbr/upjet-provider-outscale/config/routetable"
	oscsecuritygroup "github.com/outscale-vbr/upjet-provider-outscale/config/securitygroup"
	oscroute "github.com/outscale-vbr/upjet-provider-outscale/config/route"
	oscpublicip "github.com/outscale-vbr/upjet-provider-outscale/config/publicip"
	osckeypair"github.com/outscale-vbr/upjet-provider-outscale/config/keypair"
	oscloadbalancer "github.com/outscale-vbr/upjet-provider-outscale/config/loadbalancer"
	oscnat "github.com/outscale-vbr/upjet-provider-outscale/config/nat"
	oscvolume "github.com/outscale-vbr/upjet-provider-outscale/config/volume"
	oscsnapshot "github.com/outscale-vbr/upjet-provider-outscale/config/snapshot"
	oscimage "github.com/outscale-vbr/upjet-provider-outscale/config/image"
	oscnic "github.com/outscale-vbr/upjet-provider-outscale/config/nic"
	oscvpn "github.com/outscale-vbr/upjet-provider-outscale/config/vpn"
	oscclientgateway "github.com/outscale-vbr/upjet-provider-outscale/config/clientgateway"
	oscvirtualgateway "github.com/outscale-vbr/upjet-provider-outscale/config/virtualgateway"

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
