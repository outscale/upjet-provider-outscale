from osc_sdk_python import Gateway as OSCGateway
import os
import time

def createNet(gateway, ipRange):
    net = gateway.CreateNet(
        IpRange = ipRange,
    ).get('Net')
    netId = net['NetId']
    gateway.CreateTags(
        ResourceIds=[netId],
        Tags=[
            {
                "Key": "Name",
                "Value": "upjet-net"
            }
        ]
    )
    return netId

def createSubnet(gateway, ipRange, netId, subregionName):
    subnet = gateway.CreateSubnet(
        IpRange = ipRange,
        NetId = netId,
        SubregionName = subregionName, 
    ).get('Subnet')
    subnetId = subnet['SubnetId']
    gateway.CreateTags(
        ResourceIds=[subnetId],
        Tags=[
            {
                "Key": "Name",
                "Value": "upjet-subnet"
            }
        ]
    )
    return subnetId


def createVm(gateway, keypairName, subnetId, vmType, imageId):
    vms = gateway.CreateVms(
        ImageId=imageId,
        KeypairName=keypairName,
        SubnetId=subnetId,
        VmType=vmType,
    ).get('Vms')
    for vm in vms:
        vmId = vm['VmId']
        gateway.CreateTags(
            ResourceIds=[vmId],
            Tags=[
                {
                    "Key": "Name",
                    "Value": "upjet-vm"                    
                }
            ]
        )




def getNetState(gateway, tag):
    nets = gateway.ReadNets(
        Filters={
            "Tags": [tag],
        }
    )["Nets"]
    print(nets)

    for net in nets:
        print(net)
        print(net['State'])






def getSubnetState(gateway, tag):
    subnets = gateway.ReadSubnets(
        Filters={
            "Tags": [tag],
        }
    )["Subnets"]
    print(subnets)

    for subnet in subnets:
        print(subnet)
        print(subnet['State'])

def getVmState(gateway, tag):
    vms = gateway.ReadVms(
        Filters={
            "Tags": [tag],
        }
    )["Vms"]    
    print(vms)
    for vm in vms:
        print(vm)
        print(vm['VmId'])
        print(vm['State'])


def deleteNet(gateway, netId):
    res = gateway.DeleteNet(
        NetId=netId,
    )

def deleteSubnet(gateway, subnetId):
    res = gateway.DeleteSubnet(
        SubnetId=subnetId,
    )

def deleteVm(gateway, vmId):
    res = gateway.DeleteVms(
        VmIds=[vmId],
    )

def createPublicIp(gateway):
    publicIp = gateway.CreatePublicIp(
    ).get('PublicIp')
    publicIpId = publicIp['PublicIpId']
    gateway.CreateTags(
        ResourceIds=[publicIpId],
        Tags=[
            {
                "Key": "Name",
                "Value": "upjet-publicip"
            }
        ]
    )
    return publicIpId

def getPublicIp(gateway, tag):
    publicIps = gateway.ReadPublicIps(
        Filters={
            "Tags": [tag],
        }
    )["PublicIps"] 
    print(publicIps)

    for publicIp in publicIps:
        print(publicIp["PublicIpId"])

def deletePublicIp(gateway, publicIpId):
        res = gateway.DeletePublicIp(
             PublicIpId=publicIpId,
        )

def createKeypair(gateway, keypairName):
    keypair = gateway.CreateKeypair(
        KeypairName = keypairName,
    ).get('Keypair')


def getKeypair(gateway, keypairName):
    keypairs = gateway.ReadKeypairs(
        Filters={
            "KeypairNames": [keypairName],
        }
    )["Keypairs"]
    for keypair in keypairs:
        print(keypair)

def deleteKeypair(gateway, keypairName):
    res = gateway.DeleteKeypair(
        KeypairName=keypairName,
    ) 


def main():
    access_key = os.environ['OSC_ACCESS_KEY']
    secret_key = os.environ['OSC_SECRET_KEY']
    region = os.environ['OSC_REGION'] 
    gateway = OSCGateway(
        **{"access_key": access_key, "secret_key": secret_key, "region": region}
    )
    netId = createNet(gateway, "10.0.0.0/16")
    subnetId = createSubnet(gateway, "10.0.0.0/24", netId, "eu-west-2a")
    publicIpId = createPublicIp(gateway)
    keyairName = createKeypair(gateway, "upjet-key")
    vmId = createVm(gateway, keyairName, subnetId, "tinav4.c1r1p2", "ami-ec7ef91f")
    getNetState(gateway, "Name=upjet-net")
    getSubnetState(gateway, "Name=upjet-subnet")
    getPublicIp(gateway, "Name=upjet-publicip")
    getKeypair(gateway, "upjet")
    getVmState(gateway,"Name=upjet-vm" )
    time.sleep(360)
    deleteVm(gateway, vmId)
    deleteKeypair(gateway, "upjet")
    deletePublicIp(gateway,publicIpId )
    deleteSubnet(gateway, subnetId)
    deleteNet(gateway, netId)

if __name__ == "__main__":
    main()    