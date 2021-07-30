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
// gen_go_data -package accesscontextmanager -var YAML_access_level blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/accesscontextmanager/access_level.yaml

package accesscontextmanager

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/accesscontextmanager/access_level.yaml
var YAML_access_level = []byte("info:\n  title: AccessContextManager/AccessLevel\n  description: DCL Specification for the AccessContextManager AccessLevel resource\n  x-dcl-has-iam: false\npaths:\n  get:\n    description: The function used to get information about a AccessLevel\n    parameters:\n    - name: AccessLevel\n      required: true\n      description: A full instance of a AccessLevel\n  apply:\n    description: The function used to apply information about a AccessLevel\n    parameters:\n    - name: AccessLevel\n      required: true\n      description: A full instance of a AccessLevel\n  delete:\n    description: The function used to delete a AccessLevel\n    parameters:\n    - name: AccessLevel\n      required: true\n      description: A full instance of a AccessLevel\n  deleteAll:\n    description: The function used to delete all AccessLevel\n    parameters:\n    - name: policy\n      required: true\n      schema:\n        type: string\n  list:\n    description: The function used to list information about many AccessLevel\n    parameters:\n    - name: policy\n      required: true\n      schema:\n        type: string\ncomponents:\n  schemas:\n    AccessLevel:\n      title: AccessLevel\n      x-dcl-id: accessPolicies/{{policy}}/accessLevels/{{name}}\n      type: object\n      required:\n      - title\n      - name\n      properties:\n        basic:\n          type: object\n          x-dcl-go-name: Basic\n          x-dcl-go-type: AccessLevelBasic\n          description: A set of predefined conditions for the access level and a combining\n            function.\n          properties:\n            combiningFunction:\n              type: string\n              x-dcl-go-name: CombiningFunction\n              x-dcl-go-type: AccessLevelBasicCombiningFunctionEnum\n              description: How the conditions list should be combined to determine\n                if a request is granted this AccessLevel. If AND is used, each Condition\n                in conditions must be satisfied for the AccessLevel to be applied.\n                If OR is used, at least one Condition in conditions must be satisfied\n                for the AccessLevel to be applied. Defaults to AND if unspecified.\n              default: AND\n              enum:\n              - AND\n              - OR\n            conditions:\n              type: array\n              x-dcl-go-name: Conditions\n              description: A set of requirements for the AccessLevel to be granted.\n              x-dcl-send-empty: true\n              x-dcl-list-type: list\n              items:\n                type: object\n                x-dcl-go-type: AccessLevelBasicConditions\n                properties:\n                  devicePolicy:\n                    type: object\n                    x-dcl-go-name: DevicePolicy\n                    x-dcl-go-type: AccessLevelBasicConditionsDevicePolicy\n                    description: Device specific restrictions, all restrictions must\n                      hold for the Condition to be true. If not specified, all devices\n                      are allowed.\n                    properties:\n                      allowedDeviceManagementLevels:\n                        type: array\n                        x-dcl-go-name: AllowedDeviceManagementLevels\n                        description: A list of allowed device management levels. An\n                          empty list allows all management levels.\n                        x-dcl-send-empty: true\n                        x-dcl-list-type: list\n                        items:\n                          type: string\n                          x-dcl-go-type: AccessLevelBasicConditionsDevicePolicyAllowedDeviceManagementLevelsEnum\n                          enum:\n                          - MANAGEMENT_UNSPECIFIED\n                          - NONE\n                          - BASIC\n                          - COMPLETE\n                      allowedEncryptionStatuses:\n                        type: array\n                        x-dcl-go-name: AllowedEncryptionStatuses\n                        description: A list of allowed encryptions statuses. An empty\n                          list allows all statuses.\n                        x-dcl-send-empty: true\n                        x-dcl-list-type: list\n                        items:\n                          type: string\n                          x-dcl-go-type: AccessLevelBasicConditionsDevicePolicyAllowedEncryptionStatusesEnum\n                          enum:\n                          - ENCRYPTION_UNSPECIFIED\n                          - ENCRYPTION_UNSUPPORTED\n                          - UNENCRYPTED\n                          - ENCRYPTED\n                      osConstraints:\n                        type: array\n                        x-dcl-go-name: OSConstraints\n                        description: A list of allowed OS versions. An empty list\n                          allows all types and all versions.\n                        x-dcl-send-empty: true\n                        x-dcl-list-type: list\n                        items:\n                          type: object\n                          x-dcl-go-type: AccessLevelBasicConditionsDevicePolicyOSConstraints\n                          properties:\n                            minimumVersion:\n                              type: string\n                              x-dcl-go-name: MinimumVersion\n                              description: 'The minimum allowed OS version. If not\n                                set, any version of this OS satisfies the constraint.\n                                Format: \"major.minor.patch\" such as \"10.5.301\", \"9.2.1\".'\n                            osType:\n                              type: string\n                              x-dcl-go-name: OSType\n                              x-dcl-go-type: AccessLevelBasicConditionsDevicePolicyOSConstraintsOSTypeEnum\n                              description: 'The operating system type of the device.\n                                Possible values: OS_UNSPECIFIED, DESKTOP_MAC, DESKTOP_WINDOWS,\n                                DESKTOP_LINUX, DESKTOP_CHROME_OS'\n                              enum:\n                              - OS_UNSPECIFIED\n                              - DESKTOP_MAC\n                              - DESKTOP_WINDOWS\n                              - DESKTOP_LINUX\n                              - DESKTOP_CHROME_OS\n                            requireVerifiedChromeOS:\n                              type: boolean\n                              x-dcl-go-name: RequireVerifiedChromeOS\n                              description: Only allows requests from devices with\n                                a verified Chrome OS. Verifications includes requirements\n                                that the device is enterprise-managed, conformant\n                                to domain policies, and the caller has permission\n                                to call the API targeted by the request.\n                      requireAdminApproval:\n                        type: boolean\n                        x-dcl-go-name: RequireAdminApproval\n                        description: Allowed device management levels, an empty list\n                          allows all management levels.\n                      requireCorpOwned:\n                        type: boolean\n                        x-dcl-go-name: RequireCorpOwned\n                        description: Whether the device needs to be corp owned.\n                      requireScreenlock:\n                        type: boolean\n                        x-dcl-go-name: RequireScreenlock\n                        description: Whether or not screenlock is required for the\n                          DevicePolicy to be true. Defaults to false.\n                  ipSubnetworks:\n                    type: array\n                    x-dcl-go-name: IPSubnetworks\n                    description: A list of CIDR block IP subnetwork specification.\n                      May be IPv4 or IPv6. Note that for a CIDR IP address block,\n                      the specified IP address portion must be properly truncated\n                      (i.e. all the host bits must be zero) or the input is considered\n                      malformed. For example, \"192.0.2.0/24\" is accepted but \"192.0.2.1/24\"\n                      is not. Similarly, for IPv6, \"2001:db8::/32\" is accepted whereas\n                      \"2001:db8::1/32\" is not. The originating IP of a request must\n                      be in one of the listed subnets in order for this Condition\n                      to be true. If empty, all IP addresses are allowed.\n                    x-dcl-send-empty: true\n                    x-dcl-list-type: list\n                    items:\n                      type: string\n                      x-dcl-go-type: string\n                  members:\n                    type: array\n                    x-dcl-go-name: Members\n                    description: 'An allowed list of members (users, groups, service\n                      accounts). The signed-in user originating the request must be\n                      a part of one of the provided members. If not specified, a request\n                      may come from any user (logged in/not logged in, not present\n                      in any groups, etc.). Formats: `user:{emailid}`, `group:{emailid}`,\n                      `serviceAccount:{emailid}`'\n                    x-dcl-send-empty: true\n                    x-dcl-list-type: list\n                    items:\n                      type: string\n                      x-dcl-go-type: string\n                  negate:\n                    type: boolean\n                    x-dcl-go-name: Negate\n                    description: Whether to negate the Condition. If true, the Condition\n                      becomes a NAND over its non-empty fields, each field must be\n                      false for the Condition overall to be satisfied. Defaults to\n                      false.\n                  regions:\n                    type: array\n                    x-dcl-go-name: Regions\n                    description: The request must originate from one of the provided\n                      countries/regions.\n                    x-dcl-send-empty: true\n                    x-dcl-list-type: list\n                    items:\n                      type: string\n                      x-dcl-go-type: string\n                  requiredAccessLevels:\n                    type: array\n                    x-dcl-go-name: RequiredAccessLevels\n                    description: 'A list of other access levels defined in the same\n                      Policy, referenced by resource name. Referencing an AccessLevel\n                      which does not exist is an error. All access levels listed must\n                      be granted for the Condition to be true. Format: accessPolicies/{policy_id}/accessLevels/{short_name}'\n                    x-dcl-send-empty: true\n                    x-dcl-list-type: list\n                    items:\n                      type: string\n                      x-dcl-go-type: string\n        createTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: CreateTime\n          description: Time the AccessPolicy was created in UTC.\n          x-kubernetes-immutable: true\n        description:\n          type: string\n          x-dcl-go-name: Description\n          description: Description of the AccessLevel and its use. Does not affect\n            behavior.\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: 'Resource name for the Access Level. The short_name component\n            must begin with a letter and only include alphanumeric and ''_''. Format:\n            accessPolicies/{policy_id}/accessLevels/{short_name}'\n        policy:\n          type: string\n          x-dcl-go-name: Policy\n          description: The access policy for this access level.\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Accesscontextmanager/AccessPolicy\n            field: name\n            parent: true\n        title:\n          type: string\n          x-dcl-go-name: Title\n          description: Human readable title. Must be unique within the Policy.\n        updateTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: UpdateTime\n          description: Time the AccessPolicy was updated in UTC.\n          x-kubernetes-immutable: true\n")

// 12356 bytes
// MD5: 5fd3984fba773a2560e73718e3acb0ec