// Copyright 2021 Google LLC. All Rights Reserved.
// 
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// 
//     http://www.apache.org/licenses/LICENSE-2.0
// 
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// GENERATED BY gen_go_data.go
// gen_go_data -package alpha -var YAML_node_pool blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/containerazure/alpha/node_pool.yaml

package alpha

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/containerazure/alpha/node_pool.yaml
var YAML_node_pool = []byte("info:\n  title: ContainerAzure/NodePool\n  description: An Anthos node pool running on Azure.\n  x-dcl-struct-name: NodePool\n  x-dcl-has-iam: false\n  x-dcl-ref:\n    text: API reference\n    url: https://cloud.google.com/anthos/clusters/docs/multi-cloud/reference/rest/v1/projects.locations.azureClusters.azureNodePools\n  x-dcl-guides:\n  - text: Multicloud overview\n    url: https://cloud.google.com/anthos/clusters/docs/multi-cloud\npaths:\n  get:\n    description: The function used to get information about a NodePool\n    parameters:\n    - name: NodePool\n      required: true\n      description: A full instance of a NodePool\n  apply:\n    description: The function used to apply information about a NodePool\n    parameters:\n    - name: NodePool\n      required: true\n      description: A full instance of a NodePool\n  delete:\n    description: The function used to delete a NodePool\n    parameters:\n    - name: NodePool\n      required: true\n      description: A full instance of a NodePool\n  deleteAll:\n    description: The function used to delete all NodePool\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\n    - name: cluster\n      required: true\n      schema:\n        type: string\n  list:\n    description: The function used to list information about many NodePool\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\n    - name: cluster\n      required: true\n      schema:\n        type: string\ncomponents:\n  schemas:\n    NodePool:\n      title: NodePool\n      x-dcl-id: projects/{{project}}/locations/{{location}}/azureClusters/{{cluster}}/azureNodePools/{{name}}\n      x-dcl-parent-container: project\n      x-dcl-has-iam: false\n      type: object\n      required:\n      - name\n      - version\n      - config\n      - subnetId\n      - autoscaling\n      - maxPodsConstraint\n      - project\n      - location\n      - cluster\n      properties:\n        annotations:\n          type: object\n          additionalProperties:\n            type: string\n          x-dcl-go-name: Annotations\n          description: 'Optional. Annotations on the node pool. This field has the\n            same restrictions as Kubernetes annotations. The total size of all keys\n            and values combined is limited to 256k. Keys can have 2 segments: prefix\n            (optional) and name (required), separated by a slash (/). Prefix must\n            be a DNS subdomain. Name must be 63 characters or less, begin and end\n            with alphanumerics, with dashes (-), underscores (_), dots (.), and alphanumerics\n            between.'\n        autoscaling:\n          type: object\n          x-dcl-go-name: Autoscaling\n          x-dcl-go-type: NodePoolAutoscaling\n          description: Required. Autoscaler configuration for this node pool.\n          x-kubernetes-immutable: true\n          required:\n          - minNodeCount\n          - maxNodeCount\n          properties:\n            maxNodeCount:\n              type: integer\n              format: int64\n              x-dcl-go-name: MaxNodeCount\n              description: Required. Maximum number of nodes in the node pool. Must\n                be >= min_node_count.\n              x-kubernetes-immutable: true\n            minNodeCount:\n              type: integer\n              format: int64\n              x-dcl-go-name: MinNodeCount\n              description: Required. Minimum number of nodes in the node pool. Must\n                be >= 1 and <= max_node_count.\n              x-kubernetes-immutable: true\n        azureAvailabilityZone:\n          type: string\n          x-dcl-go-name: AzureAvailabilityZone\n          description: Optional. The Azure availability zone of the nodes in this\n            nodepool. When unspecified, it defaults to `1`.\n          x-kubernetes-immutable: true\n          x-dcl-server-default: true\n        cluster:\n          type: string\n          x-dcl-go-name: Cluster\n          description: The azureCluster for the resource\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Gkemulticloud/Cluster\n            field: name\n            parent: true\n        config:\n          type: object\n          x-dcl-go-name: Config\n          x-dcl-go-type: NodePoolConfig\n          description: Required. The node configuration of the node pool.\n          x-kubernetes-immutable: true\n          required:\n          - sshConfig\n          properties:\n            rootVolume:\n              type: object\n              x-dcl-go-name: RootVolume\n              x-dcl-go-type: NodePoolConfigRootVolume\n              description: Optional. Configuration related to the root volume provisioned\n                for each node pool machine. When unspecified, it defaults to a 32-GiB\n                Azure Disk.\n              x-kubernetes-immutable: true\n              x-dcl-server-default: true\n              properties:\n                sizeGib:\n                  type: integer\n                  format: int64\n                  x-dcl-go-name: SizeGib\n                  description: Optional. The size of the disk, in GiBs. When unspecified,\n                    a default value is provided. See the specific reference in the\n                    parent resource.\n                  x-kubernetes-immutable: true\n                  x-dcl-server-default: true\n            sshConfig:\n              type: object\n              x-dcl-go-name: SshConfig\n              x-dcl-go-type: NodePoolConfigSshConfig\n              description: Required. SSH configuration for how to access the node\n                pool machines.\n              x-kubernetes-immutable: true\n              required:\n              - authorizedKey\n              properties:\n                authorizedKey:\n                  type: string\n                  x-dcl-go-name: AuthorizedKey\n                  description: Required. The SSH public key data for VMs managed by\n                    Anthos. This accepts the authorized_keys file format used in OpenSSH\n                    according to the sshd(8) manual page.\n                  x-kubernetes-immutable: true\n            tags:\n              type: object\n              additionalProperties:\n                type: string\n              x-dcl-go-name: Tags\n              description: Optional. A set of tags to apply to all underlying Azure\n                resources for this node pool. This currently only includes Virtual\n                Machine Scale Sets. Specify at most 50 pairs containing alphanumerics,\n                spaces, and symbols (.+-=_:@/). Keys can be up to 127 Unicode characters.\n                Values can be up to 255 Unicode characters.\n              x-kubernetes-immutable: true\n            vmSize:\n              type: string\n              x-dcl-go-name: VmSize\n              description: 'Optional. The Azure VM size name. Example: `Standard_DS2_v2`.\n                See (/anthos/clusters/docs/azure/reference/supported-vms) for options.\n                When unspecified, it defaults to `Standard_DS2_v2`.'\n              x-kubernetes-immutable: true\n              x-dcl-server-default: true\n        createTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: CreateTime\n          readOnly: true\n          description: Output only. The time at which this node pool was created.\n          x-kubernetes-immutable: true\n        etag:\n          type: string\n          x-dcl-go-name: Etag\n          readOnly: true\n          description: Allows clients to perform consistent read-modify-writes through\n            optimistic concurrency control. May be sent on update and delete requests\n            to ensure the client has an up-to-date value before proceeding.\n          x-kubernetes-immutable: true\n        location:\n          type: string\n          x-dcl-go-name: Location\n          description: The location for the resource\n          x-kubernetes-immutable: true\n        maxPodsConstraint:\n          type: object\n          x-dcl-go-name: MaxPodsConstraint\n          x-dcl-go-type: NodePoolMaxPodsConstraint\n          description: Required. The constraint on the maximum number of pods that\n            can be run simultaneously on a node in the node pool.\n          x-kubernetes-immutable: true\n          required:\n          - maxPodsPerNode\n          properties:\n            maxPodsPerNode:\n              type: integer\n              format: int64\n              x-dcl-go-name: MaxPodsPerNode\n              description: Required. The maximum number of pods to schedule on a single\n                node.\n              x-kubernetes-immutable: true\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: The name of this resource.\n          x-kubernetes-immutable: true\n        project:\n          type: string\n          x-dcl-go-name: Project\n          description: The project for the resource\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/Project\n            field: name\n            parent: true\n        reconciling:\n          type: boolean\n          x-dcl-go-name: Reconciling\n          readOnly: true\n          description: Output only. If set, there are currently pending changes to\n            the node pool.\n          x-kubernetes-immutable: true\n        state:\n          type: string\n          x-dcl-go-name: State\n          x-dcl-go-type: NodePoolStateEnum\n          readOnly: true\n          description: 'Output only. The current state of the node pool. Possible\n            values: STATE_UNSPECIFIED, PROVISIONING, RUNNING, RECONCILING, STOPPING,\n            ERROR, DEGRADED'\n          x-kubernetes-immutable: true\n          enum:\n          - STATE_UNSPECIFIED\n          - PROVISIONING\n          - RUNNING\n          - RECONCILING\n          - STOPPING\n          - ERROR\n          - DEGRADED\n        subnetId:\n          type: string\n          x-dcl-go-name: SubnetId\n          description: Required. The ARM ID of the subnet where the node pool VMs\n            run. Make sure it's a subnet under the virtual network in the cluster\n            configuration.\n          x-kubernetes-immutable: true\n        uid:\n          type: string\n          x-dcl-go-name: Uid\n          readOnly: true\n          description: Output only. A globally unique identifier for the node pool.\n          x-kubernetes-immutable: true\n        updateTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: UpdateTime\n          readOnly: true\n          description: Output only. The time at which this node pool was last updated.\n          x-kubernetes-immutable: true\n        version:\n          type: string\n          x-dcl-go-name: Version\n          description: Required. The Kubernetes version (e.g. `1.19.10-gke.1000`)\n            running on this node pool.\n")

// 10921 bytes
// MD5: 0a73d286062588b82c35c17295166708
