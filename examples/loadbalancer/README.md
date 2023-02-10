This example will create vm with:
* Keypair
* Vm with boot disk
* Security group 
* Security group rules
* InternetService
* Route
* RouteTable
* Net
* Subnet
* User-data (base64 version of vm_startup.sh)
* LoadBalancer

See more details on [userData] (https://docs.outscale.com/en/userguide/Configuring-an-Instance-with-User-Data-and-OUTSCALE-Tags.html)
See more details on [Vm] (https://docs.outscale.com/en/userguide/About-Instances.html)

:warning: **You have to create first net, subnet and securitygroup, then change securityGroupId in sg-rule-private with sgpublic securityGroupId in sg-rule-public-outbound with sgprivate securrtyGroupId, apply securityGroupRules and then other components of the exemple.**

Once your vm is up and running, please connect to http://dns-public-lb:8080 to get hello-world.