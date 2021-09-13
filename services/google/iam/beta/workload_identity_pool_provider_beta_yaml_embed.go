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
// gen_go_data -package beta -var YAML_workload_identity_pool_provider blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/iam/beta/workload_identity_pool_provider.yaml

package beta

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/iam/beta/workload_identity_pool_provider.yaml
var YAML_workload_identity_pool_provider = []byte("info:\n  title: Iam/WorkloadIdentityPoolProvider\n  description: DCL Specification for the Iam WorkloadIdentityPoolProvider resource\n  x-dcl-has-iam: false\npaths:\n  get:\n    description: The function used to get information about a WorkloadIdentityPoolProvider\n    parameters:\n    - name: WorkloadIdentityPoolProvider\n      required: true\n      description: A full instance of a WorkloadIdentityPoolProvider\n  apply:\n    description: The function used to apply information about a WorkloadIdentityPoolProvider\n    parameters:\n    - name: WorkloadIdentityPoolProvider\n      required: true\n      description: A full instance of a WorkloadIdentityPoolProvider\n  delete:\n    description: The function used to delete a WorkloadIdentityPoolProvider\n    parameters:\n    - name: WorkloadIdentityPoolProvider\n      required: true\n      description: A full instance of a WorkloadIdentityPoolProvider\n  deleteAll:\n    description: The function used to delete all WorkloadIdentityPoolProvider\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\n    - name: workloadidentitypool\n      required: true\n      schema:\n        type: string\n  list:\n    description: The function used to list information about many WorkloadIdentityPoolProvider\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\n    - name: workloadidentitypool\n      required: true\n      schema:\n        type: string\ncomponents:\n  schemas:\n    WorkloadIdentityPoolProvider:\n      title: WorkloadIdentityPoolProvider\n      x-dcl-id: projects/{{project}}/locations/{{location}}/workloadIdentityPools/{{workload_identity_pool}}/providers/{{name}}\n      x-dcl-uses-state-hint: true\n      x-dcl-parent-container: project\n      type: object\n      required:\n      - name\n      - project\n      - location\n      - workloadIdentityPool\n      properties:\n        attributeCondition:\n          type: string\n          x-dcl-go-name: AttributeCondition\n          description: '[A Common Expression Language](https://opensource.google/projects/cel)\n            expression, in plain text, to restrict what otherwise valid authentication\n            credentials issued by the provider should not be accepted. The expression\n            must output a boolean representing whether to allow the federation. The\n            following keywords may be referenced in the expressions: * `assertion`:\n            JSON representing the authentication credential issued by the provider.\n            * `google`: The Google attributes mapped from the assertion in the `attribute_mappings`.\n            * `attribute`: The custom attributes mapped from the assertion in the\n            `attribute_mappings`. The maximum length of the attribute condition expression\n            is 4096 characters. If unspecified, all valid authentication credential\n            are accepted. The following example shows how to only allow credentials\n            with a mapped `google.groups` value of `admins`: ``` \"''admins'' in google.groups\"\n            ```'\n        attributeMapping:\n          type: object\n          additionalProperties:\n            type: string\n          x-dcl-go-name: AttributeMapping\n          description: 'Maps attributes from authentication credentials issued by\n            an external identity provider to Google Cloud attributes, such as `subject`\n            and `segment`. Each key must be a string specifying the Google Cloud IAM\n            attribute to map to. The following keys are supported: * `google.subject`:\n            The principal IAM is authenticating. You can reference this value in IAM\n            bindings. This is also the subject that appears in Cloud Logging logs.\n            Cannot exceed 127 characters. * `google.groups`: Groups the external identity\n            belongs to. You can grant groups access to resources using an IAM `principalSet`\n            binding; access applies to all members of the group. You can also provide\n            custom attributes by specifying `attribute.{custom_attribute}`, where\n            `{custom_attribute}` is the name of the custom attribute to be mapped.\n            You can define a maximum of 50 custom attributes. The maximum length of\n            a mapped attribute key is 100 characters, and the key may only contain\n            the characters [a-z0-9_]. You can reference these attributes in IAM policies\n            to define fine-grained access for a workload to Google Cloud resources.\n            For example: * `google.subject`: `principal://iam.googleapis.com/projects/{project}/locations/{location}/workloadIdentityPools/{pool}/subject/{value}`\n            * `google.groups`: `principalSet://iam.googleapis.com/projects/{project}/locations/{location}/workloadIdentityPools/{pool}/group/{value}`\n            * `attribute.{custom_attribute}`: `principalSet://iam.googleapis.com/projects/{project}/locations/{location}/workloadIdentityPools/{pool}/attribute.{custom_attribute}/{value}`\n            Each value must be a [Common Expression Language] (https://opensource.google/projects/cel)\n            function that maps an identity provider credential to the normalized attribute\n            specified by the corresponding map key. You can use the `assertion` keyword\n            in the expression to access a JSON representation of the authentication\n            credential issued by the provider. The maximum length of an attribute\n            mapping expression is 2048 characters. When evaluated, the total size\n            of all mapped attributes must not exceed 8KB. For AWS providers, if no\n            attribute mapping is defined, the following default mapping applies: ```\n            { \"google.subject\":\"assertion.arn\", \"attribute.aws_role\": \"assertion.arn.contains(''assumed-role'')\"\n            \" ? assertion.arn.extract(''{account_arn}assumed-role/'')\" \" + ''assumed-role/''\"\n            \" + assertion.arn.extract(''assumed-role/{role_name}/'')\" \" : assertion.arn\",\n            } ``` If any custom attribute mappings are defined, they must include\n            a mapping to the `google.subject` attribute. For OIDC providers, you must\n            supply a custom mapping, which must include the `google.subject` attribute.\n            For example, the following maps the `sub` claim of the incoming credential\n            to the `subject` attribute on a Google token: ``` {\"google.subject\": \"assertion.sub\"}\n            ```'\n        aws:\n          type: object\n          x-dcl-go-name: Aws\n          x-dcl-go-type: WorkloadIdentityPoolProviderAws\n          description: An Amazon Web Services identity provider.\n          x-dcl-conflicts:\n          - oidc\n          required:\n          - accountId\n          properties:\n            accountId:\n              type: string\n              x-dcl-go-name: AccountId\n              description: Required. The AWS account ID.\n            stsUri:\n              type: array\n              x-dcl-go-name: StsUri\n              description: A list of AWS STS URIs that can be used when exchanging\n                credentials. If not provided, any valid AWS STS URI is allowed. URIs\n                must use the form `https://sts.amazonaws.com` or `https://sts.{region}.amazonaws.com`,\n                where {region} is a valid AWS region. You can specify a maximum of\n                25 URIs.\n              x-dcl-send-empty: true\n              x-dcl-list-type: list\n              items:\n                type: string\n                x-dcl-go-type: string\n              x-dcl-mutable-unreadable: true\n        description:\n          type: string\n          x-dcl-go-name: Description\n          description: A description for the provider. Cannot exceed 256 characters.\n        disabled:\n          type: boolean\n          x-dcl-go-name: Disabled\n          description: Whether the provider is disabled. You cannot use a disabled\n            provider to exchange tokens. However, existing tokens still grant access.\n        displayName:\n          type: string\n          x-dcl-go-name: DisplayName\n          description: A display name for the provider. Cannot exceed 32 characters.\n        location:\n          type: string\n          x-dcl-go-name: Location\n          description: The location for the resource\n          x-kubernetes-immutable: true\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: Output only. The resource name of the provider.\n          x-kubernetes-immutable: true\n        oidc:\n          type: object\n          x-dcl-go-name: Oidc\n          x-dcl-go-type: WorkloadIdentityPoolProviderOidc\n          description: An OpenId Connect 1.0 identity provider.\n          x-dcl-conflicts:\n          - aws\n          required:\n          - issuerUri\n          properties:\n            allowedAudiences:\n              type: array\n              x-dcl-go-name: AllowedAudiences\n              description: 'Acceptable values for the `aud` field (audience) in the\n                OIDC token. Token exchange requests are rejected if the token audience\n                does not match one of the configured values. Each audience may be\n                at most 256 characters. A maximum of 10 audiences may be configured.\n                If this list is empty, the OIDC token audience must be equal to the\n                full canonical resource name of the WorkloadIdentityPoolProvider,\n                with or without the HTTPS prefix. For example: ``` //iam.googleapis.com/projects//locations//workloadIdentityPools//providers/\n                https://iam.googleapis.com/projects//locations//workloadIdentityPools//providers/\n                ```'\n              x-dcl-send-empty: true\n              x-dcl-list-type: list\n              items:\n                type: string\n                x-dcl-go-type: string\n            issuerUri:\n              type: string\n              x-dcl-go-name: IssuerUri\n              description: Required. The OIDC issuer URL. Must be an HTTPS endpoint.\n        project:\n          type: string\n          x-dcl-go-name: Project\n          description: The project for the resource\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/Project\n            field: name\n            parent: true\n        state:\n          type: string\n          x-dcl-go-name: State\n          x-dcl-go-type: WorkloadIdentityPoolProviderStateEnum\n          readOnly: true\n          description: 'Output only. The state of the provider. Possible values: STATE_UNSPECIFIED,\n            ACTIVE, DELETED'\n          x-kubernetes-immutable: true\n          enum:\n          - STATE_UNSPECIFIED\n          - ACTIVE\n          - DELETED\n        workloadIdentityPool:\n          type: string\n          x-dcl-go-name: WorkloadIdentityPool\n          description: The workloadIdentityPool for the resource\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Iam/WorkloadIdentityPool\n            field: name\n")

// 11112 bytes
// MD5: baf829a29387d331099531414276ed97
