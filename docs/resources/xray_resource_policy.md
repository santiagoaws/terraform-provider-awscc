---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_xray_resource_policy Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  This schema provides construct and validation rules for AWS-XRay Resource Policy resource parameters.
---

# awscc_xray_resource_policy (Resource)

This schema provides construct and validation rules for AWS-XRay Resource Policy resource parameters.



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `policy_document` (String) The resource policy document, which can be up to 5kb in size.
- `policy_name` (String) The name of the resource policy. Must be unique within a specific AWS account.

### Optional

- `bypass_policy_lockout_check` (Boolean) A flag to indicate whether to bypass the resource policy lockout safety check

### Read-Only

- `id` (String) Uniquely identifies the resource.

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_xray_resource_policy.example <resource ID>
```