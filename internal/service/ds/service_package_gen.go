// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package ds

import (
	"context"

	aws_sdkv2 "github.com/aws/aws-sdk-go-v2/aws"
	directoryservice_sdkv2 "github.com/aws/aws-sdk-go-v2/service/directoryservice"
	aws_sdkv1 "github.com/aws/aws-sdk-go/aws"
	session_sdkv1 "github.com/aws/aws-sdk-go/aws/session"
	directoryservice_sdkv1 "github.com/aws/aws-sdk-go/service/directoryservice"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{
		{
			Factory: newResourceTrust,
		},
	}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{
		{
			Factory:  DataSourceDirectory,
			TypeName: "aws_directory_service_directory",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  ResourceConditionalForwarder,
			TypeName: "aws_directory_service_conditional_forwarder",
		},
		{
			Factory:  ResourceDirectory,
			TypeName: "aws_directory_service_directory",
			Name:     "Directory",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  ResourceLogSubscription,
			TypeName: "aws_directory_service_log_subscription",
		},
		{
			Factory:  ResourceRadiusSettings,
			TypeName: "aws_directory_service_radius_settings",
		},
		{
			Factory:  ResourceRegion,
			TypeName: "aws_directory_service_region",
			Name:     "Region",
			Tags:     &types.ServicePackageResourceTags{},
		},
		{
			Factory:  ResourceSharedDirectory,
			TypeName: "aws_directory_service_shared_directory",
		},
		{
			Factory:  ResourceSharedDirectoryAccepter,
			TypeName: "aws_directory_service_shared_directory_accepter",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.DS
}

// NewConn returns a new AWS SDK for Go v1 client for this service package's AWS API.
func (p *servicePackage) NewConn(ctx context.Context, config map[string]any) (*directoryservice_sdkv1.DirectoryService, error) {
	sess := config[names.AttrSession].(*session_sdkv1.Session)

	return directoryservice_sdkv1.New(sess.Copy(&aws_sdkv1.Config{Endpoint: aws_sdkv1.String(config[names.AttrEndpoint].(string))})), nil
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*directoryservice_sdkv2.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws_sdkv2.Config))

	return directoryservice_sdkv2.NewFromConfig(cfg, func(o *directoryservice_sdkv2.Options) {
		if endpoint := config[names.AttrEndpoint].(string); endpoint != "" {
			o.BaseEndpoint = aws_sdkv2.String(endpoint)
		}
	}), nil
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
