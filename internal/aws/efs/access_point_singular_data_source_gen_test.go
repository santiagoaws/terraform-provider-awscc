// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package efs_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSEFSAccessPointDataSource_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::EFS::AccessPoint", "awscc_efs_access_point", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyDataSourceConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}

func TestAccAWSEFSAccessPointDataSource_NonExistent(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::EFS::AccessPoint", "awscc_efs_access_point", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config:      td.DataSourceWithNonExistentIDConfig(),
			ExpectError: regexp.MustCompile("Not Found"),
		},
	})
}
