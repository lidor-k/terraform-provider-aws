// Code generated by exported/generate/servicepackage/main.go; DO NOT EDIT.

package identitystore

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/identitystore"
	"github.com/hashicorp/terraform-provider-aws/exported/conns"
	"github.com/hashicorp/terraform-provider-aws/exported/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{
		{
			Factory:  newGroupsDataSource,
			TypeName: "aws_identitystore_groups",
			Name:     "Groups",
		},
	}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{
		{
			Factory:  dataSourceGroup,
			TypeName: "aws_identitystore_group",
			Name:     "Group",
		},
		{
			Factory:  dataSourceUser,
			TypeName: "aws_identitystore_user",
			Name:     "User",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  resourceGroup,
			TypeName: "aws_identitystore_group",
			Name:     "Group",
		},
		{
			Factory:  resourceGroupMembership,
			TypeName: "aws_identitystore_group_membership",
			Name:     "Group Membership",
		},
		{
			Factory:  resourceUser,
			TypeName: "aws_identitystore_user",
			Name:     "User",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.IdentityStore
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*identitystore.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws.Config))
	optFns := []func(*identitystore.Options){
		identitystore.WithEndpointResolverV2(newEndpointResolverV2()),
		withBaseEndpoint(config[names.AttrEndpoint].(string)),
		withExtraOptions(ctx, p, config),
	}

	return identitystore.NewFromConfig(cfg, optFns...), nil
}

// withExtraOptions returns a functional option that allows this service package to specify extra API client options.
// This option is always called after any generated options.
func withExtraOptions(ctx context.Context, sp conns.ServicePackage, config map[string]any) func(*identitystore.Options) {
	if v, ok := sp.(interface {
		withExtraOptions(context.Context, map[string]any) []func(*identitystore.Options)
	}); ok {
		optFns := v.withExtraOptions(ctx, config)

		return func(o *identitystore.Options) {
			for _, optFn := range optFns {
				optFn(o)
			}
		}
	}

	return func(*identitystore.Options) {}
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
