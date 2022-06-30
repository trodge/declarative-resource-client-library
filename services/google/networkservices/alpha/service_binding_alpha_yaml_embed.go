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
// gen_go_data -package alpha -var YAML_service_binding blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/networkservices/alpha/service_binding.yaml

package alpha

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/networkservices/alpha/service_binding.yaml
var YAML_service_binding = []byte("info:\n  title: NetworkServices/ServiceBinding\n  description: The NetworkServices ServiceBinding resource\n  x-dcl-struct-name: ServiceBinding\n  x-dcl-has-iam: false\npaths:\n  get:\n    description: The function used to get information about a ServiceBinding\n    parameters:\n    - name: ServiceBinding\n      required: true\n      description: A full instance of a ServiceBinding\n  apply:\n    description: The function used to apply information about a ServiceBinding\n    parameters:\n    - name: ServiceBinding\n      required: true\n      description: A full instance of a ServiceBinding\n  delete:\n    description: The function used to delete a ServiceBinding\n    parameters:\n    - name: ServiceBinding\n      required: true\n      description: A full instance of a ServiceBinding\n  deleteAll:\n    description: The function used to delete all ServiceBinding\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\n  list:\n    description: The function used to list information about many ServiceBinding\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\ncomponents:\n  schemas:\n    ServiceBinding:\n      title: ServiceBinding\n      x-dcl-id: projects/{{project}}/locations/{{location}}/serviceBindings/{{name}}\n      x-dcl-parent-container: project\n      x-dcl-has-create: true\n      x-dcl-has-iam: false\n      x-dcl-read-timeout: 0\n      x-dcl-apply-timeout: 0\n      x-dcl-delete-timeout: 0\n      type: object\n      required:\n      - name\n      - service\n      - project\n      - location\n      properties:\n        createTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: CreateTime\n          readOnly: true\n          description: Output only. The timestamp when the resource was created.\n          x-kubernetes-immutable: true\n        description:\n          type: string\n          x-dcl-go-name: Description\n          description: Optional. A free-text description of the resource. Max length\n            1024 characters.\n          x-kubernetes-immutable: true\n        labels:\n          type: object\n          additionalProperties:\n            type: string\n          x-dcl-go-name: Labels\n          description: Optional. Set of label tags associated with the ServiceBinding\n            resource.\n          x-kubernetes-immutable: true\n        location:\n          type: string\n          x-dcl-go-name: Location\n          description: The location for the resource\n          x-kubernetes-immutable: true\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: Required. Name of the ServiceBinding resource. It matches pattern\n            `projects/*/locations/global/serviceBindings/service_binding_name>`.\n          x-kubernetes-immutable: true\n        project:\n          type: string\n          x-dcl-go-name: Project\n          description: The project for the resource\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/Project\n            field: name\n            parent: true\n        service:\n          type: string\n          x-dcl-go-name: Service\n          description: Required. The full service directory service name of the format\n            /projects/*/locations/*/namespaces/*/services/*\n          x-kubernetes-immutable: true\n        updateTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: UpdateTime\n          readOnly: true\n          description: Output only. The timestamp when the resource was updated.\n          x-kubernetes-immutable: true\n")

// 3735 bytes
// MD5: 154b74a1954bff11fdae772f802a6fff
