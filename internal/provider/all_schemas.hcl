defaults {
  schema_cache_directory     = "../service/cloudformation/schemas"
  terraform_type_name_prefix = "awscc"
}

meta_schema {
  path = "../service/cloudformation/meta-schemas/provider.definition.schema.v1.json"
}

# 606 CloudFormation resource types schemas are available for use with the Cloud Control API.

resource_schema "aws_accessanalyzer_analyzer" {
  cloudformation_type_name = "AWS::AccessAnalyzer::Analyzer"
}
