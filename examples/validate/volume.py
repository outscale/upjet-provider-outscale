from osc_sdk_python import Gateway as OSCGateway
import os

def createVolume(gateway, size, subregionName, volumeType):
    volume = gateway.CreateVolume(Size=size, SubregionName=subregionName, VolumeType=volumeType).get('Volume')
    gateway.CreateTags(
        ResourceIds=[volumeId],
        Tags=[ 
            {
                "Key": "Name",
                "Value": "upjet-volume"
            }
        ],
    )
    volumeId = volume['VolumeId']
    return volumeId

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

def deleteVolume(gateway, volumeId):
    res = gateway.DeleteVolume(
        VolumeId=volumeId
    )

def main():
    access_key = os.environ['OSC_ACCESS_KEY']
    secret_key = os.environ['OSC_SECRET_KEY']
    region = os.environ['OSC_REGION'] 
    gateway = OSCGateway(
        **{"access_key": access_key, "secret_key": secret_key, "region": region}
    )
    volumeId = createVolume(gateway, 10, "eu-west-2a", "gp2")
    getVolumeState(gateway, "Name=upjet-volume")
    deleteVolume(gateway, volumeId)

if __name__ == "__main__":
    main()