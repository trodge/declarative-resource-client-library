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
// gen_go_data -package alpha -var YAML_budget blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/billingbudgets/alpha/budget.yaml

package alpha

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/billingbudgets/alpha/budget.yaml
var YAML_budget = []byte("info:\n  title: BillingBudgets/Budget\n  description: The BillingBudgets Budget resource\n  x-dcl-struct-name: Budget\n  x-dcl-has-iam: false\npaths:\n  get:\n    description: The function used to get information about a Budget\n    parameters:\n    - name: Budget\n      required: true\n      description: A full instance of a Budget\n  apply:\n    description: The function used to apply information about a Budget\n    parameters:\n    - name: Budget\n      required: true\n      description: A full instance of a Budget\n  delete:\n    description: The function used to delete a Budget\n    parameters:\n    - name: Budget\n      required: true\n      description: A full instance of a Budget\n  deleteAll:\n    description: The function used to delete all Budget\n    parameters:\n    - name: billingaccount\n      required: true\n      schema:\n        type: string\n  list:\n    description: The function used to list information about many Budget\n    parameters:\n    - name: billingaccount\n      required: true\n      schema:\n        type: string\ncomponents:\n  schemas:\n    Budget:\n      title: Budget\n      x-dcl-id: billingAccounts/{{billing_account}}/budgets/{{name}}\n      x-dcl-has-iam: false\n      type: object\n      required:\n      - amount\n      - billingAccount\n      properties:\n        allUpdatesRule:\n          type: object\n          x-dcl-go-name: AllUpdatesRule\n          x-dcl-go-type: BudgetAllUpdatesRule\n          description: Optional. Rules to apply to notifications sent based on budget\n            spend and thresholds.\n          properties:\n            disableDefaultIamRecipients:\n              type: boolean\n              x-dcl-go-name: DisableDefaultIamRecipients\n              description: Optional. When set to true, disables default notifications\n                sent when a threshold is exceeded. Default notifications are sent\n                to those with Billing Account Administrator and Billing Account User\n                IAM roles for the target account.\n            monitoringNotificationChannels:\n              type: array\n              x-dcl-go-name: MonitoringNotificationChannels\n              description: Optional. Targets to send notifications to when a threshold\n                is exceeded. This is in addition to default recipients who have billing\n                account IAM roles. The value is the full REST resource name of a monitoring\n                notification channel with the form `projects/{project_id}/notificationChannels/{channel_id}`.\n                A maximum of 5 channels are allowed. See https://cloud.google.com/billing/docs/how-to/budgets-notification-recipients\n                for more details.\n              x-dcl-send-empty: true\n              x-dcl-list-type: list\n              items:\n                type: string\n                x-dcl-go-type: string\n                x-dcl-references:\n                - resource: Monitoring/NotificationChannel\n                  field: name\n            pubsubTopic:\n              type: string\n              x-dcl-go-name: PubsubTopic\n              description: Optional. The name of the Pub/Sub topic where budget related\n                messages will be published, in the form `projects/{project_id}/topics/{topic_id}`.\n                Updates are sent at regular intervals to the topic. The topic needs\n                to be created before the budget is created; see https://cloud.google.com/billing/docs/how-to/budgets-programmatic-notifications\n                for more details. Caller is expected to have `pubsub.topics.setIamPolicy`\n                permission on the topic when it's set for a budget, otherwise, the\n                API call will fail with PERMISSION_DENIED. See https://cloud.google.com/billing/docs/how-to/budgets-programmatic-notifications#permissions_required_for_this_task\n                for more details on Pub/Sub roles and permissions.\n              x-dcl-references:\n              - resource: Pubsub/Topic\n                field: name\n            schemaVersion:\n              type: string\n              x-dcl-go-name: SchemaVersion\n              description: Optional. Required when NotificationsRule.pubsub_topic\n                is set. The schema version of the notification sent to NotificationsRule.pubsub_topic.\n                Only \"1.0\" is accepted. It represents the JSON schema as defined in\n                https://cloud.google.com/billing/docs/how-to/budgets-programmatic-notifications#notification_format.\n        amount:\n          type: object\n          x-dcl-go-name: Amount\n          x-dcl-go-type: BudgetAmount\n          description: Required. Budgeted amount.\n          properties:\n            lastPeriodAmount:\n              type: object\n              x-dcl-go-name: LastPeriodAmount\n              x-dcl-go-type: BudgetAmountLastPeriodAmount\n              description: Use the last period's actual spend as the budget for the\n                present period. LastPeriodAmount can only be set when the budget's\n                time period is a .\n              x-dcl-conflicts:\n              - specifiedAmount\n            specifiedAmount:\n              type: object\n              x-dcl-go-name: SpecifiedAmount\n              x-dcl-go-type: BudgetAmountSpecifiedAmount\n              description: A specified amount to use as the budget. `currency_code`\n                is optional. If specified when creating a budget, it must match the\n                currency of the billing account. If specified when updating a budget,\n                it must match the currency_code of the existing budget. The `currency_code`\n                is provided on output.\n              x-dcl-conflicts:\n              - lastPeriodAmount\n              properties:\n                currencyCode:\n                  type: string\n                  x-dcl-go-name: CurrencyCode\n                  description: The three-letter currency code defined in ISO 4217.\n                  x-kubernetes-immutable: true\n                nanos:\n                  type: integer\n                  format: int64\n                  x-dcl-go-name: Nanos\n                  description: Number of nano (10^-9) units of the amount. The value\n                    must be between -999,999,999 and +999,999,999 inclusive. If `units`\n                    is positive, `nanos` must be positive or zero. If `units` is zero,\n                    `nanos` can be positive, zero, or negative. If `units` is negative,\n                    `nanos` must be negative or zero. For example $-1.75 is represented\n                    as `units`=-1 and `nanos`=-750,000,000.\n                  x-dcl-send-empty: true\n                units:\n                  type: integer\n                  format: int64\n                  x-dcl-go-name: Units\n                  description: The whole units of the amount. For example if `currencyCode`\n                    is `\"USD\"`, then 1 unit is one US dollar.\n                  x-dcl-send-empty: true\n        billingAccount:\n          type: string\n          x-dcl-go-name: BillingAccount\n          description: The billing account of the resource\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/BillingAccount\n            field: name\n            parent: true\n        budgetFilter:\n          type: object\n          x-dcl-go-name: BudgetFilter\n          x-dcl-go-type: BudgetBudgetFilter\n          description: Optional. Filters that define which resources are used to compute\n            the actual spend against the budget amount, such as projects, services,\n            and the budget's time period, as well as other filters.\n          properties:\n            calendarPeriod:\n              type: string\n              x-dcl-go-name: CalendarPeriod\n              x-dcl-go-type: BudgetBudgetFilterCalendarPeriodEnum\n              description: 'Optional. Specifies to track usage for recurring calendar\n                period. For example, assume that CalendarPeriod.QUARTER is set. The\n                budget will track usage from April 1 to June 30, when the current\n                calendar month is April, May, June. After that, it will track usage\n                from July 1 to September 30 when the current calendar month is July,\n                August, September, so on. Possible values: CALENDAR_PERIOD_UNSPECIFIED,\n                MONTH, QUARTER, YEAR'\n              x-dcl-conflicts:\n              - customPeriod\n              x-dcl-server-default: true\n              enum:\n              - CALENDAR_PERIOD_UNSPECIFIED\n              - MONTH\n              - QUARTER\n              - YEAR\n            creditTypes:\n              type: array\n              x-dcl-go-name: CreditTypes\n              description: Optional. If Filter.credit_types_treatment is INCLUDE_SPECIFIED_CREDITS,\n                this is a list of credit types to be subtracted from gross cost to\n                determine the spend for threshold calculations. See a list of acceptable\n                credit type values. If Filter.credit_types_treatment is not INCLUDE_SPECIFIED_CREDITS,\n                this field must be empty.\n              x-dcl-send-empty: true\n              x-dcl-list-type: list\n              items:\n                type: string\n                x-dcl-go-type: string\n            creditTypesTreatment:\n              type: string\n              x-dcl-go-name: CreditTypesTreatment\n              x-dcl-go-type: BudgetBudgetFilterCreditTypesTreatmentEnum\n              description: Optional. If not set, default behavior is `INCLUDE_ALL_CREDITS`.\n              x-dcl-server-default: true\n              enum:\n              - INCLUDE_ALL_CREDITS\n              - EXCLUDE_ALL_CREDITS\n              - INCLUDE_SPECIFIED_CREDITS\n            customPeriod:\n              type: object\n              x-dcl-go-name: CustomPeriod\n              x-dcl-go-type: BudgetBudgetFilterCustomPeriod\n              description: Optional. Specifies to track usage from any start date\n                (required) to any end date (optional). This time period is static,\n                it does not recur.\n              x-dcl-conflicts:\n              - calendarPeriod\n              required:\n              - startDate\n              properties:\n                endDate:\n                  type: object\n                  x-dcl-go-name: EndDate\n                  x-dcl-go-type: BudgetBudgetFilterCustomPeriodEndDate\n                  description: Optional. The end date of the time period. Budgets\n                    with elapsed end date won't be processed. If unset, specifies\n                    to track all usage incurred since the start_date.\n                  x-kubernetes-immutable: true\n                  properties:\n                    day:\n                      type: integer\n                      format: int64\n                      x-dcl-go-name: Day\n                      description: Day of a month. Must be from 1 to 31 and valid\n                        for the year and month, or 0 to specify a year by itself or\n                        a year and month where the day isn't significant.\n                      x-kubernetes-immutable: true\n                    month:\n                      type: integer\n                      format: int64\n                      x-dcl-go-name: Month\n                      description: Month of a year. Must be from 1 to 12, or 0 to\n                        specify a year without a month and day.\n                      x-kubernetes-immutable: true\n                    year:\n                      type: integer\n                      format: int64\n                      x-dcl-go-name: Year\n                      description: Year of the date. Must be from 1 to 9999, or 0\n                        to specify a date without a year.\n                      x-kubernetes-immutable: true\n                startDate:\n                  type: object\n                  x-dcl-go-name: StartDate\n                  x-dcl-go-type: BudgetBudgetFilterCustomPeriodStartDate\n                  description: Required. The start date must be after January 1, 2017.\n                  x-kubernetes-immutable: true\n                  properties:\n                    day:\n                      type: integer\n                      format: int64\n                      x-dcl-go-name: Day\n                      description: Day of a month. Must be from 1 to 31 and valid\n                        for the year and month, or 0 to specify a year by itself or\n                        a year and month where the day isn't significant.\n                      x-kubernetes-immutable: true\n                    month:\n                      type: integer\n                      format: int64\n                      x-dcl-go-name: Month\n                      description: Month of a year. Must be from 1 to 12, or 0 to\n                        specify a year without a month and day.\n                      x-kubernetes-immutable: true\n                    year:\n                      type: integer\n                      format: int64\n                      x-dcl-go-name: Year\n                      description: Year of the date. Must be from 1 to 9999, or 0\n                        to specify a date without a year.\n                      x-kubernetes-immutable: true\n            labels:\n              type: object\n              additionalProperties:\n                type: object\n                x-dcl-go-type: BudgetBudgetFilterLabels\n                properties:\n                  values:\n                    type: array\n                    x-dcl-go-name: Values\n                    description: The values of the label\n                    x-kubernetes-immutable: true\n                    x-dcl-send-empty: true\n                    x-dcl-list-type: list\n                    items:\n                      type: string\n                      x-dcl-go-type: string\n              x-dcl-go-name: Labels\n              description: Optional. A single label and value pair specifying that\n                usage from only this set of labeled resources should be included in\n                the budget. Currently, multiple entries or multiple values per entry\n                are not allowed. If omitted, the report will include all labeled and\n                unlabeled usage.\n            projects:\n              type: array\n              x-dcl-go-name: Projects\n              description: Optional. A set of projects of the form `projects/{project}`,\n                specifying that usage from only this set of projects should be included\n                in the budget. If omitted, the report will include all usage for the\n                billing account, regardless of which project the usage occurred on.\n                Only zero or one project can be specified currently.\n              x-dcl-send-empty: true\n              x-dcl-list-type: list\n              items:\n                type: string\n                x-dcl-go-type: string\n                x-dcl-references:\n                - resource: Cloudresourcemanager/Project\n                  field: name\n            services:\n              type: array\n              x-dcl-go-name: Services\n              description: 'Optional. A set of services of the form `services/{service_id}`,\n                specifying that usage from only this set of services should be included\n                in the budget. If omitted, the report will include usage for all the\n                services. The service names are available through the Catalog API:\n                https://cloud.google.com/billing/v1/how-tos/catalog-api.'\n              x-dcl-send-empty: true\n              x-dcl-list-type: list\n              items:\n                type: string\n                x-dcl-go-type: string\n            subaccounts:\n              type: array\n              x-dcl-go-name: Subaccounts\n              description: Optional. A set of subaccounts of the form `billingAccounts/{account_id}`,\n                specifying that usage from only this set of subaccounts should be\n                included in the budget. If a subaccount is set to the name of the\n                parent account, usage from the parent account will be included. If\n                the field is omitted, the report will include usage from the parent\n                account and all subaccounts, if they exist.\n              x-dcl-send-empty: true\n              x-dcl-list-type: list\n              items:\n                type: string\n                x-dcl-go-type: string\n                x-dcl-references:\n                - resource: Cloudbilling/BillingAccount\n                  field: name\n        displayName:\n          type: string\n          x-dcl-go-name: DisplayName\n          description: User data for display name in UI. The name must be less than\n            or equal to 60 characters.\n        etag:\n          type: string\n          x-dcl-go-name: Etag\n          readOnly: true\n          description: Optional. Etag to validate that the object is unchanged for\n            a read-modify-write operation. An empty etag will cause an update to overwrite\n            other changes.\n          x-kubernetes-immutable: true\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: Output only. Resource name of the budget.\n          x-kubernetes-immutable: true\n          x-dcl-server-generated-parameter: true\n        thresholdRules:\n          type: array\n          x-dcl-go-name: ThresholdRules\n          description: Optional. Rules that trigger alerts (notifications of thresholds\n            being crossed) when spend exceeds the specified percentages of the budget.\n          x-dcl-send-empty: true\n          x-dcl-list-type: list\n          items:\n            type: object\n            x-dcl-go-type: BudgetThresholdRules\n            required:\n            - thresholdPercent\n            properties:\n              spendBasis:\n                type: string\n                x-dcl-go-name: SpendBasis\n                x-dcl-go-type: BudgetThresholdRulesSpendBasisEnum\n                description: 'Optional. The type of basis used to determine if spend\n                  has passed the threshold. Behavior defaults to CURRENT_SPEND if\n                  not set. Possible values: BASIS_UNSPECIFIED, CURRENT_SPEND, FORECASTED_SPEND'\n                enum:\n                - BASIS_UNSPECIFIED\n                - CURRENT_SPEND\n                - FORECASTED_SPEND\n              thresholdPercent:\n                type: number\n                format: double\n                x-dcl-go-name: ThresholdPercent\n                description: 'Required. Send an alert when this threshold is exceeded.\n                  This is a 1.0-based percentage, so 0.5 = 50%. Validation: non-negative\n                  number.'\n")

// 18639 bytes
// MD5: d7645fa147be99ec0fce83a9a7c7838c
