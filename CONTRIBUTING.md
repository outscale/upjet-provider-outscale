# Opening an issue

Feel free to open a [Github issue](https://github.com/outscale-vbr/upjet-outscale/issues) and explain your problem.

Please provide at least those informations:
- upjet version
- how to reproduce the issue
- please store large output (like traces) as as an attached file
- make sure your don't leak any sensible informations (credentials, ...)

# Developing the Provider

## Requirements
You'll first need [Go](http://www.golang.org) installed on your machine (version 1.19+ is *required*)

To compile the provider, run `make generate`. 

In order to test the provider, you can simply run `make run`.


In order to run the full suite of Acceptance tests, run `make testacc`.

*Note:* Acceptance tests create real resources, and often cost money to run.

*Note:* The following environment variables must be set prior to run Acceptance Tests

```sh
$ export OUTSCALE_IMAGEID="ami-xxxxxxxx"    # i.e. "ami-4a7bf2b3"
$ export OUTSCALE_ACCESSKEYID="<ACCESSKEY>" # i.e. "XXXXXXXXXXXXXXXXXXXX"
$ export OUTSCALE_SECRETKEYID="<SECRETKEY>" # i.e. "YYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYY"
$ export OUTSCALE_REGION="<REGION>"         # i.e. "eu-west-2"
$ export OUTSCALE_ACCOUNT="<ACCOUNTPID>"    # i.e. "XXXXXXXXXXXX"
```

```sh
$ make testacc
```