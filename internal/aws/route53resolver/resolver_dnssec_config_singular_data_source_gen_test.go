// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package route53resolver_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSRoute53ResolverResolverDNSSECConfigDataSource_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::Route53Resolver::ResolverDNSSECConfig", "awscc_route53resolver_resolver_dnssec_config", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config: td.DataSourceWithEmptyResourceConfig(),
			Check: resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttrPair(fmt.Sprintf("data.%s", td.ResourceName), "id", td.ResourceName, "id"),
				resource.TestCheckResourceAttrPair(fmt.Sprintf("data.%s", td.ResourceName), "arn", td.ResourceName, "arn"),
			),
		},
	})
}

func TestAccAWSRoute53ResolverResolverDNSSECConfigDataSource_NonExistent(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::Route53Resolver::ResolverDNSSECConfig", "awscc_route53resolver_resolver_dnssec_config", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config:      td.DataSourceWithNonExistentIDConfig(),
			ExpectError: regexp.MustCompile("Not Found"),
		},
	})
}
