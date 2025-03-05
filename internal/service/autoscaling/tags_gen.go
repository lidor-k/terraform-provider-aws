// Code generated by exported/generate/tags/main.go; DO NOT EDIT.
package autoscaling

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling"
	awstypes "github.com/aws/aws-sdk-go-v2/service/autoscaling/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-aws/exported/conns"
	"github.com/hashicorp/terraform-provider-aws/exported/logging"
	tfslices "github.com/hashicorp/terraform-provider-aws/exported/slices"
	tftags "github.com/hashicorp/terraform-provider-aws/exported/tags"
	"github.com/hashicorp/terraform-provider-aws/exported/tfresource"
	"github.com/hashicorp/terraform-provider-aws/exported/types/option"
	"github.com/hashicorp/terraform-provider-aws/names"
)

// findTag fetches an individual autoscaling service tag for a resource.
// Returns whether the key value and any errors. A NotFoundError is used to signal that no value was found.
// This function will optimise the handling over listTags, if possible.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func findTag(ctx context.Context, conn *autoscaling.Client, identifier, resourceType, key string, optFns ...func(*autoscaling.Options)) (*tftags.TagData, error) {
	input := autoscaling.DescribeTagsInput{
		Filters: []awstypes.Filter{
			{
				Name:   aws.String("auto-scaling-group"),
				Values: []string{identifier},
			},
			{
				Name:   aws.String(names.AttrKey),
				Values: []string{key},
			},
		},
	}

	output, err := conn.DescribeTags(ctx, &input, optFns...)

	if err != nil {
		return nil, err
	}

	listTags := KeyValueTags(ctx, output.Tags, identifier, resourceType)

	if !listTags.KeyExists(key) {
		return nil, tfresource.NewEmptyResultError(nil)
	}

	return listTags.KeyTagData(key), nil
}

// listTags lists autoscaling service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func listTags(ctx context.Context, conn *autoscaling.Client, identifier, resourceType string, optFns ...func(*autoscaling.Options)) (tftags.KeyValueTags, error) {
	input := autoscaling.DescribeTagsInput{
		Filters: []awstypes.Filter{
			{
				Name:   aws.String("auto-scaling-group"),
				Values: []string{identifier},
			},
		},
	}
	var output []awstypes.TagDescription

	pages := autoscaling.NewDescribeTagsPaginator(conn, &input)
	for pages.HasMorePages() {
		page, err := pages.NextPage(ctx, optFns...)

		if err != nil {
			return tftags.New(ctx, nil), err
		}

		for _, v := range page.Tags {
			output = append(output, v)
		}
	}

	return KeyValueTags(ctx, output, identifier, resourceType), nil
}

// ListTags lists autoscaling service tags and set them in Context.
// It is called from outside this package.
func (p *servicePackage) ListTags(ctx context.Context, meta any, identifier, resourceType string) error {
	tags, err := listTags(ctx, meta.(*conns.AWSClient).AutoScalingClient(ctx), identifier, resourceType)

	if err != nil {
		return err
	}

	if inContext, ok := tftags.FromContext(ctx); ok {
		inContext.TagsOut = option.Some(tags)
	}

	return nil
}

// []*SERVICE.Tag handling

// listOfMap returns a list of autoscaling tags in a flattened map.
//
// Compatible with setting Terraform state for strongly typed configuration blocks.
//
// This function strips tag resource identifier and type. Generally, this is
// the desired behavior so the tag schema does not require those attributes.
func listOfMap(tags tftags.KeyValueTags) []any {
	return tfslices.ApplyToAll(tags.Keys(), func(key string) any {
		return map[string]any{
			names.AttrKey:         key,
			names.AttrValue:       aws.ToString(tags.KeyValue(key)),
			"propagate_at_launch": aws.ToBool(tags.KeyAdditionalBoolValue(key, "PropagateAtLaunch")),
		}
	})
}

// Tags returns autoscaling service tags.
func Tags(tags tftags.KeyValueTags) []awstypes.Tag {
	var result []awstypes.Tag

	for _, key := range tags.Keys() {
		tag := awstypes.Tag{
			Key:               aws.String(key),
			Value:             tags.KeyValue(key),
			ResourceId:        tags.KeyAdditionalStringValue(key, "ResourceId"),
			ResourceType:      tags.KeyAdditionalStringValue(key, "ResourceType"),
			PropagateAtLaunch: tags.KeyAdditionalBoolValue(key, "PropagateAtLaunch"),
		}

		result = append(result, tag)
	}

	return result
}

// KeyValueTags creates tftags.KeyValueTags from autoscaling service tags.
//
// Accepts the following types:
//   - []awstypes.Tag
//   - []awstypes.TagDescription
//   - []any (Terraform TypeList configuration block compatible)
//   - *schema.Set (Terraform TypeSet configuration block compatible)
func KeyValueTags(ctx context.Context, tags any, identifier, resourceType string) tftags.KeyValueTags {
	switch tags := tags.(type) {
	case []awstypes.Tag:
		m := make(map[string]*tftags.TagData, len(tags))

		for _, tag := range tags {
			tagData := &tftags.TagData{
				Value: tag.Value,
			}

			tagData.AdditionalBoolFields = make(map[string]*bool)
			tagData.AdditionalBoolFields["PropagateAtLaunch"] = tag.PropagateAtLaunch
			tagData.AdditionalStringFields = make(map[string]*string)
			tagData.AdditionalStringFields["ResourceId"] = &identifier
			tagData.AdditionalStringFields["ResourceType"] = &resourceType

			m[aws.ToString(tag.Key)] = tagData
		}

		return tftags.New(ctx, m)
	case []awstypes.TagDescription:
		m := make(map[string]*tftags.TagData, len(tags))

		for _, tag := range tags {
			tagData := &tftags.TagData{
				Value: tag.Value,
			}
			tagData.AdditionalBoolFields = make(map[string]*bool)
			tagData.AdditionalBoolFields["PropagateAtLaunch"] = tag.PropagateAtLaunch
			tagData.AdditionalStringFields = make(map[string]*string)
			tagData.AdditionalStringFields["ResourceId"] = &identifier
			tagData.AdditionalStringFields["ResourceType"] = &resourceType

			m[aws.ToString(tag.Key)] = tagData
		}

		return tftags.New(ctx, m)
	case *schema.Set:
		return KeyValueTags(ctx, tags.List(), identifier, resourceType)
	case []any:
		result := make(map[string]*tftags.TagData)

		for _, tfMapRaw := range tags {
			tfMap, ok := tfMapRaw.(map[string]any)

			if !ok {
				continue
			}

			key, ok := tfMap[names.AttrKey].(string)

			if !ok {
				continue
			}

			tagData := &tftags.TagData{}

			if v, ok := tfMap[names.AttrValue].(string); ok {
				tagData.Value = &v
			}

			tagData.AdditionalBoolFields = make(map[string]*bool)
			if v, ok := tfMap["propagate_at_launch"].(bool); ok {
				tagData.AdditionalBoolFields["PropagateAtLaunch"] = &v
			}

			tagData.AdditionalStringFields = make(map[string]*string)
			tagData.AdditionalStringFields["ResourceId"] = &identifier
			tagData.AdditionalStringFields["ResourceType"] = &resourceType

			result[key] = tagData
		}

		return tftags.New(ctx, result)
	default:
		return tftags.New(ctx, nil)
	}
}

// getTagsIn returns autoscaling service tags from Context.
// nil is returned if there are no input tags.
func getTagsIn(ctx context.Context) []awstypes.Tag {
	if inContext, ok := tftags.FromContext(ctx); ok {
		if tags := Tags(inContext.TagsIn.UnwrapOrDefault()); len(tags) > 0 {
			return tags
		}
	}

	return nil
}

// setTagsOut sets autoscaling service tags in Context.
func setTagsOut(ctx context.Context, tags any, identifier, resourceType string) {
	if inContext, ok := tftags.FromContext(ctx); ok {
		inContext.TagsOut = option.Some(KeyValueTags(ctx, tags, identifier, resourceType))
	}
}

// updateTags updates autoscaling service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func updateTags(ctx context.Context, conn *autoscaling.Client, identifier, resourceType string, oldTagsSet, newTagsSet any, optFns ...func(*autoscaling.Options)) error {
	oldTags := KeyValueTags(ctx, oldTagsSet, identifier, resourceType)
	newTags := KeyValueTags(ctx, newTagsSet, identifier, resourceType)

	ctx = tflog.SetField(ctx, logging.KeyResourceId, identifier)

	removedTags := oldTags.Removed(newTags)
	removedTags = removedTags.IgnoreSystem(names.AutoScaling)
	if len(removedTags) > 0 {
		input := autoscaling.DeleteTagsInput{
			Tags: Tags(removedTags),
		}

		_, err := conn.DeleteTags(ctx, &input, optFns...)

		if err != nil {
			return fmt.Errorf("untagging resource (%s): %w", identifier, err)
		}
	}

	updatedTags := oldTags.Updated(newTags)
	updatedTags = updatedTags.IgnoreSystem(names.AutoScaling)
	if len(updatedTags) > 0 {
		input := autoscaling.CreateOrUpdateTagsInput{
			Tags: Tags(updatedTags),
		}

		_, err := conn.CreateOrUpdateTags(ctx, &input, optFns...)

		if err != nil {
			return fmt.Errorf("tagging resource (%s): %w", identifier, err)
		}
	}

	return nil
}

// UpdateTags updates autoscaling service tags.
// It is called from outside this package.
func (p *servicePackage) UpdateTags(ctx context.Context, meta any, identifier, resourceType string, oldTags, newTags any) error {
	return updateTags(ctx, meta.(*conns.AWSClient).AutoScalingClient(ctx), identifier, resourceType, oldTags, newTags)
}
