// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package databrew

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_databrew_dataset", datasetDataSourceType)
}

// datasetDataSourceType returns the Terraform awscc_databrew_dataset data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::DataBrew::Dataset resource type.
func datasetDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"format": {
			// Property: Format
			// CloudFormation resource type schema:
			// {
			//   "description": "Dataset format",
			//   "enum": [
			//     "CSV",
			//     "JSON",
			//     "PARQUET",
			//     "EXCEL"
			//   ],
			//   "type": "string"
			// }
			Description: "Dataset format",
			Type:        types.StringType,
			Computed:    true,
		},
		"format_options": {
			// Property: FormatOptions
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "Format options for dataset",
			//   "properties": {
			//     "Csv": {
			//       "additionalProperties": false,
			//       "description": "Csv options",
			//       "properties": {
			//         "Delimiter": {
			//           "maxLength": 1,
			//           "minLength": 1,
			//           "type": "string"
			//         },
			//         "HeaderRow": {
			//           "type": "boolean"
			//         }
			//       },
			//       "type": "object"
			//     },
			//     "Excel": {
			//       "additionalProperties": false,
			//       "oneOf": [
			//         {
			//           "required": [
			//             "SheetNames"
			//           ]
			//         },
			//         {
			//           "required": [
			//             "SheetIndexes"
			//           ]
			//         }
			//       ],
			//       "properties": {
			//         "HeaderRow": {
			//           "type": "boolean"
			//         },
			//         "SheetIndexes": {
			//           "insertionOrder": true,
			//           "items": {
			//             "type": "integer"
			//           },
			//           "maxItems": 1,
			//           "minItems": 1,
			//           "type": "array"
			//         },
			//         "SheetNames": {
			//           "insertionOrder": true,
			//           "items": {
			//             "type": "string"
			//           },
			//           "maxItems": 1,
			//           "minItems": 1,
			//           "type": "array"
			//         }
			//       },
			//       "type": "object"
			//     },
			//     "Json": {
			//       "additionalProperties": false,
			//       "description": "Json options",
			//       "properties": {
			//         "MultiLine": {
			//           "type": "boolean"
			//         }
			//       },
			//       "type": "object"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "Format options for dataset",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"csv": {
						// Property: Csv
						Description: "Csv options",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"delimiter": {
									// Property: Delimiter
									Type:     types.StringType,
									Computed: true,
								},
								"header_row": {
									// Property: HeaderRow
									Type:     types.BoolType,
									Computed: true,
								},
							},
						),
						Computed: true,
					},
					"excel": {
						// Property: Excel
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"header_row": {
									// Property: HeaderRow
									Type:     types.BoolType,
									Computed: true,
								},
								"sheet_indexes": {
									// Property: SheetIndexes
									Type:     types.ListType{ElemType: types.Int64Type},
									Computed: true,
								},
								"sheet_names": {
									// Property: SheetNames
									Type:     types.ListType{ElemType: types.StringType},
									Computed: true,
								},
							},
						),
						Computed: true,
					},
					"json": {
						// Property: Json
						Description: "Json options",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"multi_line": {
									// Property: MultiLine
									Type:     types.BoolType,
									Computed: true,
								},
							},
						),
						Computed: true,
					},
				},
			),
			Computed: true,
		},
		"input": {
			// Property: Input
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "Input",
			//   "properties": {
			//     "DataCatalogInputDefinition": {
			//       "additionalProperties": false,
			//       "properties": {
			//         "CatalogId": {
			//           "description": "Catalog id",
			//           "type": "string"
			//         },
			//         "DatabaseName": {
			//           "description": "Database name",
			//           "type": "string"
			//         },
			//         "TableName": {
			//           "description": "Table name",
			//           "type": "string"
			//         },
			//         "TempDirectory": {
			//           "additionalProperties": false,
			//           "description": "Input location",
			//           "properties": {
			//             "Bucket": {
			//               "type": "string"
			//             },
			//             "Key": {
			//               "type": "string"
			//             }
			//           },
			//           "required": [
			//             "Bucket"
			//           ],
			//           "type": "object"
			//         }
			//       },
			//       "type": "object"
			//     },
			//     "DatabaseInputDefinition": {
			//       "additionalProperties": false,
			//       "properties": {
			//         "DatabaseTableName": {
			//           "description": "Database table name",
			//           "type": "string"
			//         },
			//         "GlueConnectionName": {
			//           "description": "Glue connection name",
			//           "type": "string"
			//         },
			//         "QueryString": {
			//           "description": "Custom SQL to run against the provided AWS Glue connection. This SQL will be used as the input for DataBrew projects and jobs.",
			//           "type": "string"
			//         },
			//         "TempDirectory": {
			//           "additionalProperties": false,
			//           "description": "Input location",
			//           "properties": {
			//             "Bucket": {
			//               "type": "string"
			//             },
			//             "Key": {
			//               "type": "string"
			//             }
			//           },
			//           "required": [
			//             "Bucket"
			//           ],
			//           "type": "object"
			//         }
			//       },
			//       "required": [
			//         "GlueConnectionName"
			//       ],
			//       "type": "object"
			//     },
			//     "Metadata": {
			//       "additionalProperties": false,
			//       "properties": {
			//         "SourceArn": {
			//           "description": "Arn of the source of the dataset. For e.g.: AppFlow Flow ARN.",
			//           "type": "string"
			//         }
			//       },
			//       "type": "object"
			//     },
			//     "S3InputDefinition": {
			//       "additionalProperties": false,
			//       "description": "Input location",
			//       "properties": {
			//         "Bucket": {
			//           "type": "string"
			//         },
			//         "Key": {
			//           "type": "string"
			//         }
			//       },
			//       "required": [
			//         "Bucket"
			//       ],
			//       "type": "object"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "Input",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"data_catalog_input_definition": {
						// Property: DataCatalogInputDefinition
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"catalog_id": {
									// Property: CatalogId
									Description: "Catalog id",
									Type:        types.StringType,
									Computed:    true,
								},
								"database_name": {
									// Property: DatabaseName
									Description: "Database name",
									Type:        types.StringType,
									Computed:    true,
								},
								"table_name": {
									// Property: TableName
									Description: "Table name",
									Type:        types.StringType,
									Computed:    true,
								},
								"temp_directory": {
									// Property: TempDirectory
									Description: "Input location",
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"bucket": {
												// Property: Bucket
												Type:     types.StringType,
												Computed: true,
											},
											"key": {
												// Property: Key
												Type:     types.StringType,
												Computed: true,
											},
										},
									),
									Computed: true,
								},
							},
						),
						Computed: true,
					},
					"database_input_definition": {
						// Property: DatabaseInputDefinition
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"database_table_name": {
									// Property: DatabaseTableName
									Description: "Database table name",
									Type:        types.StringType,
									Computed:    true,
								},
								"glue_connection_name": {
									// Property: GlueConnectionName
									Description: "Glue connection name",
									Type:        types.StringType,
									Computed:    true,
								},
								"query_string": {
									// Property: QueryString
									Description: "Custom SQL to run against the provided AWS Glue connection. This SQL will be used as the input for DataBrew projects and jobs.",
									Type:        types.StringType,
									Computed:    true,
								},
								"temp_directory": {
									// Property: TempDirectory
									Description: "Input location",
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"bucket": {
												// Property: Bucket
												Type:     types.StringType,
												Computed: true,
											},
											"key": {
												// Property: Key
												Type:     types.StringType,
												Computed: true,
											},
										},
									),
									Computed: true,
								},
							},
						),
						Computed: true,
					},
					"metadata": {
						// Property: Metadata
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"source_arn": {
									// Property: SourceArn
									Description: "Arn of the source of the dataset. For e.g.: AppFlow Flow ARN.",
									Type:        types.StringType,
									Computed:    true,
								},
							},
						),
						Computed: true,
					},
					"s3_input_definition": {
						// Property: S3InputDefinition
						Description: "Input location",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"bucket": {
									// Property: Bucket
									Type:     types.StringType,
									Computed: true,
								},
								"key": {
									// Property: Key
									Type:     types.StringType,
									Computed: true,
								},
							},
						),
						Computed: true,
					},
				},
			),
			Computed: true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "Dataset name",
			//   "maxLength": 255,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "Dataset name",
			Type:        types.StringType,
			Computed:    true,
		},
		"path_options": {
			// Property: PathOptions
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "PathOptions",
			//   "properties": {
			//     "FilesLimit": {
			//       "additionalProperties": false,
			//       "properties": {
			//         "MaxFiles": {
			//           "description": "Maximum number of files",
			//           "type": "integer"
			//         },
			//         "Order": {
			//           "description": "Order",
			//           "enum": [
			//             "ASCENDING",
			//             "DESCENDING"
			//           ],
			//           "type": "string"
			//         },
			//         "OrderedBy": {
			//           "description": "Ordered by",
			//           "enum": [
			//             "LAST_MODIFIED_DATE"
			//           ],
			//           "type": "string"
			//         }
			//       },
			//       "required": [
			//         "MaxFiles"
			//       ],
			//       "type": "object"
			//     },
			//     "LastModifiedDateCondition": {
			//       "additionalProperties": false,
			//       "properties": {
			//         "Expression": {
			//           "description": "Filtering expression for a parameter",
			//           "maxLength": 1024,
			//           "minLength": 4,
			//           "pattern": "^[\u003e\u003c0-9A-Za-z_.,:)(!= ]+$",
			//           "type": "string"
			//         },
			//         "ValuesMap": {
			//           "insertionOrder": true,
			//           "items": {
			//             "additionalProperties": false,
			//             "description": "A key-value pair to associate expression variable names with their values",
			//             "properties": {
			//               "Value": {
			//                 "maxLength": 1024,
			//                 "minLength": 0,
			//                 "type": "string"
			//               },
			//               "ValueReference": {
			//                 "description": "Variable name",
			//                 "maxLength": 128,
			//                 "minLength": 2,
			//                 "pattern": "^:[A-Za-z0-9_]+$",
			//                 "type": "string"
			//               }
			//             },
			//             "required": [
			//               "ValueReference",
			//               "Value"
			//             ],
			//             "type": "object"
			//           },
			//           "type": "array"
			//         }
			//       },
			//       "required": [
			//         "Expression",
			//         "ValuesMap"
			//       ],
			//       "type": "object"
			//     },
			//     "Parameters": {
			//       "insertionOrder": true,
			//       "items": {
			//         "additionalProperties": false,
			//         "description": "A key-value pair to associate dataset parameter name with its definition.",
			//         "properties": {
			//           "DatasetParameter": {
			//             "additionalProperties": false,
			//             "properties": {
			//               "CreateColumn": {
			//                 "description": "Add the value of this parameter as a column in a dataset.",
			//                 "type": "boolean"
			//               },
			//               "DatetimeOptions": {
			//                 "additionalProperties": false,
			//                 "properties": {
			//                   "Format": {
			//                     "description": "Date/time format of a date parameter",
			//                     "maxLength": 100,
			//                     "minLength": 2,
			//                     "type": "string"
			//                   },
			//                   "LocaleCode": {
			//                     "description": "Locale code for a date parameter",
			//                     "maxLength": 100,
			//                     "minLength": 2,
			//                     "pattern": "^[A-Za-z0-9_\\.#@\\-]+$",
			//                     "type": "string"
			//                   },
			//                   "TimezoneOffset": {
			//                     "description": "Timezone offset",
			//                     "maxLength": 6,
			//                     "minLength": 1,
			//                     "pattern": "^(Z|[-+](\\d|\\d{2}|\\d{2}:?\\d{2}))$",
			//                     "type": "string"
			//                   }
			//                 },
			//                 "required": [
			//                   "Format"
			//                 ],
			//                 "type": "object"
			//               },
			//               "Filter": {
			//                 "additionalProperties": false,
			//                 "properties": {
			//                   "Expression": {
			//                     "description": "Filtering expression for a parameter",
			//                     "maxLength": 1024,
			//                     "minLength": 4,
			//                     "pattern": "^[\u003e\u003c0-9A-Za-z_.,:)(!= ]+$",
			//                     "type": "string"
			//                   },
			//                   "ValuesMap": {
			//                     "insertionOrder": true,
			//                     "items": {
			//                       "additionalProperties": false,
			//                       "description": "A key-value pair to associate expression variable names with their values",
			//                       "properties": {
			//                         "Value": {
			//                           "maxLength": 1024,
			//                           "minLength": 0,
			//                           "type": "string"
			//                         },
			//                         "ValueReference": {
			//                           "description": "Variable name",
			//                           "maxLength": 128,
			//                           "minLength": 2,
			//                           "pattern": "^:[A-Za-z0-9_]+$",
			//                           "type": "string"
			//                         }
			//                       },
			//                       "required": [
			//                         "ValueReference",
			//                         "Value"
			//                       ],
			//                       "type": "object"
			//                     },
			//                     "type": "array"
			//                   }
			//                 },
			//                 "required": [
			//                   "Expression",
			//                   "ValuesMap"
			//                 ],
			//                 "type": "object"
			//               },
			//               "Name": {
			//                 "description": "Parameter name",
			//                 "maxLength": 255,
			//                 "minLength": 1,
			//                 "type": "string"
			//               },
			//               "Type": {
			//                 "description": "Parameter type",
			//                 "enum": [
			//                   "String",
			//                   "Number",
			//                   "Datetime"
			//                 ],
			//                 "type": "string"
			//               }
			//             },
			//             "required": [
			//               "Name",
			//               "Type"
			//             ],
			//             "type": "object"
			//           },
			//           "PathParameterName": {
			//             "description": "Parameter name",
			//             "maxLength": 255,
			//             "minLength": 1,
			//             "type": "string"
			//           }
			//         },
			//         "required": [
			//           "PathParameterName",
			//           "DatasetParameter"
			//         ],
			//         "type": "object"
			//       },
			//       "type": "array"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "PathOptions",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"files_limit": {
						// Property: FilesLimit
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"max_files": {
									// Property: MaxFiles
									Description: "Maximum number of files",
									Type:        types.Int64Type,
									Computed:    true,
								},
								"order": {
									// Property: Order
									Description: "Order",
									Type:        types.StringType,
									Computed:    true,
								},
								"ordered_by": {
									// Property: OrderedBy
									Description: "Ordered by",
									Type:        types.StringType,
									Computed:    true,
								},
							},
						),
						Computed: true,
					},
					"last_modified_date_condition": {
						// Property: LastModifiedDateCondition
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"expression": {
									// Property: Expression
									Description: "Filtering expression for a parameter",
									Type:        types.StringType,
									Computed:    true,
								},
								"values_map": {
									// Property: ValuesMap
									Attributes: tfsdk.ListNestedAttributes(
										map[string]tfsdk.Attribute{
											"value": {
												// Property: Value
												Type:     types.StringType,
												Computed: true,
											},
											"value_reference": {
												// Property: ValueReference
												Description: "Variable name",
												Type:        types.StringType,
												Computed:    true,
											},
										},
										tfsdk.ListNestedAttributesOptions{},
									),
									Computed: true,
								},
							},
						),
						Computed: true,
					},
					"parameters": {
						// Property: Parameters
						Attributes: tfsdk.ListNestedAttributes(
							map[string]tfsdk.Attribute{
								"dataset_parameter": {
									// Property: DatasetParameter
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"create_column": {
												// Property: CreateColumn
												Description: "Add the value of this parameter as a column in a dataset.",
												Type:        types.BoolType,
												Computed:    true,
											},
											"datetime_options": {
												// Property: DatetimeOptions
												Attributes: tfsdk.SingleNestedAttributes(
													map[string]tfsdk.Attribute{
														"format": {
															// Property: Format
															Description: "Date/time format of a date parameter",
															Type:        types.StringType,
															Computed:    true,
														},
														"locale_code": {
															// Property: LocaleCode
															Description: "Locale code for a date parameter",
															Type:        types.StringType,
															Computed:    true,
														},
														"timezone_offset": {
															// Property: TimezoneOffset
															Description: "Timezone offset",
															Type:        types.StringType,
															Computed:    true,
														},
													},
												),
												Computed: true,
											},
											"filter": {
												// Property: Filter
												Attributes: tfsdk.SingleNestedAttributes(
													map[string]tfsdk.Attribute{
														"expression": {
															// Property: Expression
															Description: "Filtering expression for a parameter",
															Type:        types.StringType,
															Computed:    true,
														},
														"values_map": {
															// Property: ValuesMap
															Attributes: tfsdk.ListNestedAttributes(
																map[string]tfsdk.Attribute{
																	"value": {
																		// Property: Value
																		Type:     types.StringType,
																		Computed: true,
																	},
																	"value_reference": {
																		// Property: ValueReference
																		Description: "Variable name",
																		Type:        types.StringType,
																		Computed:    true,
																	},
																},
																tfsdk.ListNestedAttributesOptions{},
															),
															Computed: true,
														},
													},
												),
												Computed: true,
											},
											"name": {
												// Property: Name
												Description: "Parameter name",
												Type:        types.StringType,
												Computed:    true,
											},
											"type": {
												// Property: Type
												Description: "Parameter type",
												Type:        types.StringType,
												Computed:    true,
											},
										},
									),
									Computed: true,
								},
								"path_parameter_name": {
									// Property: PathParameterName
									Description: "Parameter name",
									Type:        types.StringType,
									Computed:    true,
								},
							},
							tfsdk.ListNestedAttributesOptions{},
						),
						Computed: true,
					},
				},
			),
			Computed: true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "A key-value pair to associate with a resource.",
			//     "properties": {
			//       "Key": {
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Value": {
			//         "maxLength": 256,
			//         "minLength": 0,
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Value",
			//       "Key"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array",
			//   "uniqueItems": false
			// }
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Type:     types.StringType,
						Computed: true,
					},
					"value": {
						// Property: Value
						Type:     types.StringType,
						Computed: true,
					},
				},
				tfsdk.ListNestedAttributesOptions{},
			),
			Computed: true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::DataBrew::Dataset",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::DataBrew::Dataset").WithTerraformTypeName("awscc_databrew_dataset")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"bucket":                        "Bucket",
		"catalog_id":                    "CatalogId",
		"create_column":                 "CreateColumn",
		"csv":                           "Csv",
		"data_catalog_input_definition": "DataCatalogInputDefinition",
		"database_input_definition":     "DatabaseInputDefinition",
		"database_name":                 "DatabaseName",
		"database_table_name":           "DatabaseTableName",
		"dataset_parameter":             "DatasetParameter",
		"datetime_options":              "DatetimeOptions",
		"delimiter":                     "Delimiter",
		"excel":                         "Excel",
		"expression":                    "Expression",
		"files_limit":                   "FilesLimit",
		"filter":                        "Filter",
		"format":                        "Format",
		"format_options":                "FormatOptions",
		"glue_connection_name":          "GlueConnectionName",
		"header_row":                    "HeaderRow",
		"input":                         "Input",
		"json":                          "Json",
		"key":                           "Key",
		"last_modified_date_condition":  "LastModifiedDateCondition",
		"locale_code":                   "LocaleCode",
		"max_files":                     "MaxFiles",
		"metadata":                      "Metadata",
		"multi_line":                    "MultiLine",
		"name":                          "Name",
		"order":                         "Order",
		"ordered_by":                    "OrderedBy",
		"parameters":                    "Parameters",
		"path_options":                  "PathOptions",
		"path_parameter_name":           "PathParameterName",
		"query_string":                  "QueryString",
		"s3_input_definition":           "S3InputDefinition",
		"sheet_indexes":                 "SheetIndexes",
		"sheet_names":                   "SheetNames",
		"source_arn":                    "SourceArn",
		"table_name":                    "TableName",
		"tags":                          "Tags",
		"temp_directory":                "TempDirectory",
		"timezone_offset":               "TimezoneOffset",
		"type":                          "Type",
		"value":                         "Value",
		"value_reference":               "ValueReference",
		"values_map":                    "ValuesMap",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
