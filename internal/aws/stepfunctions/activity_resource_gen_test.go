// Code generated by generators/resource/main.go; DO NOT EDIT.

package stepfunctions_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSStepFunctionsActivity_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::StepFunctions::Activity", "awscc_stepfunctions_activity", "test")

	td.ResourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}
