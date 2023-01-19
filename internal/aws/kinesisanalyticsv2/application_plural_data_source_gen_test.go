// Code generated by generators/plural-data-source/main.go; DO NOT EDIT.

package kinesisanalyticsv2_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSKinesisAnalyticsV2ApplicationsDataSource_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::KinesisAnalyticsV2::Application", "awscc_kinesisanalyticsv2_applications", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config: td.EmptyDataSourceConfig(),
			Check: resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttrSet(fmt.Sprintf("data.%s", td.ResourceName), "ids.#"),
			),
		},
	})
}
