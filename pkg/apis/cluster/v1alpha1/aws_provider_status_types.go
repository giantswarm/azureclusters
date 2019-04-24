package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AWSProviderStatus is the structure put into the provider status of the
// Cluster API's Cluster type. There it is tracked as serialized raw extension.
//
//     kind: AWSProviderStatus
//     apiVersion: cluster.giantswarm.io/v1alpha1
//     metadata:
//       name: 8y5kc
//     cluster:
//       conditions:
//       - lastTransitionTime: "2019-03-25T17:10:09.333633991Z"
//         type: Created
//       id: 8y5kc
//       versions:
//       - lastTransitionTime: "2019-03-25T17:10:09.995948706Z"
//         version: 4.9.0
//     provider:
//       network:
//         cidr: 10.1.6.0/24
//
type AWSProviderStatus struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Cluster           AWSProviderStatusCluster  `json:"cluster" yaml:"cluster"`
	Provider          AWSProviderStatusProvider `json:"provider" yaml:"provider"`
}

type AWSProviderStatusCluster struct {
	Conditions []AWSProviderStatusClusterCondition `json:"conditions" yaml:"conditions"`
	ID         string                              `json:"id" yaml:"id"`
	Versions   []AWSProviderStatusClusterVersion   `json:"versions" yaml:"versions"`
}

type AWSProviderStatusClusterCondition struct {
	LastTransitionTime DeepCopyTime `json:"lastTransitionTime" yaml:"lastTransitionTime"`
	Type               string       `json:"type" yaml:"type"`
}

type AWSProviderStatusClusterVersion struct {
	LastTransitionTime DeepCopyTime `json:"lastTransitionTime" yaml:"lastTransitionTime"`
	Version            string       `json:"version" yaml:"version"`
}

type AWSProviderStatusProvider struct {
	Network AWSProviderStatusProviderNetwork `json:"network" yaml:"network"`
}

type AWSProviderStatusProviderNetwork struct {
	CIDR string `json:"cidr" yaml:"cidr"`
}
