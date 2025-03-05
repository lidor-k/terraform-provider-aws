// Code generated by exported/generate/tags/main.go; DO NOT EDIT.
package resourceexplorer2

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/resourceexplorer2"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-provider-aws/exported/conns"
	"github.com/hashicorp/terraform-provider-aws/exported/logging"
	tftags "github.com/hashicorp/terraform-provider-aws/exported/tags"
	"github.com/hashicorp/terraform-provider-aws/exported/types/option"
	"github.com/hashicorp/terraform-provider-aws/names"
)

// listTags lists resourceexplorer2 service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func listTags(ctx context.Context, conn *resourceexplorer2.Client, identifier string, optFns ...func(*resourceexplorer2.Options)) (tftags.KeyValueTags, error) {
	input := resourceexplorer2.ListTagsForResourceInput{
		ResourceArn: aws.String(identifier),
	}

	output, err := conn.ListTagsForResource(ctx, &input, optFns...)

	if err != nil {
		return tftags.New(ctx, nil), err
	}

	return KeyValueTags(ctx, output.Tags), nil
}

// ListTags lists resourceexplorer2 service tags and set them in Context.
// It is called from outside this package.
func (p *servicePackage) ListTags(ctx context.Context, meta any, identifier string) error {
	tags, err := listTags(ctx, meta.(*conns.AWSClient).ResourceExplorer2Client(ctx), identifier)

	if err != nil {
		return err
	}

	if inContext, ok := tftags.FromContext(ctx); ok {
		inContext.TagsOut = option.Some(tags)
	}

	return nil
}

// map[string]string handling

// Tags returns resourceexplorer2 service tags.
func Tags(tags tftags.KeyValueTags) map[string]string {
	return tags.Map()
}

// KeyValueTags creates tftags.KeyValueTags from resourceexplorer2 service tags.
func KeyValueTags(ctx context.Context, tags map[string]string) tftags.KeyValueTags {
	return tftags.New(ctx, tags)
}

// getTagsIn returns resourceexplorer2 service tags from Context.
// nil is returned if there are no input tags.
func getTagsIn(ctx context.Context) map[string]string {
	if inContext, ok := tftags.FromContext(ctx); ok {
		if tags := Tags(inContext.TagsIn.UnwrapOrDefault()); len(tags) > 0 {
			return tags
		}
	}

	return nil
}

// setTagsOut sets resourceexplorer2 service tags in Context.
func setTagsOut(ctx context.Context, tags map[string]string) {
	if inContext, ok := tftags.FromContext(ctx); ok {
		inContext.TagsOut = option.Some(KeyValueTags(ctx, tags))
	}
}

// updateTags updates resourceexplorer2 service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func updateTags(ctx context.Context, conn *resourceexplorer2.Client, identifier string, oldTagsMap, newTagsMap any, optFns ...func(*resourceexplorer2.Options)) error {
	oldTags := tftags.New(ctx, oldTagsMap)
	newTags := tftags.New(ctx, newTagsMap)

	ctx = tflog.SetField(ctx, logging.KeyResourceId, identifier)

	removedTags := oldTags.Removed(newTags)
	removedTags = removedTags.IgnoreSystem(names.ResourceExplorer2)
	if len(removedTags) > 0 {
		input := resourceexplorer2.UntagResourceInput{
			ResourceArn: aws.String(identifier),
			TagKeys:     removedTags.Keys(),
		}

		_, err := conn.UntagResource(ctx, &input, optFns...)

		if err != nil {
			return fmt.Errorf("untagging resource (%s): %w", identifier, err)
		}
	}

	updatedTags := oldTags.Updated(newTags)
	updatedTags = updatedTags.IgnoreSystem(names.ResourceExplorer2)
	if len(updatedTags) > 0 {
		input := resourceexplorer2.TagResourceInput{
			ResourceArn: aws.String(identifier),
			Tags:        Tags(updatedTags),
		}

		_, err := conn.TagResource(ctx, &input, optFns...)

		if err != nil {
			return fmt.Errorf("tagging resource (%s): %w", identifier, err)
		}
	}

	return nil
}

// UpdateTags updates resourceexplorer2 service tags.
// It is called from outside this package.
func (p *servicePackage) UpdateTags(ctx context.Context, meta any, identifier string, oldTags, newTags any) error {
	return updateTags(ctx, meta.(*conns.AWSClient).ResourceExplorer2Client(ctx), identifier, oldTags, newTags)
}
