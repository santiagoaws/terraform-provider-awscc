// Code generated by generators/resource/main.go; DO NOT EDIT.

package lookoutmetrics_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSLookoutMetricsAlert_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::LookoutMetrics::Alert", "awscc_lookoutmetrics_alert", "test")

	td.ResourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}
