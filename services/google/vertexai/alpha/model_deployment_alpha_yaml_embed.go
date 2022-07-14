// Copyright 2022 Google LLC. All Rights Reserved.
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
// gen_go_data -package alpha -var YAML_model_deployment blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/vertexai/alpha/model_deployment.yaml

package alpha

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/vertexai/alpha/model_deployment.yaml
var YAML_model_deployment = []byte("info:\n  title: VertexAI/ModelDeployment\n  description: The VertexAI ModelDeployment resource\n  x-dcl-struct-name: ModelDeployment\n  x-dcl-has-iam: false\npaths:\n  get:\n    description: The function used to get information about a ModelDeployment\n    parameters:\n    - name: ModelDeployment\n      required: true\n      description: A full instance of a ModelDeployment\n  apply:\n    description: The function used to apply information about a ModelDeployment\n    parameters:\n    - name: ModelDeployment\n      required: true\n      description: A full instance of a ModelDeployment\n  delete:\n    description: The function used to delete a ModelDeployment\n    parameters:\n    - name: ModelDeployment\n      required: true\n      description: A full instance of a ModelDeployment\n  deleteAll:\n    description: The function used to delete all ModelDeployment\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\n    - name: endpoint\n      required: true\n      schema:\n        type: string\n  list:\n    description: The function used to list information about many ModelDeployment\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\n    - name: endpoint\n      required: true\n      schema:\n        type: string\ncomponents:\n  schemas:\n    ModelDeployment:\n      title: ModelDeployment\n      x-dcl-id: projects/{{project}}/locations/{{location}}/endpoints/{{endpoint}}/models/{{model}}\n      x-dcl-parent-container: project\n      x-dcl-has-create: true\n      x-dcl-has-iam: false\n      x-dcl-read-timeout: 0\n      x-dcl-apply-timeout: 0\n      x-dcl-delete-timeout: 0\n      type: object\n      required:\n      - model\n      - dedicatedResources\n      - endpoint\n      properties:\n        dedicatedResources:\n          type: object\n          x-dcl-go-name: DedicatedResources\n          x-dcl-go-type: ModelDeploymentDedicatedResources\n          description: A description of resources that are dedicated to the DeployedModel,\n            and that need a higher degree of manual configuration.\n          x-kubernetes-immutable: true\n          required:\n          - machineSpec\n          - minReplicaCount\n          properties:\n            machineSpec:\n              type: object\n              x-dcl-go-name: MachineSpec\n              x-dcl-go-type: ModelDeploymentDedicatedResourcesMachineSpec\n              description: Required. Immutable. The specification of a single machine\n                used by the prediction.\n              x-kubernetes-immutable: true\n              required:\n              - machineType\n              properties:\n                machineType:\n                  type: string\n                  x-dcl-go-name: MachineType\n                  description: 'Immutable. The type of the machine. See the [list\n                    of machine types supported for prediction](https://cloud.google.com/vertex-ai/docs/predictions/configure-compute#machine-types)\n                    See the [list of machine types supported for custom training](https://cloud.google.com/vertex-ai/docs/training/configure-compute#machine-types).\n                    For DeployedModel this field is optional, and the default value\n                    is `n1-standard-2`. For BatchPredictionJob or as part of WorkerPoolSpec\n                    this field is required. TODO(rsurowka): Try to better unify the\n                    required vs optional.'\n                  x-kubernetes-immutable: true\n            maxReplicaCount:\n              type: integer\n              format: int64\n              x-dcl-go-name: MaxReplicaCount\n              description: Immutable. The maximum number of replicas this DeployedModel\n                may be deployed on when the traffic against it increases. If the requested\n                value is too large, the deployment will error, but if deployment succeeds\n                then the ability to scale the model to that many replicas is guaranteed\n                (barring service outages). If traffic against the DeployedModel increases\n                beyond what its replicas at maximum may handle, a portion of the traffic\n                will be dropped. If this value is not provided, will use min_replica_count\n                as the default value. The value of this field impacts the charge against\n                Vertex CPU and GPU quotas. Specifically, you will be charged for max_replica_count\n                * number of cores in the selected machine type) and (max_replica_count\n                * number of GPUs per replica in the selected machine type).\n              x-kubernetes-immutable: true\n              x-dcl-server-default: true\n            minReplicaCount:\n              type: integer\n              format: int64\n              x-dcl-go-name: MinReplicaCount\n              description: Required. Immutable. The minimum number of machine replicas\n                this DeployedModel will be always deployed on. This value must be\n                greater than or equal to 1. If traffic against the DeployedModel increases,\n                it may dynamically be deployed onto more replicas, and as traffic\n                decreases, some of these extra replicas may be freed.\n              x-kubernetes-immutable: true\n        deployedModelId:\n          type: string\n          x-dcl-go-name: DeployedModelId\n          readOnly: true\n          description: The deployed ID of the model in the endpoint\n          x-kubernetes-immutable: true\n        endpoint:\n          type: string\n          x-dcl-go-name: Endpoint\n          description: The name of the endpoint to deploy to\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Vertexai/Endpoint\n            field: name\n            parent: true\n        location:\n          type: string\n          x-dcl-go-name: Location\n          description: The location of the endpoint\n          x-kubernetes-immutable: true\n          x-dcl-extract-if-empty: true\n        model:\n          type: string\n          x-dcl-go-name: Model\n          description: The name of the model to deploy\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Vertexai/Model\n            field: name\n        project:\n          type: string\n          x-dcl-go-name: Project\n          description: The project of the endpoint\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/Project\n            field: name\n            parent: true\n          x-dcl-extract-if-empty: true\n")

// 6682 bytes
// MD5: 9a554d981d65bacd3be9e377ce6b127a
