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

def createSecurityGroup(gateway, netId, securityGroupName, description):
    securityGroup = gateway.CreateSecurityGroup(
        Description = description,
        NetId = netId,
        SecurityGroupName = securityGroupName, 
    ).get('SecurityGroup')
    securityGroupId = securityGroup['SecurityGroupId']
    gateway.CreateTags(
        ResourceIds=[securityGroupId],
        Tags=[
            {
                "Key": "Name",
                "Value": "upjet-sg"
            }
        ]
    )
    return securityGroupId

def createSecurityGroupRule(gateway, flow, fromPortRange, toPortRange, ipProtocol, ipRange, securityGroupId):
    securityGroup = gateway.CreateSecurityGroupRule(
    Flow="Inbound",
    Rules=[ 
        { 
            "FromPortRange": int(fromPortRange), 
            "IpProtocol": ipProtocol, 
            "IpRanges": [ipRange], 
            "ToPortRange": int(toPortRange) 
        } 
    ],
    SecurityGroupId=securityGroupId,
).get('SecurityGroup')
    print(securityGroup)


def createInternetService(gateway):
    internetService = gateway.CreateInternetService(

    ).get('InternetService')
    internetServiceId = internetService['InternetServiceId']

    gateway.CreateTags(
        ResourceIds=[internetServiceId],
        Tags=[
            {
                "Key": "Name",
                "Value": "upjet-igw"
            }
        ]
    )
    return internetServiceId 

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

def getInternetServiceState(gateway, tag):
    internetServices = gateway.ReadInternetServices(
        Filters={
            "Tags": [tag],
        }
    )["InternetServices"]
    print(internetServices)
    for internetService in internetServices:
        print(internetService)
        print(internetService['InternetServiceId'])

def getSecurityGroupState(gateway, tag):
    securityGroups = gateway.ReadSecurityGroups(
        Filters={
            "Tags": [tag],
        }
    )["SecurityGroups"]
    print(securityGroups)
    for securityGroup in securityGroups:
        print(securityGroup)
        print(securityGroup['SecurityGroupId'])

def getSecurityGroupRuleState(gateway,  flow, securityGroupId):
    securityGroups = gateway.ReadSecurityGroups(
        Filters={
            "SecurityGroupIds": [securityGroupId],
        }
    )["SecurityGroups"]
    for securityGroup in securityGroups:
        if flow == "Inbound":
            for inboundRule in securityGroup['InboundRules']:
                print(inboundRule)
        else:
            for outboundRule in securityGroup['OutboundRules']:
                print(outboundRule)


def deleteNet(gateway, netId):
    res = gateway.DeleteNet(
        NetId=netId,
    )

def deleteRouteTable(gateway, routeTableId):
    res = gateway.DeleteRouteTable(
        RouteTableId=routeTableId,
    )

def deleteSubnet(gateway, subnetId):
    res = gateway.DeleteSubnet(
        SubnetId=subnetId,
    )

def deleteInternetService(gateway, internetServiceId):
    res = gateway.DeleteInternetService(
        InternetServiceId=internetServiceId
    )

def deleteSecurityGroup(gateway, securityGroupId):
    res = gateway.DeleteSecurityGroup(
        SecurityGroupId=securityGroupId,
    )

def deleteSecurityGroupRule(gateway, flow, fromPortRange, toPortRange, ipProtocol, ipRange, securityGroupId):
    res = gateway.DeleteSecurityGroupRule(
        Flow=flow,
        Rules=[
            {
                "FromPortRange": int(fromPortRange),
                "IpProtocol": ipProtocol,
                "IpRanges": [ipRange],
                "ToPortRange": int(toPortRange),
            }
        ],
        SecurityGroupId=securityGroupId,
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
    securityGroupId = createSecurityGroup(gateway, netId, "public-sg", "my public sg")
    createSecurityGroupRule(gateway, "Inbound", "8080", "8080", "tcp", "0.0.0.0/0", securityGroupId)
    internetServiceId = createInternetService(gateway)
    getNetState(gateway, "Name=upjet-net")
    getInternetServiceState(gateway,"Name=upjet-igw")
    getSubnetState(gateway, "Name=upjet-subnet")
    getSecurityGroupState(gateway, "Name=upjet-sg")
    getSecurityGroupRuleState(gateway, "Inbound", securityGroupId)
    deleteInternetService(gateway, internetServiceId)
    deleteSecurityGroupRule(gateway, "Inbound", "8080", "8080", "tcp", "0.0.0.0/0", securityGroupId)
    deleteSecurityGroup(gateway, securityGroupId)
    deleteSubnet(gateway, subnetId)
    deleteNet(gateway, netId)

if __name__ == "__main__":
    main()    

