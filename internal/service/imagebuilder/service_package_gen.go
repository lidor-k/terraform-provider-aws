// Code generated by exported/generate/servicepackage/main.go; DO NOT EDIT.

package imagebuilder

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/imagebuilder"
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
			Factory:  newLifecyclePolicyResource,
			TypeName: "aws_imagebuilder_lifecycle_policy",
			Name:     "Lifecycle Policy",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
	}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{
		{
			Factory:  dataSourceComponent,
			TypeName: "aws_imagebuilder_component",
			Name:     "Component",
			Tags:     &types.ServicePackageResourceTags{},
		},
		{
			Factory:  dataSourceComponents,
			TypeName: "aws_imagebuilder_components",
			Name:     "Components",
		},
		{
			Factory:  dataSourceContainerRecipe,
			TypeName: "aws_imagebuilder_container_recipe",
			Name:     "Container Recipe",
			Tags:     &types.ServicePackageResourceTags{},
		},
		{
			Factory:  dataSourceContainerRecipes,
			TypeName: "aws_imagebuilder_container_recipes",
			Name:     "Container Recipes",
		},
		{
			Factory:  dataSourceDistributionConfiguration,
			TypeName: "aws_imagebuilder_distribution_configuration",
			Name:     "Distribution Configuration",
			Tags:     &types.ServicePackageResourceTags{},
		},
		{
			Factory:  dataSourceDistributionConfigurations,
			TypeName: "aws_imagebuilder_distribution_configurations",
			Name:     "Distribution Configurations",
		},
		{
			Factory:  dataSourceImage,
			TypeName: "aws_imagebuilder_image",
			Name:     "Image",
			Tags:     &types.ServicePackageResourceTags{},
		},
		{
			Factory:  dataSourceImagePipeline,
			TypeName: "aws_imagebuilder_image_pipeline",
			Name:     "Image Pipeline",
			Tags:     &types.ServicePackageResourceTags{},
		},
		{
			Factory:  dataSourceImagePipelines,
			TypeName: "aws_imagebuilder_image_pipelines",
			Name:     "Image Pipelines",
		},
		{
			Factory:  dataSourceImageRecipe,
			TypeName: "aws_imagebuilder_image_recipe",
			Name:     "Image Recipe",
			Tags:     &types.ServicePackageResourceTags{},
		},
		{
			Factory:  dataSourceImageRecipes,
			TypeName: "aws_imagebuilder_image_recipes",
			Name:     "Image Recipes",
		},
		{
			Factory:  dataSourceInfrastructureConfiguration,
			TypeName: "aws_imagebuilder_infrastructure_configuration",
			Name:     "Infrastructure Configuration",
			Tags:     &types.ServicePackageResourceTags{},
		},
		{
			Factory:  dataSourceInfrastructureConfigurations,
			TypeName: "aws_imagebuilder_infrastructure_configurations",
			Name:     "Infrastructure Configurations",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  resourceComponent,
			TypeName: "aws_imagebuilder_component",
			Name:     "Component",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  resourceContainerRecipe,
			TypeName: "aws_imagebuilder_container_recipe",
			Name:     "Container Recipe",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  resourceDistributionConfiguration,
			TypeName: "aws_imagebuilder_distribution_configuration",
			Name:     "Distribution Configuration",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  resourceImage,
			TypeName: "aws_imagebuilder_image",
			Name:     "Image",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  resourceImagePipeline,
			TypeName: "aws_imagebuilder_image_pipeline",
			Name:     "Image Pipeline",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  resourceImageRecipe,
			TypeName: "aws_imagebuilder_image_recipe",
			Name:     "Image Recipe",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  resourceInfrastructureConfiguration,
			TypeName: "aws_imagebuilder_infrastructure_configuration",
			Name:     "Infrastructure Configuration",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  resourceWorkflow,
			TypeName: "aws_imagebuilder_workflow",
			Name:     "Workflow",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.ImageBuilder
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*imagebuilder.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws.Config))
	optFns := []func(*imagebuilder.Options){
		imagebuilder.WithEndpointResolverV2(newEndpointResolverV2()),
		withBaseEndpoint(config[names.AttrEndpoint].(string)),
		withExtraOptions(ctx, p, config),
	}

	return imagebuilder.NewFromConfig(cfg, optFns...), nil
}

// withExtraOptions returns a functional option that allows this service package to specify extra API client options.
// This option is always called after any generated options.
func withExtraOptions(ctx context.Context, sp conns.ServicePackage, config map[string]any) func(*imagebuilder.Options) {
	if v, ok := sp.(interface {
		withExtraOptions(context.Context, map[string]any) []func(*imagebuilder.Options)
	}); ok {
		optFns := v.withExtraOptions(ctx, config)

		return func(o *imagebuilder.Options) {
			for _, optFn := range optFns {
				optFn(o)
			}
		}
	}

	return func(*imagebuilder.Options) {}
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
