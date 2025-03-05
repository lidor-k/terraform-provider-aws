// Code generated by exported/generate/servicepackage/main.go; DO NOT EDIT.

package computeoptimizer

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/computeoptimizer"
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
			Factory:  newEnrollmentStatusResource,
			TypeName: "aws_computeoptimizer_enrollment_status",
			Name:     "Enrollment Status",
		},
		{
			Factory:  newRecommendationPreferencesResource,
			TypeName: "aws_computeoptimizer_recommendation_preferences",
			Name:     "Recommendation Preferences",
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
	return names.ComputeOptimizer
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*computeoptimizer.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws.Config))
	optFns := []func(*computeoptimizer.Options){
		computeoptimizer.WithEndpointResolverV2(newEndpointResolverV2()),
		withBaseEndpoint(config[names.AttrEndpoint].(string)),
		withExtraOptions(ctx, p, config),
	}

	return computeoptimizer.NewFromConfig(cfg, optFns...), nil
}

// withExtraOptions returns a functional option that allows this service package to specify extra API client options.
// This option is always called after any generated options.
func withExtraOptions(ctx context.Context, sp conns.ServicePackage, config map[string]any) func(*computeoptimizer.Options) {
	if v, ok := sp.(interface {
		withExtraOptions(context.Context, map[string]any) []func(*computeoptimizer.Options)
	}); ok {
		optFns := v.withExtraOptions(ctx, config)

		return func(o *computeoptimizer.Options) {
			for _, optFn := range optFns {
				optFn(o)
			}
		}
	}

	return func(*computeoptimizer.Options) {}
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
