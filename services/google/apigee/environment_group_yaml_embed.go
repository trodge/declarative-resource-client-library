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
// gen_go_data -package apigee -var YAML_environment_group blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/apigee/environment_group.yaml

package apigee

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/apigee/environment_group.yaml
var YAML_environment_group = []byte("info:\n  title: Apigee/EnvironmentGroup\n  description: The Apigee EnvironmentGroup resource\n  x-dcl-struct-name: EnvironmentGroup\n  x-dcl-has-iam: false\npaths:\n  get:\n    description: The function used to get information about a EnvironmentGroup\n    parameters:\n    - name: EnvironmentGroup\n      required: true\n      description: A full instance of a EnvironmentGroup\n  apply:\n    description: The function used to apply information about a EnvironmentGroup\n    parameters:\n    - name: EnvironmentGroup\n      required: true\n      description: A full instance of a EnvironmentGroup\n  delete:\n    description: The function used to delete a EnvironmentGroup\n    parameters:\n    - name: EnvironmentGroup\n      required: true\n      description: A full instance of a EnvironmentGroup\n  deleteAll:\n    description: The function used to delete all EnvironmentGroup\n    parameters:\n    - name: apigeeorganization\n      required: true\n      schema:\n        type: string\n  list:\n    description: The function used to list information about many EnvironmentGroup\n    parameters:\n    - name: apigeeorganization\n      required: true\n      schema:\n        type: string\ncomponents:\n  schemas:\n    EnvironmentGroup:\n      title: EnvironmentGroup\n      x-dcl-id: organizations/{{apigee_organization}}/envgroups/{{name}}\n      x-dcl-has-create: true\n      x-dcl-has-iam: false\n      x-dcl-read-timeout: 0\n      x-dcl-apply-timeout: 0\n      x-dcl-delete-timeout: 0\n      type: object\n      required:\n      - name\n      - hostnames\n      - apigeeOrganization\n      properties:\n        apigeeOrganization:\n          type: string\n          x-dcl-go-name: ApigeeOrganization\n          description: The apigee organization for the resource\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Apigee/Organization\n            field: name\n            parent: true\n        createdAt:\n          type: integer\n          format: int64\n          x-dcl-go-name: CreatedAt\n          readOnly: true\n          x-kubernetes-immutable: true\n        hostnames:\n          type: array\n          x-dcl-go-name: Hostnames\n          description: Required. Host names for this environment group.\n          x-dcl-send-empty: true\n          x-dcl-list-type: list\n          items:\n            type: string\n            x-dcl-go-type: string\n        lastModifiedAt:\n          type: integer\n          format: int64\n          x-dcl-go-name: LastModifiedAt\n          readOnly: true\n          x-kubernetes-immutable: true\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: ID of the environment group.\n          x-kubernetes-immutable: true\n        state:\n          type: string\n          x-dcl-go-name: State\n          x-dcl-go-type: EnvironmentGroupStateEnum\n          readOnly: true\n          description: 'Output only. State of the environment. Values other than ACTIVE\n            means the resource is not ready to use. Possible values: STATE_UNSPECIFIED,\n            CREATING, ACTIVE, DELETING'\n          x-kubernetes-immutable: true\n          enum:\n          - STATE_UNSPECIFIED\n          - CREATING\n          - ACTIVE\n          - DELETING\n")

// 3169 bytes
// MD5: 6ec5282dc9ed723f88529bffaedcc6f4
