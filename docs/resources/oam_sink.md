---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_oam_sink Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource Type definition for AWS::Oam::Sink
---

# awscc_oam_sink (Resource)

Resource Type definition for AWS::Oam::Sink



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) The name of the ObservabilityAccessManager Sink.

### Optional

- `policy` (Map of String) The policy of this ObservabilityAccessManager Sink.
- `tags` (Map of String) Tags to apply to the sink

### Read-Only

- `arn` (String) The Amazon resource name (ARN) of the ObservabilityAccessManager Sink
- `id` (String) Uniquely identifies the resource.

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_oam_sink.example <resource ID>
```
