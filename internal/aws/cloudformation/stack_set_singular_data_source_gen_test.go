// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package cloudformation_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSCloudFormationStackSetDataSource_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::CloudFormation::StackSet", "awscc_cloudformation_stack_set", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyDataSourceConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}

func TestAccAWSCloudFormationStackSetDataSource_NonExistent(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::CloudFormation::StackSet", "awscc_cloudformation_stack_set", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config:      td.DataSourceWithNonExistentIDConfig(),
			ExpectError: regexp.MustCompile("Not Found"),
		},
	})
}
