# HelmChartRepository CRD

An operator bundle which delivers a `HelmChartRepository` cluster-scoped 

```
apiVersion: helm.openshift.io/v1alpha1
kind: HelmChartRepository
metadata:
  name: example-helmchartrepository
spec:
  # ...
  # ...
```

### Compile

```
go build -o bin/helmchartrepository cmd/manager/main.go
```

### Build Operator Image

```
docker build -t quay.io/redhat-developer/helm-chart-repository-configuration-api:latest .
```

### Build Bundle Image

```
operator-sdk bundle create quay.io/redhat-developer/helm-chart-repository-configuration-api-bundle:lates
```

### Build Index Image

```
opm index add --bundles quay.io/redhat-developer/helm-chart-repository-configuration-api-bundle:latest --tag quay.io/redhat-developer/helm-chart-repository-configuration-api-index:0.0.1 --build-tool=docker
```


### Make it show up on in-cluster OperatorHub !

```

```
