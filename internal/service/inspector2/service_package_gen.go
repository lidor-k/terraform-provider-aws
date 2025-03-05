// Code generated by exported/generate/servicepackage/main.go; DO NOT EDIT.

package inspector2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/inspector2"
	"github.com/hashicorp/terraform-provider-aws/exported/conns"
	"github.com/hashicorp/terraform-provider-aws/exported/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  resourceDelegatedAdminAccount,
			TypeName: "aws_inspector2_delegated_admin_account",
			Name:     "Delegated Admin Account",
		},
		{
			Factory:  ResourceEnabler,
			TypeName: "aws_inspector2_enabler",
			Name:     "Enabler",
		},
		{
			Factory:  resourceMemberAssociation,
			TypeName: "aws_inspector2_member_association",
			Name:     "Member Association",
		},
		{
			Factory:  resourceOrganizationConfiguration,
			TypeName: "aws_inspector2_organization_configuration",
			Name:     "Organization Configuration",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.Inspector2
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*inspector2.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws.Config))
	optFns := []func(*inspector2.Options){
		inspector2.WithEndpointResolverV2(newEndpointResolverV2()),
		withBaseEndpoint(config[names.AttrEndpoint].(string)),
		withExtraOptions(ctx, p, config),
	}

	return inspector2.NewFromConfig(cfg, optFns...), nil
}

// withExtraOptions returns a functional option that allows this service package to specify extra API client options.
// This option is always called after any generated options.
func withExtraOptions(ctx context.Context, sp conns.ServicePackage, config map[string]any) func(*inspector2.Options) {
	if v, ok := sp.(interface {
		withExtraOptions(context.Context, map[string]any) []func(*inspector2.Options)
	}); ok {
		optFns := v.withExtraOptions(ctx, config)

		return func(o *inspector2.Options) {
			for _, optFn := range optFns {
				optFn(o)
			}
		}
	}

	return func(*inspector2.Options) {}
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
