// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package ec2_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSEC2TransitGatewayMulticastGroupSourceDataSource_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::EC2::TransitGatewayMulticastGroupSource", "awscc_ec2_transit_gateway_multicast_group_source", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyDataSourceConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}

func TestAccAWSEC2TransitGatewayMulticastGroupSourceDataSource_NonExistent(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::EC2::TransitGatewayMulticastGroupSource", "awscc_ec2_transit_gateway_multicast_group_source", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config:      td.DataSourceWithNonExistentIDConfig(),
			ExpectError: regexp.MustCompile("Not Found"),
		},
	})
}
