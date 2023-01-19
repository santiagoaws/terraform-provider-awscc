// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package applicationinsights_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSApplicationInsightsApplicationDataSource_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::ApplicationInsights::Application", "awscc_applicationinsights_application", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyDataSourceConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}

func TestAccAWSApplicationInsightsApplicationDataSource_NonExistent(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::ApplicationInsights::Application", "awscc_applicationinsights_application", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config:      td.DataSourceWithNonExistentIDConfig(),
			ExpectError: regexp.MustCompile("Not Found"),
		},
	})
}
