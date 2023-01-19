// Code generated by generators/resource/main.go; DO NOT EDIT.

package ecr_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSECRPublicRepository_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::ECR::PublicRepository", "awscc_ecr_public_repository", "test")

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

func TestAccAWSECRPublicRepository_disappears(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::ECR::PublicRepository", "awscc_ecr_public_repository", "test")

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
