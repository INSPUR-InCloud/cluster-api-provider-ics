package cloudprovider

var AntreaDefaultApplyConfigs = []string {
`apiVersion: v1
kind: ServiceAccount
metadata:
  labels:
    app: antrea
  name: antctl
  namespace: kube-system
`,
`apiVersion: v1
kind: ServiceAccount
metadata:
  labels:
    app: antrea
  name: antrea-agent
  namespace: kube-system
`,
`apiVersion: v1
kind: ServiceAccount
metadata:
  labels:
    app: antrea
  name: antrea-controller
  namespace: kube-system
`,
`apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  labels:
    app: antrea
    rbac.authorization.k8s.io/aggregate-to-admin: "true"
    rbac.authorization.k8s.io/aggregate-to-edit: "true"
  name: aggregate-antrea-clustergroups-edit
rules:
- apiGroups:
  - core.antrea.tanzu.vmware.com
  resources:
  - clustergroups
  verbs:
  - get
  - list
  - watch
  - create
  - update
  - patch
  - delete
`,
`apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  labels:
    app: antrea
    rbac.authorization.k8s.io/aggregate-to-view: "true"
  name: aggregate-antrea-clustergroups-view
rules:
- apiGroups:
  - core.antrea.tanzu.vmware.com
  resources:
  - clustergroups
  verbs:
  - get
  - list
  - watch
`,
`apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  labels:
    app: antrea
    rbac.authorization.k8s.io/aggregate-to-admin: "true"
    rbac.authorization.k8s.io/aggregate-to-edit: "true"
  name: aggregate-antrea-policies-edit
rules:
- apiGroups:
  - security.antrea.tanzu.vmware.com
  resources:
  - clusternetworkpolicies
  - networkpolicies
  verbs:
  - get
  - list
  - watch
  - create
  - update
  - patch
  - delete
`,
`apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  labels:
    app: antrea
    rbac.authorization.k8s.io/aggregate-to-view: "true"
  name: aggregate-antrea-policies-view
rules:
- apiGroups:
  - security.antrea.tanzu.vmware.com
  resources:
  - clusternetworkpolicies
  - networkpolicies
  verbs:
  - get
  - list
  - watch
`,
`apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  labels:
    app: antrea
    rbac.authorization.k8s.io/aggregate-to-admin: "true"
    rbac.authorization.k8s.io/aggregate-to-edit: "true"
  name: aggregate-traceflows-edit
rules:
- apiGroups:
  - ops.antrea.tanzu.vmware.com
  resources:
  - traceflows
  verbs:
  - get
  - list
  - watch
  - create
  - update
  - patch
  - delete
`,
`apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  labels:
    app: antrea
    rbac.authorization.k8s.io/aggregate-to-view: "true"
  name: aggregate-traceflows-view
rules:
- apiGroups:
  - ops.antrea.tanzu.vmware.com
  resources:
  - traceflows
  verbs:
  - get
  - list
  - watch
`,
`apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  labels:
    app: antrea
  name: antctl
rules:
- apiGroups:
  - controlplane.antrea.tanzu.vmware.com
  - networking.antrea.tanzu.vmware.com
  resources:
  - networkpolicies
  - appliedtogroups
  - addressgroups
  verbs:
  - get
  - list
- apiGroups:
  - stats.antrea.tanzu.vmware.com
  resources:
  - networkpolicystats
  - antreaclusternetworkpolicystats
  - antreanetworkpolicystats
  verbs:
  - get
  - list
- apiGroups:
  - system.antrea.tanzu.vmware.com
  resources:
  - controllerinfos
  - agentinfos
  verbs:
  - get
- apiGroups:
  - system.antrea.tanzu.vmware.com
  resources:
  - supportbundles
  verbs:
  - get
  - post
- apiGroups:
  - system.antrea.tanzu.vmware.com
  resources:
  - supportbundles/download
  verbs:
  - get
- nonResourceURLs:
  - /agentinfo
  - /addressgroups
  - /appliedtogroups
  - /loglevel
  - /networkpolicies
  - /ovsflows
  - /ovstracing
  - /podinterfaces
  verbs:
  - get
`,
`apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  labels:
    app: antrea
  name: antrea-agent
rules:
- apiGroups:
  - ""
  resources:
  - nodes
  verbs:
  - get
  - watch
  - list
- apiGroups:
  - ""
  resources:
  - pods
  verbs:
  - get
  - watch
  - list
  - patch
- apiGroups:
  - ""
  resources:
  - endpoints
  - services
  verbs:
  - get
  - watch
  - list
- apiGroups:
  - discovery.k8s.io
  resources:
  - endpointslices
  verbs:
  - get
  - watch
  - list
- apiGroups:
  - clusterinformation.antrea.tanzu.vmware.com
  resources:
  - antreaagentinfos
  verbs:
  - get
  - create
  - update
  - delete
- apiGroups:
  - controlplane.antrea.tanzu.vmware.com
  - networking.antrea.tanzu.vmware.com
  resources:
  - networkpolicies
  - appliedtogroups
  - addressgroups
  verbs:
  - get
  - watch
  - list
- apiGroups:
  - controlplane.antrea.tanzu.vmware.com
  resources:
  - nodestatssummaries
  verbs:
  - create
- apiGroups:
  - controlplane.antrea.tanzu.vmware.com
  resources:
  - networkpolicies/status
  verbs:
  - create
  - get
- apiGroups:
  - authentication.k8s.io
  resources:
  - tokenreviews
  verbs:
  - create
- apiGroups:
  - authorization.k8s.io
  resources:
  - subjectaccessreviews
  verbs:
  - create
- apiGroups:
  - ""
  resourceNames:
  - extension-apiserver-authentication
  resources:
  - configmaps
  verbs:
  - get
  - list
  - watch
- apiGroups:
  - ""
  resourceNames:
  - antrea-ca
  resources:
  - configmaps
  verbs:
  - get
  - watch
  - list
- apiGroups:
  - ops.antrea.tanzu.vmware.com
  resources:
  - traceflows
  - traceflows/status
  verbs:
  - get
  - watch
  - list
  - update
  - patch
  - create
  - delete
`,
`apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  labels:
    app: antrea
  name: antrea-controller
rules:
- apiGroups:
  - ""
  resources:
  - nodes
  - pods
  - namespaces
  - services
  verbs:
  - get
  - watch
  - list
- apiGroups:
  - networking.k8s.io
  resources:
  - networkpolicies
  verbs:
  - get
  - watch
  - list
- apiGroups:
  - clusterinformation.antrea.tanzu.vmware.com
  resources:
  - antreacontrollerinfos
  verbs:
  - get
  - create
  - update
  - delete
- apiGroups:
  - clusterinformation.antrea.tanzu.vmware.com
  resources:
  - antreaagentinfos
  verbs:
  - list
  - delete
- apiGroups:
  - authentication.k8s.io
  resources:
  - tokenreviews
  verbs:
  - create
- apiGroups:
  - authorization.k8s.io
  resources:
  - subjectaccessreviews
  verbs:
  - create
- apiGroups:
  - ""
  resourceNames:
  - extension-apiserver-authentication
  resources:
  - configmaps
  verbs:
  - get
  - list
  - watch
- apiGroups:
  - ""
  resourceNames:
  - antrea-ca
  resources:
  - configmaps
  verbs:
  - get
  - update
- apiGroups:
  - apiregistration.k8s.io
  resourceNames:
  - v1alpha1.stats.antrea.tanzu.vmware.com
  - v1beta1.system.antrea.tanzu.vmware.com
  - v1beta2.controlplane.antrea.tanzu.vmware.com
  - v1beta1.controlplane.antrea.tanzu.vmware.com
  - v1beta1.networking.antrea.tanzu.vmware.com
  resources:
  - apiservices
  verbs:
  - get
  - update
- apiGroups:
  - admissionregistration.k8s.io
  resourceNames:
  - crdmutator.antrea.tanzu.vmware.com
  - crdvalidator.antrea.tanzu.vmware.com
  resources:
  - mutatingwebhookconfigurations
  - validatingwebhookconfigurations
  verbs:
  - get
  - update
- apiGroups:
  - security.antrea.tanzu.vmware.com
  resources:
  - clusternetworkpolicies
  - networkpolicies
  verbs:
  - get
  - watch
  - list
- apiGroups:
  - security.antrea.tanzu.vmware.com
  resources:
  - clusternetworkpolicies/status
  - networkpolicies/status
  verbs:
  - update
- apiGroups:
  - security.antrea.tanzu.vmware.com
  resources:
  - tiers
  verbs:
  - get
  - watch
  - list
  - create
  - update
- apiGroups:
  - ops.antrea.tanzu.vmware.com
  resources:
  - traceflows
  - traceflows/status
  verbs:
  - get
  - watch
  - list
  - update
  - patch
  - create
  - delete
- apiGroups:
  - core.antrea.tanzu.vmware.com
  resources:
  - externalentities
  - clustergroups
  verbs:
  - get
  - watch
  - list
- apiGroups:
  - core.antrea.tanzu.vmware.com
  resources:
  - clustergroups/status
  verbs:
  - update
`,
`apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  labels:
    app: antrea
  name: antctl
  namespace: kube-system
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: antctl
subjects:
- kind: ServiceAccount
  name: antctl
  namespace: kube-system
`,
`apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  labels:
    app: antrea
  name: antrea-agent
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: antrea-agent
subjects:
- kind: ServiceAccount
  name: antrea-agent
  namespace: kube-system
`,
`apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  labels:
    app: antrea
  name: antrea-controller
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: antrea-controller
subjects:
- kind: ServiceAccount
  name: antrea-controller
  namespace: kube-system
`,
`apiVersion: v1
kind: ConfigMap
metadata:
  labels:
    app: antrea
  name: antrea-ca
  namespace: kube-system
`,
`apiVersion: v1
data:
  antrea-agent.conf: |
    # FeatureGates is a map of feature names to bools that enable or disable experimental features.
    featureGates:
    # Enable AntreaProxy which provides ServiceLB for in-cluster Services in antrea-agent.
    # It should be enabled on Windows, otherwise NetworkPolicy will not take effect on
    # Service traffic.
    #  AntreaProxy: true

    # Enable EndpointSlice support in AntreaProxy. Don't enable this feature unless that EndpointSlice
    # API version v1beta1 is supported and set as enabled in Kubernetes. If AntreaProxy is not enabled,
    # this flag will not take effect.
    #  EndpointSlice: false

    # Enable traceflow which provides packet tracing feature to diagnose network issue.
    #  Traceflow: true

    # Enable NodePortLocal feature to make the pods reachable externally through NodePort
    #  NodePortLocal: false

    # Enable Antrea ClusterNetworkPolicy feature to complement K8s NetworkPolicy for cluster admins
    # to define security policies which apply to the entire cluster, and Antrea NetworkPolicy
    # feature that supports priorities, rule actions and externalEntities in the future.
    #  AntreaPolicy: false

    # Enable flowexporter which exports polled conntrack connections as IPFIX flow records from each
    # agent to a configured collector.
    #  FlowExporter: false

    # Enable collecting and exposing NetworkPolicy statistics.
    #  NetworkPolicyStats: false

    # Name of the OpenVSwitch bridge antrea-agent will create and use.
    # Make sure it doesn't conflict with your existing OpenVSwitch bridges.
    #ovsBridge: br-int

    # Datapath type to use for the OpenVSwitch bridge created by Antrea. Supported values are:
    # - system
    # - netdev
    # 'system' is the default value and corresponds to the kernel datapath. Use 'netdev' to run
    # OVS in userspace mode. Userspace mode requires the tun device driver to be available.
    #ovsDatapathType: system

    # Name of the interface antrea-agent will create and use for host <--> pod communication.
    # Make sure it doesn't conflict with your existing interfaces.
    #hostGateway: antrea-gw0

    # Determines how traffic is encapsulated. It has the following options:
    # encap(default):    Inter-node Pod traffic is always encapsulated and Pod to external network
    #                    traffic is SNAT'd.
    # noEncap:           Inter-node Pod traffic is not encapsulated; Pod to external network traffic is
    #                    SNAT'd if noSNAT is not set to true. Underlying network must be capable of
    #                    supporting Pod traffic across IP subnets.
    # hybrid:            noEncap if source and destination Nodes are on the same subnet, otherwise encap.
    # networkPolicyOnly: Antrea enforces NetworkPolicy only, and utilizes CNI chaining and delegates Pod
    #                    IPAM and connectivity to the primary CNI.
    #
    #trafficEncapMode: encap

    # Whether or not to SNAT (using the Node IP) the egress traffic from a Pod to the external network.
    # This option is for the noEncap traffic mode only, and the default value is false. In the noEncap
    # mode, if the cluster's Pod CIDR is reachable from the external network, then the Pod traffic to
    # the external network needs not be SNAT'd. In the networkPolicyOnly mode, antrea-agent never
    # performs SNAT and this option will be ignored; for other modes it must be set to false.
    #noSNAT: false

    # Tunnel protocols used for encapsulating traffic across Nodes. Supported values:
    # - geneve (default)
    # - vxlan
    # - gre
    # - stt
    #tunnelType: geneve

    # Default MTU to use for the host gateway interface and the network interface of each Pod.
    # If omitted, antrea-agent will discover the MTU of the Node's primary interface and
    # also adjust MTU to accommodate for tunnel encapsulation overhead (if applicable).
    #defaultMTU: 1450

    # Whether or not to enable IPsec encryption of tunnel traffic. IPsec encryption is only supported
    # for the GRE tunnel type.
    #enableIPSecTunnel: false

    # ClusterIP CIDR range for Services. It's required when AntreaProxy is not enabled, and should be
    # set to the same value as the one specified by --service-cluster-ip-range for kube-apiserver. When
    # AntreaProxy is enabled, this parameter is not needed and will be ignored if provided.
    #serviceCIDR: 10.96.0.0/12

    # ClusterIP CIDR range for IPv6 Services. It's required when using kube-proxy to provide IPv6 Service in a Dual-Stack
    # cluster or an IPv6 only cluster. The value should be the same as the configuration for kube-apiserver specified by
    # --service-cluster-ip-range. When AntreaProxy is enabled, this parameter is not needed.
    # No default value for this field.
    #serviceCIDRv6:

    # The port for the antrea-agent APIServer to serve on.
    # Note that if it's set to another value, the containerPort of the api port of the
    # antrea-agent container must be set to the same value.
    #apiPort: 10350

    # Enable metrics exposure via Prometheus. Initializes Prometheus metrics listener.
    #enablePrometheusMetrics: true

    # Provide the IPFIX collector address as a string with format <HOST>:[<PORT>][:<PROTO>].
    # HOST can either be the DNS name or the IP of the Flow Collector. For example,
    # "flow-aggregator.flow-aggregator.svc" can be provided as DNS name to connect
    # to the Antrea Flow Aggregator service. If IP, it can be either IPv4 or IPv6.
    # However, IPv6 address should be wrapped with [].
    # If PORT is empty, we default to 4739, the standard IPFIX port.
    # If no PROTO is given, we consider "tcp" as default. We support "tcp" and "udp"
    # L4 transport protocols.
    #flowCollectorAddr: "flow-aggregator.flow-aggregator.svc:4739:tcp"

    # Provide flow poll interval as a duration string. This determines how often the flow exporter dumps connections from the conntrack module.
    # Flow poll interval should be greater than or equal to 1s (one second).
    # Valid time units are "ns", "us" (or "μs"), "ms", "s", "m", "h".
    #flowPollInterval: "5s"

    # Provide flow export frequency, which is the number of poll cycles elapsed before flow exporter exports flow records to
    # the flow collector.
    # Flow export frequency should be greater than or equal to 1.
    #flowExportFrequency: 12

    # Enable TLS communication from flow exporter to flow aggregator.
    #enableTLSToFlowAggregator: true

    # Provide the port range used by NodePortLocal. When the NodePortLocal feature is enabled, a port from that range will be assigned
    # whenever a Pod's container defines a specific port to be exposed (each container can define a list of ports as pod.spec.containers[].ports),
    # and all Node traffic directed to that port will be forwarded to the Pod.
    #nplPortRange: 40000-41000

    # Provide the address of Kubernetes apiserver, to override any value provided in kubeconfig or InClusterConfig.
    # Defaults to "". It must be a host string, a host:port pair, or a URL to the base of the apiserver.
    #kubeAPIServerOverride: ""

    # Comma-separated list of Cipher Suites. If omitted, the default Go Cipher Suites will be used.
    # https://golang.org/pkg/crypto/tls/#pkg-constants
    # Note that TLS1.3 Cipher Suites cannot be added to the list. But the apiserver will always
    # prefer TLS1.3 Cipher Suites whenever possible.
    #tlsCipherSuites:

    # TLS min version from: VersionTLS10, VersionTLS11, VersionTLS12, VersionTLS13.
    #tlsMinVersion:
  antrea-cni.conflist: |
    {
        "cniVersion":"0.3.0",
        "name": "antrea",
        "plugins": [
            {
                "type": "antrea",
                "ipam": {
                    "type": "host-local"
                }
            },
            {
                "type": "portmap",
                "capabilities": {"portMappings": true}
            },
            {
                "type": "bandwidth",
                "capabilities": {"bandwidth": true}
            }
        ]
    }
  antrea-controller.conf: |
    # FeatureGates is a map of feature names to bools that enable or disable experimental features.
    featureGates:
    # Enable traceflow which provides packet tracing feature to diagnose network issue.
    #  Traceflow: true

    # Enable Antrea ClusterNetworkPolicy feature to complement K8s NetworkPolicy for cluster admins
    # to define security policies which apply to the entire cluster, and Antrea NetworkPolicy
    # feature that supports priorities, rule actions and externalEntities in the future.
    #  AntreaPolicy: false

    # Enable collecting and exposing NetworkPolicy statistics.
    #  NetworkPolicyStats: false

    # The port for the antrea-controller APIServer to serve on.
    # Note that if it's set to another value, the containerPort of the api port of the
    # antrea-controller container must be set to the same value.
    #apiPort: 10349

    # Enable metrics exposure via Prometheus. Initializes Prometheus metrics listener.
    #enablePrometheusMetrics: true

    # Indicates whether to use auto-generated self-signed TLS certificate.
    # If false, A Secret named "antrea-controller-tls" must be provided with the following keys:
    #   ca.crt: <CA certificate>
    #   tls.crt: <TLS certificate>
    #   tls.key: <TLS private key>
    # And the Secret must be mounted to directory "/var/run/antrea/antrea-controller-tls" of the
    # antrea-controller container.
    #selfSignedCert: true

    # Comma-separated list of Cipher Suites. If omitted, the default Go Cipher Suites will be used.
    # https://golang.org/pkg/crypto/tls/#pkg-constants
    # Note that TLS1.3 Cipher Suites cannot be added to the list. But the apiserver will always
    # prefer TLS1.3 Cipher Suites whenever possible.
    #tlsCipherSuites:

    # TLS min version from: VersionTLS10, VersionTLS11, VersionTLS12, VersionTLS13.
    #tlsMinVersion:
kind: ConfigMap
metadata:
  annotations: {}
  labels:
    app: antrea
  name: antrea-config-md64tc85t9
  namespace: kube-system
`,
`apiVersion: v1
kind: Service
metadata:
  labels:
    app: antrea
  name: antrea
  namespace: kube-system
spec:
  ports:
  - port: 443
    protocol: TCP
    targetPort: api
  selector:
    app: antrea
    component: antrea-controller
`,
`apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    app: antrea
    component: antrea-controller
  name: antrea-controller
  namespace: kube-system
spec:
  replicas: 1
  selector:
    matchLabels:
      app: antrea
      component: antrea-controller
  strategy:
    type: Recreate
  template:
    metadata:
      labels:
        app: antrea
        component: antrea-controller
    spec:
      containers:
      - args:
        - --config
        - /etc/antrea/antrea-controller.conf
        - --logtostderr=false
        - --log_dir=/var/log/antrea
        - --alsologtostderr
        - --log_file_max_size=100
        - --log_file_max_num=4
        - --v=0
        command:
        - antrea-controller
        env:
        - name: POD_NAME
          valueFrom:
            fieldRef:
              fieldPath: metadata.name
        - name: POD_NAMESPACE
          valueFrom:
            fieldRef:
              fieldPath: metadata.namespace
        - name: NODE_NAME
          valueFrom:
            fieldRef:
              fieldPath: spec.nodeName
        - name: SERVICEACCOUNT_NAME
          valueFrom:
            fieldRef:
              fieldPath: spec.serviceAccountName
        image: k8s.gcr.io.ics.com/antrea/antrea-ubuntu:v0.13.1
        livenessProbe:
          failureThreshold: 5
          httpGet:
            host: 127.0.0.1
            path: /livez
            port: api
            scheme: HTTPS
          periodSeconds: 10
          timeoutSeconds: 5
        name: antrea-controller
        ports:
        - containerPort: 10349
          name: api
          protocol: TCP
        readinessProbe:
          failureThreshold: 5
          httpGet:
            host: 127.0.0.1
            path: /readyz
            port: api
            scheme: HTTPS
          initialDelaySeconds: 5
          periodSeconds: 10
          timeoutSeconds: 5
        resources:
          requests:
            cpu: 200m
        volumeMounts:
        - mountPath: /etc/antrea/antrea-controller.conf
          name: antrea-config
          readOnly: true
          subPath: antrea-controller.conf
        - mountPath: /var/run/antrea/antrea-controller-tls
          name: antrea-controller-tls
        - mountPath: /var/log/antrea
          name: host-var-log-antrea
      hostNetwork: true
      nodeSelector:
        kubernetes.io/os: linux
      priorityClassName: system-cluster-critical
      serviceAccountName: antrea-controller
      tolerations:
      - key: CriticalAddonsOnly
        operator: Exists
      - effect: NoSchedule
        key: node-role.kubernetes.io/master
      volumes:
      - configMap:
          name: antrea-config-md64tc85t9
        name: antrea-config
      - name: antrea-controller-tls
        secret:
          defaultMode: 256
          optional: true
          secretName: antrea-controller-tls
      - hostPath:
          path: /var/log/antrea
          type: DirectoryOrCreate
        name: host-var-log-antrea
`,
`apiVersion: apiregistration.k8s.io/v1
kind: APIService
metadata:
  labels:
    app: antrea
  name: v1alpha1.stats.antrea.tanzu.vmware.com
spec:
  group: stats.antrea.tanzu.vmware.com
  groupPriorityMinimum: 100
  service:
    name: antrea
    namespace: kube-system
  version: v1alpha1
  versionPriority: 100
`,
`apiVersion: apiregistration.k8s.io/v1
kind: APIService
metadata:
  labels:
    app: antrea
  name: v1beta1.controlplane.antrea.tanzu.vmware.com
spec:
  group: controlplane.antrea.tanzu.vmware.com
  groupPriorityMinimum: 100
  service:
    name: antrea
    namespace: kube-system
  version: v1beta1
  versionPriority: 100
`,
`apiVersion: apiregistration.k8s.io/v1
kind: APIService
metadata:
  labels:
    app: antrea
  name: v1beta1.networking.antrea.tanzu.vmware.com
spec:
  group: networking.antrea.tanzu.vmware.com
  groupPriorityMinimum: 100
  service:
    name: antrea
    namespace: kube-system
  version: v1beta1
  versionPriority: 100
`,
`apiVersion: apiregistration.k8s.io/v1
kind: APIService
metadata:
  labels:
    app: antrea
  name: v1beta1.system.antrea.tanzu.vmware.com
spec:
  group: system.antrea.tanzu.vmware.com
  groupPriorityMinimum: 100
  service:
    name: antrea
    namespace: kube-system
  version: v1beta1
  versionPriority: 100
`,
`apiVersion: apiregistration.k8s.io/v1
kind: APIService
metadata:
  labels:
    app: antrea
  name: v1beta2.controlplane.antrea.tanzu.vmware.com
spec:
  group: controlplane.antrea.tanzu.vmware.com
  groupPriorityMinimum: 100
  service:
    name: antrea
    namespace: kube-system
  version: v1beta2
  versionPriority: 100
`,
`apiVersion: apps/v1
kind: DaemonSet
metadata:
  labels:
    app: antrea
    component: antrea-agent
  name: antrea-agent
  namespace: kube-system
spec:
  selector:
    matchLabels:
      app: antrea
      component: antrea-agent
  template:
    metadata:
      labels:
        app: antrea
        component: antrea-agent
    spec:
      containers:
      - args:
        - --config
        - /etc/antrea/antrea-agent.conf
        - --logtostderr=false
        - --log_dir=/var/log/antrea
        - --alsologtostderr
        - --log_file_max_size=100
        - --log_file_max_num=4
        - --v=0
        command:
        - antrea-agent
        env:
        - name: POD_NAME
          valueFrom:
            fieldRef:
              fieldPath: metadata.name
        - name: POD_NAMESPACE
          valueFrom:
            fieldRef:
              fieldPath: metadata.namespace
        - name: NODE_NAME
          valueFrom:
            fieldRef:
              fieldPath: spec.nodeName
        image: k8s.gcr.io.ics.com/antrea/antrea-ubuntu:v0.13.1
        livenessProbe:
          exec:
            command:
            - /bin/sh
            - -c
            - container_liveness_probe agent
          failureThreshold: 5
          initialDelaySeconds: 5
          periodSeconds: 10
          timeoutSeconds: 5
        name: antrea-agent
        ports:
        - containerPort: 10350
          name: api
          protocol: TCP
        readinessProbe:
          failureThreshold: 5
          httpGet:
            host: 127.0.0.1
            path: /readyz
            port: api
            scheme: HTTPS
          initialDelaySeconds: 5
          periodSeconds: 10
          timeoutSeconds: 5
        resources:
          requests:
            cpu: 200m
        securityContext:
          privileged: true
        volumeMounts:
        - mountPath: /etc/antrea/antrea-agent.conf
          name: antrea-config
          readOnly: true
          subPath: antrea-agent.conf
        - mountPath: /var/run/antrea
          name: host-var-run-antrea
        - mountPath: /var/run/openvswitch
          name: host-var-run-antrea
          subPath: openvswitch
        - mountPath: /var/lib/cni
          name: host-var-run-antrea
          subPath: cni
        - mountPath: /var/log/antrea
          name: host-var-log-antrea
        - mountPath: /host/proc
          name: host-proc
          readOnly: true
        - mountPath: /host/var/run/netns
          mountPropagation: HostToContainer
          name: host-var-run-netns
          readOnly: true
        - mountPath: /run/xtables.lock
          name: xtables-lock
      - args:
        - --log_file_max_size=100
        - --log_file_max_num=4
        command:
        - start_ovs
        image: k8s.gcr.io.ics.com/antrea/antrea-ubuntu:v0.13.1
        livenessProbe:
          exec:
            command:
            - /bin/sh
            - -c
            - timeout 10 container_liveness_probe ovs
          failureThreshold: 5
          initialDelaySeconds: 5
          periodSeconds: 10
          timeoutSeconds: 10
        name: antrea-ovs
        resources:
          requests:
            cpu: 200m
        securityContext:
          capabilities:
            add:
            - SYS_NICE
            - NET_ADMIN
            - SYS_ADMIN
            - IPC_LOCK
        volumeMounts:
        - mountPath: /var/run/openvswitch
          name: host-var-run-antrea
          subPath: openvswitch
        - mountPath: /var/log/openvswitch
          name: host-var-log-antrea
          subPath: openvswitch
      dnsPolicy: ClusterFirstWithHostNet
      hostNetwork: true
      initContainers:
      - command:
        - install_cni
        image: k8s.gcr.io.ics.com/antrea/antrea-ubuntu:v0.13.1
        name: install-cni
        resources:
          requests:
            cpu: 100m
        securityContext:
          capabilities:
            add:
            - SYS_MODULE
        volumeMounts:
        - mountPath: /etc/antrea/antrea-cni.conflist
          name: antrea-config
          readOnly: true
          subPath: antrea-cni.conflist
        - mountPath: /host/etc/cni/net.d
          name: host-cni-conf
        - mountPath: /host/opt/cni/bin
          name: host-cni-bin
        - mountPath: /lib/modules
          name: host-lib-modules
          readOnly: true
        - mountPath: /var/run/antrea
          name: host-var-run-antrea
      nodeSelector:
        kubernetes.io/os: linux
      priorityClassName: system-node-critical
      serviceAccountName: antrea-agent
      tolerations:
      - key: CriticalAddonsOnly
        operator: Exists
      - effect: NoSchedule
        operator: Exists
      - effect: NoExecute
        operator: Exists
      volumes:
      - configMap:
          name: antrea-config-md64tc85t9
        name: antrea-config
      - hostPath:
          path: /etc/cni/net.d
        name: host-cni-conf
      - hostPath:
          path: /opt/cni/bin
        name: host-cni-bin
      - hostPath:
          path: /proc
        name: host-proc
      - hostPath:
          path: /var/run/netns
        name: host-var-run-netns
      - hostPath:
          path: /var/run/antrea
          type: DirectoryOrCreate
        name: host-var-run-antrea
      - hostPath:
          path: /var/log/antrea
          type: DirectoryOrCreate
        name: host-var-log-antrea
      - hostPath:
          path: /lib/modules
        name: host-lib-modules
      - hostPath:
          path: /run/xtables.lock
          type: FileOrCreate
        name: xtables-lock
  updateStrategy:
    type: RollingUpdate
`,
`apiVersion: admissionregistration.k8s.io/v1
kind: MutatingWebhookConfiguration
metadata:
  labels:
    app: antrea
  name: crdmutator.antrea.tanzu.vmware.com
webhooks:
- admissionReviewVersions:
  - v1
  - v1beta1
  clientConfig:
    service:
      name: antrea
      namespace: kube-system
      path: /mutate/acnp
  name: acnpmutator.antrea.tanzu.vmware.com
  rules:
  - apiGroups:
    - security.antrea.tanzu.vmware.com
    apiVersions:
    - v1alpha1
    operations:
    - CREATE
    - UPDATE
    resources:
    - clusternetworkpolicies
    scope: Cluster
  sideEffects: None
  timeoutSeconds: 5
- admissionReviewVersions:
  - v1
  - v1beta1
  clientConfig:
    service:
      name: antrea
      namespace: kube-system
      path: /mutate/anp
  name: anpmutator.antrea.tanzu.vmware.com
  rules:
  - apiGroups:
    - security.antrea.tanzu.vmware.com
    apiVersions:
    - v1alpha1
    operations:
    - CREATE
    - UPDATE
    resources:
    - networkpolicies
    scope: Namespaced
  sideEffects: None
  timeoutSeconds: 5
`,
`apiVersion: admissionregistration.k8s.io/v1
kind: ValidatingWebhookConfiguration
metadata:
  labels:
    app: antrea
  name: crdvalidator.antrea.tanzu.vmware.com
webhooks:
- admissionReviewVersions:
  - v1
  - v1beta1
  clientConfig:
    service:
      name: antrea
      namespace: kube-system
      path: /validate/tier
  name: tiervalidator.antrea.tanzu.vmware.com
  rules:
  - apiGroups:
    - security.antrea.tanzu.vmware.com
    apiVersions:
    - v1alpha1
    operations:
    - CREATE
    - UPDATE
    - DELETE
    resources:
    - tiers
    scope: Cluster
  sideEffects: None
  timeoutSeconds: 5
- admissionReviewVersions:
  - v1
  - v1beta1
  clientConfig:
    service:
      name: antrea
      namespace: kube-system
      path: /validate/acnp
  name: acnpvalidator.antrea.tanzu.vmware.com
  rules:
  - apiGroups:
    - security.antrea.tanzu.vmware.com
    apiVersions:
    - v1alpha1
    operations:
    - CREATE
    - UPDATE
    resources:
    - clusternetworkpolicies
    scope: Cluster
  sideEffects: None
  timeoutSeconds: 5
- admissionReviewVersions:
  - v1
  - v1beta1
  clientConfig:
    service:
      name: antrea
      namespace: kube-system
      path: /validate/anp
  name: anpvalidator.antrea.tanzu.vmware.com
  rules:
  - apiGroups:
    - security.antrea.tanzu.vmware.com
    apiVersions:
    - v1alpha1
    operations:
    - CREATE
    - UPDATE
    resources:
    - networkpolicies
    scope: Namespaced
  sideEffects: None
  timeoutSeconds: 5
- admissionReviewVersions:
  - v1
  - v1beta1
  clientConfig:
    service:
      name: antrea
      namespace: kube-system
      path: /validate/clustergroup
  name: clustergroupvalidator.antrea.tanzu.vmware.com
  rules:
  - apiGroups:
    - core.antrea.tanzu.vmware.com
    apiVersions:
    - v1alpha2
    operations:
    - CREATE
    - UPDATE
    - DELETE
    resources:
    - clustergroups
    scope: Cluster
  sideEffects: None
  timeoutSeconds: 5
`,
}