from kubernetes import client, config, utils
from osc_sdk_python import Gateway as OSCGateway
import subprocess
import os
import time

WAIT_FOR_ANY_TIMEOUT=600


def execute_bash_cmd(cmd, full_local_path, my_env):
    result = subprocess.run(cmd, stdout=subprocess.PIPE, stderr=subprocess.PIPE, text=True, cwd=full_local_path, env=my_env )
    print(result.stdout)
    print(result.stderr)

def checkNetState(gateway, tag, netName, my_env):
    nets = gateway.ReadNets(
        Filters={
            "Tags": [tag],
        }
    )["Nets"]
    for net in nets:
        if net['State'] == "available":
            netId = getNetId(netName, my_env).replace("'", "")
            print("Find netId from net object", netId)
            print("Find netId", net['NetId'])
            if net['NetId'] == netId:
                print("Find the same netId", netId)
                return True
            else: 
                print("Find not the same netId", netId)
                return False
        else:
            return False

def checkVmState(gateway, tag, vmName, my_env):
    vms = gateway.ReadVms(
        Filters={
            "Tags": [tag],
        }
    )["Vms"]
    print(vms)
    for vm in vms:
        if vm['State'] == 'running':
            vmId = getVmId(vmName, my_env).replace("'","")
            print("Find vmId from vm object", vmId)
            print("Find vmId", vm['VmId'])
            if vm['VmId'] == vmId:
                print("Find the same vmId", vmId)
                return True
            else:
                print("Find not the same vmId", vmId)
                return False
        else:
            return False




def checkVolumeState(gateway, tag, volumeName, my_env):
    volumes = gateway.ReadVolumes(
        Filters={ 
            "Tags": [tag], 
        },
    )["Volumes"]
    for volume in volumes:

        if volume['State'] == "available":
            volumeId = getVolumeId(volumeName, my_env).replace("'", "")
            print("Find volumeId from volume object", volumeId)
            print("Find volumeId", volume['VolumeId'])
            if volume['VolumeId'] == volumeId:
                print("Find the same volumeId", volumeId)
                return True
            else: 
                print("Find not the same volumeId", volumeId)
                return False
        else:
            return False
        
def checkSnapshotState(gateway, tag, snapshotName, my_env):
    snapshots = gateway.ReadSnapshots(
        Filters={
            "Tags": [tag],
        }
    )["Snapshots"]
    for snapshot in snapshots:
        if snapshot['State'] == "completed":
            snapshotId = getSnapshotId(snapshotName, my_env).replace("'", "")
            print("Find snapshotId from volume object", snapshotId)
            print("Find snapshotId", snapshot['SnapshotId'])
            if snapshot['SnapshotId'] == snapshotId:
                print("Find the same snapshotId", snapshotId)
                return True
            else: 
                print("Find not the same snapshotId", snapshotId)
                return False
        else:
            return False

def checkSubnetState(gateway, tag, subnetName, my_env):
    subnets = gateway.ReadSubnets(
        Filters={
            "Tags": [tag],
        }
    )["Subnets"]
    print(subnets)

    for subnet in subnets:
        if subnet['State'] == 'available':
            subnetId = getSubnetId(subnetName, my_env).replace("'", "")
            print("Find subnetId from subnet object", subnetId)
            print("Find subnetId", subnet['SubnetId'])
            if subnet['SubnetId'] == subnetId:
                print("Find the same subneetId", subnetId)
                return True
            else:
                print("Find not the same subnetId", subnetId)
                return False

def checkInternetServiceState(gateway, tag, internetServiceName, my_env):
    internetServices = gateway.ReadInternetServices(
        Filters={
            "Tags": [tag],
        }
    )["InternetServices"]
    print(internetServices)

    for internetService in internetServices:
        internetServiceId = getInternetServiceId(internetServiceName, my_env).replace("'", "")
        print("Find internetServiceId from internetService object", internetServiceId)
        print("Find internetServiceId", internetService['InternetServiceId'])
        if internetService['InternetServiceId'] == internetServiceId:
            print("Find the same internetServiceId", internetServiceId)
            return True
        else:
            print("Find not the same internetServiceId", internetServiceId)
            return False


def checkRouteTableState(gateway, tag, routeTableName, my_env):
    routeTables = gateway.ReadRouteTables(
        Filters={
            "Tags": [tag],
        }
    )["RouteTables"]
    print("find ok", routeTables)


    for routeTable in routeTables:
        routeTableId = getRouteTableId(routeTableName, my_env).replace("'", "")
        print("Find routeTableId from routeTable object", routeTableId)
        print("Find routeTableId", routeTable['RouteTableId'])
        if routeTable['RouteTableId'] == routeTableId:
            print("Find the same routeTableId", routeTableId)
            return True
        else:
            print("Find not the same routeTableId", routeTableId)
            return False
        
def checkRouteState(gateway, routeTableId, routeName, my_env):
    routeTables = gateway.ReadRouteTables(
        Filters={
            "RouteTableIds": [routeTableId],
        }
    )["RouteTables"]
    for routeTable in routeTables:
        for route in routeTable['Routes']:
            if route['State'] == 'active':
                return True
            else:
                return False

                
def checkSecurityGroupState(gateway, tag, securityGroupName, my_env):
    securityGroups = gateway.ReadSecurityGroups(
        Filters={
            "Tags": [tag],
        }
    )["SecurityGroups"]
    print("find ok", securityGroups)


    for securityGroup in securityGroups:
        securityGroupId = getSecurityGroupId(securityGroupName, my_env).replace("'", "")
        print("Find securityGroupId from securityGroup object", securityGroupId)
        print("Find securityGroupId", securityGroup['SecurityGroupId'])
        if securityGroup['SecurityGroupId'] == securityGroupId:
            print("Find the same securityGroupId", securityGroupId)
            return True
        else:
            print("Find not the same securityGroupId", securityGroupId)
            return False

def checkSecurityGroupRuleState(gateway, flow, fromPortRange, toPortRange, securityGroupId, securityGroupRuleName, my_env):
    securityGroups = gateway.ReadSecurityGroups(
        Filters={
            "SecurityGroupIds": [securityGroupId],
        }
    )["SecurityGroups"]
    for securityGroup in securityGroups:
        securityGroupFromSecurityRuleId = getSecurityGroupFromSecurityGroupId(securityGroupRuleName, my_env).replace("'", "")
        print("Find securityGroupFromSecurityRuleId from securityGroupRule object", securityGroupFromSecurityRuleId)
        print("Find securityGroupId", securityGroup['SecurityGroupId'])
        if flow == "Inbound":
            for inboundRule in securityGroup['InboundRules']:

                if inboundRule['FromPortRange'] == int(fromPortRange) and inboundRule['ToPortRange'] == int(toPortRange) and securityGroupFromSecurityRuleId == securityGroup['SecurityGroupId']: 
                    print(inboundRule['FromPortRange'], inboundRule['ToPortRange'], securityGroupFromSecurityRuleId )
                    return True
                else:
                    return False 

        else:
            for outboundRule in securityGroup['OutboundRules']:               
                if outboundRule['FromPortRange'] == int(fromPortRange) and outboundRule['ToPortRange'] == int(toPortRange)  and securityGroupFromSecurityRuleId == securityGroup['SecurityGroupId']:
                    return True
                else:
                    return False                 

def checkNetAvailable(gateway, tag, time_sleep, netName, my_env):
    timeout = time.time() + WAIT_FOR_ANY_TIMEOUT
    while time.time() < timeout:
        netState = checkNetState(gateway, tag, netName, my_env)
        if netState:
            print("Net is now available")
            break
        else:
            print("Net is not yet available")
        time.sleep(time_sleep)

def checkVmAvailable(gateway, tag, time_sleep, vmName, my_env):
    timeout = time.time() + WAIT_FOR_ANY_TIMEOUT
    while time.time() < timeout:
        vmState = checkVmState(gateway, tag, vmName, my_env)
        if vmState:
            print("Vm is now available")
            break
        else:
            print("Vm is not yet available")
        time.sleep(time_sleep)

def checkSubnetAvailable(gateway, tag, time_sleep, subnetName, my_env):
    timeout = time.time() + WAIT_FOR_ANY_TIMEOUT
    while time.time() < timeout:
        subnetState = checkSubnetState(gateway, tag, subnetName, my_env)
        if subnetState:
            print("Subnet is not available")
            break
        else:
            print("Subnet is not yet available")
        time.sleep(time_sleep)

def checkInternetServiceAvailable(gateway, tag, time_sleep, internetServiceName, my_env):
    timeout = time.time() + WAIT_FOR_ANY_TIMEOUT
    while time.time() < timeout:
        internetServiceState = checkInternetServiceState(gateway, tag, internetServiceName, my_env)
        if internetServiceState:
            print("InternetService is available")
            break
        else:
            print("InternetService is not available")
        time.sleep(time_sleep)

def checkVolumeAvailable(gateway, tag, time_sleep, volumeName, my_env):
    timeout = time.time() + WAIT_FOR_ANY_TIMEOUT
    while time.time() < timeout:
        volumeState = checkVolumeState(gateway, tag, volumeName, my_env)
        if volumeState:
            print("Volume is available")
            break
        else:
            print("Volume is not yet available")
        time.sleep(time_sleep)

def checkRouteTableAvailable(gateway, tag, time_sleep, volumeName, my_env):
    timeout = time.time() + WAIT_FOR_ANY_TIMEOUT
    while time.time() < timeout:
        routeTableState = checkRouteTableState(gateway, tag, volumeName, my_env)
        if routeTableState:
            print("RouteTable is available")
            break
        else:
            print("RouteTable is not yet available")
        time.sleep(time_sleep)

def checkSnapshotAvailable(gateway, tag, time_sleep, snapshotName, my_env):
    timeout = time.time() + WAIT_FOR_ANY_TIMEOUT
    while time.time() < timeout:
        snapshotState = checkSnapshotState(gateway, tag, snapshotName, my_env)
        if snapshotState:
            print("Snapshot is available")
            break
        else:
            print("Snapshot is not yet available")
        time.sleep(time_sleep)

        
def checkKeypair(gateway, keypairName):
    keypairs = gateway.ReadKeypairs(
        Filters={
            "KeypairNames": [keypairName],
        }
    )["Keypairs"]
    if len(keypairs) != 0:
        return True
    else:
        return False


def checkKeypairAvailable(gateway, keypairName, time_sleep):
    timeout = time.time() + WAIT_FOR_ANY_TIMEOUT
    while time.time() < timeout:
        keypairState = checkKeypair(gateway, keypairName)
        if keypairState:
            print("Keypair is available")
            break
        else:
            print("Keypair is not available")
        time.sleep(time_sleep)

def checkSecurityGroupAvailable(gateway, tag, time_sleep, securityGroupName, my_env):
    timeout = time.time() + WAIT_FOR_ANY_TIMEOUT
    while time.time() < timeout:
        securityGroupState = checkSecurityGroupState(gateway, tag, securityGroupName, my_env)
        if securityGroupState:
            print("SecurityGroup is available")
            break
        else:
            print("SecurityGroup is not available")
        time.sleep(time_sleep)

def checkSecurityGroupRuleAvailable(gateway, time_sleep, flow, fromPortRange, toPortRange, securityGroupId, securityGroupRuleName, my_env):
    timeout = time.time() + WAIT_FOR_ANY_TIMEOUT
    while time.time() < timeout:
        securityGroupRuleState = checkSecurityGroupRuleState(gateway, flow, fromPortRange, toPortRange, securityGroupId, securityGroupRuleName, my_env)
        if securityGroupRuleState:
            print("SecurityGroupRule is available")
            break
        else:
            print("SecurityGroupRule is not available")
        time.sleep(time_sleep)

def checkRouteAvailable(gateway, time_sleep, routeTableId, routeName, my_env):
    timeout = time.time() + WAIT_FOR_ANY_TIMEOUT
    while time.time() < timeout:
        routeState = checkRouteState(gateway, routeTableId, routeName, my_env)
        if routeState:
            print("Route is available")
            break
        else:
            print("Route is not available")
        time.sleep(time_sleep)

def getNetId(netName, my_env):
    result = subprocess.run(["kubectl", "get",  "net", netName, "-o=jsonpath='{.status.atProvider.netId}'"], stdout=subprocess.PIPE, stderr=subprocess.PIPE, text=True, env=my_env )
    return result.stdout


def getSubnetId(subnetName, my_env):
    result = subprocess.run(["kubectl", "get",  "subnet", subnetName, "-o=jsonpath='{.status.atProvider.subnetId}'"], stdout=subprocess.PIPE, stderr=subprocess.PIPE, text=True, env=my_env )
    return result.stdout

def getVolumeId(volumeName, my_env):
    result = subprocess.run(["kubectl", "get",  "volume", volumeName, "-o=jsonpath='{.status.atProvider.volumeId}'"], stdout=subprocess.PIPE, stderr=subprocess.PIPE, text=True, env=my_env )
    return result.stdout

def getSnapshotId(snapshotName, my_env):
    result = subprocess.run(["kubectl", "get",  "snapshot", snapshotName, "-o=jsonpath='{.status.atProvider.snapshotId}'"], stdout=subprocess.PIPE, stderr=subprocess.PIPE, text=True, env=my_env )
    return result.stdout

def getInternetServiceId(internetServiceName, my_env):
    result = subprocess.run(["kubectl", "get",  "internetservice", internetServiceName, "-o=jsonpath='{.status.atProvider.internetServiceId}'"], stdout=subprocess.PIPE, stderr=subprocess.PIPE, text=True, env=my_env )
    return result.stdout

def getRouteTableId(routeTableName, my_env):
    result = subprocess.run(["kubectl", "get",  "routeTable", routeTableName, "-o=jsonpath='{.status.atProvider.routeTableId}'"], stdout=subprocess.PIPE, stderr=subprocess.PIPE, text=True, env=my_env )
    return result.stdout

def getSecurityGroupId(securityGroupName, my_env):
    result = subprocess.run(["kubectl", "get",  "securityGroup", securityGroupName, "-o=jsonpath='{.status.atProvider.securityGroupId}'"], stdout=subprocess.PIPE, stderr=subprocess.PIPE, text=True, env=my_env )
    return result.stdout

def getSecurityGroupFromSecurityGroupId(securityGroupRuleName, my_env):
    result = subprocess.run(["kubectl", "get",  "securityGroupRule", securityGroupRuleName, "-o=jsonpath='{.status.atProvider.id}'"], stdout=subprocess.PIPE, stderr=subprocess.PIPE, text=True, env=my_env )
    return result.stdout

def getVmId(vmName, my_env):
    result = subprocess.run(["kubectl", "get",  "vm", vmName, "-o=jsonpath='{.status.atProvider.id}'"], stdout=subprocess.PIPE, stderr=subprocess.PIPE, text=True, env=my_env )
    return result.stdout

def main():
    my_env = os.environ.copy()
    my_env['KUBECONFIG'] = "/root/.kube/config"
    access_key = os.environ['OSC_ACCESS_KEY']
    secret_key=os.environ['OSC_SECRET_KEY']
    region=os.environ['OSC_REGION']
    gateway = OSCGateway(
        **{"access_key": access_key, "secret_key": secret_key, "region": region}
    )

    # Check Net_Vm
    execute_bash_cmd(["kubectl", "apply", "-f", "net_vm/"], "./examples", my_env )
    checkNetAvailable(gateway, "Name=cp-net", 5, "cp-net", my_env)
    checkSubnetAvailable(gateway, "Name=cp-public-subnet", 5, "cp-public-subnet", my_env)
    checkKeypairAvailable(gateway, "cp-keypair", 5)
    checkInternetServiceAvailable(gateway, "Name=cp-igw", 5, "cp-igw", my_env)
    checkRouteTableAvailable(gateway, "Name=cp-public-rtb", 5, "cp-public-rtb", my_env)
    routeTableId = getRouteTableId("cp-public-rtb", my_env).replace("'", "")
    checkRouteAvailable(gateway,  5, routeTableId, "cp-public-route", my_env)
    checkSecurityGroupAvailable(gateway, "Name=cp-public-sg", 5, "cp-public-sg", my_env)
    securityGroupId = getSecurityGroupId("cp-public-sg", my_env).replace("'", "")
    checkSecurityGroupRuleAvailable(gateway, 5, "Inbound", "8080", "8080", securityGroupId, "cp-public-sg-rule", my_env)
    checkVmAvailable(gateway, "Name=cp-vm", 5, "cp-vm", my_env)
    execute_bash_cmd(["kubectl", "delete", "-f", "net_vm/"], "./examples", my_env )

    # Check Snapshot
    execute_bash_cmd(["kubectl", "apply", "-f", "snapshot/"], "./examples", my_env )
    checkVolumeAvailable(gateway, "Name=cp-volume", 5, "volume", my_env)
    checkSnapshotAvailable(gateway, "Name=cp-snap", 5, "snap", my_env)
    execute_bash_cmd(["kubectl", "delete", "-f", "snapshot/"], "./examples", my_env )

if __name__ == "__main__":
    main()