#!/bin/bash
set -euo pipefail
IFS=$'\n\t'

for version in v1 v1beta1; do
  # Using kubebuilder comments, create new CRDs from CR definitions in source files
  echo "Generating all $version CRDs"
  ./tools/bin/controller-gen \
    crd \
    paths=./pkg/apis/... \
    output:dir="../config/crd/$version" \
    crd:crdVersions="$version"
  ./tools/bin/controller-gen \
    crd \
    paths=sigs.k8s.io/cluster-api/api/v1alpha2 \
    output:dir="../config/crd/$version" \
    crd:crdVersions="$version"

  # We only want Cluster and MachineDeployment for now, so delete the other two CAPI CRDs.
  rm ../config/crd/$version/cluster.x-k8s.io_machines.yaml
  rm ../config/crd/$version/cluster.x-k8s.io_machinesets.yaml

  # Add .metadata.name validation to Release CRD using kustomize since
  # kubebuilder comments can't modify metav1.ObjectMeta
  echo "Kustomizing $version CRDs"
  for crd in "../config/crd/patches/$version"/*; do
    ./tools/bin/kustomize --load_restrictor LoadRestrictionsNone build \
      "$crd" \
      -o "../config/crd/$version/$(basename "$crd").yaml"
  done
done
