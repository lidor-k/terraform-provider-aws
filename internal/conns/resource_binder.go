package conns

import "github.com/aws/aws-sdk-go-v2/aws"

var resource_binder_aws_config aws.Config = aws.Config{}

func SetResourceBinderAWSConfig(config aws.Config) {
	resource_binder_aws_config = config
}

func ResourceBinderAWSConfig() *aws.Config {
	return &resource_binder_aws_config
}
