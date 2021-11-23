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
// gen_go_data -package cloudidentity -var YAML_membership blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/cloudidentity/membership.yaml

package cloudidentity

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/cloudidentity/membership.yaml
var YAML_membership = []byte("info:\n  title: Cloudidentity/Membership\n  description: The Cloudidentity Membership resource\n  x-dcl-has-iam: false\npaths:\n  get:\n    description: The function used to get information about a Membership\n    parameters:\n    - name: Membership\n      required: true\n      description: A full instance of a Membership\n  apply:\n    description: The function used to apply information about a Membership\n    parameters:\n    - name: Membership\n      required: true\n      description: A full instance of a Membership\n  delete:\n    description: The function used to delete a Membership\n    parameters:\n    - name: Membership\n      required: true\n      description: A full instance of a Membership\n  deleteAll:\n    description: The function used to delete all Membership\n    parameters:\n    - name: group\n      required: true\n      schema:\n        type: string\n  list:\n    description: The function used to list information about many Membership\n    parameters:\n    - name: group\n      required: true\n      schema:\n        type: string\ncomponents:\n  schemas:\n    Membership:\n      title: Membership\n      x-dcl-id: groups/{{group}}/memberships/{{name}}\n      x-dcl-has-iam: false\n      type: object\n      required:\n      - preferredMemberKey\n      - roles\n      - group\n      properties:\n        createTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: CreateTime\n          readOnly: true\n          description: Output only. The time when the `Membership` was created.\n          x-kubernetes-immutable: true\n        deliverySetting:\n          type: string\n          x-dcl-go-name: DeliverySetting\n          x-dcl-go-type: MembershipDeliverySettingEnum\n          readOnly: true\n          description: 'Output only. Delivery setting associated with the membership.\n            Possible values: DELIVERY_SETTING_UNSPECIFIED, ALL_MAIL, DIGEST, DAILY,\n            NONE, DISABLED'\n          x-kubernetes-immutable: true\n          enum:\n          - DELIVERY_SETTING_UNSPECIFIED\n          - ALL_MAIL\n          - DIGEST\n          - DAILY\n          - NONE\n          - DISABLED\n        displayName:\n          type: object\n          x-dcl-go-name: DisplayName\n          x-dcl-go-type: MembershipDisplayName\n          readOnly: true\n          description: Output only. The display name of this member, if available\n          x-kubernetes-immutable: true\n          properties:\n            familyName:\n              type: string\n              x-dcl-go-name: FamilyName\n              readOnly: true\n              description: Output only. Member's family name\n              x-kubernetes-immutable: true\n            fullName:\n              type: string\n              x-dcl-go-name: FullName\n              readOnly: true\n              description: Output only. Localized UTF-16 full name for the member.\n                Localization is done based on the language in the request and the\n                language of the stored display name.\n              x-kubernetes-immutable: true\n            givenName:\n              type: string\n              x-dcl-go-name: GivenName\n              readOnly: true\n              description: Output only. Member's given name\n              x-kubernetes-immutable: true\n        group:\n          type: string\n          x-dcl-go-name: Group\n          description: The group for the resource\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Cloudidentity/Group\n            field: name\n            parent: true\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: Output only. The [resource name](https://cloud.google.com/apis/design/resource_names)\n            of the `Membership`. Shall be of the form `groups/{group}/memberships/{membership}`.\n          x-dcl-server-generated-parameter: true\n        preferredMemberKey:\n          type: object\n          x-dcl-go-name: PreferredMemberKey\n          x-dcl-go-type: MembershipPreferredMemberKey\n          description: Required. Immutable. The `EntityKey` of the member.\n          x-kubernetes-immutable: true\n          required:\n          - id\n          properties:\n            id:\n              type: string\n              x-dcl-go-name: Id\n              description: The ID of the entity. For Google-managed entities, the\n                `id` must be the email address of a group or user. For external-identity-mapped\n                entities, the `id` must be a string conforming to the Identity Source's\n                requirements. Must be unique within a `namespace`.\n              x-kubernetes-immutable: true\n            namespace:\n              type: string\n              x-dcl-go-name: Namespace\n              description: The namespace in which the entity exists. If not specified,\n                the `EntityKey` represents a Google-managed entity such as a Google\n                user or a Google Group. If specified, the `EntityKey` represents an\n                external-identity-mapped group. The namespace must correspond to an\n                identity source created in Admin Console and must be in the form of\n                `identitysources/{identity_source_id}`.\n              x-kubernetes-immutable: true\n        roles:\n          type: array\n          x-dcl-go-name: Roles\n          description: The `MembershipRole`s that apply to the `Membership`. If unspecified,\n            defaults to a single `MembershipRole` with `name` `MEMBER`. Must not contain\n            duplicate `MembershipRole`s with the same `name`.\n          x-dcl-send-empty: true\n          x-dcl-list-type: set\n          items:\n            type: object\n            x-dcl-go-type: MembershipRoles\n            required:\n            - name\n            properties:\n              expiryDetail:\n                type: object\n                x-dcl-go-name: ExpiryDetail\n                x-dcl-go-type: MembershipRolesExpiryDetail\n                description: The expiry details of the `MembershipRole`. Expiry details\n                  are only supported for `MEMBER` `MembershipRoles`. May be set if\n                  `name` is `MEMBER`. Must not be set if `name` is any other value.\n                x-dcl-send-empty: true\n                properties:\n                  expireTime:\n                    type: string\n                    format: date-time\n                    x-dcl-go-name: ExpireTime\n                    description: The time at which the `MembershipRole` will expire.\n              name:\n                type: string\n                x-dcl-go-name: Name\n              restrictionEvaluations:\n                type: object\n                x-dcl-go-name: RestrictionEvaluations\n                x-dcl-go-type: MembershipRolesRestrictionEvaluations\n                description: Evaluations of restrictions applied to parent group on\n                  this membership.\n                properties:\n                  memberRestrictionEvaluation:\n                    type: object\n                    x-dcl-go-name: MemberRestrictionEvaluation\n                    x-dcl-go-type: MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation\n                    description: Evaluation of the member restriction applied to this\n                      membership. Empty if the user lacks permission to view the restriction\n                      evaluation.\n                    properties:\n                      state:\n                        type: string\n                        x-dcl-go-name: State\n                        x-dcl-go-type: MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationStateEnum\n                        readOnly: true\n                        description: 'Output only. The current state of the restriction\n                          Possible values: ENCRYPTION_STATE_UNSPECIFIED, UNSUPPORTED_BY_DEVICE,\n                          ENCRYPTED, NOT_ENCRYPTED'\n                        enum:\n                        - ENCRYPTION_STATE_UNSPECIFIED\n                        - UNSUPPORTED_BY_DEVICE\n                        - ENCRYPTED\n                        - NOT_ENCRYPTED\n        type:\n          type: string\n          x-dcl-go-name: Type\n          x-dcl-go-type: MembershipTypeEnum\n          readOnly: true\n          description: 'Output only. The type of the membership. Possible values:\n            OWNER_TYPE_UNSPECIFIED, OWNER_TYPE_CUSTOMER, OWNER_TYPE_PARTNER'\n          x-kubernetes-immutable: true\n          enum:\n          - OWNER_TYPE_UNSPECIFIED\n          - OWNER_TYPE_CUSTOMER\n          - OWNER_TYPE_PARTNER\n        updateTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: UpdateTime\n          readOnly: true\n          description: Output only. The time when the `Membership` was last updated.\n          x-kubernetes-immutable: true\n")

// 8749 bytes
// MD5: 974943a2dabfa77f7c0c60f00366a4b0
