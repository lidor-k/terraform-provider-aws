// Code generated by exported/generate/tags/main.go; DO NOT EDIT.
package kms

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	awstypes "github.com/aws/aws-sdk-go-v2/service/kms/types"
	"github.com/hashicorp/aws-sdk-go-base/v2/tfawserr"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-provider-aws/exported/conns"
	"github.com/hashicorp/terraform-provider-aws/exported/logging"
	tftags "github.com/hashicorp/terraform-provider-aws/exported/tags"
	"github.com/hashicorp/terraform-provider-aws/exported/tfresource"
	"github.com/hashicorp/terraform-provider-aws/exported/types/option"
	"github.com/hashicorp/terraform-provider-aws/names"
)

// listTags lists kms service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func listTags(ctx context.Context, conn *kms.Client, identifier string, optFns ...func(*kms.Options)) (tftags.KeyValueTags, error) {
	input := kms.ListResourceTagsInput{
		KeyId: aws.String(identifier),
	}
	var output []awstypes.Tag

	pages := kms.NewListResourceTagsPaginator(conn, &input)
	for pages.HasMorePages() {
		page, err := pages.NextPage(ctx, optFns...)

		if tfawserr.ErrCodeEquals(err, "NotFoundException") {
			return nil, &retry.NotFoundError{
				LastError:   err,
				LastRequest: &input,
			}
		}

		if err != nil {
			return tftags.New(ctx, nil), err
		}

		for _, v := range page.Tags {
			output = append(output, v)
		}
	}

	return KeyValueTags(ctx, output), nil
}

// ListTags lists kms service tags and set them in Context.
// It is called from outside this package.
func (p *servicePackage) ListTags(ctx context.Context, meta any, identifier string) error {
	tags, err := listTags(ctx, meta.(*conns.AWSClient).KMSClient(ctx), identifier)

	if err != nil {
		return err
	}

	if inContext, ok := tftags.FromContext(ctx); ok {
		inContext.TagsOut = option.Some(tags)
	}

	return nil
}

// []*SERVICE.Tag handling

// Tags returns kms service tags.
func Tags(tags tftags.KeyValueTags) []awstypes.Tag {
	result := make([]awstypes.Tag, 0, len(tags))

	for k, v := range tags.Map() {
		tag := awstypes.Tag{
			TagKey:   aws.String(k),
			TagValue: aws.String(v),
		}

		result = append(result, tag)
	}

	return result
}

// KeyValueTags creates tftags.KeyValueTags from kms service tags.
func KeyValueTags(ctx context.Context, tags []awstypes.Tag) tftags.KeyValueTags {
	m := make(map[string]*string, len(tags))

	for _, tag := range tags {
		m[aws.ToString(tag.TagKey)] = tag.TagValue
	}

	return tftags.New(ctx, m)
}

// getTagsIn returns kms service tags from Context.
// nil is returned if there are no input tags.
func getTagsIn(ctx context.Context) []awstypes.Tag {
	if inContext, ok := tftags.FromContext(ctx); ok {
		if tags := Tags(inContext.TagsIn.UnwrapOrDefault()); len(tags) > 0 {
			return tags
		}
	}

	return nil
}

// setTagsOut sets kms service tags in Context.
func setTagsOut(ctx context.Context, tags []awstypes.Tag) {
	if inContext, ok := tftags.FromContext(ctx); ok {
		inContext.TagsOut = option.Some(KeyValueTags(ctx, tags))
	}
}

// updateTags updates kms service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func updateTags(ctx context.Context, conn *kms.Client, identifier string, oldTagsMap, newTagsMap any, optFns ...func(*kms.Options)) error {
	oldTags := tftags.New(ctx, oldTagsMap)
	newTags := tftags.New(ctx, newTagsMap)

	ctx = tflog.SetField(ctx, logging.KeyResourceId, identifier)

	removedTags := oldTags.Removed(newTags)
	removedTags = removedTags.IgnoreSystem(names.KMS)
	if len(removedTags) > 0 {
		input := kms.UntagResourceInput{
			KeyId:   aws.String(identifier),
			TagKeys: removedTags.Keys(),
		}

		_, err := conn.UntagResource(ctx, &input, optFns...)

		if err != nil {
			return fmt.Errorf("untagging resource (%s): %w", identifier, err)
		}
	}

	updatedTags := oldTags.Updated(newTags)
	updatedTags = updatedTags.IgnoreSystem(names.KMS)
	if len(updatedTags) > 0 {
		input := kms.TagResourceInput{
			KeyId: aws.String(identifier),
			Tags:  Tags(updatedTags),
		}

		_, err := conn.TagResource(ctx, &input, optFns...)

		if err != nil {
			return fmt.Errorf("tagging resource (%s): %w", identifier, err)
		}
	}

	if len(removedTags) > 0 || len(updatedTags) > 0 {
		if err := waitTagsPropagated(ctx, conn, identifier, newTags, optFns...); err != nil {
			return fmt.Errorf("waiting for resource (%s) tag propagation: %w", identifier, err)
		}
	}

	return nil
}

// UpdateTags updates kms service tags.
// It is called from outside this package.
func (p *servicePackage) UpdateTags(ctx context.Context, meta any, identifier string, oldTags, newTags any) error {
	return updateTags(ctx, meta.(*conns.AWSClient).KMSClient(ctx), identifier, oldTags, newTags)
}

// waitTagsPropagated waits for kms service tags to be propagated.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func waitTagsPropagated(ctx context.Context, conn *kms.Client, id string, tags tftags.KeyValueTags, optFns ...func(*kms.Options)) error {
	tflog.Debug(ctx, "Waiting for tag propagation", map[string]any{
		names.AttrTags: tags,
	})

	checkFunc := func() (bool, error) {
		output, err := listTags(ctx, conn, id, optFns...)

		if tfresource.NotFound(err) {
			return false, nil
		}

		if err != nil {
			return false, err
		}

		if inContext, ok := tftags.FromContext(ctx); ok {
			tags = tags.IgnoreConfig(inContext.IgnoreConfig)
			output = output.IgnoreConfig(inContext.IgnoreConfig)
		}

		return output.Equal(tags), nil
	}
	opts := tfresource.WaitOpts{
		ContinuousTargetOccurence: 5,
		MinTimeout:                1 * time.Second,
	}

	return tfresource.WaitUntil(ctx, 10*time.Minute, checkFunc, opts)
}
