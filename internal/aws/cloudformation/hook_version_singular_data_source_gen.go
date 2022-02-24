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
	registry.AddDataSourceTypeFactory("awscc_cloudformation_hook_version", hookVersionDataSourceType)
}

// hookVersionDataSourceType returns the Terraform awscc_cloudformation_hook_version data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::CloudFormation::HookVersion resource type.
func hookVersionDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the type, here the HookVersion. This is used to uniquely identify a HookVersion resource",
			//   "pattern": "^arn:aws[A-Za-z0-9-]{0,64}:cloudformation:[A-Za-z0-9-]{1,64}:([0-9]{12})?:type/hook/.+$",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the type, here the HookVersion. This is used to uniquely identify a HookVersion resource",
			Type:        types.StringType,
			Computed:    true,
		},
		"execution_role_arn": {
			// Property: ExecutionRoleArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the IAM execution role to use to register the type. If your resource type calls AWS APIs in any of its handlers, you must create an IAM execution role that includes the necessary permissions to call those AWS APIs, and provision that execution role in your account. CloudFormation then assumes that execution role to provide your resource type with the appropriate credentials.",
			//   "maxLength": 256,
			//   "pattern": "arn:.+:iam::[0-9]{12}:role/.+",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the IAM execution role to use to register the type. If your resource type calls AWS APIs in any of its handlers, you must create an IAM execution role that includes the necessary permissions to call those AWS APIs, and provision that execution role in your account. CloudFormation then assumes that execution role to provide your resource type with the appropriate credentials.",
			Type:        types.StringType,
			Computed:    true,
		},
		"is_default_version": {
			// Property: IsDefaultVersion
			// CloudFormation resource type schema:
			// {
			//   "description": "Indicates if this type version is the current default version",
			//   "type": "boolean"
			// }
			Description: "Indicates if this type version is the current default version",
			Type:        types.BoolType,
			Computed:    true,
		},
		"logging_config": {
			// Property: LoggingConfig
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "Specifies logging configuration information for a type.",
			//   "properties": {
			//     "LogGroupName": {
			//       "description": "The Amazon CloudWatch log group to which CloudFormation sends error logging information when invoking the type's handlers.",
			//       "maxLength": 512,
			//       "minLength": 1,
			//       "pattern": "^[\\.\\-_/#A-Za-z0-9]+$",
			//       "type": "string"
			//     },
			//     "LogRoleArn": {
			//       "description": "The ARN of the role that CloudFormation should assume when sending log entries to CloudWatch logs.",
			//       "maxLength": 256,
			//       "minLength": 1,
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "Specifies logging configuration information for a type.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"log_group_name": {
						// Property: LogGroupName
						Description: "The Amazon CloudWatch log group to which CloudFormation sends error logging information when invoking the type's handlers.",
						Type:        types.StringType,
						Computed:    true,
					},
					"log_role_arn": {
						// Property: LogRoleArn
						Description: "The ARN of the role that CloudFormation should assume when sending log entries to CloudWatch logs.",
						Type:        types.StringType,
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
		"schema_handler_package": {
			// Property: SchemaHandlerPackage
			// CloudFormation resource type schema:
			// {
			//   "description": "A url to the S3 bucket containing the schema handler package that contains the schema, event handlers, and associated files for the type you want to register.\n\nFor information on generating a schema handler package for the type you want to register, see submit in the CloudFormation CLI User Guide.",
			//   "maxLength": 4096,
			//   "type": "string"
			// }
			Description: "A url to the S3 bucket containing the schema handler package that contains the schema, event handlers, and associated files for the type you want to register.\n\nFor information on generating a schema handler package for the type you want to register, see submit in the CloudFormation CLI User Guide.",
			Type:        types.StringType,
			Computed:    true,
		},
		"type_arn": {
			// Property: TypeArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the type without the versionID.",
			//   "pattern": "^arn:aws[A-Za-z0-9-]{0,64}:cloudformation:[A-Za-z0-9-]{1,64}:([0-9]{12})?:type/hook/.+$",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the type without the versionID.",
			Type:        types.StringType,
			Computed:    true,
		},
		"type_name": {
			// Property: TypeName
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the type being registered.\n\nWe recommend that type names adhere to the following pattern: company_or_organization::service::type.",
			//   "pattern": "^[A-Za-z0-9]{2,64}::[A-Za-z0-9]{2,64}::[A-Za-z0-9]{2,64}$",
			//   "type": "string"
			// }
			Description: "The name of the type being registered.\n\nWe recommend that type names adhere to the following pattern: company_or_organization::service::type.",
			Type:        types.StringType,
			Computed:    true,
		},
		"version_id": {
			// Property: VersionId
			// CloudFormation resource type schema:
			// {
			//   "description": "The ID of the version of the type represented by this hook instance.",
			//   "pattern": "^[A-Za-z0-9-]{1,128}$",
			//   "type": "string"
			// }
			Description: "The ID of the version of the type represented by this hook instance.",
			Type:        types.StringType,
			Computed:    true,
		},
		"visibility": {
			// Property: Visibility
			// CloudFormation resource type schema:
			// {
			//   "description": "The scope at which the type is visible and usable in CloudFormation operations.\n\nValid values include:\n\nPRIVATE: The type is only visible and usable within the account in which it is registered. Currently, AWS CloudFormation marks any types you register as PRIVATE.\n\nPUBLIC: The type is publically visible and usable within any Amazon account.",
			//   "enum": [
			//     "PUBLIC",
			//     "PRIVATE"
			//   ],
			//   "type": "string"
			// }
			Description: "The scope at which the type is visible and usable in CloudFormation operations.\n\nValid values include:\n\nPRIVATE: The type is only visible and usable within the account in which it is registered. Currently, AWS CloudFormation marks any types you register as PRIVATE.\n\nPUBLIC: The type is publically visible and usable within any Amazon account.",
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
		Description: "Data Source schema for AWS::CloudFormation::HookVersion",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::CloudFormation::HookVersion").WithTerraformTypeName("awscc_cloudformation_hook_version")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                    "Arn",
		"execution_role_arn":     "ExecutionRoleArn",
		"is_default_version":     "IsDefaultVersion",
		"log_group_name":         "LogGroupName",
		"log_role_arn":           "LogRoleArn",
		"logging_config":         "LoggingConfig",
		"schema_handler_package": "SchemaHandlerPackage",
		"type_arn":               "TypeArn",
		"type_name":              "TypeName",
		"version_id":             "VersionId",
		"visibility":             "Visibility",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
