// Code generated by generators/resource/main.go; DO NOT EDIT.

package ssmincidents_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSSSMIncidentsResponsePlan_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::SSMIncidents::ResponsePlan", "awscc_ssmincidents_response_plan", "test")

	td.ResourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}
