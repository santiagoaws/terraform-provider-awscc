---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_iot_fleet_metric Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  An aggregated metric of certain devices in your fleet
---

# awscc_iot_fleet_metric (Resource)

An aggregated metric of certain devices in your fleet



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **metric_name** (String) The name of the fleet metric

### Optional

- **aggregation_field** (String) The aggregation field to perform aggregation and metric emission
- **aggregation_type** (Attributes) Aggregation types supported by Fleet Indexing (see [below for nested schema](#nestedatt--aggregation_type))
- **description** (String) The description of a fleet metric
- **index_name** (String) The index name of a fleet metric
- **period** (Number) The period of metric emission in seconds
- **query_string** (String) The Fleet Indexing query used by a fleet metric
- **query_version** (String) The version of a Fleet Indexing query used by a fleet metric
- **tags** (Attributes Set, Max: 50) An array of key-value pairs to apply to this resource (see [below for nested schema](#nestedatt--tags))
- **unit** (String) The unit of data points emitted by a fleet metric

### Read-Only

- **creation_date** (Number) The creation date of a fleet metric
- **id** (String) Uniquely identifies the resource.
- **last_modified_date** (Number) The last modified date of a fleet metric
- **metric_arn** (String) The Amazon Resource Number (ARN) of a fleet metric metric
- **version** (Number) The version of a fleet metric

<a id="nestedatt--aggregation_type"></a>
### Nested Schema for `aggregation_type`

Optional:

- **name** (String) Fleet Indexing aggregation type names such as Statistics, Percentiles and Cardinality
- **values** (List of String) Fleet Indexing aggregation type values


<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Optional:

- **key** (String) The tag's key
- **value** (String) The tag's value

