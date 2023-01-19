// Code generated by generators/resource/main.go; DO NOT EDIT.

package refactorspaces_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSRefactorSpacesApplication_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::RefactorSpaces::Application", "awscc_refactorspaces_application", "test")

	td.ResourceTest(t, []resource.TestStep{
		{
			Config: td.EmptyConfig(),
			Check: resource.ComposeTestCheckFunc(
				td.CheckExistsInAWS(),
			),
		},
		{
			ResourceName:      td.ResourceName,
			ImportState:       true,
			ImportStateVerify: true,
		},
	})
}

func TestAccAWSRefactorSpacesApplication_disappears(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::RefactorSpaces::Application", "awscc_refactorspaces_application", "test")

	td.ResourceTest(t, []resource.TestStep{
		{
			Config: td.EmptyConfig(),
			Check: resource.ComposeTestCheckFunc(
				td.CheckExistsInAWS(),
				td.DeleteResource(),
			),
			ExpectNonEmptyPlan: true,
		},
	})
}
