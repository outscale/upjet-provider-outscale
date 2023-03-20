from osc_sdk_python import Gateway as OSCGateway
import os
import time

def createVolume(gateway, size, subregionName, volumeType):
    volume = gateway.CreateVolume(Size=size, SubregionName=subregionName, VolumeType=volumeType).get('Volume')
    volumeId = volume['VolumeId']
    gateway.CreateTags(
        ResourceIds=[volumeId],
        Tags=[ 
            {
                "Key": "Name",
                "Value": "upjet-volume"
            }
        ],
    )
    return volumeId

def createSnapshot(gateway, volumeId):
    snapshot = gateway.CreateSnapshot(
        VolumeId=volumeId,
    ).get('Snapshot')
    snapshotId = snapshot['SnapshotId']
    gateway.CreateTags(
        ResourceIds=[snapshotId],
        Tags=[ 
            {
                "Key": "Name",
                "Value": "upjet-snapshot"
            }
        ],
    )
    return snapshotId

def getVolumeState(gateway, tag):
    volumes = gateway.ReadVolumes(
        Filters={ 
            "Tags": [tag], 
        },
    )["Volumes"]
    print(volumes)

    for volume in volumes:
        print(volume)
        print(volume['State'])

def getSnapshoState(gateway, tag):
    snapshots = gateway.ReadSnapshots(
        Filters={
            "Tags": [tag],
        }
    )["Snapshots"]
    print(snapshots)

    for snapshot in snapshots:
        print(snapshot)
        print(snapshot['SnapshotId'])

def deleteVolume(gateway, volumeId):
    res = gateway.DeleteVolume(
        VolumeId=volumeId
    )

def deleteSnapshot(gateway, snapshotId):
    res = gateway.DeleteSnapshot(
        SnapshotId=snapshotId
    )

def main():
    access_key = os.environ['OSC_ACCESS_KEY']
    secret_key = os.environ['OSC_SECRET_KEY']
    region = os.environ['OSC_REGION'] 
    gateway = OSCGateway(
        **{"access_key": access_key, "secret_key": secret_key, "region": region}
    )
    volumeId = createVolume(gateway, 10, "eu-west-2a", "gp2")
    time.sleep(5)
    snapshotId = createSnapshot(gateway, volumeId)
    getVolumeState(gateway, "Name=upjet-volume")
    getSnapshoState(gateway, "Name=upjet-snapshot")
    deleteVolume(gateway, volumeId)
    deleteSnapshot(gateway, snapshotId)

if __name__ == "__main__":
    main()


