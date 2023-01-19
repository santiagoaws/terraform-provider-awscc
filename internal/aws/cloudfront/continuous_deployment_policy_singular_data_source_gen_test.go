// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package cloudfront_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSCloudFrontContinuousDeploymentPolicyDataSource_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::CloudFront::ContinuousDeploymentPolicy", "awscc_cloudfront_continuous_deployment_policy", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyDataSourceConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}

func TestAccAWSCloudFrontContinuousDeploymentPolicyDataSource_NonExistent(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::CloudFront::ContinuousDeploymentPolicy", "awscc_cloudfront_continuous_deployment_policy", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config:      td.DataSourceWithNonExistentIDConfig(),
			ExpectError: regexp.MustCompile("Not Found"),
		},
	})
}
