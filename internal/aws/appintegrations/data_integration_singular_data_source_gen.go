// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package appintegrations

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_appintegrations_data_integration", dataIntegrationDataSourceType)
}

// dataIntegrationDataSourceType returns the Terraform awscc_appintegrations_data_integration data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::AppIntegrations::DataIntegration resource type.
func dataIntegrationDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"data_integration_arn": {
			// Property: DataIntegrationArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the data integration.",
			//   "maxLength": 2048,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the data integration.",
			Type:        types.StringType,
			Computed:    true,
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "description": "The data integration description.",
			//   "maxLength": 1000,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "The data integration description.",
			Type:        types.StringType,
			Computed:    true,
		},
		"id": {
			// Property: Id
			// CloudFormation resource type schema:
			// {
			//   "description": "The unique identifer of the data integration.",
			//   "maxLength": 255,
			//   "minLength": 1,
			//   "pattern": "[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}",
			//   "type": "string"
			// }
			Description: "The unique identifer of the data integration.",
			Type:        types.StringType,
			Computed:    true,
		},
		"kms_key": {
			// Property: KmsKey
			// CloudFormation resource type schema:
			// {
			//   "description": "The KMS key of the data integration.",
			//   "maxLength": 255,
			//   "minLength": 1,
			//   "pattern": ".*\\S.*",
			//   "type": "string"
			// }
			Description: "The KMS key of the data integration.",
			Type:        types.StringType,
			Computed:    true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the data integration.",
			//   "maxLength": 255,
			//   "minLength": 1,
			//   "pattern": "^[a-zA-Z0-9/\\._\\-]+$",
			//   "type": "string"
			// }
			Description: "The name of the data integration.",
			Type:        types.StringType,
			Computed:    true,
		},
		"schedule_config": {
			// Property: ScheduleConfig
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "The name of the data and how often it should be pulled from the source.",
			//   "properties": {
			//     "FirstExecutionFrom": {
			//       "description": "The start date for objects to import in the first flow run. Epoch or ISO timestamp format is supported.",
			//       "maxLength": 255,
			//       "minLength": 1,
			//       "pattern": ".*\\S.*",
			//       "type": "string"
			//     },
			//     "Object": {
			//       "description": "The name of the object to pull from the data source.",
			//       "maxLength": 255,
			//       "minLength": 1,
			//       "pattern": "^[a-zA-Z0-9/\\._\\-]+$",
			//       "type": "string"
			//     },
			//     "ScheduleExpression": {
			//       "description": "How often the data should be pulled from data source.",
			//       "maxLength": 255,
			//       "minLength": 1,
			//       "pattern": ".*\\S.*",
			//       "type": "string"
			//     }
			//   },
			//   "required": [
			//     "FirstExecutionFrom",
			//     "Object",
			//     "ScheduleExpression"
			//   ],
			//   "type": "object"
			// }
			Description: "The name of the data and how often it should be pulled from the source.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"first_execution_from": {
						// Property: FirstExecutionFrom
						Description: "The start date for objects to import in the first flow run. Epoch or ISO timestamp format is supported.",
						Type:        types.StringType,
						Computed:    true,
					},
					"object": {
						// Property: Object
						Description: "The name of the object to pull from the data source.",
						Type:        types.StringType,
						Computed:    true,
					},
					"schedule_expression": {
						// Property: ScheduleExpression
						Description: "How often the data should be pulled from data source.",
						Type:        types.StringType,
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
		"source_uri": {
			// Property: SourceURI
			// CloudFormation resource type schema:
			// {
			//   "description": "The URI of the data source.",
			//   "maxLength": 255,
			//   "minLength": 1,
			//   "pattern": "^\\w+\\:\\/\\/\\w+\\/[\\w/!@#+=.-]+$",
			//   "type": "string"
			// }
			Description: "The URI of the data source.",
			Type:        types.StringType,
			Computed:    true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "The tags (keys and values) associated with the data integration.",
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "A label for tagging DataIntegration resources",
			//     "properties": {
			//       "Key": {
			//         "description": "A key to identify the tag.",
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "pattern": "",
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "Corresponding tag value for the key.",
			//         "maxLength": 256,
			//         "minLength": 0,
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Key",
			//       "Value"
			//     ],
			//     "type": "object"
			//   },
			//   "maxItems": 200,
			//   "minItems": 0,
			//   "type": "array"
			// }
			Description: "The tags (keys and values) associated with the data integration.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "A key to identify the tag.",
						Type:        types.StringType,
						Computed:    true,
					},
					"value": {
						// Property: Value
						Description: "Corresponding tag value for the key.",
						Type:        types.StringType,
						Computed:    true,
					},
				},
				tfsdk.ListNestedAttributesOptions{},
			),
			Computed: true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::AppIntegrations::DataIntegration",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::AppIntegrations::DataIntegration").WithTerraformTypeName("awscc_appintegrations_data_integration")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"data_integration_arn": "DataIntegrationArn",
		"description":          "Description",
		"first_execution_from": "FirstExecutionFrom",
		"id":                   "Id",
		"key":                  "Key",
		"kms_key":              "KmsKey",
		"name":                 "Name",
		"object":               "Object",
		"schedule_config":      "ScheduleConfig",
		"schedule_expression":  "ScheduleExpression",
		"source_uri":           "SourceURI",
		"tags":                 "Tags",
		"value":                "Value",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
