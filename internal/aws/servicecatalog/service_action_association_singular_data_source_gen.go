// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package servicecatalog

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_servicecatalog_service_action_association", serviceActionAssociationDataSourceType)
}

// serviceActionAssociationDataSourceType returns the Terraform awscc_servicecatalog_service_action_association data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::ServiceCatalog::ServiceActionAssociation resource type.
func serviceActionAssociationDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"product_id": {
			// Property: ProductId
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 100,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"provisioning_artifact_id": {
			// Property: ProvisioningArtifactId
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 100,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"service_action_id": {
			// Property: ServiceActionId
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 100,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::ServiceCatalog::ServiceActionAssociation",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::ServiceCatalog::ServiceActionAssociation").WithTerraformTypeName("awscc_servicecatalog_service_action_association")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"product_id":               "ProductId",
		"provisioning_artifact_id": "ProvisioningArtifactId",
		"service_action_id":        "ServiceActionId",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}