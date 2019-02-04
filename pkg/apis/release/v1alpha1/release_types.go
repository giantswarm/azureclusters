package v1alpha1

import (
	"github.com/giantswarm/to"
	apiextensionsv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	kindRelease = "Release"

	ChangelogKindAdded      ReleaseChangelogKind = "added"
	ChangelogKindChanged    ReleaseChangelogKind = "changed"
	ChangelogKindDeprecated ReleaseChangelogKind = "deprecated"
	ChangelogKindFixed      ReleaseChangelogKind = "fixed"
	ChangelogKindRemoved    ReleaseChangelogKind = "removed"
	ChangelogKindSecurity   ReleaseChangelogKind = "security"
)

type ReleaseChangelogKind string

var releaseCRDValidation = &apiextensionsv1beta1.CustomResourceValidation{
	// See http://json-schema.org/learn.
	OpenAPIV3Schema: &apiextensionsv1beta1.JSONSchemaProps{
		Properties: map[string]apiextensionsv1beta1.JSONSchemaProps{
			"spec": {
				Type: "object",
				Properties: map[string]apiextensionsv1beta1.JSONSchemaProps{
					"changelog": {
						Type: "array",
						Items: &apiextensionsv1beta1.JSONSchemaPropsOrArray{
							Schema: &apiextensionsv1beta1.JSONSchemaProps{
								Type: "object",
								Properties: map[string]apiextensionsv1beta1.JSONSchemaProps{
									"component": {
										Type:      "string",
										MinLength: to.Int64P(3),
									},
									"description": {
										Type:      "string",
										MinLength: to.Int64P(3),
									},
									"kind": {
										Enum: []apiextensionsv1beta1.JSON{
											{Raw: []byte("added")},
											{Raw: []byte("changed")},
											{Raw: []byte("deprecated")},
											{Raw: []byte("fixed")},
											{Raw: []byte("removed")},
											{Raw: []byte("security")},
										},
									},
								},
								Required: []string{
									"component",
									"description",
									"kind",
								},
							},
						},
						MinItems: to.Int64P(1),
					},
					"components": {
						Type: "array",
						Items: &apiextensionsv1beta1.JSONSchemaPropsOrArray{
							Schema: &apiextensionsv1beta1.JSONSchemaProps{
								Type: "object",
								Properties: map[string]apiextensionsv1beta1.JSONSchemaProps{
									"name": {
										Type:      "string",
										MinLength: to.Int64P(3),
									},
									"version": {
										Type:      "string",
										MinLength: to.Int64P(5),
									},
								},
							},
						},
						MinItems: to.Int64P(1),
					},
					"parentVersion": {
						Type:    "string",
						Pattern: `^\d+\.\d+\.\d+$`,
					},
					"version": {
						Type:      "string",
						MinLength: to.Int64P(5),
					},
				},
				Required: []string{
					"changelog",
					"components",
					"parentVersion",
					"version",
				},
			},
			"status": {
				Type: "object",
				Properties: map[string]apiextensionsv1beta1.JSONSchemaProps{
					"cycle": NewReleaseCycleCRD().Spec.Validation.OpenAPIV3Schema.Properties["spec"],
				},
			},
		},
	},
}

func NewReleaseCRD() *apiextensionsv1beta1.CustomResourceDefinition {
	return &apiextensionsv1beta1.CustomResourceDefinition{
		TypeMeta: metav1.TypeMeta{
			APIVersion: apiextensionsv1beta1.SchemeGroupVersion.String(),
			Kind:       "CustomResourceDefinition",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: "releases.release.giantswarm.io",
		},
		Spec: apiextensionsv1beta1.CustomResourceDefinitionSpec{
			Group:   "release.giantswarm.io",
			Scope:   "Cluster",
			Version: "v1alpha1",
			Names: apiextensionsv1beta1.CustomResourceDefinitionNames{
				Kind:     "Release",
				Plural:   "releases",
				Singular: "release",
			},
			Subresources: &apiextensionsv1beta1.CustomResourceSubresources{
				Status: &apiextensionsv1beta1.CustomResourceSubresourceStatus{},
			},
			Validation: releaseCRDValidation,
		},
	}
}

func NewReleaseTypeMeta() metav1.TypeMeta {
	return metav1.TypeMeta{
		APIVersion: version,
		Kind:       kindRelease,
	}
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type Release struct {
	metav1.TypeMeta   `json:",inline" yaml:",inline"`
	metav1.ObjectMeta `json:"metadata" yaml:"metadata"`
	Spec              ReleaseSpec   `json:"spec" yaml:"spec"`
	Status            ReleaseStatus `json:"status,omitempty" yaml:"status,omitempty"`
}

type ReleaseSpec struct {
	// Changelog is the changelog since ParentVersion.
	Changelog []ReleaseSpecChangelogEntry `json:"changelog" yaml:"changelog"`
	// Components describes components managing this release.
	Components []ReleaseSpecComponent `json:"components" yaml:"components"`
	// ParentVersion is a version from which the changes in changelog are
	// described. We need that because we may introduce bug fixes after
	// next major release and then taking previous semver version may not
	// render correct changelog. This should always be in the semver format
	// without the "v" prefix.
	ParentVersion string `json:"parentVersion" yaml:"parentVersion"`
	// Version is the version of the release. Releases with semver version
	// (without the "v" prefix) are taken from control-plane AppCatalog.
	// All other releases are taken from control-plane-test AppCatalog.
	Version string `json:"version" yaml:"version"`
}

type ReleaseSpecChangelogEntry struct {
	// Component name.
	Component string `json:"component" yaml:"component"`
	// Description of the component changes expressed in full sentence.
	Description string `json:"description" yaml:"description"`
	// Kind of the component changes. It can be one of: "added", "changed",
	// "deprecated", "fixed", "removed", "security".
	Kind ReleaseChangelogKind `json:"kind" yaml:"kind"`
}

type ReleaseSpecChangelogEntryKind string

type ReleaseSpecComponent struct {
	// Name of the release component.
	Name string `json:"name" yaml:"name"`
	// Version of the release component.
	Version string `json:"version" yaml:"version"`
}

type ReleaseStatus struct {
	// Cycle is the most recent observed copy of the specification of the
	// ReleaseCycle CR referencing this Release CR.
	Cycle ReleaseCycleSpec `json:"cycle,omitempty" yaml:"cycle,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ReleaseList struct {
	metav1.TypeMeta `json:",inline" yaml:",inline"`
	metav1.ListMeta `json:"metadata" yaml:"metadata"`
	Items           []Release `json:"items" yaml:"items"`
}
