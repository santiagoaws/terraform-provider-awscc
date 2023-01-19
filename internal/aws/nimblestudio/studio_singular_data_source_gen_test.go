// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package nimblestudio_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSNimbleStudioStudioDataSource_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::NimbleStudio::Studio", "awscc_nimblestudio_studio", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyDataSourceConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}

func TestAccAWSNimbleStudioStudioDataSource_NonExistent(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::NimbleStudio::Studio", "awscc_nimblestudio_studio", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config:      td.DataSourceWithNonExistentIDConfig(),
			ExpectError: regexp.MustCompile("Not Found"),
		},
	})
}
