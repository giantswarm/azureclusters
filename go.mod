module github.com/giantswarm/apiextensions

go 1.13

require (
	github.com/go-openapi/errors v0.19.4
	github.com/google/go-cmp v0.4.0
	k8s.io/api v0.17.0
	k8s.io/apiextensions-apiserver v0.17.0
	k8s.io/apimachinery v0.17.0
	k8s.io/client-go v0.16.6
	k8s.io/code-generator v0.16.6
	sigs.k8s.io/cluster-api v0.0.0
	sigs.k8s.io/controller-tools v0.0.0
	sigs.k8s.io/yaml v1.2.0
)

replace (
	golang.org/x/tools => golang.org/x/tools v0.0.0-20190821162956-65e3620a7ae7
	k8s.io/api => k8s.io/api v0.16.6
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.16.6
	k8s.io/apimachinery => k8s.io/apimachinery v0.16.6
	k8s.io/client-go => k8s.io/client-go v0.16.6
	k8s.io/code-generator => k8s.io/code-generator v0.16.6
	sigs.k8s.io/cluster-api => sigs.k8s.io/cluster-api v0.2.10
	sigs.k8s.io/controller-tools => sigs.k8s.io/controller-tools v0.2.4
)
