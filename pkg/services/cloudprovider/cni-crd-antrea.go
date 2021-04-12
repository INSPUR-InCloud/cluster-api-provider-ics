package cloudprovider

var AntreaCustomResources = []string {
`apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  labels:
    app: antrea
  name: antreaagentinfos.clusterinformation.antrea.tanzu.vmware.com
spec:
  group: clusterinformation.antrea.tanzu.vmware.com
  names:
    kind: AntreaAgentInfo
    plural: antreaagentinfos
    shortNames:
    - aai
    singular: antreaagentinfo
  scope: Cluster
  versions:
  - name: v1beta1
    schema:
      openAPIV3Schema:
        type: object
        x-kubernetes-preserve-unknown-fields: true
    served: true
    storage: true
`,
`apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  labels:
    app: antrea
  name: antreacontrollerinfos.clusterinformation.antrea.tanzu.vmware.com
spec:
  group: clusterinformation.antrea.tanzu.vmware.com
  names:
    kind: AntreaControllerInfo
    plural: antreacontrollerinfos
    shortNames:
    - aci
    singular: antreacontrollerinfo
  scope: Cluster
  versions:
  - name: v1beta1
    schema:
      openAPIV3Schema:
        type: object
        x-kubernetes-preserve-unknown-fields: true
    served: true
    storage: true
`,
`apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  labels:
    app: antrea
  name: clustergroups.core.antrea.tanzu.vmware.com
spec:
  group: core.antrea.tanzu.vmware.com
  names:
    kind: ClusterGroup
    plural: clustergroups
    shortNames:
    - cg
    singular: group
  scope: Cluster
  versions:
  - name: v1alpha2
    schema:
      openAPIV3Schema:
        properties:
          spec:
            properties:
              externalEntitySelector:
                x-kubernetes-preserve-unknown-fields: true
              ipBlock:
                properties:
                  cidr:
                    format: cidr
                    type: string
                type: object
              namespaceSelector:
                x-kubernetes-preserve-unknown-fields: true
              podSelector:
                x-kubernetes-preserve-unknown-fields: true
              serviceReference:
                properties:
                  name:
                    type: string
                  namespace:
                    type: string
                type: object
            type: object
          status:
            properties:
              conditions:
                items:
                  properties:
                    lastTransitionTime:
                      type: string
                    status:
                      type: string
                    type:
                      type: string
                  type: object
                type: array
            type: object
        type: object
    served: true
    storage: true
    subresources:
      status: {}
`,
`apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  labels:
    app: antrea
  name: clusternetworkpolicies.security.antrea.tanzu.vmware.com
spec:
  group: security.antrea.tanzu.vmware.com
  names:
    kind: ClusterNetworkPolicy
    plural: clusternetworkpolicies
    shortNames:
    - cnp
    - acnp
    singular: clusternetworkpolicy
  scope: Cluster
  versions:
  - additionalPrinterColumns:
    - description: The Tier to which this ClusterNetworkPolicy belongs to.
      jsonPath: .spec.tier
      name: Tier
      type: string
    - description: The Priority of this ClusterNetworkPolicy relative to other policies.
      format: float
      jsonPath: .spec.priority
      name: Priority
      type: number
    - description: The total number of Nodes that should realize the NetworkPolicy.
      format: int32
      jsonPath: .status.desiredNodesRealized
      name: Desired Nodes
      type: number
    - description: The number of Nodes that have realized the NetworkPolicy.
      format: int32
      jsonPath: .status.currentNodesRealized
      name: Current Nodes
      type: number
    - jsonPath: .metadata.creationTimestamp
      name: Age
      type: date
    name: v1alpha1
    schema:
      openAPIV3Schema:
        properties:
          spec:
            properties:
              appliedTo:
                items:
                  properties:
                    group:
                      type: string
                    namespaceSelector:
                      x-kubernetes-preserve-unknown-fields: true
                    podSelector:
                      x-kubernetes-preserve-unknown-fields: true
                  type: object
                type: array
              egress:
                items:
                  properties:
                    action:
                      enum:
                      - Allow
                      - Drop
                      type: string
                    appliedTo:
                      items:
                        properties:
                          group:
                            type: string
                          namespaceSelector:
                            x-kubernetes-preserve-unknown-fields: true
                          podSelector:
                            x-kubernetes-preserve-unknown-fields: true
                        type: object
                      type: array
                    enableLogging:
                      type: boolean
                    name:
                      type: string
                    ports:
                      items:
                        properties:
                          endPort:
                            type: integer
                          port:
                            x-kubernetes-int-or-string: true
                          protocol:
                            type: string
                        type: object
                      type: array
                    to:
                      items:
                        properties:
                          group:
                            type: string
                          ipBlock:
                            properties:
                              cidr:
                                format: cidr
                                type: string
                            type: object
                          namespaceSelector:
                            x-kubernetes-preserve-unknown-fields: true
                          podSelector:
                            x-kubernetes-preserve-unknown-fields: true
                        type: object
                      type: array
                  required:
                  - action
                  type: object
                type: array
              ingress:
                items:
                  properties:
                    action:
                      enum:
                      - Allow
                      - Drop
                      type: string
                    appliedTo:
                      items:
                        properties:
                          group:
                            type: string
                          namespaceSelector:
                            x-kubernetes-preserve-unknown-fields: true
                          podSelector:
                            x-kubernetes-preserve-unknown-fields: true
                        type: object
                      type: array
                    enableLogging:
                      type: boolean
                    from:
                      items:
                        properties:
                          group:
                            type: string
                          ipBlock:
                            properties:
                              cidr:
                                format: cidr
                                type: string
                            type: object
                          namespaceSelector:
                            x-kubernetes-preserve-unknown-fields: true
                          podSelector:
                            x-kubernetes-preserve-unknown-fields: true
                        type: object
                      type: array
                    name:
                      type: string
                    ports:
                      items:
                        properties:
                          endPort:
                            type: integer
                          port:
                            x-kubernetes-int-or-string: true
                          protocol:
                            type: string
                        type: object
                      type: array
                  required:
                  - action
                  type: object
                type: array
              priority:
                format: float
                maximum: 10000
                minimum: 1
                type: number
              tier:
                type: string
            required:
            - priority
            type: object
          status:
            properties:
              currentNodesRealized:
                type: integer
              desiredNodesRealized:
                type: integer
              observedGeneration:
                type: integer
              phase:
                type: string
            type: object
        type: object
    served: true
    storage: true
    subresources:
      status: {}
`,
`apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  labels:
    app: antrea
  name: externalentities.core.antrea.tanzu.vmware.com
spec:
  group: core.antrea.tanzu.vmware.com
  names:
    kind: ExternalEntity
    plural: externalentities
    shortNames:
    - ee
    singular: externalentity
  scope: Namespaced
  versions:
  - name: v1alpha2
    schema:
      openAPIV3Schema:
        properties:
          spec:
            properties:
              endpoints:
                items:
                  properties:
                    ip:
                      pattern: ^(((([1]?\d)?\d|2[0-4]\d|25[0-5])\.){3}(([1]?\d)?\d|2[0-4]\d|25[0-5]))|([\da-fA-F]{1,4}(\:[\da-fA-F]{1,4}){7})|(([\da-fA-F]{1,4}:){0,5}::([\da-fA-F]{1,4}:){0,5}[\da-fA-F]{1,4})$
                      type: string
                    name:
                      type: string
                  type: object
                type: array
              externalNode:
                type: string
              ports:
                items:
                  properties:
                    name:
                      type: string
                    port:
                      x-kubernetes-int-or-string: true
                    protocol:
                      type: string
                  type: object
                type: array
            type: object
        type: object
    served: true
    storage: true
  - name: v1alpha1
    schema:
      openAPIV3Schema:
        type: object
    served: false
    storage: false
`,
`apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  labels:
    app: antrea
  name: networkpolicies.security.antrea.tanzu.vmware.com
spec:
  group: security.antrea.tanzu.vmware.com
  names:
    kind: NetworkPolicy
    plural: networkpolicies
    shortNames:
    - netpol
    - anp
    singular: networkpolicy
  scope: Namespaced
  versions:
  - additionalPrinterColumns:
    - description: The Tier to which this Antrea NetworkPolicy belongs to.
      jsonPath: .spec.tier
      name: Tier
      type: string
    - description: The Priority of this Antrea NetworkPolicy relative to other policies.
      format: float
      jsonPath: .spec.priority
      name: Priority
      type: number
    - description: The total number of Nodes that should realize the NetworkPolicy.
      format: int32
      jsonPath: .status.desiredNodesRealized
      name: Desired Nodes
      type: number
    - description: The number of Nodes that have realized the NetworkPolicy.
      format: int32
      jsonPath: .status.currentNodesRealized
      name: Current Nodes
      type: number
    - jsonPath: .metadata.creationTimestamp
      name: Age
      type: date
    name: v1alpha1
    schema:
      openAPIV3Schema:
        properties:
          spec:
            properties:
              appliedTo:
                items:
                  properties:
                    podSelector:
                      type: object
                      x-kubernetes-preserve-unknown-fields: true
                  type: object
                type: array
              egress:
                items:
                  properties:
                    action:
                      enum:
                      - Allow
                      - Drop
                      type: string
                    appliedTo:
                      items:
                        properties:
                          podSelector:
                            type: object
                            x-kubernetes-preserve-unknown-fields: true
                        type: object
                      type: array
                    enableLogging:
                      type: boolean
                    name:
                      type: string
                    ports:
                      items:
                        properties:
                          endPort:
                            type: integer
                          port:
                            x-kubernetes-int-or-string: true
                          protocol:
                            type: string
                        type: object
                      type: array
                    to:
                      items:
                        properties:
                          externalEntitySelector:
                            x-kubernetes-preserve-unknown-fields: true
                          ipBlock:
                            properties:
                              cidr:
                                format: cidr
                                type: string
                            type: object
                          namespaceSelector:
                            x-kubernetes-preserve-unknown-fields: true
                          podSelector:
                            x-kubernetes-preserve-unknown-fields: true
                        type: object
                      type: array
                  required:
                  - action
                  type: object
                type: array
              ingress:
                items:
                  properties:
                    action:
                      enum:
                      - Allow
                      - Drop
                      type: string
                    appliedTo:
                      items:
                        properties:
                          podSelector:
                            type: object
                            x-kubernetes-preserve-unknown-fields: true
                        type: object
                      type: array
                    enableLogging:
                      type: boolean
                    from:
                      items:
                        properties:
                          externalEntitySelector:
                            x-kubernetes-preserve-unknown-fields: true
                          ipBlock:
                            properties:
                              cidr:
                                format: cidr
                                type: string
                            type: object
                          namespaceSelector:
                            x-kubernetes-preserve-unknown-fields: true
                          podSelector:
                            x-kubernetes-preserve-unknown-fields: true
                        type: object
                      type: array
                    name:
                      type: string
                    ports:
                      items:
                        properties:
                          endPort:
                            type: integer
                          port:
                            x-kubernetes-int-or-string: true
                          protocol:
                            type: string
                        type: object
                      type: array
                  required:
                  - action
                  type: object
                type: array
              priority:
                format: float
                maximum: 10000
                minimum: 1
                type: number
              tier:
                type: string
            required:
            - priority
            type: object
          status:
            properties:
              currentNodesRealized:
                type: integer
              desiredNodesRealized:
                type: integer
              observedGeneration:
                type: integer
              phase:
                type: string
            type: object
        type: object
    served: true
    storage: true
    subresources:
      status: {}
`,
`apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  labels:
    app: antrea
  name: tiers.security.antrea.tanzu.vmware.com
spec:
  group: security.antrea.tanzu.vmware.com
  names:
    kind: Tier
    plural: tiers
    shortNames:
    - tr
    singular: tier
  scope: Cluster
  versions:
  - additionalPrinterColumns:
    - description: The Priority of this Tier relative to other Tiers.
      jsonPath: .spec.priority
      name: Priority
      type: integer
    - jsonPath: .metadata.creationTimestamp
      name: Age
      type: date
    name: v1alpha1
    schema:
      openAPIV3Schema:
        properties:
          spec:
            properties:
              description:
                type: string
              priority:
                maximum: 255
                minimum: 0
                type: integer
            required:
            - priority
            type: object
        type: object
    served: true
    storage: true
`,
`apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  labels:
    app: antrea
  name: traceflows.ops.antrea.tanzu.vmware.com
spec:
  group: ops.antrea.tanzu.vmware.com
  names:
    kind: Traceflow
    plural: traceflows
    shortNames:
    - tf
    singular: traceflow
  scope: Cluster
  versions:
  - additionalPrinterColumns:
    - description: The phase of the Traceflow.
      jsonPath: .status.phase
      name: Phase
      type: string
    - description: The name of the source Pod.
      jsonPath: .spec.source.pod
      name: Source-Pod
      priority: 10
      type: string
    - description: The name of the destination Pod.
      jsonPath: .spec.destination.pod
      name: Destination-Pod
      priority: 10
      type: string
    - description: The IP address of the destination.
      jsonPath: .spec.destination.ip
      name: Destination-IP
      priority: 10
      type: string
    - jsonPath: .metadata.creationTimestamp
      name: Age
      type: date
    name: v1alpha1
    schema:
      openAPIV3Schema:
        properties:
          spec:
            properties:
              destination:
                oneOf:
                - required:
                  - pod
                  - namespace
                - required:
                  - service
                  - namespace
                - required:
                  - ip
                properties:
                  ip:
                    pattern: ^(((([1]?\d)?\d|2[0-4]\d|25[0-5])\.){3}(([1]?\d)?\d|2[0-4]\d|25[0-5]))|([\da-fA-F]{1,4}(\:[\da-fA-F]{1,4}){7})|(([\da-fA-F]{1,4}:){0,5}::([\da-fA-F]{1,4}:){0,5}[\da-fA-F]{1,4})$
                    type: string
                  namespace:
                    type: string
                  pod:
                    type: string
                  service:
                    type: string
                type: object
              packet:
                properties:
                  ipHeader:
                    properties:
                      flags:
                        type: integer
                      protocol:
                        type: integer
                      srcIP:
                        pattern: ^(((([1]?\d)?\d|2[0-4]\d|25[0-5])\.){3}(([1]?\d)?\d|2[0-4]\d|25[0-5]))|([\da-fA-F]{1,4}(\:[\da-fA-F]{1,4}){7})|(([\da-fA-F]{1,4}:){0,5}::([\da-fA-F]{1,4}:){0,5}[\da-fA-F]{1,4})$
                        type: string
                      ttl:
                        type: integer
                    type: object
                  ipv6Header:
                    properties:
                      hopLimit:
                        type: integer
                      nextHeader:
                        type: integer
                      srcIP:
                        format: ipv6
                        type: string
                    type: object
                  transportHeader:
                    properties:
                      icmp:
                        properties:
                          id:
                            type: integer
                          sequence:
                            type: integer
                        type: object
                      tcp:
                        properties:
                          dstPort:
                            type: integer
                          flags:
                            type: integer
                          srcPort:
                            type: integer
                        type: object
                      udp:
                        properties:
                          dstPort:
                            type: integer
                          srcPort:
                            type: integer
                        type: object
                    type: object
                type: object
              source:
                properties:
                  namespace:
                    type: string
                  pod:
                    type: string
                required:
                - pod
                - namespace
                type: object
            required:
            - source
            - destination
            type: object
          status:
            properties:
              dataplaneTag:
                type: integer
              phase:
                type: string
              reason:
                type: string
              results:
                items:
                  properties:
                    node:
                      type: string
                    observations:
                      items:
                        properties:
                          action:
                            type: string
                          component:
                            type: string
                          componentInfo:
                            type: string
                          dstMAC:
                            type: string
                          networkPolicy:
                            type: string
                          pod:
                            type: string
                          translatedDstIP:
                            type: string
                          translatedSrcIP:
                            type: string
                          ttl:
                            type: integer
                          tunnelDstIP:
                            type: string
                        type: object
                      type: array
                    role:
                      type: string
                    timestamp:
                      type: integer
                  type: object
                type: array
            type: object
        required:
        - spec
        type: object
    served: true
    storage: true
    subresources:
      status: {}
`,
}