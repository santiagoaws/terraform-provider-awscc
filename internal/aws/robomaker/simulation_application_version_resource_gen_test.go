// Code generated by generators/resource/main.go; DO NOT EDIT.

package robomaker_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSRoboMakerSimulationApplicationVersion_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::RoboMaker::SimulationApplicationVersion", "awscc_robomaker_simulation_application_version", "test")

	td.ResourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}
