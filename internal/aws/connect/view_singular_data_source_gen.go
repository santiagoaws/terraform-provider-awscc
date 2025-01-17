// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package connect

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_connect_view", viewDataSource)
}

// viewDataSource returns the Terraform awscc_connect_view data source.
// This Terraform data source corresponds to the CloudFormation AWS::Connect::View resource.
func viewDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Actions
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The actions of the view in an array.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "maxLength": 255,
		//	    "minLength": 1,
		//	    "pattern": "^([\\p{L}\\p{N}_.:\\/=+\\-@]+[\\p{L}\\p{Z}\\p{N}_.:\\/=+\\-@]*)$",
		//	    "type": "string"
		//	  },
		//	  "maxItems": 1000,
		//	  "type": "array"
		//	}
		"actions": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "The actions of the view in an array.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The description of the view.",
		//	  "maxLength": 4096,
		//	  "minLength": 0,
		//	  "pattern": "^([\\p{L}\\p{N}_.:\\/=+\\-@,]+[\\p{L}\\p{Z}\\p{N}_.:\\/=+\\-@,]*)$",
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The description of the view.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: InstanceArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the instance.",
		//	  "maxLength": 100,
		//	  "minLength": 1,
		//	  "pattern": "^arn:aws[-a-z0-9]*:connect:[-a-z0-9]*:[0-9]{12}:instance/[-a-zA-Z0-9]*$",
		//	  "type": "string"
		//	}
		"instance_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the instance.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the view.",
		//	  "maxLength": 512,
		//	  "minLength": 1,
		//	  "pattern": "^([\\p{L}\\p{N}_.:\\/=+\\-@]+[\\p{L}\\p{Z}\\p{N}_.:\\/=+\\-@]*)$",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the view.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "One or more tags.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A key-value pair to associate with a resource.",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "pattern": "",
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value for the tag. . You can specify a value that is maximum of 256 Unicode characters",
		//	        "maxLength": 256,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Key",
		//	      "Value"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "maxItems": 50,
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"tags": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value for the tag. . You can specify a value that is maximum of 256 Unicode characters",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "One or more tags.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Template
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The template of the view as JSON.",
		//	  "type": "object"
		//	}
		"template": schema.MapAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "The template of the view as JSON.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ViewArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the view.",
		//	  "maxLength": 255,
		//	  "minLength": 1,
		//	  "pattern": "^arn:aws[-a-z0-9]*:connect:[-a-z0-9]*:[0-9]{12}:instance/[-a-zA-Z0-9]*/view/[-:$a-zA-Z0-9]*$",
		//	  "type": "string"
		//	}
		"view_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the view.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ViewContentSha256
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The view content hash.",
		//	  "pattern": "^[a-zA-Z0-9]{64}$",
		//	  "type": "string"
		//	}
		"view_content_sha_256": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The view content hash.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ViewId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The view id of the view.",
		//	  "maxLength": 500,
		//	  "minLength": 1,
		//	  "pattern": "^[a-zA-Z0-9\\_\\-:\\/$]+$",
		//	  "type": "string"
		//	}
		"view_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The view id of the view.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::Connect::View",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Connect::View").WithTerraformTypeName("awscc_connect_view")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"actions":              "Actions",
		"description":          "Description",
		"instance_arn":         "InstanceArn",
		"key":                  "Key",
		"name":                 "Name",
		"tags":                 "Tags",
		"template":             "Template",
		"value":                "Value",
		"view_arn":             "ViewArn",
		"view_content_sha_256": "ViewContentSha256",
		"view_id":              "ViewId",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
