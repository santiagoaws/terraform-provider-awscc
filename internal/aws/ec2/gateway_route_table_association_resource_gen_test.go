// Code generated by generators/resource/main.go; DO NOT EDIT.

package ec2_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSEC2GatewayRouteTableAssociation_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::EC2::GatewayRouteTableAssociation", "awscc_ec2_gateway_route_table_association", "test")

	td.ResourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}
