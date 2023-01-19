// Code generated by generators/resource/main.go; DO NOT EDIT.

package mediaconnect_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSMediaConnectFlow_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::MediaConnect::Flow", "awscc_mediaconnect_flow", "test")

	td.ResourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}
