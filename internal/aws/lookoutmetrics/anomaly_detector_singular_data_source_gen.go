// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package lookoutmetrics

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_lookoutmetrics_anomaly_detector", anomalyDetectorDataSourceType)
}

// anomalyDetectorDataSourceType returns the Terraform awscc_lookoutmetrics_anomaly_detector data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::LookoutMetrics::AnomalyDetector resource type.
func anomalyDetectorDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"anomaly_detector_config": {
			// Property: AnomalyDetectorConfig
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "Configuration options for the AnomalyDetector",
			//   "properties": {
			//     "AnomalyDetectorFrequency": {
			//       "description": "Frequency of anomaly detection",
			//       "enum": [
			//         "PT5M",
			//         "PT10M",
			//         "PT1H",
			//         "P1D"
			//       ],
			//       "type": "string"
			//     }
			//   },
			//   "required": [
			//     "AnomalyDetectorFrequency"
			//   ],
			//   "type": "object"
			// }
			Description: "Configuration options for the AnomalyDetector",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"anomaly_detector_frequency": {
						// Property: AnomalyDetectorFrequency
						Description: "Frequency of anomaly detection",
						Type:        types.StringType,
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
		"anomaly_detector_description": {
			// Property: AnomalyDetectorDescription
			// CloudFormation resource type schema:
			// {
			//   "description": "A description for the AnomalyDetector.",
			//   "maxLength": 256,
			//   "pattern": ".*\\S.*",
			//   "type": "string"
			// }
			Description: "A description for the AnomalyDetector.",
			Type:        types.StringType,
			Computed:    true,
		},
		"anomaly_detector_name": {
			// Property: AnomalyDetectorName
			// CloudFormation resource type schema:
			// {
			//   "description": "Name for the Amazon Lookout for Metrics Anomaly Detector",
			//   "maxLength": 63,
			//   "minLength": 1,
			//   "pattern": "^[a-zA-Z0-9][a-zA-Z0-9\\-_]*",
			//   "type": "string"
			// }
			Description: "Name for the Amazon Lookout for Metrics Anomaly Detector",
			Type:        types.StringType,
			Computed:    true,
		},
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 256,
			//   "pattern": "arn:([a-z\\d-]+):.*:.*:.*:.+",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"kms_key_arn": {
			// Property: KmsKeyArn
			// CloudFormation resource type schema:
			// {
			//   "description": "KMS key used to encrypt the AnomalyDetector data",
			//   "maxLength": 2048,
			//   "minLength": 20,
			//   "pattern": "arn:aws.*:kms:.*:[0-9]{12}:key/.*",
			//   "type": "string"
			// }
			Description: "KMS key used to encrypt the AnomalyDetector data",
			Type:        types.StringType,
			Computed:    true,
		},
		"metric_set_list": {
			// Property: MetricSetList
			// CloudFormation resource type schema:
			// {
			//   "description": "List of metric sets for anomaly detection",
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "DimensionList": {
			//         "description": "Dimensions for this MetricSet.",
			//         "insertionOrder": false,
			//         "items": {
			//           "description": "Name of a column in the data.",
			//           "maxLength": 63,
			//           "minLength": 1,
			//           "pattern": "^[a-zA-Z0-9][a-zA-Z0-9\\-_]*",
			//           "type": "string"
			//         },
			//         "minItems": 0,
			//         "type": "array"
			//       },
			//       "MetricList": {
			//         "description": "Metrics captured by this MetricSet.",
			//         "insertionOrder": false,
			//         "items": {
			//           "additionalProperties": false,
			//           "properties": {
			//             "AggregationFunction": {
			//               "description": "Operator used to aggregate metric values",
			//               "enum": [
			//                 "AVG",
			//                 "SUM"
			//               ],
			//               "type": "string"
			//             },
			//             "MetricName": {
			//               "description": "Name of a column in the data.",
			//               "maxLength": 63,
			//               "minLength": 1,
			//               "pattern": "^[a-zA-Z0-9][a-zA-Z0-9\\-_]*",
			//               "type": "string"
			//             },
			//             "Namespace": {
			//               "maxLength": 255,
			//               "minLength": 1,
			//               "pattern": "[^:].*",
			//               "type": "string"
			//             }
			//           },
			//           "required": [
			//             "MetricName",
			//             "AggregationFunction"
			//           ],
			//           "type": "object"
			//         },
			//         "minItems": 1,
			//         "type": "array"
			//       },
			//       "MetricSetDescription": {
			//         "description": "A description for the MetricSet.",
			//         "maxLength": 256,
			//         "pattern": ".*\\S.*",
			//         "type": "string"
			//       },
			//       "MetricSetFrequency": {
			//         "description": "A frequency period to aggregate the data",
			//         "enum": [
			//           "PT5M",
			//           "PT10M",
			//           "PT1H",
			//           "P1D"
			//         ],
			//         "type": "string"
			//       },
			//       "MetricSetName": {
			//         "description": "The name of the MetricSet.",
			//         "maxLength": 63,
			//         "minLength": 1,
			//         "pattern": "^[a-zA-Z0-9][a-zA-Z0-9\\-_]*",
			//         "type": "string"
			//       },
			//       "MetricSource": {
			//         "additionalProperties": false,
			//         "properties": {
			//           "AppFlowConfig": {
			//             "additionalProperties": false,
			//             "properties": {
			//               "FlowName": {
			//                 "maxLength": 256,
			//                 "pattern": "[a-zA-Z0-9][\\w!@#.-]+",
			//                 "type": "string"
			//               },
			//               "RoleArn": {
			//                 "maxLength": 256,
			//                 "pattern": "arn:([a-z\\d-]+):.*:.*:.*:.+",
			//                 "type": "string"
			//               }
			//             },
			//             "required": [
			//               "RoleArn",
			//               "FlowName"
			//             ],
			//             "type": "object"
			//           },
			//           "CloudwatchConfig": {
			//             "additionalProperties": false,
			//             "properties": {
			//               "RoleArn": {
			//                 "maxLength": 256,
			//                 "pattern": "arn:([a-z\\d-]+):.*:.*:.*:.+",
			//                 "type": "string"
			//               }
			//             },
			//             "required": [
			//               "RoleArn"
			//             ],
			//             "type": "object"
			//           },
			//           "RDSSourceConfig": {
			//             "additionalProperties": false,
			//             "properties": {
			//               "DBInstanceIdentifier": {
			//                 "maxLength": 63,
			//                 "minLength": 1,
			//                 "pattern": "",
			//                 "type": "string"
			//               },
			//               "DatabaseHost": {
			//                 "maxLength": 253,
			//                 "minLength": 1,
			//                 "pattern": ".*\\S.*",
			//                 "type": "string"
			//               },
			//               "DatabaseName": {
			//                 "maxLength": 64,
			//                 "minLength": 1,
			//                 "pattern": "[a-zA-Z0-9_]+",
			//                 "type": "string"
			//               },
			//               "DatabasePort": {
			//                 "maximum": 65535,
			//                 "minimum": 1,
			//                 "type": "integer"
			//               },
			//               "RoleArn": {
			//                 "maxLength": 256,
			//                 "pattern": "arn:([a-z\\d-]+):.*:.*:.*:.+",
			//                 "type": "string"
			//               },
			//               "SecretManagerArn": {
			//                 "maxLength": 256,
			//                 "pattern": "arn:([a-z\\d-]+):.*:.*:secret:AmazonLookoutMetrics-.+",
			//                 "type": "string"
			//               },
			//               "TableName": {
			//                 "maxLength": 100,
			//                 "minLength": 1,
			//                 "pattern": "^[a-zA-Z][a-zA-Z0-9_]*$",
			//                 "type": "string"
			//               },
			//               "VpcConfiguration": {
			//                 "additionalProperties": false,
			//                 "properties": {
			//                   "SecurityGroupIdList": {
			//                     "items": {
			//                       "maxLength": 255,
			//                       "minLength": 1,
			//                       "pattern": "[-0-9a-zA-Z]+",
			//                       "type": "string"
			//                     },
			//                     "type": "array"
			//                   },
			//                   "SubnetIdList": {
			//                     "items": {
			//                       "maxLength": 255,
			//                       "pattern": "[\\-0-9a-zA-Z]+",
			//                       "type": "string"
			//                     },
			//                     "type": "array"
			//                   }
			//                 },
			//                 "required": [
			//                   "SubnetIdList",
			//                   "SecurityGroupIdList"
			//                 ],
			//                 "type": "object"
			//               }
			//             },
			//             "required": [
			//               "DBInstanceIdentifier",
			//               "DatabaseHost",
			//               "DatabasePort",
			//               "SecretManagerArn",
			//               "DatabaseName",
			//               "TableName",
			//               "RoleArn",
			//               "VpcConfiguration"
			//             ],
			//             "type": "object"
			//           },
			//           "RedshiftSourceConfig": {
			//             "additionalProperties": false,
			//             "properties": {
			//               "ClusterIdentifier": {
			//                 "maxLength": 63,
			//                 "minLength": 1,
			//                 "pattern": "",
			//                 "type": "string"
			//               },
			//               "DatabaseHost": {
			//                 "maxLength": 253,
			//                 "minLength": 1,
			//                 "pattern": ".*\\S.*",
			//                 "type": "string"
			//               },
			//               "DatabaseName": {
			//                 "maxLength": 100,
			//                 "minLength": 1,
			//                 "pattern": "[a-z0-9]+",
			//                 "type": "string"
			//               },
			//               "DatabasePort": {
			//                 "maximum": 65535,
			//                 "minimum": 1,
			//                 "type": "integer"
			//               },
			//               "RoleArn": {
			//                 "maxLength": 256,
			//                 "pattern": "arn:([a-z\\d-]+):.*:.*:.*:.+",
			//                 "type": "string"
			//               },
			//               "SecretManagerArn": {
			//                 "maxLength": 256,
			//                 "pattern": "arn:([a-z\\d-]+):.*:.*:secret:AmazonLookoutMetrics-.+",
			//                 "type": "string"
			//               },
			//               "TableName": {
			//                 "maxLength": 100,
			//                 "minLength": 1,
			//                 "pattern": "^[a-zA-Z][a-zA-Z0-9_]*$",
			//                 "type": "string"
			//               },
			//               "VpcConfiguration": {
			//                 "additionalProperties": false,
			//                 "properties": {
			//                   "SecurityGroupIdList": {
			//                     "items": {
			//                       "maxLength": 255,
			//                       "minLength": 1,
			//                       "pattern": "[-0-9a-zA-Z]+",
			//                       "type": "string"
			//                     },
			//                     "type": "array"
			//                   },
			//                   "SubnetIdList": {
			//                     "items": {
			//                       "maxLength": 255,
			//                       "pattern": "[\\-0-9a-zA-Z]+",
			//                       "type": "string"
			//                     },
			//                     "type": "array"
			//                   }
			//                 },
			//                 "required": [
			//                   "SubnetIdList",
			//                   "SecurityGroupIdList"
			//                 ],
			//                 "type": "object"
			//               }
			//             },
			//             "required": [
			//               "ClusterIdentifier",
			//               "DatabaseHost",
			//               "DatabasePort",
			//               "SecretManagerArn",
			//               "DatabaseName",
			//               "TableName",
			//               "RoleArn",
			//               "VpcConfiguration"
			//             ],
			//             "type": "object"
			//           },
			//           "S3SourceConfig": {
			//             "additionalProperties": false,
			//             "properties": {
			//               "FileFormatDescriptor": {
			//                 "additionalProperties": false,
			//                 "properties": {
			//                   "CsvFormatDescriptor": {
			//                     "additionalProperties": false,
			//                     "properties": {
			//                       "Charset": {
			//                         "maxLength": 63,
			//                         "pattern": "^[a-zA-Z0-9][a-zA-Z0-9\\-_]*",
			//                         "type": "string"
			//                       },
			//                       "ContainsHeader": {
			//                         "type": "boolean"
			//                       },
			//                       "Delimiter": {
			//                         "maxLength": 1,
			//                         "pattern": "[^\\r\\n]",
			//                         "type": "string"
			//                       },
			//                       "FileCompression": {
			//                         "enum": [
			//                           "NONE",
			//                           "GZIP"
			//                         ],
			//                         "type": "string"
			//                       },
			//                       "HeaderList": {
			//                         "items": {
			//                           "description": "Name of a column in the data.",
			//                           "maxLength": 63,
			//                           "minLength": 1,
			//                           "pattern": "^[a-zA-Z0-9][a-zA-Z0-9\\-_]*",
			//                           "type": "string"
			//                         },
			//                         "type": "array"
			//                       },
			//                       "QuoteSymbol": {
			//                         "maxLength": 1,
			//                         "pattern": "[^\\r\\n]|^$",
			//                         "type": "string"
			//                       }
			//                     },
			//                     "type": "object"
			//                   },
			//                   "JsonFormatDescriptor": {
			//                     "additionalProperties": false,
			//                     "properties": {
			//                       "Charset": {
			//                         "maxLength": 63,
			//                         "pattern": "^[a-zA-Z0-9][a-zA-Z0-9\\-_]*",
			//                         "type": "string"
			//                       },
			//                       "FileCompression": {
			//                         "enum": [
			//                           "NONE",
			//                           "GZIP"
			//                         ],
			//                         "type": "string"
			//                       }
			//                     },
			//                     "type": "object"
			//                   }
			//                 },
			//                 "type": "object"
			//               },
			//               "HistoricalDataPathList": {
			//                 "items": {
			//                   "maxLength": 1024,
			//                   "pattern": "^s3://[a-z0-9].+$",
			//                   "type": "string"
			//                 },
			//                 "maxItems": 1,
			//                 "minItems": 1,
			//                 "type": "array"
			//               },
			//               "RoleArn": {
			//                 "maxLength": 256,
			//                 "pattern": "arn:([a-z\\d-]+):.*:.*:.*:.+",
			//                 "type": "string"
			//               },
			//               "TemplatedPathList": {
			//                 "items": {
			//                   "maxLength": 1024,
			//                   "pattern": "^s3://[a-zA-Z0-9_\\-\\/ {}=]+$",
			//                   "type": "string"
			//                 },
			//                 "maxItems": 1,
			//                 "minItems": 1,
			//                 "type": "array"
			//               }
			//             },
			//             "required": [
			//               "RoleArn",
			//               "FileFormatDescriptor"
			//             ],
			//             "type": "object"
			//           }
			//         },
			//         "type": "object"
			//       },
			//       "Offset": {
			//         "description": "Offset, in seconds, between the frequency interval and the time at which the metrics are available.",
			//         "maximum": 432000,
			//         "minimum": 0,
			//         "type": "integer"
			//       },
			//       "TimestampColumn": {
			//         "additionalProperties": false,
			//         "properties": {
			//           "ColumnFormat": {
			//             "description": "A timestamp format for the timestamps in the dataset",
			//             "maxLength": 63,
			//             "pattern": ".*\\S.*",
			//             "type": "string"
			//           },
			//           "ColumnName": {
			//             "description": "Name of a column in the data.",
			//             "maxLength": 63,
			//             "minLength": 1,
			//             "pattern": "^[a-zA-Z0-9][a-zA-Z0-9\\-_]*",
			//             "type": "string"
			//           }
			//         },
			//         "type": "object"
			//       },
			//       "Timezone": {
			//         "maxLength": 60,
			//         "pattern": ".*\\S.*",
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "MetricSetName",
			//       "MetricList",
			//       "MetricSource"
			//     ],
			//     "type": "object"
			//   },
			//   "maxItems": 1,
			//   "minItems": 1,
			//   "type": "array"
			// }
			Description: "List of metric sets for anomaly detection",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"dimension_list": {
						// Property: DimensionList
						Description: "Dimensions for this MetricSet.",
						Type:        types.ListType{ElemType: types.StringType},
						Computed:    true,
					},
					"metric_list": {
						// Property: MetricList
						Description: "Metrics captured by this MetricSet.",
						Attributes: tfsdk.ListNestedAttributes(
							map[string]tfsdk.Attribute{
								"aggregation_function": {
									// Property: AggregationFunction
									Description: "Operator used to aggregate metric values",
									Type:        types.StringType,
									Computed:    true,
								},
								"metric_name": {
									// Property: MetricName
									Description: "Name of a column in the data.",
									Type:        types.StringType,
									Computed:    true,
								},
								"namespace": {
									// Property: Namespace
									Type:     types.StringType,
									Computed: true,
								},
							},
							tfsdk.ListNestedAttributesOptions{},
						),
						Computed: true,
					},
					"metric_set_description": {
						// Property: MetricSetDescription
						Description: "A description for the MetricSet.",
						Type:        types.StringType,
						Computed:    true,
					},
					"metric_set_frequency": {
						// Property: MetricSetFrequency
						Description: "A frequency period to aggregate the data",
						Type:        types.StringType,
						Computed:    true,
					},
					"metric_set_name": {
						// Property: MetricSetName
						Description: "The name of the MetricSet.",
						Type:        types.StringType,
						Computed:    true,
					},
					"metric_source": {
						// Property: MetricSource
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"app_flow_config": {
									// Property: AppFlowConfig
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"flow_name": {
												// Property: FlowName
												Type:     types.StringType,
												Computed: true,
											},
											"role_arn": {
												// Property: RoleArn
												Type:     types.StringType,
												Computed: true,
											},
										},
									),
									Computed: true,
								},
								"cloudwatch_config": {
									// Property: CloudwatchConfig
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"role_arn": {
												// Property: RoleArn
												Type:     types.StringType,
												Computed: true,
											},
										},
									),
									Computed: true,
								},
								"rds_source_config": {
									// Property: RDSSourceConfig
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"db_instance_identifier": {
												// Property: DBInstanceIdentifier
												Type:     types.StringType,
												Computed: true,
											},
											"database_host": {
												// Property: DatabaseHost
												Type:     types.StringType,
												Computed: true,
											},
											"database_name": {
												// Property: DatabaseName
												Type:     types.StringType,
												Computed: true,
											},
											"database_port": {
												// Property: DatabasePort
												Type:     types.Int64Type,
												Computed: true,
											},
											"role_arn": {
												// Property: RoleArn
												Type:     types.StringType,
												Computed: true,
											},
											"secret_manager_arn": {
												// Property: SecretManagerArn
												Type:     types.StringType,
												Computed: true,
											},
											"table_name": {
												// Property: TableName
												Type:     types.StringType,
												Computed: true,
											},
											"vpc_configuration": {
												// Property: VpcConfiguration
												Attributes: tfsdk.SingleNestedAttributes(
													map[string]tfsdk.Attribute{
														"security_group_id_list": {
															// Property: SecurityGroupIdList
															Type:     types.ListType{ElemType: types.StringType},
															Computed: true,
														},
														"subnet_id_list": {
															// Property: SubnetIdList
															Type:     types.ListType{ElemType: types.StringType},
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
								"redshift_source_config": {
									// Property: RedshiftSourceConfig
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"cluster_identifier": {
												// Property: ClusterIdentifier
												Type:     types.StringType,
												Computed: true,
											},
											"database_host": {
												// Property: DatabaseHost
												Type:     types.StringType,
												Computed: true,
											},
											"database_name": {
												// Property: DatabaseName
												Type:     types.StringType,
												Computed: true,
											},
											"database_port": {
												// Property: DatabasePort
												Type:     types.Int64Type,
												Computed: true,
											},
											"role_arn": {
												// Property: RoleArn
												Type:     types.StringType,
												Computed: true,
											},
											"secret_manager_arn": {
												// Property: SecretManagerArn
												Type:     types.StringType,
												Computed: true,
											},
											"table_name": {
												// Property: TableName
												Type:     types.StringType,
												Computed: true,
											},
											"vpc_configuration": {
												// Property: VpcConfiguration
												Attributes: tfsdk.SingleNestedAttributes(
													map[string]tfsdk.Attribute{
														"security_group_id_list": {
															// Property: SecurityGroupIdList
															Type:     types.ListType{ElemType: types.StringType},
															Computed: true,
														},
														"subnet_id_list": {
															// Property: SubnetIdList
															Type:     types.ListType{ElemType: types.StringType},
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
								"s3_source_config": {
									// Property: S3SourceConfig
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"file_format_descriptor": {
												// Property: FileFormatDescriptor
												Attributes: tfsdk.SingleNestedAttributes(
													map[string]tfsdk.Attribute{
														"csv_format_descriptor": {
															// Property: CsvFormatDescriptor
															Attributes: tfsdk.SingleNestedAttributes(
																map[string]tfsdk.Attribute{
																	"charset": {
																		// Property: Charset
																		Type:     types.StringType,
																		Computed: true,
																	},
																	"contains_header": {
																		// Property: ContainsHeader
																		Type:     types.BoolType,
																		Computed: true,
																	},
																	"delimiter": {
																		// Property: Delimiter
																		Type:     types.StringType,
																		Computed: true,
																	},
																	"file_compression": {
																		// Property: FileCompression
																		Type:     types.StringType,
																		Computed: true,
																	},
																	"header_list": {
																		// Property: HeaderList
																		Type:     types.ListType{ElemType: types.StringType},
																		Computed: true,
																	},
																	"quote_symbol": {
																		// Property: QuoteSymbol
																		Type:     types.StringType,
																		Computed: true,
																	},
																},
															),
															Computed: true,
														},
														"json_format_descriptor": {
															// Property: JsonFormatDescriptor
															Attributes: tfsdk.SingleNestedAttributes(
																map[string]tfsdk.Attribute{
																	"charset": {
																		// Property: Charset
																		Type:     types.StringType,
																		Computed: true,
																	},
																	"file_compression": {
																		// Property: FileCompression
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
											"historical_data_path_list": {
												// Property: HistoricalDataPathList
												Type:     types.ListType{ElemType: types.StringType},
												Computed: true,
											},
											"role_arn": {
												// Property: RoleArn
												Type:     types.StringType,
												Computed: true,
											},
											"templated_path_list": {
												// Property: TemplatedPathList
												Type:     types.ListType{ElemType: types.StringType},
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
					"offset": {
						// Property: Offset
						Description: "Offset, in seconds, between the frequency interval and the time at which the metrics are available.",
						Type:        types.Int64Type,
						Computed:    true,
					},
					"timestamp_column": {
						// Property: TimestampColumn
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"column_format": {
									// Property: ColumnFormat
									Description: "A timestamp format for the timestamps in the dataset",
									Type:        types.StringType,
									Computed:    true,
								},
								"column_name": {
									// Property: ColumnName
									Description: "Name of a column in the data.",
									Type:        types.StringType,
									Computed:    true,
								},
							},
						),
						Computed: true,
					},
					"timezone": {
						// Property: Timezone
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
		Description: "Data Source schema for AWS::LookoutMetrics::AnomalyDetector",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::LookoutMetrics::AnomalyDetector").WithTerraformTypeName("awscc_lookoutmetrics_anomaly_detector")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"aggregation_function":         "AggregationFunction",
		"anomaly_detector_config":      "AnomalyDetectorConfig",
		"anomaly_detector_description": "AnomalyDetectorDescription",
		"anomaly_detector_frequency":   "AnomalyDetectorFrequency",
		"anomaly_detector_name":        "AnomalyDetectorName",
		"app_flow_config":              "AppFlowConfig",
		"arn":                          "Arn",
		"charset":                      "Charset",
		"cloudwatch_config":            "CloudwatchConfig",
		"cluster_identifier":           "ClusterIdentifier",
		"column_format":                "ColumnFormat",
		"column_name":                  "ColumnName",
		"contains_header":              "ContainsHeader",
		"csv_format_descriptor":        "CsvFormatDescriptor",
		"database_host":                "DatabaseHost",
		"database_name":                "DatabaseName",
		"database_port":                "DatabasePort",
		"db_instance_identifier":       "DBInstanceIdentifier",
		"delimiter":                    "Delimiter",
		"dimension_list":               "DimensionList",
		"file_compression":             "FileCompression",
		"file_format_descriptor":       "FileFormatDescriptor",
		"flow_name":                    "FlowName",
		"header_list":                  "HeaderList",
		"historical_data_path_list":    "HistoricalDataPathList",
		"json_format_descriptor":       "JsonFormatDescriptor",
		"kms_key_arn":                  "KmsKeyArn",
		"metric_list":                  "MetricList",
		"metric_name":                  "MetricName",
		"metric_set_description":       "MetricSetDescription",
		"metric_set_frequency":         "MetricSetFrequency",
		"metric_set_list":              "MetricSetList",
		"metric_set_name":              "MetricSetName",
		"metric_source":                "MetricSource",
		"namespace":                    "Namespace",
		"offset":                       "Offset",
		"quote_symbol":                 "QuoteSymbol",
		"rds_source_config":            "RDSSourceConfig",
		"redshift_source_config":       "RedshiftSourceConfig",
		"role_arn":                     "RoleArn",
		"s3_source_config":             "S3SourceConfig",
		"secret_manager_arn":           "SecretManagerArn",
		"security_group_id_list":       "SecurityGroupIdList",
		"subnet_id_list":               "SubnetIdList",
		"table_name":                   "TableName",
		"templated_path_list":          "TemplatedPathList",
		"timestamp_column":             "TimestampColumn",
		"timezone":                     "Timezone",
		"vpc_configuration":            "VpcConfiguration",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
