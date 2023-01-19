// Code generated by generators/plural-data-source/main.go; DO NOT EDIT.

package wafv2_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSWAFv2LoggingConfigurationsDataSource_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::WAFv2::LoggingConfiguration", "awscc_wafv2_logging_configurations", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config: td.EmptyDataSourceConfig(),
			Check: resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttrSet(fmt.Sprintf("data.%s", td.ResourceName), "ids.#"),
			),
		},
	})
}
