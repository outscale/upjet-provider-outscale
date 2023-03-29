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

def createRouteTable(gateway, netId):
    routeTable = gateway.CreateRouteTable(
        NetId = netId,
    ).get('RouteTable')
    routeTableId = routeTable['RouteTableId']
    gateway.CreateTags(
        ResourceIds=[routeTableId],
        Tags=[
            {
                "Key": "Name",
                "Value": "upjet-routetable"
            }
        ]
    )
    return routeTableId    

def createRoute(gateway, gatewayId, destinationIpRange, routeTableId):
    routeTable = gateway.CreateRoute(
        GatewayId = gatewayId,
        DestinationIpRange = destinationIpRange,
        RouteTableId = routeTableId
    ).get('RouteTable')
    routeTableId = routeTable['RouteTableId']
    return routeTableId



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

def linkInternetService(gateway, internetServiceId, netId):
    res = gateway.LinkInternetService(
        InternetServiceId=internetServiceId,
        NetId=netId,
    )

def unlinkInternetService(gateway, internetServiceId, netId):
    res = gateway.UnlinkInternetService(
        InternetServiceId=internetServiceId,
        NetId=netId,
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

def getRouteTableState(gateway, tag):
    routeTables = gateway.ReadRouteTables(
        Filters={
            "Tags": [tag],
        }
    )["RouteTables"]
    print(routeTables)

    for routeTable in routeTables:
        print(routeTable)
        

def getRouteState(gateway, routeTableId):
    routeTables = gateway.ReadRouteTables(
        Filters={
            "RouteTableIds": [routeTableId],
        }
    )["RouteTables"]
    for routeTable in routeTables:
        for route in routeTable['Routes']:
            print(route)
            print(route['State'])


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

def deleteRoute(gateway, routeTableId, destinationIpRange):
    res = gateway.DeleteRoute(
        RouteTableId=routeTableId,
        DestinationIpRange=destinationIpRange,
        
    )

def deleteInternetService(gateway, internetServiceId):
    res = gateway.DeleteInternetService(
        InternetServiceId=internetServiceId
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
    routeTableId = createRouteTable(gateway, netId)
    internetServiceId = createInternetService(gateway)
    linkInternetService(gateway, internetServiceId, netId)
    routeTableIdFromRoute = createRoute(gateway, internetServiceId, "0.0.0.0/0", routeTableId)
    getNetState(gateway, "Name=upjet-net")
    getInternetServiceState(gateway,"Name=upjet-igw")
    getRouteTableState(gateway, "Name=upjet-routetable")
    getSubnetState(gateway, "Name=upjet-subnet")
    getRouteState(gateway, routeTableIdFromRoute)
    unlinkInternetService(gateway, internetServiceId, netId)
    deleteInternetService(gateway, internetServiceId)
    deleteRouteTable(gateway, routeTableId)
    deleteSubnet(gateway, subnetId)
    deleteNet(gateway, netId)

if __name__ == "__main__":
    main()    