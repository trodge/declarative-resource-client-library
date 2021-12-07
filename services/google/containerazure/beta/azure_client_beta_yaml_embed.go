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
// gen_go_data -package beta -var YAML_azure_client blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/containerazure/beta/azure_client.yaml

package beta

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/containerazure/beta/azure_client.yaml
var YAML_azure_client = []byte("info:\n  title: ContainerAzure/Client\n  description: The ContainerAzure Client resource\n  x-dcl-struct-name: AzureClient\n  x-dcl-has-iam: false\npaths:\n  get:\n    description: The function used to get information about a Client\n    parameters:\n    - name: Client\n      required: true\n      description: A full instance of a Client\n  apply:\n    description: The function used to apply information about a Client\n    parameters:\n    - name: Client\n      required: true\n      description: A full instance of a Client\n  delete:\n    description: The function used to delete a Client\n    parameters:\n    - name: Client\n      required: true\n      description: A full instance of a Client\n  deleteAll:\n    description: The function used to delete all Client\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\n  list:\n    description: The function used to list information about many Client\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\ncomponents:\n  schemas:\n    Client:\n      title: AzureClient\n      x-dcl-id: projects/{{project}}/locations/{{location}}/azureClients/{{name}}\n      x-dcl-parent-container: project\n      x-dcl-has-iam: false\n      type: object\n      required:\n      - name\n      - tenantId\n      - applicationId\n      - project\n      - location\n      properties:\n        applicationId:\n          type: string\n          x-dcl-go-name: ApplicationId\n          description: Required. The Azure Active Directory Application ID.\n          x-kubernetes-immutable: true\n        certificate:\n          type: string\n          x-dcl-go-name: Certificate\n          readOnly: true\n          description: Output only. The PEM encoded x509 certificate.\n          x-kubernetes-immutable: true\n        createTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: CreateTime\n          readOnly: true\n          description: Output only. The time at which this resource was created.\n          x-kubernetes-immutable: true\n        location:\n          type: string\n          x-dcl-go-name: Location\n          description: The location for the resource\n          x-kubernetes-immutable: true\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: The name of this resource. `AzureClient` resource names are\n            formatted as `projects/<project-number>/locations/<region>/azureClients/<client-id>`.\n            See (https://cloud.google.com/apis/design/resource_names) for more details\n            on Google Cloud resource names.\n          x-kubernetes-immutable: true\n        project:\n          type: string\n          x-dcl-go-name: Project\n          description: The project for the resource\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/Project\n            field: name\n            parent: true\n        tenantId:\n          type: string\n          x-dcl-go-name: TenantId\n          description: Required. The Azure Active Directory Tenant ID.\n          x-kubernetes-immutable: true\n        uid:\n          type: string\n          x-dcl-go-name: Uid\n          readOnly: true\n          description: Output only. A globally unique identifier for the client.\n          x-kubernetes-immutable: true\n")

// 3451 bytes
// MD5: 8cca0fa90f0ae0059c356121567a9d67