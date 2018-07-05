package v1alpha1

type ClusterGuestConfig struct {
	// DNSZone for guest cluster is supplemented with host prefixes for
	// specific services such as Kubernetes API or Etcd. In general this DNS
	// Zone should start with `k8s` like for example
	// `k8s.cluster.example.com.`.
	DNSZone        string                            `json:"dnsZone" yaml:"dnsZone"`
	ID             string                            `json:"id" yaml:"id"`
	Name           string                            `json:"name,omitempty" yaml:"name,omitempty"`
	Owner          string                            `json:"owner,omitempty" yaml:"owner,omitempty"`
	ReleaseVersion string                            `json:"releaseVersion,omitempty" yaml:"releaseVersion,omitempty"`
	VersionBundles []ClusterGuestConfigVersionBundle `json:"versionBundles,omitempty" yaml:"versionBundles,omitempty"`
}

type ClusterGuestConfigNode struct {
	CPUCores      int     `json:"cpuCores,omitempty" yaml:"cpuCores,omitempty"`
	MemorySizeGB  float64 `json:"memorySizeGB,omitempty" yaml:"memorySizeGB,omitempty"`
	StorageSizeGB float64 `json:"storageSizeGB,omitempty" yaml:"storageSizeGB,omitempty"`
}

type ClusterGuestConfigVersionBundle struct {
	Name    string `json:"name" yaml:"name"`
	Version string `json:"version" yaml:"version"`
}
