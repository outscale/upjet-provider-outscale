# Q/A Upjet

## Crossplane

Crossplane is a projet to create infrastructure using k8s. ([Crossplane][Crossplane]) 
Upjet is a project to generate crossplane provider.([Upjet][Upjet])


## Provider

There is a explanation about how to generate a provider. ([Generating a provider][Generating a provider])

We cover almost every call available at the beginning of 2023.

Each time you do a PR, everything is regenerated.

There is a github action which can be integrated in terraform to generated new version of upjet when there are a new version of terraform. ([Connect ci terraform with upjet][Connect ci terraform with upjet])

You can used xrd which is the main purpose of the projet. ([Example of xrd][Example of xrd]) 

What is a xrd ? ([xrd][xrd]).

## Test

In  order to create e2etest, we use uptest. ([uptest][uptest])

## Acronym

|      |       Acronym                 |
|------|-------------------------------|
|  XRD | Composite Resource Definition |   
|  XR  | Composite Resource            |
|  XRC | Claims                        |  
|  CRD | Custom Resource Definition    |


## Concept

|                  | Concept                                                                                                                                                                         |
|------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
|   Pod            | A Kubernetes pod is a collection of one  or more  Linux containers, and  a Kubernetes  application                                                                              |
|   XRD            |  A custom API specification.                                                                                                                                                    |
|   XR             | Created by using the custom API defined in a Composite Resource Definition. XRs use the Composition template to create new managed resources.                                   |
|   XRC            | Like a Composite Resource, but with namespace scoping.                                                                                                                          |
|  Crossplane Pods | The base Crossplane installation consists of two pods, the crossplane pod and the crossplane-rbac-manager pod. Both pods install in the crossplane-system namespace by default. |       



## Environment

A K8s cluster as you wants.
You can use kind but for me i create a cluster with one worker and one master using this project ([osc-k8s-rke-cluster][osc-k8s-rke-cluster])

A account on outscale on eu-west-2/cloud-gouv or us-east-2. If you want on another region, please create omi on those region.


<!-- References -->
[Upjet]: https://github.com/crossplane/upjet
[Crossplane]: https://www.crossplane.io/
[Generating a provider]: https://github.com/crossplane/upjet/blob/main/docs/generating-a-provider.md
[Connect ci terraform with upjet]: https://github.com/outscale-vbr/test/blob/main/.github/workflows/launch.yaml
[Example of xrd]: https://github.com/outscale/upjet-provider-outscale/tree/main/examples-xrd/netvm
[xrd]: https://docs.crossplane.io/latest/concepts/composite-resource-definitions/
[uptest]: https://github.com/upbound/uptest
[osc-k8s-rke-cluster]: https://github.com/outscale/osc-k8s-rke-cluster