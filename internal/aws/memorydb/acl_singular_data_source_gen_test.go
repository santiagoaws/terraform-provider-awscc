// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package memorydb_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSMemoryDBACLDataSource_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::MemoryDB::ACL", "awscc_memorydb_acl", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyDataSourceConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}

func TestAccAWSMemoryDBACLDataSource_NonExistent(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::MemoryDB::ACL", "awscc_memorydb_acl", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config:      td.DataSourceWithNonExistentIDConfig(),
			ExpectError: regexp.MustCompile("Not Found"),
		},
	})
}
