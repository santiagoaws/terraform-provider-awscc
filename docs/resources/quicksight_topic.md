---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_quicksight_topic Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Definition of the AWS::QuickSight::Topic Resource Type.
---

# awscc_quicksight_topic (Resource)

Definition of the AWS::QuickSight::Topic Resource Type.



<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `aws_account_id` (String)
- `data_sets` (Attributes List) (see [below for nested schema](#nestedatt--data_sets))
- `description` (String)
- `name` (String)
- `topic_id` (String)

### Read-Only

- `arn` (String)
- `id` (String) Uniquely identifies the resource.

<a id="nestedatt--data_sets"></a>
### Nested Schema for `data_sets`

Required:

- `dataset_arn` (String)

Optional:

- `calculated_fields` (Attributes List) (see [below for nested schema](#nestedatt--data_sets--calculated_fields))
- `columns` (Attributes List) (see [below for nested schema](#nestedatt--data_sets--columns))
- `data_aggregation` (Attributes) (see [below for nested schema](#nestedatt--data_sets--data_aggregation))
- `dataset_description` (String)
- `dataset_name` (String)
- `filters` (Attributes List) (see [below for nested schema](#nestedatt--data_sets--filters))
- `named_entities` (Attributes List) (see [below for nested schema](#nestedatt--data_sets--named_entities))

<a id="nestedatt--data_sets--calculated_fields"></a>
### Nested Schema for `data_sets.calculated_fields`

Required:

- `calculated_field_name` (String)
- `expression` (String)

Optional:

- `aggregation` (String)
- `allowed_aggregations` (List of String)
- `calculated_field_description` (String)
- `calculated_field_synonyms` (List of String)
- `cell_value_synonyms` (Attributes List) (see [below for nested schema](#nestedatt--data_sets--calculated_fields--cell_value_synonyms))
- `column_data_role` (String)
- `comparative_order` (Attributes) (see [below for nested schema](#nestedatt--data_sets--calculated_fields--comparative_order))
- `default_formatting` (Attributes) (see [below for nested schema](#nestedatt--data_sets--calculated_fields--default_formatting))
- `is_included_in_topic` (Boolean)
- `never_aggregate_in_filter` (Boolean)
- `not_allowed_aggregations` (List of String)
- `semantic_type` (Attributes) (see [below for nested schema](#nestedatt--data_sets--calculated_fields--semantic_type))
- `time_granularity` (String)

<a id="nestedatt--data_sets--calculated_fields--cell_value_synonyms"></a>
### Nested Schema for `data_sets.calculated_fields.cell_value_synonyms`

Optional:

- `cell_value` (String)
- `synonyms` (List of String)


<a id="nestedatt--data_sets--calculated_fields--comparative_order"></a>
### Nested Schema for `data_sets.calculated_fields.comparative_order`

Optional:

- `specifed_order` (List of String)
- `treat_undefined_specified_values` (String)
- `use_ordering` (String)


<a id="nestedatt--data_sets--calculated_fields--default_formatting"></a>
### Nested Schema for `data_sets.calculated_fields.default_formatting`

Optional:

- `display_format` (String)
- `display_format_options` (Attributes) (see [below for nested schema](#nestedatt--data_sets--calculated_fields--default_formatting--display_format_options))

<a id="nestedatt--data_sets--calculated_fields--default_formatting--display_format_options"></a>
### Nested Schema for `data_sets.calculated_fields.default_formatting.display_format_options`

Optional:

- `blank_cell_format` (String)
- `currency_symbol` (String)
- `date_format` (String)
- `decimal_separator` (String)
- `fraction_digits` (Number)
- `grouping_separator` (String)
- `negative_format` (Attributes) (see [below for nested schema](#nestedatt--data_sets--calculated_fields--default_formatting--display_format_options--negative_format))
- `prefix` (String)
- `suffix` (String)
- `unit_scaler` (String)
- `use_blank_cell_format` (Boolean)
- `use_grouping` (Boolean)

<a id="nestedatt--data_sets--calculated_fields--default_formatting--display_format_options--negative_format"></a>
### Nested Schema for `data_sets.calculated_fields.default_formatting.display_format_options.negative_format`

Optional:

- `prefix` (String)
- `suffix` (String)




<a id="nestedatt--data_sets--calculated_fields--semantic_type"></a>
### Nested Schema for `data_sets.calculated_fields.semantic_type`

Optional:

- `falsey_cell_value` (String)
- `falsey_cell_value_synonyms` (List of String)
- `sub_type_name` (String)
- `truthy_cell_value` (String)
- `truthy_cell_value_synonyms` (List of String)
- `type_name` (String)
- `type_parameters` (Map of String)



<a id="nestedatt--data_sets--columns"></a>
### Nested Schema for `data_sets.columns`

Required:

- `column_name` (String)

Optional:

- `aggregation` (String)
- `allowed_aggregations` (List of String)
- `cell_value_synonyms` (Attributes List) (see [below for nested schema](#nestedatt--data_sets--columns--cell_value_synonyms))
- `column_data_role` (String)
- `column_description` (String)
- `column_friendly_name` (String)
- `column_synonyms` (List of String)
- `comparative_order` (Attributes) (see [below for nested schema](#nestedatt--data_sets--columns--comparative_order))
- `default_formatting` (Attributes) (see [below for nested schema](#nestedatt--data_sets--columns--default_formatting))
- `is_included_in_topic` (Boolean)
- `never_aggregate_in_filter` (Boolean)
- `not_allowed_aggregations` (List of String)
- `semantic_type` (Attributes) (see [below for nested schema](#nestedatt--data_sets--columns--semantic_type))
- `time_granularity` (String)

<a id="nestedatt--data_sets--columns--cell_value_synonyms"></a>
### Nested Schema for `data_sets.columns.cell_value_synonyms`

Optional:

- `cell_value` (String)
- `synonyms` (List of String)


<a id="nestedatt--data_sets--columns--comparative_order"></a>
### Nested Schema for `data_sets.columns.comparative_order`

Optional:

- `specifed_order` (List of String)
- `treat_undefined_specified_values` (String)
- `use_ordering` (String)


<a id="nestedatt--data_sets--columns--default_formatting"></a>
### Nested Schema for `data_sets.columns.default_formatting`

Optional:

- `display_format` (String)
- `display_format_options` (Attributes) (see [below for nested schema](#nestedatt--data_sets--columns--default_formatting--display_format_options))

<a id="nestedatt--data_sets--columns--default_formatting--display_format_options"></a>
### Nested Schema for `data_sets.columns.default_formatting.display_format_options`

Optional:

- `blank_cell_format` (String)
- `currency_symbol` (String)
- `date_format` (String)
- `decimal_separator` (String)
- `fraction_digits` (Number)
- `grouping_separator` (String)
- `negative_format` (Attributes) (see [below for nested schema](#nestedatt--data_sets--columns--default_formatting--display_format_options--negative_format))
- `prefix` (String)
- `suffix` (String)
- `unit_scaler` (String)
- `use_blank_cell_format` (Boolean)
- `use_grouping` (Boolean)

<a id="nestedatt--data_sets--columns--default_formatting--display_format_options--negative_format"></a>
### Nested Schema for `data_sets.columns.default_formatting.display_format_options.negative_format`

Optional:

- `prefix` (String)
- `suffix` (String)




<a id="nestedatt--data_sets--columns--semantic_type"></a>
### Nested Schema for `data_sets.columns.semantic_type`

Optional:

- `falsey_cell_value` (String)
- `falsey_cell_value_synonyms` (List of String)
- `sub_type_name` (String)
- `truthy_cell_value` (String)
- `truthy_cell_value_synonyms` (List of String)
- `type_name` (String)
- `type_parameters` (Map of String)



<a id="nestedatt--data_sets--data_aggregation"></a>
### Nested Schema for `data_sets.data_aggregation`

Optional:

- `dataset_row_date_granularity` (String)
- `default_date_column_name` (String)


<a id="nestedatt--data_sets--filters"></a>
### Nested Schema for `data_sets.filters`

Required:

- `filter_name` (String)
- `operand_field_name` (String)

Optional:

- `category_filter` (Attributes) (see [below for nested schema](#nestedatt--data_sets--filters--category_filter))
- `date_range_filter` (Attributes) (see [below for nested schema](#nestedatt--data_sets--filters--date_range_filter))
- `filter_class` (String)
- `filter_description` (String)
- `filter_synonyms` (List of String)
- `filter_type` (String)
- `numeric_equality_filter` (Attributes) (see [below for nested schema](#nestedatt--data_sets--filters--numeric_equality_filter))
- `numeric_range_filter` (Attributes) (see [below for nested schema](#nestedatt--data_sets--filters--numeric_range_filter))
- `relative_date_filter` (Attributes) (see [below for nested schema](#nestedatt--data_sets--filters--relative_date_filter))

<a id="nestedatt--data_sets--filters--category_filter"></a>
### Nested Schema for `data_sets.filters.category_filter`

Optional:

- `category_filter_function` (String)
- `category_filter_type` (String)
- `constant` (Attributes) (see [below for nested schema](#nestedatt--data_sets--filters--category_filter--constant))
- `inverse` (Boolean)

<a id="nestedatt--data_sets--filters--category_filter--constant"></a>
### Nested Schema for `data_sets.filters.category_filter.inverse`

Optional:

- `collective_constant` (Attributes) (see [below for nested schema](#nestedatt--data_sets--filters--category_filter--inverse--collective_constant))
- `constant_type` (String)
- `singular_constant` (String)

<a id="nestedatt--data_sets--filters--category_filter--inverse--collective_constant"></a>
### Nested Schema for `data_sets.filters.category_filter.inverse.collective_constant`

Optional:

- `value_list` (List of String)




<a id="nestedatt--data_sets--filters--date_range_filter"></a>
### Nested Schema for `data_sets.filters.date_range_filter`

Optional:

- `constant` (Attributes) (see [below for nested schema](#nestedatt--data_sets--filters--date_range_filter--constant))
- `inclusive` (Boolean)

<a id="nestedatt--data_sets--filters--date_range_filter--constant"></a>
### Nested Schema for `data_sets.filters.date_range_filter.inclusive`

Optional:

- `constant_type` (String)
- `range_constant` (Attributes) (see [below for nested schema](#nestedatt--data_sets--filters--date_range_filter--inclusive--range_constant))

<a id="nestedatt--data_sets--filters--date_range_filter--inclusive--range_constant"></a>
### Nested Schema for `data_sets.filters.date_range_filter.inclusive.range_constant`

Optional:

- `maximum` (String)
- `minimum` (String)




<a id="nestedatt--data_sets--filters--numeric_equality_filter"></a>
### Nested Schema for `data_sets.filters.numeric_equality_filter`

Optional:

- `aggregation` (String)
- `constant` (Attributes) (see [below for nested schema](#nestedatt--data_sets--filters--numeric_equality_filter--constant))

<a id="nestedatt--data_sets--filters--numeric_equality_filter--constant"></a>
### Nested Schema for `data_sets.filters.numeric_equality_filter.constant`

Optional:

- `constant_type` (String)
- `singular_constant` (String)



<a id="nestedatt--data_sets--filters--numeric_range_filter"></a>
### Nested Schema for `data_sets.filters.numeric_range_filter`

Optional:

- `aggregation` (String)
- `constant` (Attributes) (see [below for nested schema](#nestedatt--data_sets--filters--numeric_range_filter--constant))
- `inclusive` (Boolean)

<a id="nestedatt--data_sets--filters--numeric_range_filter--constant"></a>
### Nested Schema for `data_sets.filters.numeric_range_filter.inclusive`

Optional:

- `constant_type` (String)
- `range_constant` (Attributes) (see [below for nested schema](#nestedatt--data_sets--filters--numeric_range_filter--inclusive--range_constant))

<a id="nestedatt--data_sets--filters--numeric_range_filter--inclusive--range_constant"></a>
### Nested Schema for `data_sets.filters.numeric_range_filter.inclusive.range_constant`

Optional:

- `maximum` (String)
- `minimum` (String)




<a id="nestedatt--data_sets--filters--relative_date_filter"></a>
### Nested Schema for `data_sets.filters.relative_date_filter`

Optional:

- `constant` (Attributes) (see [below for nested schema](#nestedatt--data_sets--filters--relative_date_filter--constant))
- `relative_date_filter_function` (String)
- `time_granularity` (String)

<a id="nestedatt--data_sets--filters--relative_date_filter--constant"></a>
### Nested Schema for `data_sets.filters.relative_date_filter.time_granularity`

Optional:

- `constant_type` (String)
- `singular_constant` (String)




<a id="nestedatt--data_sets--named_entities"></a>
### Nested Schema for `data_sets.named_entities`

Required:

- `entity_name` (String)

Optional:

- `definition` (Attributes List) (see [below for nested schema](#nestedatt--data_sets--named_entities--definition))
- `entity_description` (String)
- `entity_synonyms` (List of String)
- `semantic_entity_type` (Attributes) (see [below for nested schema](#nestedatt--data_sets--named_entities--semantic_entity_type))

<a id="nestedatt--data_sets--named_entities--definition"></a>
### Nested Schema for `data_sets.named_entities.definition`

Optional:

- `field_name` (String)
- `metric` (Attributes) (see [below for nested schema](#nestedatt--data_sets--named_entities--definition--metric))
- `property_name` (String)
- `property_role` (String)
- `property_usage` (String)

<a id="nestedatt--data_sets--named_entities--definition--metric"></a>
### Nested Schema for `data_sets.named_entities.definition.property_usage`

Optional:

- `aggregation` (String)
- `aggregation_function_parameters` (Map of String)



<a id="nestedatt--data_sets--named_entities--semantic_entity_type"></a>
### Nested Schema for `data_sets.named_entities.semantic_entity_type`

Optional:

- `sub_type_name` (String)
- `type_name` (String)
- `type_parameters` (Map of String)

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_quicksight_topic.example <resource ID>
```
