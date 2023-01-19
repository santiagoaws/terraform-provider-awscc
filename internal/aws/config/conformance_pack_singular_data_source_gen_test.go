// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package config_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSConfigConformancePackDataSource_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::Config::ConformancePack", "awscc_config_conformance_pack", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyDataSourceConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}

func TestAccAWSConfigConformancePackDataSource_NonExistent(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::Config::ConformancePack", "awscc_config_conformance_pack", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config:      td.DataSourceWithNonExistentIDConfig(),
			ExpectError: regexp.MustCompile("Not Found"),
		},
	})
}
