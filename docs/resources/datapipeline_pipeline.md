---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_datapipeline_pipeline Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  An example resource schema demonstrating some basic constructs and validation rules.
---

# awscc_datapipeline_pipeline (Resource)

An example resource schema demonstrating some basic constructs and validation rules.



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) The name of the pipeline.

### Optional

- `activate` (Boolean) Indicates whether to validate and start the pipeline or stop an active pipeline. By default, the value is set to true.
- `description` (String) A description of the pipeline.
- `parameter_objects` (Attributes List) The parameter objects used with the pipeline. (see [below for nested schema](#nestedatt--parameter_objects))
- `parameter_values` (Attributes List) The parameter values used with the pipeline. (see [below for nested schema](#nestedatt--parameter_values))
- `pipeline_objects` (Attributes List) The objects that define the pipeline. These objects overwrite the existing pipeline definition. Not all objects, fields, and values can be updated. For information about restrictions, see Editing Your Pipeline in the AWS Data Pipeline Developer Guide. (see [below for nested schema](#nestedatt--pipeline_objects))
- `pipeline_tags` (Attributes List) A list of arbitrary tags (key-value pairs) to associate with the pipeline, which you can use to control permissions. For more information, see Controlling Access to Pipelines and Resources in the AWS Data Pipeline Developer Guide. (see [below for nested schema](#nestedatt--pipeline_tags))

### Read-Only

- `id` (String) Uniquely identifies the resource.
- `pipeline_id` (String)

<a id="nestedatt--parameter_objects"></a>
### Nested Schema for `parameter_objects`

Required:

- `attributes` (Attributes List) The attributes of the parameter object. (see [below for nested schema](#nestedatt--parameter_objects--attributes))
- `id` (String) The ID of the parameter object.

<a id="nestedatt--parameter_objects--attributes"></a>
### Nested Schema for `parameter_objects.attributes`

Required:

- `key` (String) The field identifier.
- `string_value` (String) The field value, expressed as a String.



<a id="nestedatt--parameter_values"></a>
### Nested Schema for `parameter_values`

Required:

- `id` (String) The ID of the parameter value.
- `string_value` (String) The field value, expressed as a String.


<a id="nestedatt--pipeline_objects"></a>
### Nested Schema for `pipeline_objects`

Required:

- `fields` (Attributes List) Key-value pairs that define the properties of the object. (see [below for nested schema](#nestedatt--pipeline_objects--fields))
- `id` (String) The ID of the object.
- `name` (String) The name of the object.

<a id="nestedatt--pipeline_objects--fields"></a>
### Nested Schema for `pipeline_objects.fields`

Required:

- `key` (String) Specifies the name of a field for a particular object. To view valid values for a particular field, see Pipeline Object Reference in the AWS Data Pipeline Developer Guide.

Optional:

- `ref_value` (String) A field value that you specify as an identifier of another object in the same pipeline definition.
- `string_value` (String) A field value that you specify as a string. To view valid values for a particular field, see Pipeline Object Reference in the AWS Data Pipeline Developer Guide.



<a id="nestedatt--pipeline_tags"></a>
### Nested Schema for `pipeline_tags`

Required:

- `key` (String) The key name of a tag.
- `value` (String) The value to associate with the key name.

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_datapipeline_pipeline.example <resource ID>
```