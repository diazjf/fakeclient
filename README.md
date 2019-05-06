# Kubernetes FakeClient Example

This repository houses an example of how to use Go-Client, and more specifically, how to write unit-tests for functions which use the Kubernetes Go-Client.

This repository is meant to go along with the [Fake it Until You Make it: Unit Tests with Go-Client Fake Client](https://sched.co/MPgE) Presentation for KubeCon Europe 2019.

## Pre-Requisites

In order to run this application, you must first install a couple of dependancies.

You can install [dep](https://github.com/golang/dep) and the just run the
following command:

```bash
dep ensure
```

**Note:** These dependancies were generated for Kubernetes version 1.14. You may need different dependancies depending on your Kubernetes version.

## Setting Up Minikube

You can install [minikube](https://kubernetes.io/docs/tasks/tools/install-minikube) and then run:

```bash
minikube start
```

This will generate a config in `$HOME/.kube/config`.

## Usage

Once the dependancies have been installed, simply run:

```bash
go run internal/main.go
```

If you want to use another Kubernetes Cluster, then you can pass the `-kubeconfig` flag:

```bash
go run internal/main.go -kubeconfig <path-to-kubeconfig>
```

**Note:** You may need additional setup depending on your cloud provider. For example
extra [authentication](https://github.com/kubernetes/client-go/tree/master/plugin/pkg/client/auth) to the clusters.