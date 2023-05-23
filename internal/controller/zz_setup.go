/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	clientgateway "github.com/outscale/upjet-provider-outscale/internal/controller/clientgateway/clientgateway"
	dhcpoption "github.com/outscale/upjet-provider-outscale/internal/controller/dhcpoption/dhcpoption"
	exporttask "github.com/outscale/upjet-provider-outscale/internal/controller/image/exporttask"
	image "github.com/outscale/upjet-provider-outscale/internal/controller/image/image"
	launchpermission "github.com/outscale/upjet-provider-outscale/internal/controller/image/launchpermission"
	internetservice "github.com/outscale/upjet-provider-outscale/internal/controller/internetservice/internetservice"
	internetservicelink "github.com/outscale/upjet-provider-outscale/internal/controller/internetservice/internetservicelink"
	keypair "github.com/outscale/upjet-provider-outscale/internal/controller/keypair/keypair"
	loadbalancer "github.com/outscale/upjet-provider-outscale/internal/controller/loadbalancer/loadbalancer"
	loadbalancerattributes "github.com/outscale/upjet-provider-outscale/internal/controller/loadbalancer/loadbalancerattributes"
	loadbalancerlisternerrule "github.com/outscale/upjet-provider-outscale/internal/controller/loadbalancer/loadbalancerlisternerrule"
	loadbalancerpolicy "github.com/outscale/upjet-provider-outscale/internal/controller/loadbalancer/loadbalancerpolicy"
	loadbalancervms "github.com/outscale/upjet-provider-outscale/internal/controller/loadbalancer/loadbalancervms"
	nat "github.com/outscale/upjet-provider-outscale/internal/controller/nat/nat"
	net "github.com/outscale/upjet-provider-outscale/internal/controller/net/net"
	netattributes "github.com/outscale/upjet-provider-outscale/internal/controller/net/netattributes"
	netpeering "github.com/outscale/upjet-provider-outscale/internal/controller/net/netpeering"
	netpeeringacceptation "github.com/outscale/upjet-provider-outscale/internal/controller/net/netpeeringacceptation"
	link "github.com/outscale/upjet-provider-outscale/internal/controller/nic/link"
	nic "github.com/outscale/upjet-provider-outscale/internal/controller/nic/nic"
	nicprivateip "github.com/outscale/upjet-provider-outscale/internal/controller/nic/nicprivateip"
	providerconfig "github.com/outscale/upjet-provider-outscale/internal/controller/providerconfig"
	publicip "github.com/outscale/upjet-provider-outscale/internal/controller/publicip/publicip"
	publiciplink "github.com/outscale/upjet-provider-outscale/internal/controller/publicip/publiciplink"
	route "github.com/outscale/upjet-provider-outscale/internal/controller/route/route"
	routetable "github.com/outscale/upjet-provider-outscale/internal/controller/routetable/routetable"
	routetablelink "github.com/outscale/upjet-provider-outscale/internal/controller/routetable/routetablelink"
	securitygroup "github.com/outscale/upjet-provider-outscale/internal/controller/securitygroup/securitygroup"
	securitygrouprule "github.com/outscale/upjet-provider-outscale/internal/controller/securitygroup/securitygrouprule"
	snapshot "github.com/outscale/upjet-provider-outscale/internal/controller/snapshot/snapshot"
	snapshotattributes "github.com/outscale/upjet-provider-outscale/internal/controller/snapshot/snapshotattributes"
	snapshotexporttask "github.com/outscale/upjet-provider-outscale/internal/controller/snapshot/snapshotexporttask"
	subnet "github.com/outscale/upjet-provider-outscale/internal/controller/subnet/subnet"
	virtualgateway "github.com/outscale/upjet-provider-outscale/internal/controller/virtualgateway/virtualgateway"
	vm "github.com/outscale/upjet-provider-outscale/internal/controller/vm/vm"
	volume "github.com/outscale/upjet-provider-outscale/internal/controller/volume/volume"
	volumeslink "github.com/outscale/upjet-provider-outscale/internal/controller/volume/volumeslink"
	vpnconnection "github.com/outscale/upjet-provider-outscale/internal/controller/vpn/vpnconnection"
	vpnconnectionroute "github.com/outscale/upjet-provider-outscale/internal/controller/vpn/vpnconnectionroute"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		clientgateway.Setup,
		dhcpoption.Setup,
		exporttask.Setup,
		image.Setup,
		launchpermission.Setup,
		internetservice.Setup,
		internetservicelink.Setup,
		keypair.Setup,
		loadbalancer.Setup,
		loadbalancerattributes.Setup,
		loadbalancerlisternerrule.Setup,
		loadbalancerpolicy.Setup,
		loadbalancervms.Setup,
		nat.Setup,
		net.Setup,
		netattributes.Setup,
		netpeering.Setup,
		netpeeringacceptation.Setup,
		link.Setup,
		nic.Setup,
		nicprivateip.Setup,
		providerconfig.Setup,
		publicip.Setup,
		publiciplink.Setup,
		route.Setup,
		routetable.Setup,
		routetablelink.Setup,
		securitygroup.Setup,
		securitygrouprule.Setup,
		snapshot.Setup,
		snapshotattributes.Setup,
		snapshotexporttask.Setup,
		subnet.Setup,
		virtualgateway.Setup,
		vm.Setup,
		volume.Setup,
		volumeslink.Setup,
		vpnconnection.Setup,
		vpnconnectionroute.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
