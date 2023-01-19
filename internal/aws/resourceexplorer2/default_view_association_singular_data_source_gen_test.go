// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package resourceexplorer2_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSResourceExplorer2DefaultViewAssociationDataSource_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::ResourceExplorer2::DefaultViewAssociation", "awscc_resourceexplorer2_default_view_association", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyDataSourceConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}

func TestAccAWSResourceExplorer2DefaultViewAssociationDataSource_NonExistent(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::ResourceExplorer2::DefaultViewAssociation", "awscc_resourceexplorer2_default_view_association", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config:      td.DataSourceWithNonExistentIDConfig(),
			ExpectError: regexp.MustCompile("Not Found"),
		},
	})
}
