from osc_sdk_python import Gateway as OSCGateway
import os

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



def deleteNet(gateway, netId):
    res = gateway.DeleteNet(
        NetId=netId,
    )

def deleteSubnet(gateway, subnetId):
    res = gateway.DeleteSubnet(
        SubnetId=subnetId,
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
    getNetState(gateway, "Name=upjet-net")
    getSubnetState(gateway, "Name=upjet-subnet")
    deleteSubnet(gateway, subnetId)
    deleteNet(gateway, netId)

if __name__ == "__main__":
    main()    