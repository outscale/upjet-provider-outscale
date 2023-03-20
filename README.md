# Provider Upjet-Provider-Outscale

`upjet-provider-outscale` is a [Crossplane](https://crossplane.io/) provider that
is built using [Upjet](https://github.com/upbound/upjet) code
generation tools and exposes XRM-conformant managed resources for the
Upjet-Provider-Outscale API.

(Rajouter une image, regader sur https://github.com/outscale/terraform-provider-outscale/blob/master/README.md?plain=1)

# Requirements
To run upjet, please have:
* [Kubernetes](https://github.com/upbound/upjet)
* [Crossplane](https://github.com/upbound/upjet)


## Getting Started
### Install Outscale Provider
You can use declarative installation:
```
cat <<EOF | kubectl apply -f -
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: upjet-provider-outscale
spec:
  package: vbaer/upjet-provider-outscale:v0.1.0
EOF
```
### Using the provider
You must have the following environment variable:
```
export OSC_ACCESS_KEY=<your-access-key>
export OSC_SECRET_KEY=<your-secret-access-key>
export OSC_REGION=<your-region>
make providerconfig
```
// (make credentials with providerConfig)

## Issues and contributions
Check [CONTRIBUTING.md](./CONTRIBUTING.md) for more details.
## Developing

Run code-generation pipeline:
```console
go run cmd/generator/main.go "$PWD"
```

Run against a Kubernetes cluster:

```console
make run
```

Build, push, and install:

```console
make all
```

Build binary:

```console
make build
```

## Report a Bug

For filing bugs, suggesting improvements, or requesting new features, please
open an [issue](https://github.com/outscale-vbr/upjet-provider-outscale/issues).
