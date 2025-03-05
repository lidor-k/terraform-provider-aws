// Code generated by exported/generate/tags/main.go; DO NOT EDIT.
package dynamodb

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-provider-aws/exported/logging"
	tftags "github.com/hashicorp/terraform-provider-aws/exported/tags"
	"github.com/hashicorp/terraform-provider-aws/exported/tfresource"
	"github.com/hashicorp/terraform-provider-aws/names"
)

// updateTagsResource updates dynamodb service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func updateTagsResource(ctx context.Context, conn *dynamodb.Client, identifier string, oldTagsMap, newTagsMap any, optFns ...func(*dynamodb.Options)) error {
	oldTags := tftags.New(ctx, oldTagsMap)
	newTags := tftags.New(ctx, newTagsMap)

	ctx = tflog.SetField(ctx, logging.KeyResourceId, identifier)

	removedTags := oldTags.Removed(newTags)
	removedTags = removedTags.IgnoreSystem(names.DynamoDB)
	if len(removedTags) > 0 {
		input := dynamodb.UntagResourceInput{
			ResourceArn: aws.String(identifier),
			TagKeys:     removedTags.Keys(),
		}

		_, err := conn.UntagResource(ctx, &input, optFns...)

		if err != nil {
			return fmt.Errorf("untagging resource (%s): %w", identifier, err)
		}
	}

	updatedTags := oldTags.Updated(newTags)
	updatedTags = updatedTags.IgnoreSystem(names.DynamoDB)
	if len(updatedTags) > 0 {
		input := dynamodb.TagResourceInput{
			ResourceArn: aws.String(identifier),
			Tags:        Tags(updatedTags),
		}

		_, err := conn.TagResource(ctx, &input, optFns...)

		if err != nil {
			return fmt.Errorf("tagging resource (%s): %w", identifier, err)
		}
	}

	if len(removedTags) > 0 || len(updatedTags) > 0 {
		if err := waitTagsPropagedForResource(ctx, conn, identifier, newTags, optFns...); err != nil {
			return fmt.Errorf("waiting for resource (%s) tag propagation: %w", identifier, err)
		}
	}

	return nil
}

// waitTagsPropagedForResource waits for dynamodb service tags to be propagated.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func waitTagsPropagedForResource(ctx context.Context, conn *dynamodb.Client, id string, tags tftags.KeyValueTags, optFns ...func(*dynamodb.Options)) error {
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

		return output.ContainsAll(tags), nil
	}
	opts := tfresource.WaitOpts{
		ContinuousTargetOccurence: 2,
		MinTimeout:                1 * time.Second,
	}

	return tfresource.WaitUntil(ctx, 2*time.Minute, checkFunc, opts)
}
