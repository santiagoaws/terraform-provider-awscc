// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package sso

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_sso_assignment", assignmentDataSourceType)
}

// assignmentDataSourceType returns the Terraform awscc_sso_assignment data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::SSO::Assignment resource type.
func assignmentDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"instance_arn": {
			// Property: InstanceArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The sso instance that the permission set is owned.",
			//   "maxLength": 1224,
			//   "minLength": 10,
			//   "pattern": "arn:(aws|aws-us-gov|aws-cn|aws-iso|aws-iso-b):sso:::instance/(sso)?ins-[a-zA-Z0-9-.]{16}",
			//   "type": "string"
			// }
			Description: "The sso instance that the permission set is owned.",
			Type:        types.StringType,
			Computed:    true,
		},
		"permission_set_arn": {
			// Property: PermissionSetArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The permission set that the assignemt will be assigned",
			//   "maxLength": 1224,
			//   "minLength": 10,
			//   "pattern": "arn:(aws|aws-us-gov|aws-cn|aws-iso|aws-iso-b):sso:::permissionSet/(sso)?ins-[a-zA-Z0-9-.]{16}/ps-[a-zA-Z0-9-./]{16}",
			//   "type": "string"
			// }
			Description: "The permission set that the assignemt will be assigned",
			Type:        types.StringType,
			Computed:    true,
		},
		"principal_id": {
			// Property: PrincipalId
			// CloudFormation resource type schema:
			// {
			//   "description": "The assignee's identifier, user id/group id",
			//   "maxLength": 47,
			//   "minLength": 1,
			//   "pattern": "^([0-9a-f]{10}-|)[A-Fa-f0-9]{8}-[A-Fa-f0-9]{4}-[A-Fa-f0-9]{4}-[A-Fa-f0-9]{4}-[A-Fa-f0-9]{12}$",
			//   "type": "string"
			// }
			Description: "The assignee's identifier, user id/group id",
			Type:        types.StringType,
			Computed:    true,
		},
		"principal_type": {
			// Property: PrincipalType
			// CloudFormation resource type schema:
			// {
			//   "description": "The assignee's type, user/group",
			//   "enum": [
			//     "USER",
			//     "GROUP"
			//   ],
			//   "type": "string"
			// }
			Description: "The assignee's type, user/group",
			Type:        types.StringType,
			Computed:    true,
		},
		"target_id": {
			// Property: TargetId
			// CloudFormation resource type schema:
			// {
			//   "description": "The account id to be provisioned.",
			//   "pattern": "\\d{12}",
			//   "type": "string"
			// }
			Description: "The account id to be provisioned.",
			Type:        types.StringType,
			Computed:    true,
		},
		"target_type": {
			// Property: TargetType
			// CloudFormation resource type schema:
			// {
			//   "description": "The type of resource to be provsioned to, only aws account now",
			//   "enum": [
			//     "AWS_ACCOUNT"
			//   ],
			//   "type": "string"
			// }
			Description: "The type of resource to be provsioned to, only aws account now",
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
		Description: "Data Source schema for AWS::SSO::Assignment",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::SSO::Assignment").WithTerraformTypeName("awscc_sso_assignment")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"instance_arn":       "InstanceArn",
		"permission_set_arn": "PermissionSetArn",
		"principal_id":       "PrincipalId",
		"principal_type":     "PrincipalType",
		"target_id":          "TargetId",
		"target_type":        "TargetType",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
