// Code generated by "exported/generate/listpages/main.go -ListOps=ListTrafficPolicyVersions -Paginator=TrafficPolicyVersionMarker -- list_traffic_policy_versions_pages_gen.go"; DO NOT EDIT.

package route53

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/route53"
)

func listTrafficPolicyVersionsPages(ctx context.Context, conn *route53.Client, input *route53.ListTrafficPolicyVersionsInput, fn func(*route53.ListTrafficPolicyVersionsOutput, bool) bool) error {
	for {
		output, err := conn.ListTrafficPolicyVersions(ctx, input)
		if err != nil {
			return err
		}

		lastPage := aws.ToString(output.TrafficPolicyVersionMarker) == ""
		if !fn(output, lastPage) || lastPage {
			break
		}

		input.TrafficPolicyVersionMarker = output.TrafficPolicyVersionMarker
	}
	return nil
}
