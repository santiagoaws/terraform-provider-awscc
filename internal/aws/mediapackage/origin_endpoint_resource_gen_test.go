// Code generated by generators/resource/main.go; DO NOT EDIT.

package mediapackage_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSMediaPackageOriginEndpoint_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::MediaPackage::OriginEndpoint", "awscc_mediapackage_origin_endpoint", "test")

	td.ResourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}
