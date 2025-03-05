// Code generated by exported/generate/servicepackage/main.go; DO NOT EDIT.

package securitylake

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/securitylake"
	"github.com/hashicorp/terraform-provider-aws/exported/conns"
	"github.com/hashicorp/terraform-provider-aws/exported/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{
		{
			Factory:  newAWSLogSourceResource,
			TypeName: "aws_securitylake_aws_log_source",
			Name:     "AWS Log Source",
		},
		{
			Factory:  newCustomLogSourceResource,
			TypeName: "aws_securitylake_custom_log_source",
			Name:     "Custom Log Source",
		},
		{
			Factory:  newDataLakeResource,
			TypeName: "aws_securitylake_data_lake",
			Name:     "Data Lake",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  newSubscriberResource,
			TypeName: "aws_securitylake_subscriber",
			Name:     "Subscriber",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  newSubscriberNotificationResource,
			TypeName: "aws_securitylake_subscriber_notification",
			Name:     "Subscriber Notification",
		},
	}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{}
}

func (p *servicePackage) ServicePackageName() string {
	return names.SecurityLake
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*securitylake.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws.Config))
	optFns := []func(*securitylake.Options){
		securitylake.WithEndpointResolverV2(newEndpointResolverV2()),
		withBaseEndpoint(config[names.AttrEndpoint].(string)),
		withExtraOptions(ctx, p, config),
	}

	return securitylake.NewFromConfig(cfg, optFns...), nil
}

// withExtraOptions returns a functional option that allows this service package to specify extra API client options.
// This option is always called after any generated options.
func withExtraOptions(ctx context.Context, sp conns.ServicePackage, config map[string]any) func(*securitylake.Options) {
	if v, ok := sp.(interface {
		withExtraOptions(context.Context, map[string]any) []func(*securitylake.Options)
	}); ok {
		optFns := v.withExtraOptions(ctx, config)

		return func(o *securitylake.Options) {
			for _, optFn := range optFns {
				optFn(o)
			}
		}
	}

	return func(*securitylake.Options) {}
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
