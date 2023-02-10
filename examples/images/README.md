This example will create image (OMI) in public cloud with:
* Keypair
* Image
* Vm
* User-data (base64 version of vm_startup.sh)

See more details on [userData] (https://docs.outscale.com/en/userguide/Configuring-an-Instance-with-User-Data-and-OUTSCALE-Tags.html)
See more details on [Vm] (https://docs.outscale.com/en/userguide/About-Instances.html)
See more details on [Omi] (https://docs.outscale.com/en/userguide/OUTSCALE-Machine-Images-(OMIs).html)

Once your vm is up and running, please connect to http://dns-public-lb:8080 to get hello-world.