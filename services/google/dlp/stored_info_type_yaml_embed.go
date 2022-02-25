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
// gen_go_data -package dlp -var YAML_stored_info_type blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/dlp/stored_info_type.yaml

package dlp

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/dlp/stored_info_type.yaml
var YAML_stored_info_type = []byte("info:\n  title: Dlp/StoredInfoType\n  description: The Dlp StoredInfoType resource\n  x-dcl-struct-name: StoredInfoType\n  x-dcl-has-iam: false\npaths:\n  get:\n    description: The function used to get information about a StoredInfoType\n    parameters:\n    - name: StoredInfoType\n      required: true\n      description: A full instance of a StoredInfoType\n  apply:\n    description: The function used to apply information about a StoredInfoType\n    parameters:\n    - name: StoredInfoType\n      required: true\n      description: A full instance of a StoredInfoType\n  delete:\n    description: The function used to delete a StoredInfoType\n    parameters:\n    - name: StoredInfoType\n      required: true\n      description: A full instance of a StoredInfoType\n  deleteAll:\n    description: The function used to delete all StoredInfoType\n    parameters:\n    - name: location\n      required: true\n      schema:\n        type: string\n    - name: parent\n      required: true\n      schema:\n        type: string\n  list:\n    description: The function used to list information about many StoredInfoType\n    parameters:\n    - name: location\n      required: true\n      schema:\n        type: string\n    - name: parent\n      required: true\n      schema:\n        type: string\ncomponents:\n  schemas:\n    StoredInfoType:\n      title: StoredInfoType\n      x-dcl-id: '{{parent}}/storedInfoTypes/{{name}}'\n      x-dcl-locations:\n      - region\n      x-dcl-has-iam: false\n      type: object\n      required:\n      - name\n      - parent\n      properties:\n        description:\n          type: string\n          x-dcl-go-name: Description\n          description: Description of the StoredInfoType (max 256 characters).\n          x-kubernetes-immutable: true\n        dictionary:\n          type: object\n          x-dcl-go-name: Dictionary\n          x-dcl-go-type: StoredInfoTypeDictionary\n          description: Store dictionary-based CustomInfoType.\n          x-kubernetes-immutable: true\n          x-dcl-conflicts:\n          - largeCustomDictionary\n          - regex\n          properties:\n            cloudStoragePath:\n              type: object\n              x-dcl-go-name: CloudStoragePath\n              x-dcl-go-type: StoredInfoTypeDictionaryCloudStoragePath\n              description: Newline-delimited file of words in Cloud Storage. Only\n                a single file is accepted.\n              x-kubernetes-immutable: true\n              x-dcl-conflicts:\n              - wordList\n              required:\n              - path\n              properties:\n                path:\n                  type: string\n                  x-dcl-go-name: Path\n                  description: 'A url representing a file or path (no wildcards) in\n                    Cloud Storage. Example: gs://[BUCKET_NAME]/dictionary.txt'\n                  x-kubernetes-immutable: true\n                  x-dcl-references:\n                  - resource: Storage/Object\n                    field: selfLink\n            wordList:\n              type: object\n              x-dcl-go-name: WordList\n              x-dcl-go-type: StoredInfoTypeDictionaryWordList\n              description: List of words or phrases to search for.\n              x-kubernetes-immutable: true\n              x-dcl-conflicts:\n              - cloudStoragePath\n              required:\n              - words\n              properties:\n                words:\n                  type: array\n                  x-dcl-go-name: Words\n                  description: Words or phrases defining the dictionary. The dictionary\n                    must contain at least one phrase and every phrase must contain\n                    at least 2 characters that are letters or digits. [required]\n                  x-kubernetes-immutable: true\n                  x-dcl-send-empty: true\n                  x-dcl-list-type: list\n                  items:\n                    type: string\n                    x-dcl-go-type: string\n        displayName:\n          type: string\n          x-dcl-go-name: DisplayName\n          description: Display name of the StoredInfoType (max 256 characters).\n          x-kubernetes-immutable: true\n        largeCustomDictionary:\n          type: object\n          x-dcl-go-name: LargeCustomDictionary\n          x-dcl-go-type: StoredInfoTypeLargeCustomDictionary\n          description: StoredInfoType where findings are defined by a dictionary of\n            phrases.\n          x-kubernetes-immutable: true\n          x-dcl-conflicts:\n          - dictionary\n          - regex\n          properties:\n            bigQueryField:\n              type: object\n              x-dcl-go-name: BigQueryField\n              x-dcl-go-type: StoredInfoTypeLargeCustomDictionaryBigQueryField\n              description: Field in a BigQuery table where each cell represents a\n                dictionary phrase.\n              x-kubernetes-immutable: true\n              x-dcl-conflicts:\n              - cloudStorageFileSet\n              properties:\n                field:\n                  type: object\n                  x-dcl-go-name: Field\n                  x-dcl-go-type: StoredInfoTypeLargeCustomDictionaryBigQueryFieldField\n                  description: Designated field in the BigQuery table.\n                  x-kubernetes-immutable: true\n                  properties:\n                    name:\n                      type: string\n                      x-dcl-go-name: Name\n                      description: Name describing the field.\n                      x-kubernetes-immutable: true\n                table:\n                  type: object\n                  x-dcl-go-name: Table\n                  x-dcl-go-type: StoredInfoTypeLargeCustomDictionaryBigQueryFieldTable\n                  description: Source table of the field.\n                  x-kubernetes-immutable: true\n                  properties:\n                    datasetId:\n                      type: string\n                      x-dcl-go-name: DatasetId\n                      description: Dataset ID of the table.\n                      x-kubernetes-immutable: true\n                      x-dcl-references:\n                      - resource: Bigquery/Dataset\n                        field: name\n                    projectId:\n                      type: string\n                      x-dcl-go-name: ProjectId\n                      description: The Google Cloud Platform project ID of the project\n                        containing the table. If omitted, project ID is inferred from\n                        the API call.\n                      x-kubernetes-immutable: true\n                      x-dcl-references:\n                      - resource: Cloudresourcemanager/Project\n                        field: name\n                    tableId:\n                      type: string\n                      x-dcl-go-name: TableId\n                      description: Name of the table.\n                      x-kubernetes-immutable: true\n                      x-dcl-references:\n                      - resource: Bigquery/Table\n                        field: name\n            cloudStorageFileSet:\n              type: object\n              x-dcl-go-name: CloudStorageFileSet\n              x-dcl-go-type: StoredInfoTypeLargeCustomDictionaryCloudStorageFileSet\n              description: Set of files containing newline-delimited lists of dictionary\n                phrases.\n              x-kubernetes-immutable: true\n              x-dcl-conflicts:\n              - bigQueryField\n              required:\n              - url\n              properties:\n                url:\n                  type: string\n                  x-dcl-go-name: Url\n                  description: The url, in the format `gs:///`. Trailing wildcard\n                    in the path is allowed.\n                  x-kubernetes-immutable: true\n            outputPath:\n              type: object\n              x-dcl-go-name: OutputPath\n              x-dcl-go-type: StoredInfoTypeLargeCustomDictionaryOutputPath\n              description: Location to store dictionary artifacts in Google Cloud\n                Storage. These files will only be accessible by project owners and\n                the DLP API. If any of these artifacts are modified, the dictionary\n                is considered invalid and can no longer be used.\n              x-kubernetes-immutable: true\n              required:\n              - path\n              properties:\n                path:\n                  type: string\n                  x-dcl-go-name: Path\n                  description: 'A url representing a file or path (no wildcards) in\n                    Cloud Storage. Example: gs://[BUCKET_NAME]/dictionary.txt'\n                  x-kubernetes-immutable: true\n                  x-dcl-references:\n                  - resource: Storage/Object\n                    field: selfLink\n        location:\n          type: string\n          x-dcl-go-name: Location\n          description: The location of the resource\n          x-kubernetes-immutable: true\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: Resource name.\n          x-kubernetes-immutable: true\n        parent:\n          type: string\n          x-dcl-go-name: Parent\n          description: The parent of the resource.\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/Organization\n            field: name\n            parent: true\n          - resource: Cloudresourcemanager/Project\n            field: name\n            parent: true\n        regex:\n          type: object\n          x-dcl-go-name: Regex\n          x-dcl-go-type: StoredInfoTypeRegex\n          description: Store regular expression-based StoredInfoType.\n          x-kubernetes-immutable: true\n          x-dcl-conflicts:\n          - largeCustomDictionary\n          - dictionary\n          required:\n          - pattern\n          properties:\n            groupIndexes:\n              type: array\n              x-dcl-go-name: GroupIndexes\n              description: The index of the submatch to extract as findings. When\n                not specified, the entire match is returned. No more than 3 may be\n                included.\n              x-kubernetes-immutable: true\n              x-dcl-send-empty: true\n              x-dcl-list-type: list\n              items:\n                type: integer\n                format: int64\n                x-dcl-go-type: int64\n            pattern:\n              type: string\n              x-dcl-go-name: Pattern\n              description: Pattern defining the regular expression. Its syntax (https://github.com/google/re2/wiki/Syntax)\n                can be found under the google/re2 repository on GitHub.\n              x-kubernetes-immutable: true\n")

// 10661 bytes
// MD5: 44b275be2cc27c9eaeb3b34f4f629b5c
