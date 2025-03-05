// Code generated by exported/generate/servicepackage/main.go; DO NOT EDIT.

package servicequotas

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/servicequotas"
	"github.com/hashicorp/terraform-provider-aws/exported/conns"
	"github.com/hashicorp/terraform-provider-aws/exported/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{
		{
			Factory:  newDataSourceTemplates,
			TypeName: "aws_servicequotas_templates",
			Name:     "Templates",
		},
	}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{
		{
			Factory:  newResourceTemplate,
			TypeName: "aws_servicequotas_template",
			Name:     "Template",
		},
		{
			Factory:  newResourceTemplateAssociation,
			TypeName: "aws_servicequotas_template_association",
			Name:     "Template Association",
		},
	}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{
		{
			Factory:  DataSourceService,
			TypeName: "aws_servicequotas_service",
			Name:     "Service",
		},
		{
			Factory:  DataSourceServiceQuota,
			TypeName: "aws_servicequotas_service_quota",
			Name:     "Service Quota",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  ResourceServiceQuota,
			TypeName: "aws_servicequotas_service_quota",
			Name:     "Service Quota",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.ServiceQuotas
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*servicequotas.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws.Config))
	optFns := []func(*servicequotas.Options){
		servicequotas.WithEndpointResolverV2(newEndpointResolverV2()),
		withBaseEndpoint(config[names.AttrEndpoint].(string)),
		withExtraOptions(ctx, p, config),
	}

	return servicequotas.NewFromConfig(cfg, optFns...), nil
}

// withExtraOptions returns a functional option that allows this service package to specify extra API client options.
// This option is always called after any generated options.
func withExtraOptions(ctx context.Context, sp conns.ServicePackage, config map[string]any) func(*servicequotas.Options) {
	if v, ok := sp.(interface {
		withExtraOptions(context.Context, map[string]any) []func(*servicequotas.Options)
	}); ok {
		optFns := v.withExtraOptions(ctx, config)

		return func(o *servicequotas.Options) {
			for _, optFn := range optFns {
				optFn(o)
			}
		}
	}

	return func(*servicequotas.Options) {}
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
