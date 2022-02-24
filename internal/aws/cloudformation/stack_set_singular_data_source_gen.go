// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package cloudformation

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_cloudformation_stack_set", stackSetDataSourceType)
}

// stackSetDataSourceType returns the Terraform awscc_cloudformation_stack_set data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::CloudFormation::StackSet resource type.
func stackSetDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"administration_role_arn": {
			// Property: AdministrationRoleARN
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Number (ARN) of the IAM role to use to create this stack set. Specify an IAM role only if you are using customized administrator roles to control which users or groups can manage specific stack sets within the same administrator account.",
			//   "maxLength": 2048,
			//   "minLength": 20,
			//   "type": "string"
			// }
			Description: "The Amazon Resource Number (ARN) of the IAM role to use to create this stack set. Specify an IAM role only if you are using customized administrator roles to control which users or groups can manage specific stack sets within the same administrator account.",
			Type:        types.StringType,
			Computed:    true,
		},
		"auto_deployment": {
			// Property: AutoDeployment
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "Describes whether StackSets automatically deploys to AWS Organizations accounts that are added to the target organization or organizational unit (OU). Specify only if PermissionModel is SERVICE_MANAGED.",
			//   "properties": {
			//     "Enabled": {
			//       "description": "If set to true, StackSets automatically deploys additional stack instances to AWS Organizations accounts that are added to a target organization or organizational unit (OU) in the specified Regions. If an account is removed from a target organization or OU, StackSets deletes stack instances from the account in the specified Regions.",
			//       "type": "boolean"
			//     },
			//     "RetainStacksOnAccountRemoval": {
			//       "description": "If set to true, stack resources are retained when an account is removed from a target organization or OU. If set to false, stack resources are deleted. Specify only if Enabled is set to True.",
			//       "type": "boolean"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "Describes whether StackSets automatically deploys to AWS Organizations accounts that are added to the target organization or organizational unit (OU). Specify only if PermissionModel is SERVICE_MANAGED.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"enabled": {
						// Property: Enabled
						Description: "If set to true, StackSets automatically deploys additional stack instances to AWS Organizations accounts that are added to a target organization or organizational unit (OU) in the specified Regions. If an account is removed from a target organization or OU, StackSets deletes stack instances from the account in the specified Regions.",
						Type:        types.BoolType,
						Computed:    true,
					},
					"retain_stacks_on_account_removal": {
						// Property: RetainStacksOnAccountRemoval
						Description: "If set to true, stack resources are retained when an account is removed from a target organization or OU. If set to false, stack resources are deleted. Specify only if Enabled is set to True.",
						Type:        types.BoolType,
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
		"call_as": {
			// Property: CallAs
			// CloudFormation resource type schema:
			// {
			//   "description": "Specifies the AWS account that you are acting from. By default, SELF is specified. For self-managed permissions, specify SELF; for service-managed permissions, if you are signed in to the organization's management account, specify SELF. If you are signed in to a delegated administrator account, specify DELEGATED_ADMIN.",
			//   "enum": [
			//     "SELF",
			//     "DELEGATED_ADMIN"
			//   ],
			//   "type": "string"
			// }
			Description: "Specifies the AWS account that you are acting from. By default, SELF is specified. For self-managed permissions, specify SELF; for service-managed permissions, if you are signed in to the organization's management account, specify SELF. If you are signed in to a delegated administrator account, specify DELEGATED_ADMIN.",
			Type:        types.StringType,
			Computed:    true,
		},
		"capabilities": {
			// Property: Capabilities
			// CloudFormation resource type schema:
			// {
			//   "description": "In some cases, you must explicitly acknowledge that your stack set template contains certain capabilities in order for AWS CloudFormation to create the stack set and related stack instances.",
			//   "insertionOrder": false,
			//   "items": {
			//     "enum": [
			//       "CAPABILITY_IAM",
			//       "CAPABILITY_NAMED_IAM",
			//       "CAPABILITY_AUTO_EXPAND"
			//     ],
			//     "type": "string"
			//   },
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "In some cases, you must explicitly acknowledge that your stack set template contains certain capabilities in order for AWS CloudFormation to create the stack set and related stack instances.",
			Type:        types.SetType{ElemType: types.StringType},
			Computed:    true,
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "description": "A description of the stack set. You can use the description to identify the stack set's purpose or other important information.",
			//   "maxLength": 1024,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "A description of the stack set. You can use the description to identify the stack set's purpose or other important information.",
			Type:        types.StringType,
			Computed:    true,
		},
		"execution_role_name": {
			// Property: ExecutionRoleName
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the IAM execution role to use to create the stack set. If you do not specify an execution role, AWS CloudFormation uses the AWSCloudFormationStackSetExecutionRole role for the stack set operation.",
			//   "maxLength": 64,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "The name of the IAM execution role to use to create the stack set. If you do not specify an execution role, AWS CloudFormation uses the AWSCloudFormationStackSetExecutionRole role for the stack set operation.",
			Type:        types.StringType,
			Computed:    true,
		},
		"managed_execution": {
			// Property: ManagedExecution
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "Describes whether StackSets performs non-conflicting operations concurrently and queues conflicting operations.",
			//   "properties": {
			//     "Active": {
			//       "description": "When true, StackSets performs non-conflicting operations concurrently and queues conflicting operations. After conflicting operations finish, StackSets starts queued operations in request order.",
			//       "type": "boolean"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "Describes whether StackSets performs non-conflicting operations concurrently and queues conflicting operations.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"active": {
						// Property: Active
						Description: "When true, StackSets performs non-conflicting operations concurrently and queues conflicting operations. After conflicting operations finish, StackSets starts queued operations in request order.",
						Type:        types.BoolType,
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
		"operation_preferences": {
			// Property: OperationPreferences
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "The user-specified preferences for how AWS CloudFormation performs a stack set operation.",
			//   "properties": {
			//     "FailureToleranceCount": {
			//       "minimum": 0,
			//       "type": "integer"
			//     },
			//     "FailureTolerancePercentage": {
			//       "maximum": 100,
			//       "minimum": 0,
			//       "type": "integer"
			//     },
			//     "MaxConcurrentCount": {
			//       "minimum": 1,
			//       "type": "integer"
			//     },
			//     "MaxConcurrentPercentage": {
			//       "maximum": 100,
			//       "minimum": 0,
			//       "type": "integer"
			//     },
			//     "RegionConcurrencyType": {
			//       "description": "The concurrency type of deploying StackSets operations in regions, could be in parallel or one region at a time",
			//       "enum": [
			//         "SEQUENTIAL",
			//         "PARALLEL"
			//       ],
			//       "type": "string"
			//     },
			//     "RegionOrder": {
			//       "items": {
			//         "pattern": "^[a-zA-Z0-9-]{1,128}$",
			//         "type": "string"
			//       },
			//       "type": "array"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "The user-specified preferences for how AWS CloudFormation performs a stack set operation.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"failure_tolerance_count": {
						// Property: FailureToleranceCount
						Type:     types.Int64Type,
						Computed: true,
					},
					"failure_tolerance_percentage": {
						// Property: FailureTolerancePercentage
						Type:     types.Int64Type,
						Computed: true,
					},
					"max_concurrent_count": {
						// Property: MaxConcurrentCount
						Type:     types.Int64Type,
						Computed: true,
					},
					"max_concurrent_percentage": {
						// Property: MaxConcurrentPercentage
						Type:     types.Int64Type,
						Computed: true,
					},
					"region_concurrency_type": {
						// Property: RegionConcurrencyType
						Description: "The concurrency type of deploying StackSets operations in regions, could be in parallel or one region at a time",
						Type:        types.StringType,
						Computed:    true,
					},
					"region_order": {
						// Property: RegionOrder
						Type:     types.ListType{ElemType: types.StringType},
						Computed: true,
					},
				},
			),
			Computed: true,
		},
		"parameters": {
			// Property: Parameters
			// CloudFormation resource type schema:
			// {
			//   "description": "The input parameters for the stack set template.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "ParameterKey": {
			//         "description": "The key associated with the parameter. If you don't specify a key and value for a particular parameter, AWS CloudFormation uses the default value that is specified in your template.",
			//         "type": "string"
			//       },
			//       "ParameterValue": {
			//         "description": "The input value associated with the parameter.",
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "ParameterKey",
			//       "ParameterValue"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "The input parameters for the stack set template.",
			Attributes: tfsdk.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"parameter_key": {
						// Property: ParameterKey
						Description: "The key associated with the parameter. If you don't specify a key and value for a particular parameter, AWS CloudFormation uses the default value that is specified in your template.",
						Type:        types.StringType,
						Computed:    true,
					},
					"parameter_value": {
						// Property: ParameterValue
						Description: "The input value associated with the parameter.",
						Type:        types.StringType,
						Computed:    true,
					},
				},
				tfsdk.SetNestedAttributesOptions{},
			),
			Computed: true,
		},
		"permission_model": {
			// Property: PermissionModel
			// CloudFormation resource type schema:
			// {
			//   "description": "Describes how the IAM roles required for stack set operations are created. By default, SELF-MANAGED is specified.",
			//   "enum": [
			//     "SERVICE_MANAGED",
			//     "SELF_MANAGED"
			//   ],
			//   "type": "string"
			// }
			Description: "Describes how the IAM roles required for stack set operations are created. By default, SELF-MANAGED is specified.",
			Type:        types.StringType,
			Computed:    true,
		},
		"stack_instances_group": {
			// Property: StackInstancesGroup
			// CloudFormation resource type schema:
			// {
			//   "description": "A group of stack instances with parameters in some specific accounts and regions.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "Stack instances in some specific accounts and Regions.",
			//     "properties": {
			//       "DeploymentTargets": {
			//         "additionalProperties": false,
			//         "description": " The AWS OrganizationalUnitIds or Accounts for which to create stack instances in the specified Regions.",
			//         "properties": {
			//           "Accounts": {
			//             "description": "AWS accounts that you want to create stack instances in the specified Region(s) for.",
			//             "insertionOrder": false,
			//             "items": {
			//               "description": "AWS account that you want to create stack instances in the specified Region(s) for.",
			//               "pattern": "^[0-9]{12}$",
			//               "type": "string"
			//             },
			//             "minItems": 1,
			//             "type": "array",
			//             "uniqueItems": true
			//           },
			//           "OrganizationalUnitIds": {
			//             "description": "The organization root ID or organizational unit (OU) IDs to which StackSets deploys.",
			//             "insertionOrder": false,
			//             "items": {
			//               "pattern": "^(ou-[a-z0-9]{4,32}-[a-z0-9]{8,32}|r-[a-z0-9]{4,32})$",
			//               "type": "string"
			//             },
			//             "minItems": 1,
			//             "type": "array",
			//             "uniqueItems": true
			//           }
			//         },
			//         "type": "object"
			//       },
			//       "ParameterOverrides": {
			//         "description": "A list of stack set parameters whose values you want to override in the selected stack instances.",
			//         "insertionOrder": false,
			//         "items": {
			//           "additionalProperties": false,
			//           "properties": {
			//             "ParameterKey": {
			//               "description": "The key associated with the parameter. If you don't specify a key and value for a particular parameter, AWS CloudFormation uses the default value that is specified in your template.",
			//               "type": "string"
			//             },
			//             "ParameterValue": {
			//               "description": "The input value associated with the parameter.",
			//               "type": "string"
			//             }
			//           },
			//           "required": [
			//             "ParameterKey",
			//             "ParameterValue"
			//           ],
			//           "type": "object"
			//         },
			//         "type": "array",
			//         "uniqueItems": true
			//       },
			//       "Regions": {
			//         "description": "The names of one or more Regions where you want to create stack instances using the specified AWS account(s).",
			//         "insertionOrder": false,
			//         "items": {
			//           "pattern": "^[a-zA-Z0-9-]{1,128}$",
			//           "type": "string"
			//         },
			//         "minItems": 1,
			//         "type": "array",
			//         "uniqueItems": true
			//       }
			//     },
			//     "required": [
			//       "DeploymentTargets",
			//       "Regions"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "A group of stack instances with parameters in some specific accounts and regions.",
			Attributes: tfsdk.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"deployment_targets": {
						// Property: DeploymentTargets
						Description: " The AWS OrganizationalUnitIds or Accounts for which to create stack instances in the specified Regions.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"accounts": {
									// Property: Accounts
									Description: "AWS accounts that you want to create stack instances in the specified Region(s) for.",
									Type:        types.SetType{ElemType: types.StringType},
									Computed:    true,
								},
								"organizational_unit_ids": {
									// Property: OrganizationalUnitIds
									Description: "The organization root ID or organizational unit (OU) IDs to which StackSets deploys.",
									Type:        types.SetType{ElemType: types.StringType},
									Computed:    true,
								},
							},
						),
						Computed: true,
					},
					"parameter_overrides": {
						// Property: ParameterOverrides
						Description: "A list of stack set parameters whose values you want to override in the selected stack instances.",
						Attributes: tfsdk.SetNestedAttributes(
							map[string]tfsdk.Attribute{
								"parameter_key": {
									// Property: ParameterKey
									Description: "The key associated with the parameter. If you don't specify a key and value for a particular parameter, AWS CloudFormation uses the default value that is specified in your template.",
									Type:        types.StringType,
									Computed:    true,
								},
								"parameter_value": {
									// Property: ParameterValue
									Description: "The input value associated with the parameter.",
									Type:        types.StringType,
									Computed:    true,
								},
							},
							tfsdk.SetNestedAttributesOptions{},
						),
						Computed: true,
					},
					"regions": {
						// Property: Regions
						Description: "The names of one or more Regions where you want to create stack instances using the specified AWS account(s).",
						Type:        types.SetType{ElemType: types.StringType},
						Computed:    true,
					},
				},
				tfsdk.SetNestedAttributesOptions{},
			),
			Computed: true,
		},
		"stack_set_id": {
			// Property: StackSetId
			// CloudFormation resource type schema:
			// {
			//   "description": "The ID of the stack set that you're creating.",
			//   "type": "string"
			// }
			Description: "The ID of the stack set that you're creating.",
			Type:        types.StringType,
			Computed:    true,
		},
		"stack_set_name": {
			// Property: StackSetName
			// CloudFormation resource type schema:
			// {
			//   "description": "The name to associate with the stack set. The name must be unique in the Region where you create your stack set.",
			//   "maxLength": 128,
			//   "pattern": "^[a-zA-Z][a-zA-Z0-9\\-]{0,127}$",
			//   "type": "string"
			// }
			Description: "The name to associate with the stack set. The name must be unique in the Region where you create your stack set.",
			Type:        types.StringType,
			Computed:    true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "The key-value pairs to associate with this stack set and the stacks created from it. AWS CloudFormation also propagates these tags to supported resources that are created in the stacks. A maximum number of 50 tags can be specified.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "Tag type enables you to specify a key-value pair that can be used to store information about an AWS CloudFormation StackSet.",
			//     "properties": {
			//       "Key": {
			//         "description": "A string used to identify this tag. You can specify a maximum of 127 characters for a tag key.",
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "pattern": "",
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "A string containing the value for this tag. You can specify a maximum of 256 characters for a tag value.",
			//         "maxLength": 256,
			//         "minLength": 1,
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Key",
			//       "Value"
			//     ],
			//     "type": "object"
			//   },
			//   "maxItems": 50,
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "The key-value pairs to associate with this stack set and the stacks created from it. AWS CloudFormation also propagates these tags to supported resources that are created in the stacks. A maximum number of 50 tags can be specified.",
			Attributes: tfsdk.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "A string used to identify this tag. You can specify a maximum of 127 characters for a tag key.",
						Type:        types.StringType,
						Computed:    true,
					},
					"value": {
						// Property: Value
						Description: "A string containing the value for this tag. You can specify a maximum of 256 characters for a tag value.",
						Type:        types.StringType,
						Computed:    true,
					},
				},
				tfsdk.SetNestedAttributesOptions{},
			),
			Computed: true,
		},
		"template_body": {
			// Property: TemplateBody
			// CloudFormation resource type schema:
			// {
			//   "description": "The structure that contains the template body, with a minimum length of 1 byte and a maximum length of 51,200 bytes.",
			//   "maxLength": 51200,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "The structure that contains the template body, with a minimum length of 1 byte and a maximum length of 51,200 bytes.",
			Type:        types.StringType,
			Computed:    true,
		},
		"template_url": {
			// Property: TemplateURL
			// CloudFormation resource type schema:
			// {
			//   "description": "Location of file containing the template body. The URL must point to a template (max size: 460,800 bytes) that is located in an Amazon S3 bucket.",
			//   "maxLength": 1024,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "Location of file containing the template body. The URL must point to a template (max size: 460,800 bytes) that is located in an Amazon S3 bucket.",
			Type:        types.StringType,
			Computed:    true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::CloudFormation::StackSet",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::CloudFormation::StackSet").WithTerraformTypeName("awscc_cloudformation_stack_set")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"accounts":                         "Accounts",
		"active":                           "Active",
		"administration_role_arn":          "AdministrationRoleARN",
		"auto_deployment":                  "AutoDeployment",
		"call_as":                          "CallAs",
		"capabilities":                     "Capabilities",
		"deployment_targets":               "DeploymentTargets",
		"description":                      "Description",
		"enabled":                          "Enabled",
		"execution_role_name":              "ExecutionRoleName",
		"failure_tolerance_count":          "FailureToleranceCount",
		"failure_tolerance_percentage":     "FailureTolerancePercentage",
		"key":                              "Key",
		"managed_execution":                "ManagedExecution",
		"max_concurrent_count":             "MaxConcurrentCount",
		"max_concurrent_percentage":        "MaxConcurrentPercentage",
		"operation_preferences":            "OperationPreferences",
		"organizational_unit_ids":          "OrganizationalUnitIds",
		"parameter_key":                    "ParameterKey",
		"parameter_overrides":              "ParameterOverrides",
		"parameter_value":                  "ParameterValue",
		"parameters":                       "Parameters",
		"permission_model":                 "PermissionModel",
		"region_concurrency_type":          "RegionConcurrencyType",
		"region_order":                     "RegionOrder",
		"regions":                          "Regions",
		"retain_stacks_on_account_removal": "RetainStacksOnAccountRemoval",
		"stack_instances_group":            "StackInstancesGroup",
		"stack_set_id":                     "StackSetId",
		"stack_set_name":                   "StackSetName",
		"tags":                             "Tags",
		"template_body":                    "TemplateBody",
		"template_url":                     "TemplateURL",
		"value":                            "Value",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
