// Code generated by generators/resource/main.go; DO NOT EDIT.

package scheduler_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSSchedulerScheduleGroup_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::Scheduler::ScheduleGroup", "awscc_scheduler_schedule_group", "test")

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

func TestAccAWSSchedulerScheduleGroup_disappears(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::Scheduler::ScheduleGroup", "awscc_scheduler_schedule_group", "test")

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
