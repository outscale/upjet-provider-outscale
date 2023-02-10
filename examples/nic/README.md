This example will create vm with FNI with:
* Keypair
* Vm with boot disk
* Security group 
* Security group rules
* FNI
* InternetService
* Route
* RouteTable
* User-data (base64 version of vm_startup.sh)
* Net
* Subnet

See more details on [userData] (https://docs.outscale.com/en/userguide/Configuring-an-Instance-with-User-Data-and-OUTSCALE-Tags.html)
See more details on [Vm] (https://docs.outscale.com/en/userguide/About-Instances.html)
See more details on [FNI] (https://docs.outscale.com/fr/userguide/Cr%C3%A9er-une-FNI.html)

Once your vm is up and running, please connect to http://dns-public-lb:8080 to get hello-world.