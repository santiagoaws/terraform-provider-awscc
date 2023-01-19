// Code generated by generators/plural-data-source/main.go; DO NOT EDIT.

package cloudfront_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSCloudFrontOriginRequestPoliciesDataSource_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::CloudFront::OriginRequestPolicy", "awscc_cloudfront_origin_request_policies", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config: td.EmptyDataSourceConfig(),
			Check: resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttrSet(fmt.Sprintf("data.%s", td.ResourceName), "ids.#"),
			),
		},
	})
}
