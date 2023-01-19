// Code generated by generators/resource/main.go; DO NOT EDIT.

package lightsail_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSLightsailLoadBalancer_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::Lightsail::LoadBalancer", "awscc_lightsail_load_balancer", "test")

	td.ResourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}
