package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	ClusterVersionLimit = 5
)

const (
	StatusClusterStatusFalse = "False"
	StatusClusterStatusTrue  = "True"
)

const (
	StatusClusterTypeCreated  = "Created"
	StatusClusterTypeCreating = "Creating"
)

const (
	StatusClusterTypeDeleted  = "Deleted"
	StatusClusterTypeDeleting = "Deleting"
)

const (
	StatusClusterTypeUpdated  = "Updated"
	StatusClusterTypeUpdating = "Updating"
)

type StatusCluster struct {
	// +kubebuilder:validation:Optional
	// +nullable
	// Conditions is a list of status information expressing the current
	// conditional state of a guest cluster. This may reflect the status of the
	// guest cluster being updating or being up to date.
	Conditions []StatusClusterCondition `json:"conditions,omitempty" yaml:"conditions,omitempty"`
	// +kubebuilder:validation:Optional
	Network StatusClusterNetwork `json:"network" yaml:"network"`
	// +kubebuilder:validation:Optional
	// +nullable
	// Nodes is a list of guest cluster node information reflecting the current
	// state of the guest cluster nodes.
	Nodes []StatusClusterNode `json:"nodes,omitempty" yaml:"nodes,omitempty"`
	// +kubebuilder:validation:Optional
	// Resources is a list of arbitrary conditions of operatorkit resource
	// implementations.
	Resources []StatusClusterResource `json:"resources,omitempty" yaml:"resources,omitempty"`
	// +kubebuilder:validation:Optional
	Scaling StatusClusterScaling `json:"scaling" yaml:"scaling"`
	// +kubebuilder:validation:Optional
	// +nullable
	// Versions is a list that acts like a historical track record of versions a
	// guest cluster went through. A version is only added to the list as soon as
	// the guest cluster successfully migrated to the version added here.
	Versions []StatusClusterVersion `json:"versions,omitempty" yaml:"versions,omitempty"`
}

// StatusClusterCondition expresses the conditions in which a guest cluster may
// is.
type StatusClusterCondition struct {
	// +kubebuilder:validation:Optional
	// +nullable
	// LastTransitionTime is the last time the condition transitioned from one
	// status to another.
	LastTransitionTime metav1.Time `json:"lastTransitionTime,omitempty" yaml:"lastTransitionTime,omitempty"`
	// +kubebuilder:validation:Optional
	// Status may be True, False or Unknown.
	Status string `json:"status" yaml:"status"`
	// +kubebuilder:validation:Optional
	// Type may be Creating, Created, Scaling, Scaled, Draining, Drained,
	// Updating, Updated, Deleting, Deleted.
	Type string `json:"type" yaml:"type"`
}

// StatusClusterNetwork expresses the network segment that is allocated for a
// guest cluster.
type StatusClusterNetwork struct {
	// +kubebuilder:validation:Optional
	CIDR string `json:"cidr" yaml:"cidr"`
}

// StatusClusterNode holds information about a guest cluster node.
type StatusClusterNode struct {
	// +kubebuilder:validation:Optional
	// +nullable
	// Labels contains the kubernetes labels for corresponding node.
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
	// +kubebuilder:validation:Optional
	// +nullable
	// LastTransitionTime is the last time the condition transitioned from one
	// status to another.
	LastTransitionTime metav1.Time `json:"lastTransitionTime,omitempty" yaml:"lastTransitionTime,omitempty"`
	// +kubebuilder:validation:Optional
	// Name referrs to a tenant cluster node name.
	Name string `json:"name" yaml:"name"`
	// +kubebuilder:validation:Optional
	// Version referrs to the version used by the node as mandated by the provider
	// operator.
	Version string `json:"version" yaml:"version"`
}

// Resource is structure holding arbitrary conditions of operatorkit resource
// implementations. Imagine an operator implements an instance resource. This
// resource may operates sequentially but has to operate based on a certain
// system state it manages. So it tracks the status as needed here specific to
// its own implementation and means in order to fulfil its premise.
type StatusClusterResource struct {
	// +kubebuilder:validation:Optional
	// +nullable
	Conditions []StatusClusterResourceCondition `json:"conditions,omitempty" yaml:"conditions,omitempty"`
	// +kubebuilder:validation:Optional
	Name string `json:"name" yaml:"name"`
}

// StatusClusterResourceCondition expresses the conditions in which an
// operatorkit resource may is.
type StatusClusterResourceCondition struct {
	// +kubebuilder:validation:Optional
	// +nullable
	// LastTransitionTime is the last time the condition transitioned from one
	// status to another.
	LastTransitionTime metav1.Time `json:"lastTransitionTime,omitempty" yaml:"lastTransitionTime,omitempty"`
	// +kubebuilder:validation:Optional
	// Status may be True, False or Unknown.
	Status string `json:"status" yaml:"status"`
	// +kubebuilder:validation:Optional
	// Type may be anything an operatorkit resource may define.
	Type string `json:"type" yaml:"type"`
}

// StatusClusterScaling expresses the current status of desired number of
// worker nodes in guest cluster.
type StatusClusterScaling struct {
	// +kubebuilder:validation:Optional
	DesiredCapacity int `json:"desiredCapacity" yaml:"desiredCapacity"`
}

// StatusClusterVersion expresses the versions in which a guest cluster was and
// may still be.
type StatusClusterVersion struct {
	// +kubebuilder:validation:Optional
	// +nullable
	// TODO date is deprecated due to LastTransitionTime
	// This can be removed ones the new properties are properly used in all tenant
	// clusters.
	//
	//     https://github.com/giantswarm/giantswarm/issues/3988
	//
	Date metav1.Time `json:"date,omitempty" yaml:"date,omitempty"`
	// +kubebuilder:validation:Optional
	// +nullable
	// LastTransitionTime is the last time the condition transitioned from one
	// status to another.
	LastTransitionTime metav1.Time `json:"lastTransitionTime,omitempty" yaml:"lastTransitionTime,omitempty"`
	// +kubebuilder:validation:Optional
	// Semver is some semver version, e.g. 1.0.0.
	Semver string `json:"semver" yaml:"semver"`
}

// DeepCopyInto implements the deep copy magic the k8s codegen is not able to
// generate out of the box.
func (in *StatusClusterVersion) DeepCopyInto(out *StatusClusterVersion) {
	*out = *in
}
