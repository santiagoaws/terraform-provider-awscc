// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package iotsitewise_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSIoTSiteWiseAssetModelDataSource_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::IoTSiteWise::AssetModel", "awscc_iotsitewise_asset_model", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyDataSourceConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}

func TestAccAWSIoTSiteWiseAssetModelDataSource_NonExistent(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::IoTSiteWise::AssetModel", "awscc_iotsitewise_asset_model", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config:      td.DataSourceWithNonExistentIDConfig(),
			ExpectError: regexp.MustCompile("Not Found"),
		},
	})
}
