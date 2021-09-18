// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package route53recoveryreadiness

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_route53recoveryreadiness_recovery_group", recoveryGroupDataSourceType)
}

// recoveryGroupDataSourceType returns the Terraform awscc_route53recoveryreadiness_recovery_group data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::Route53RecoveryReadiness::RecoveryGroup resource type.
func recoveryGroupDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"cells": {
			// Property: Cells
			// CloudFormation resource type schema:
			// {
			//   "description": "A list of the cell Amazon Resource Names (ARNs) in the recovery group.",
			//   "insertionOrder": false,
			//   "items": {
			//     "maxLength": 256,
			//     "minLength": 1,
			//     "type": "string"
			//   },
			//   "maxItems": 5,
			//   "type": "array"
			// }
			Description: "A list of the cell Amazon Resource Names (ARNs) in the recovery group.",
			Type:        types.ListType{ElemType: types.StringType},
			Computed:    true,
		},
		"recovery_group_arn": {
			// Property: RecoveryGroupArn
			// CloudFormation resource type schema:
			// {
			//   "description": "A collection of tags associated with a resource.",
			//   "maxLength": 256,
			//   "type": "string"
			// }
			Description: "A collection of tags associated with a resource.",
			Type:        types.StringType,
			Computed:    true,
		},
		"recovery_group_name": {
			// Property: RecoveryGroupName
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the recovery group to create.",
			//   "maxLength": 64,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The name of the recovery group to create.",
			Type:        types.StringType,
			Computed:    true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "A collection of tags associated with a resource.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Key": {
			//         "type": "string"
			//       },
			//       "Value": {
			//         "insertionOrder": false,
			//         "items": {
			//           "maxItems": 50,
			//           "type": "string"
			//         },
			//         "type": "array"
			//       }
			//     },
			//     "required": [
			//       "Value",
			//       "Key"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array"
			// }
			Description: "A collection of tags associated with a resource.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Type:     types.StringType,
						Computed: true,
					},
					"value": {
						// Property: Value
						Type:     types.ListType{ElemType: types.StringType},
						Computed: true,
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
		Description: "Data Source schema for AWS::Route53RecoveryReadiness::RecoveryGroup",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Route53RecoveryReadiness::RecoveryGroup").WithTerraformTypeName("awscc_route53recoveryreadiness_recovery_group")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"cells":               "Cells",
		"key":                 "Key",
		"recovery_group_arn":  "RecoveryGroupArn",
		"recovery_group_name": "RecoveryGroupName",
		"tags":                "Tags",
		"value":               "Value",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}