package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/giantswarm/apiextensions/pkg/apis/core"
	"github.com/giantswarm/apiextensions/pkg/key"
	"github.com/giantswarm/apiextensions/pkg/serialization"
)

// NewKVMClusterConfigCR returns a KVMClusterConfig custom resource.
func NewKVMClusterConfigCR(name string) *KVMClusterConfig {
	cr := KVMClusterConfig{}
	cr.TypeMeta, cr.ObjectMeta = key.NewMeta(SchemeGroupVersion, core.KindKVMClusterConfig, name, "")
	return &cr
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=giantswarm;kvm
// +kubebuilder:storageversion

type KVMClusterConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              KVMClusterConfigSpec `json:"spec"`
}

type KVMClusterConfigSpec struct {
	Guest         KVMClusterConfigSpecGuest         `json:"guest"`
	VersionBundle KVMClusterConfigSpecVersionBundle `json:"versionBundle"`
}

type KVMClusterConfigSpecGuest struct {
	ClusterGuestConfig `json:",inline"`
	Masters            []KVMClusterConfigSpecGuestMaster `json:"masters,omitempty"`
	Workers            []KVMClusterConfigSpecGuestWorker `json:"workers,omitempty"`
}

type KVMClusterConfigSpecGuestMaster struct {
	KVMClusterConfigSpecGuestNode `json:",inline"`
}

type KVMClusterConfigSpecGuestWorker struct {
	KVMClusterConfigSpecGuestNode `json:",inline"`
	Labels                        map[string]string `json:"labels"`
}

// TODO: change MemorySizeGB and StorageSizeGB to resource.Quantity
type KVMClusterConfigSpecGuestNode struct {
	ID            string              `json:"id"`
	CPUCores      int                 `json:"cpuCores,omitempty"`
	MemorySizeGB  serialization.Float `json:"memorySizeGB,omitempty"`
	StorageSizeGB serialization.Float `json:"storageSizeGB,omitempty"`
}

type KVMClusterConfigSpecVersionBundle struct {
	Version string `json:"version"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type KVMClusterConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []KVMClusterConfig `json:"items"`
}
