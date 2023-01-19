// Code generated by generators/plural-data-source/main.go; DO NOT EDIT.

package iotsitewise_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSIoTSiteWiseAssetModelsDataSource_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::IoTSiteWise::AssetModel", "awscc_iotsitewise_asset_models", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config: td.EmptyDataSourceConfig(),
			Check: resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttrSet(fmt.Sprintf("data.%s", td.ResourceName), "ids.#"),
			),
		},
	})
}
