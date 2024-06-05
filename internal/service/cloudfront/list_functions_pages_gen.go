// Code generated by "internal/generate/listpages/main.go -AWSSDKVersion=2 -ListOps=ListFunctions -InputPaginator=Marker -OutputPaginator=FunctionList.NextMarker -- list_functions_pages_gen.go"; DO NOT EDIT.

package cloudfront

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront"
)

func listFunctionsPages(ctx context.Context, conn *cloudfront.Client, input *cloudfront.ListFunctionsInput, fn func(*cloudfront.ListFunctionsOutput, bool) bool) error {
	for {
		output, err := conn.ListFunctions(ctx, input)
		if err != nil {
			return err
		}

		lastPage := aws.ToString(output.FunctionList.NextMarker) == ""
		if !fn(output, lastPage) || lastPage {
			break
		}

		input.Marker = output.FunctionList.NextMarker
	}
	return nil
}
