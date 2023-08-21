# Kube Controller

A backend service that communicates with the Kubernetes API server to manage NetworkPolicies.

It can run inside or outside of a cluster.

It exposes a API to get cluster's NetworPolicies, Namespaces and Pods, and also allows to create and edit NetworkPolicies.

## Development

### Requirements

- [Go](https://golang.org/doc/install)
- [Docker](https://docs.docker.com/get-docker/)
- [Kubernetes](https://kubernetes.io/docs/tasks/tools/install-kubectl/)
- [air (for development)](https://github.com/cosmtrek/air)

### Running Locally

Install dependencies:

```bash
# Only needed once
make install
```

If running outside of a cluster, the program uses your kubeconfig file to connect to the cluster. If you don't have a kubeconfig file, you can create one by running `kubectl config view --raw > ~/.kube/config`.

```bash
# Run the program (optionally set PORT={:port_number})
make run
```

If you want hot reloading, you can use air:

```bash
# Run the program with air (optionally set PORT={:port_number})
make air
```

### Running in a Cluster

When running inside a cluster, the program uses the Service Account token mounted inside the Pod at the `/var/run/secrets/kubernetes.io/serviceaccount`

If you have RBAC enabled on your cluster, use the following snippet to create role binding which will grant the default service account view permissions.

```bash
kubectl create clusterrolebinding default-view --clusterrole=view --serviceaccount=default:default
```

To deploy the program to the cluster, run the following commands:


```bash
# install dependencies
make install

# compile the program (optionally set TARGET_OS={linux|darwin|windows})
make compile

# build the docker image (optionally set DOCKER_TAG={tag})
make build

# push the docker image (optionally set DOCKER_TAG={tag}, DOCKER_REGISTRY={registry})
make push

# deploy the program to the cluster (optionally set DEPLOYMENT_FILE={file})
make apply

# restart the deployment if needed (optionally set DOCKER_TAG={tag})
make restart
```

### Testing

TBD

## API endpoints

See [api.md](api.md)