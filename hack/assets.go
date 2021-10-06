package main

import (
	"github.com/giantswarm/apiextensions/v3/pkg/crd"
)

var upstreamReleaseAssets = []crd.ReleaseAssetFileDefinition{
	// aws
	{
		Owner:    "kubernetes-sigs",
		Repo:     "cluster-api",
		Version:  "v0.3.14",
		Files:    []string{"cluster-api-components.yaml"},
		Provider: "aws",
	},
	{
		Owner:   "kubernetes-sigs",
		Repo:    "cluster-api-provider-aws",
		Version: "v0.6.5",
		Files: []string{
			"eks-bootstrap-components.yaml",
			"eks-controlplane-components.yaml",
			"infrastructure-components.yaml",
		},
		Provider: "aws",
	},
	// azure
	{
		Owner:    "kubernetes-sigs",
		Repo:     "cluster-api",
		Version:  "v0.4.1",
		Files:    []string{"cluster-api-components.yaml"},
		Provider: "azure",
	},
	{
		Owner:    "kubernetes-sigs",
		Repo:     "cluster-api-provider-azure",
		Version:  "v0.4.12",
		Files:    []string{"infrastructure-components.yaml"},
		Provider: "azure",
	},
	{
		Owner:    "Azure",
		Repo:     "aad-pod-identity",
		Version:  "v1.8.0",
		Files:    []string{"deployment.yaml"},
		Provider: "azure",
	},
	// kvm
	{
		Owner:    "kubernetes-sigs",
		Repo:     "cluster-api",
		Version:  "v0.4.1",
		Files:    []string{"cluster-api-components.yaml"},
		Provider: "kvm",
	},
	// vsphere
	{
		Owner:    "kubernetes-sigs",
		Repo:     "cluster-api",
		Version:  "v0.4.1",
		Files:    []string{"cluster-api-components.yaml"},
		Provider: "vsphere",
	},
	{
		Owner:    "kubernetes-sigs",
		Repo:     "cluster-api-provider-vsphere",
		Version:  "v0.8.1",
		Files:    []string{"infrastructure-components.yaml"},
		Provider: "vsphere",
	},
}
