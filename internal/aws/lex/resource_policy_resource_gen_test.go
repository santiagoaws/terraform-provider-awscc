// Code generated by generators/resource/main.go; DO NOT EDIT.

package lex_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSLexResourcePolicy_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::Lex::ResourcePolicy", "awscc_lex_resource_policy", "test")

	td.ResourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}
