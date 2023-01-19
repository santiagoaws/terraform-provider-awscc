// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package macie_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSMacieCustomDataIdentifierDataSource_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::Macie::CustomDataIdentifier", "awscc_macie_custom_data_identifier", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyDataSourceConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}

func TestAccAWSMacieCustomDataIdentifierDataSource_NonExistent(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::Macie::CustomDataIdentifier", "awscc_macie_custom_data_identifier", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config:      td.DataSourceWithNonExistentIDConfig(),
			ExpectError: regexp.MustCompile("Not Found"),
		},
	})
}
