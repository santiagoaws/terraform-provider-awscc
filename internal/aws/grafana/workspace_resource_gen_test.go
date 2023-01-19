// Code generated by generators/resource/main.go; DO NOT EDIT.

package grafana_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSGrafanaWorkspace_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::Grafana::Workspace", "awscc_grafana_workspace", "test")

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

func TestAccAWSGrafanaWorkspace_disappears(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::Grafana::Workspace", "awscc_grafana_workspace", "test")

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
