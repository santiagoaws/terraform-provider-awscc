// Code generated by generators/plural-data-source/main.go; DO NOT EDIT.

package kms_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSKMSReplicaKeysDataSource_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::KMS::ReplicaKey", "awscc_kms_replica_keys", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config: td.EmptyDataSourceConfig(),
			Check: resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttrSet(fmt.Sprintf("data.%s", td.ResourceName), "ids.#"),
			),
		},
	})
}
