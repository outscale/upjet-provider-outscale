from osc_sdk_python import Gateway as OSCGateway
import os

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

def main():
    access_key = os.environ['OSC_ACCESS_KEY']
    secret_key = os.environ['OSC_SECRET_KEY']
    region = os.environ['OSC_REGION'] 
    gateway = OSCGateway(
        **{"access_key": access_key, "secret_key": secret_key, "region": region}
    )
    publicIpId = createPublicIp(gateway)
    getPublicIp(gateway, "Name=upjet-publicip")
    deletePublicIp(gateway,publicIpId )

if __name__ == "__main__":
    main()    
