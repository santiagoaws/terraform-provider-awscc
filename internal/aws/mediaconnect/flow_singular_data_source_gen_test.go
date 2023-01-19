// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package mediaconnect_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSMediaConnectFlowDataSource_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::MediaConnect::Flow", "awscc_mediaconnect_flow", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyDataSourceConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}

func TestAccAWSMediaConnectFlowDataSource_NonExistent(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::MediaConnect::Flow", "awscc_mediaconnect_flow", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config:      td.DataSourceWithNonExistentIDConfig(),
			ExpectError: regexp.MustCompile("Not Found"),
		},
	})
}
