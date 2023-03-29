from osc_sdk_python import Gateway as OSCGateway
import os

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
    keyairName = createKeypair(gateway, "upjet")
    getKeypair(gateway, "upjet")
    deleteKeypair(gateway, "upjet")

if __name__ == "__main__":
    main()  