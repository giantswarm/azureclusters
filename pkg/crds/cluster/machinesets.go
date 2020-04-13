package cluster

import (
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	"sigs.k8s.io/yaml"
)

const machinesetsYAML = `
---
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  annotations:
    controller-gen.kubebuilder.io/version: (devel)
  creationTimestamp: null
  name: machinesets.cluster.x-k8s.io
spec:
  group: cluster.x-k8s.io
  names:
    categories:
    - cluster-api
    kind: MachineSet
    listKind: MachineSetList
    plural: machinesets
    shortNames:
    - ms
    singular: machineset
  scope: Namespaced
  subresources:
    scale:
      labelSelectorPath: .status.selector
      specReplicasPath: .spec.replicas
      statusReplicasPath: .status.replicas
    status: {}
  validation:
    openAPIV3Schema:
      description: MachineSet is the Schema for the machinesets API
      properties:
        apiVersion:
          description: 'APIVersion defines the versioned schema of this representation
            of an object. Servers should convert recognized schemas to the latest
            internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
          type: string
        kind:
          description: 'Kind is a string value representing the REST resource this
            object represents. Servers may infer this from the endpoint the client
            submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
          type: string
        metadata:
          type: object
        spec:
          description: MachineSetSpec defines the desired state of MachineSet
          properties:
            deletePolicy:
              description: DeletePolicy defines the policy used to identify nodes
                to delete when downscaling. Defaults to "Random".  Valid values are
                "Random, "Newest", "Oldest"
              enum:
              - Random
              - Newest
              - Oldest
              type: string
            minReadySeconds:
              description: MinReadySeconds is the minimum number of seconds for which
                a newly created machine should be ready. Defaults to 0 (machine will
                be considered available as soon as it is ready)
              format: int32
              type: integer
            replicas:
              description: Replicas is the number of desired replicas. This is a pointer
                to distinguish between explicit zero and unspecified. Defaults to
                1.
              format: int32
              type: integer
            selector:
              description: 'Selector is a label query over machines that should match
                the replica count. Label keys and values that must match in order
                to be controlled by this MachineSet. It must match the machine template''s
                labels. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#label-selectors'
              properties:
                matchExpressions:
                  description: matchExpressions is a list of label selector requirements.
                    The requirements are ANDed.
                  items:
                    description: A label selector requirement is a selector that contains
                      values, a key, and an operator that relates the key and values.
                    properties:
                      key:
                        description: key is the label key that the selector applies
                          to.
                        type: string
                      operator:
                        description: operator represents a key's relationship to a
                          set of values. Valid operators are In, NotIn, Exists and
                          DoesNotExist.
                        type: string
                      values:
                        description: values is an array of string values. If the operator
                          is In or NotIn, the values array must be non-empty. If the
                          operator is Exists or DoesNotExist, the values array must
                          be empty. This array is replaced during a strategic merge
                          patch.
                        items:
                          type: string
                        type: array
                    required:
                    - key
                    - operator
                    type: object
                  type: array
                matchLabels:
                  additionalProperties:
                    type: string
                  description: matchLabels is a map of {key,value} pairs. A single
                    {key,value} in the matchLabels map is equivalent to an element
                    of matchExpressions, whose key field is "key", the operator is
                    "In", and the values array contains only "value". The requirements
                    are ANDed.
                  type: object
              type: object
            template:
              description: Template is the object that describes the machine that
                will be created if insufficient replicas are detected. Object references
                to custom resources resources are treated as templates.
              properties:
                metadata:
                  description: 'Standard object''s metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata'
                  properties:
                    annotations:
                      additionalProperties:
                        type: string
                      description: 'Annotations is an unstructured key value map stored
                        with a resource that may be set by external tools to store
                        and retrieve arbitrary metadata. They are not queryable and
                        should be preserved when modifying objects. More info: http://kubernetes.io/docs/user-guide/annotations'
                      type: object
                    generateName:
                      description: "GenerateName is an optional prefix, used by the
                        server, to generate a unique name ONLY IF the Name field has
                        not been provided. If this field is used, the name returned
                        to the client will be different than the name passed. This
                        value will also be combined with a unique suffix. The provided
                        value has the same validation rules as the Name field, and
                        may be truncated by the length of the suffix required to make
                        the value unique on the server. \n If this field is specified
                        and the generated name exists, the server will NOT return
                        a 409 - instead, it will either return 201 Created or 500
                        with Reason ServerTimeout indicating a unique name could not
                        be found in the time allotted, and the client should retry
                        (optionally after the time indicated in the Retry-After header).
                        \n Applied only if Name is not specified. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#idempotency"
                      type: string
                    labels:
                      additionalProperties:
                        type: string
                      description: 'Map of string keys and values that can be used
                        to organize and categorize (scope and select) objects. May
                        match selectors of replication controllers and services. More
                        info: http://kubernetes.io/docs/user-guide/labels'
                      type: object
                    name:
                      description: 'Name must be unique within a namespace. Is required
                        when creating resources, although some resources may allow
                        a client to request the generation of an appropriate name
                        automatically. Name is primarily intended for creation idempotence
                        and configuration definition. Cannot be updated. More info:
                        http://kubernetes.io/docs/user-guide/identifiers#names'
                      type: string
                    namespace:
                      description: "Namespace defines the space within each name must
                        be unique. An empty namespace is equivalent to the \"default\"
                        namespace, but \"default\" is the canonical representation.
                        Not all objects are required to be scoped to a namespace -
                        the value of this field for those objects will be empty. \n
                        Must be a DNS_LABEL. Cannot be updated. More info: http://kubernetes.io/docs/user-guide/namespaces"
                      type: string
                    ownerReferences:
                      description: List of objects depended by this object. If ALL
                        objects in the list have been deleted, this object will be
                        garbage collected. If this object is managed by a controller,
                        then an entry in this list will point to this controller,
                        with the controller field set to true. There cannot be more
                        than one managing controller.
                      items:
                        description: OwnerReference contains enough information to
                          let you identify an owning object. An owning object must
                          be in the same namespace as the dependent, or be cluster-scoped,
                          so there is no namespace field.
                        properties:
                          apiVersion:
                            description: API version of the referent.
                            type: string
                          blockOwnerDeletion:
                            description: If true, AND if the owner has the "foregroundDeletion"
                              finalizer, then the owner cannot be deleted from the
                              key-value store until this reference is removed. Defaults
                              to false. To set this field, a user needs "delete" permission
                              of the owner, otherwise 422 (Unprocessable Entity) will
                              be returned.
                            type: boolean
                          controller:
                            description: If true, this reference points to the managing
                              controller.
                            type: boolean
                          kind:
                            description: 'Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
                            type: string
                          name:
                            description: 'Name of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#names'
                            type: string
                          uid:
                            description: 'UID of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#uids'
                            type: string
                        required:
                        - apiVersion
                        - kind
                        - name
                        - uid
                        type: object
                      type: array
                  type: object
                spec:
                  description: 'Specification of the desired behavior of the machine.
                    More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status'
                  properties:
                    bootstrap:
                      description: Bootstrap is a reference to a local struct which
                        encapsulates fields to configure the Machine’s bootstrapping
                        mechanism.
                      properties:
                        configRef:
                          description: ConfigRef is a reference to a bootstrap provider-specific
                            resource that holds configuration details. The reference
                            is optional to allow users/operators to specify Bootstrap.Data
                            without the need of a controller.
                          properties:
                            apiVersion:
                              description: API version of the referent.
                              type: string
                            fieldPath:
                              description: 'If referring to a piece of an object instead
                                of an entire object, this string should contain a
                                valid JSON/Go field access statement, such as desiredState.manifest.containers[2].
                                For example, if the object reference is to a container
                                within a pod, this would take on a value like: "spec.containers{name}"
                                (where "name" refers to the name of the container
                                that triggered the event) or if no container name
                                is specified "spec.containers[2]" (container with
                                index 2 in this pod). This syntax is chosen only to
                                have some well-defined way of referencing a part of
                                an object. TODO: this design is not final and this
                                field is subject to change in the future.'
                              type: string
                            kind:
                              description: 'Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
                              type: string
                            name:
                              description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
                              type: string
                            namespace:
                              description: 'Namespace of the referent. More info:
                                https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
                              type: string
                            resourceVersion:
                              description: 'Specific resourceVersion to which this
                                reference is made, if any. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency'
                              type: string
                            uid:
                              description: 'UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids'
                              type: string
                          type: object
                        data:
                          description: Data contains the bootstrap data, such as cloud-init
                            details scripts. If nil, the Machine should remain in
                            the Pending state.
                          type: string
                      type: object
                    infrastructureRef:
                      description: InfrastructureRef is a required reference to a
                        custom resource offered by an infrastructure provider.
                      properties:
                        apiVersion:
                          description: API version of the referent.
                          type: string
                        fieldPath:
                          description: 'If referring to a piece of an object instead
                            of an entire object, this string should contain a valid
                            JSON/Go field access statement, such as desiredState.manifest.containers[2].
                            For example, if the object reference is to a container
                            within a pod, this would take on a value like: "spec.containers{name}"
                            (where "name" refers to the name of the container that
                            triggered the event) or if no container name is specified
                            "spec.containers[2]" (container with index 2 in this pod).
                            This syntax is chosen only to have some well-defined way
                            of referencing a part of an object. TODO: this design
                            is not final and this field is subject to change in the
                            future.'
                          type: string
                        kind:
                          description: 'Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
                          type: string
                        name:
                          description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
                          type: string
                        namespace:
                          description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
                          type: string
                        resourceVersion:
                          description: 'Specific resourceVersion to which this reference
                            is made, if any. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency'
                          type: string
                        uid:
                          description: 'UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids'
                          type: string
                      type: object
                    metadata:
                      description: 'DEPRECATED: ObjectMeta has no function and isn''t
                        used anywhere.'
                      properties:
                        annotations:
                          additionalProperties:
                            type: string
                          description: 'Annotations is an unstructured key value map
                            stored with a resource that may be set by external tools
                            to store and retrieve arbitrary metadata. They are not
                            queryable and should be preserved when modifying objects.
                            More info: http://kubernetes.io/docs/user-guide/annotations'
                          type: object
                        generateName:
                          description: "GenerateName is an optional prefix, used by
                            the server, to generate a unique name ONLY IF the Name
                            field has not been provided. If this field is used, the
                            name returned to the client will be different than the
                            name passed. This value will also be combined with a unique
                            suffix. The provided value has the same validation rules
                            as the Name field, and may be truncated by the length
                            of the suffix required to make the value unique on the
                            server. \n If this field is specified and the generated
                            name exists, the server will NOT return a 409 - instead,
                            it will either return 201 Created or 500 with Reason ServerTimeout
                            indicating a unique name could not be found in the time
                            allotted, and the client should retry (optionally after
                            the time indicated in the Retry-After header). \n Applied
                            only if Name is not specified. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#idempotency"
                          type: string
                        labels:
                          additionalProperties:
                            type: string
                          description: 'Map of string keys and values that can be
                            used to organize and categorize (scope and select) objects.
                            May match selectors of replication controllers and services.
                            More info: http://kubernetes.io/docs/user-guide/labels'
                          type: object
                        name:
                          description: 'Name must be unique within a namespace. Is
                            required when creating resources, although some resources
                            may allow a client to request the generation of an appropriate
                            name automatically. Name is primarily intended for creation
                            idempotence and configuration definition. Cannot be updated.
                            More info: http://kubernetes.io/docs/user-guide/identifiers#names'
                          type: string
                        namespace:
                          description: "Namespace defines the space within each name
                            must be unique. An empty namespace is equivalent to the
                            \"default\" namespace, but \"default\" is the canonical
                            representation. Not all objects are required to be scoped
                            to a namespace - the value of this field for those objects
                            will be empty. \n Must be a DNS_LABEL. Cannot be updated.
                            More info: http://kubernetes.io/docs/user-guide/namespaces"
                          type: string
                        ownerReferences:
                          description: List of objects depended by this object. If
                            ALL objects in the list have been deleted, this object
                            will be garbage collected. If this object is managed by
                            a controller, then an entry in this list will point to
                            this controller, with the controller field set to true.
                            There cannot be more than one managing controller.
                          items:
                            description: OwnerReference contains enough information
                              to let you identify an owning object. An owning object
                              must be in the same namespace as the dependent, or be
                              cluster-scoped, so there is no namespace field.
                            properties:
                              apiVersion:
                                description: API version of the referent.
                                type: string
                              blockOwnerDeletion:
                                description: If true, AND if the owner has the "foregroundDeletion"
                                  finalizer, then the owner cannot be deleted from
                                  the key-value store until this reference is removed.
                                  Defaults to false. To set this field, a user needs
                                  "delete" permission of the owner, otherwise 422
                                  (Unprocessable Entity) will be returned.
                                type: boolean
                              controller:
                                description: If true, this reference points to the
                                  managing controller.
                                type: boolean
                              kind:
                                description: 'Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
                                type: string
                              name:
                                description: 'Name of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#names'
                                type: string
                              uid:
                                description: 'UID of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#uids'
                                type: string
                            required:
                            - apiVersion
                            - kind
                            - name
                            - uid
                            type: object
                          type: array
                      type: object
                    providerID:
                      description: ProviderID is the identification ID of the machine
                        provided by the provider. This field must match the provider
                        ID as seen on the node object corresponding to this machine.
                        This field is required by higher level consumers of cluster-api.
                        Example use case is cluster autoscaler with cluster-api as
                        provider. Clean-up logic in the autoscaler compares machines
                        to nodes to find out machines at provider which could not
                        get registered as Kubernetes nodes. With cluster-api as a
                        generic out-of-tree provider for autoscaler, this field is
                        required by autoscaler to be able to have a provider view
                        of the list of machines. Another list of nodes is queried
                        from the k8s apiserver and then a comparison is done to find
                        out unregistered machines and are marked for delete. This
                        field will be set by the actuators and consumed by higher
                        level entities like autoscaler that will be interfacing with
                        cluster-api as generic provider.
                      type: string
                    version:
                      description: Version defines the desired Kubernetes version.
                        This field is meant to be optionally used by bootstrap providers.
                      type: string
                  required:
                  - bootstrap
                  - infrastructureRef
                  type: object
              type: object
          required:
          - selector
          type: object
        status:
          description: MachineSetStatus defines the observed state of MachineSet
          properties:
            availableReplicas:
              description: The number of available replicas (ready for at least minReadySeconds)
                for this MachineSet.
              format: int32
              type: integer
            errorMessage:
              type: string
            errorReason:
              description: "In the event that there is a terminal problem reconciling
                the replicas, both ErrorReason and ErrorMessage will be set. ErrorReason
                will be populated with a succinct value suitable for machine interpretation,
                while ErrorMessage will contain a more verbose string suitable for
                logging and human consumption. \n These fields should not be set for
                transitive errors that a controller faces that are expected to be
                fixed automatically over time (like service outages), but instead
                indicate that something is fundamentally wrong with the MachineTemplate's
                spec or the configuration of the machine controller, and that manual
                intervention is required. Examples of terminal errors would be invalid
                combinations of settings in the spec, values that are unsupported
                by the machine controller, or the responsible machine controller itself
                being critically misconfigured. \n Any transient errors that occur
                during the reconciliation of Machines can be added as events to the
                MachineSet object and/or logged in the controller's output."
              type: string
            fullyLabeledReplicas:
              description: The number of replicas that have labels matching the labels
                of the machine template of the MachineSet.
              format: int32
              type: integer
            observedGeneration:
              description: ObservedGeneration reflects the generation of the most
                recently observed MachineSet.
              format: int64
              type: integer
            readyReplicas:
              description: The number of ready replicas for this MachineSet. A machine
                is considered ready when the node has been created and is "Ready".
              format: int32
              type: integer
            replicas:
              description: Replicas is the most recently observed number of replicas.
              format: int32
              type: integer
            selector:
              description: 'Selector is the same as the label selector but in the
                string format to avoid introspection by clients. The string will be
                in the same format as the query-param syntax. More info about label
                selectors: http://kubernetes.io/docs/user-guide/labels#label-selectors'
              type: string
          required:
          - replicas
          type: object
      type: object
  version: v1alpha2
  versions:
  - name: v1alpha2
    served: true
    storage: true
status:
  acceptedNames:
    kind: ""
    plural: ""
  conditions: []
  storedVersions: []
`

func NewMachineSetCRD() *v1beta1.CustomResourceDefinition {
	var crd v1beta1.CustomResourceDefinition
	_ = yaml.Unmarshal([]byte(machinesetsYAML), &crd)
	return &crd
}
