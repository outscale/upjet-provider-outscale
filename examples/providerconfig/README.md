In order to use crossplane please:

Install [crossplane] (https://www.techtarget.com/searchitoperations/tutorial/Step-by-step-guide-to-working-with-Crossplane-and-Kubernetes)

Create your secret based on secret.yaml.tmpl using your AK/SK.
```
kubectl apply -f secret.yaml
```

Deploy provider Config:
```
kubectl apply -f providerconfig.yaml
```

Deploy Outscale provider:

```
kubectl apply -f provider.yaml
```
