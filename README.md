# Kubernetes FakeClient Example

This repository houses an example of how to use Go-Client, and more specifically, how to write unit-tests for functions which use the Kubernetes Go-Client.

This repository is meant to go along with the [Fake it Until You Make it: Unit Tests with Go-Client Fake Client](https://sched.co/MPgE) Presentation for KubeCon Europe 2019.

You can also follow along with the recorded session on [youtube](https://youtu.be/reDCJYbxtRg).

## Pre-Requisites

In order to run this application, you must first install a couple of dependancies.

You can install [dep](https://github.com/golang/dep) and the just run the
following command:

```bash
dep ensure
```

**Note:** These dependancies were generated for Kubernetes version 1.14. The code may
only work with that Kubernetes version.

## Setting Up Minikube and Secrets

You can install [minikube](https://kubernetes.io/docs/tasks/tools/install-minikube) and then run:

```bash
minikube start
```

This will generate a config in `$HOME/.kube/config`.

Now we can generate a few secrets to test it out:

```bash
kubectl create secret generic my-login-information --from-literal=username=us3r1 --from-literal=password=p4ssw0rd2

kubectl create secret generic my-api-key --from-literal=apikey=ABCDEFGHIJKLMNOP
```

## Usage

Once the dependancies have been installed, and minikube has been setup, you can simply run:

```bash
go run internal/main.go
```

If you want to use another Kubernetes Cluster, then you can pass the `-kubeconfig` flag:

```bash
go run internal/main.go -kubeconfig <path-to-kubeconfig>
```

This will output something similar to the following:

```bash
....

my-api-key
{
  "apikey": "ABCDEFGHIJKLMNOP"
}

my-login-information
{
  "password": "p4ssw0rd1",
  "username": "admin"
}
```

**Note:** You may need additional setup depending on your cloud provider. For example
extra [authentication](https://github.com/kubernetes/client-go/tree/master/plugin/pkg/client/auth) to the clusters.

## Unit Tests

You can run the Unit-Tests by running:

```bash
go test ./...
```
