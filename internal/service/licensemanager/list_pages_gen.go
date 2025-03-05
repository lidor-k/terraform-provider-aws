// Code generated by "exported/generate/listpages/main.go -ListOps=ListLicenseConfigurations,ListLicenseSpecificationsForResource,ListReceivedLicenses,ListDistributedGrants,ListReceivedGrants"; DO NOT EDIT.

package licensemanager

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/licensemanager"
)

func listDistributedGrantsPages(ctx context.Context, conn *licensemanager.Client, input *licensemanager.ListDistributedGrantsInput, fn func(*licensemanager.ListDistributedGrantsOutput, bool) bool) error {
	for {
		output, err := conn.ListDistributedGrants(ctx, input)
		if err != nil {
			return err
		}

		lastPage := aws.ToString(output.NextToken) == ""
		if !fn(output, lastPage) || lastPage {
			break
		}

		input.NextToken = output.NextToken
	}
	return nil
}
func listLicenseConfigurationsPages(ctx context.Context, conn *licensemanager.Client, input *licensemanager.ListLicenseConfigurationsInput, fn func(*licensemanager.ListLicenseConfigurationsOutput, bool) bool) error {
	for {
		output, err := conn.ListLicenseConfigurations(ctx, input)
		if err != nil {
			return err
		}

		lastPage := aws.ToString(output.NextToken) == ""
		if !fn(output, lastPage) || lastPage {
			break
		}

		input.NextToken = output.NextToken
	}
	return nil
}
func listLicenseSpecificationsForResourcePages(ctx context.Context, conn *licensemanager.Client, input *licensemanager.ListLicenseSpecificationsForResourceInput, fn func(*licensemanager.ListLicenseSpecificationsForResourceOutput, bool) bool) error {
	for {
		output, err := conn.ListLicenseSpecificationsForResource(ctx, input)
		if err != nil {
			return err
		}

		lastPage := aws.ToString(output.NextToken) == ""
		if !fn(output, lastPage) || lastPage {
			break
		}

		input.NextToken = output.NextToken
	}
	return nil
}
func listReceivedGrantsPages(ctx context.Context, conn *licensemanager.Client, input *licensemanager.ListReceivedGrantsInput, fn func(*licensemanager.ListReceivedGrantsOutput, bool) bool) error {
	for {
		output, err := conn.ListReceivedGrants(ctx, input)
		if err != nil {
			return err
		}

		lastPage := aws.ToString(output.NextToken) == ""
		if !fn(output, lastPage) || lastPage {
			break
		}

		input.NextToken = output.NextToken
	}
	return nil
}
func listReceivedLicensesPages(ctx context.Context, conn *licensemanager.Client, input *licensemanager.ListReceivedLicensesInput, fn func(*licensemanager.ListReceivedLicensesOutput, bool) bool) error {
	for {
		output, err := conn.ListReceivedLicenses(ctx, input)
		if err != nil {
			return err
		}

		lastPage := aws.ToString(output.NextToken) == ""
		if !fn(output, lastPage) || lastPage {
			break
		}

		input.NextToken = output.NextToken
	}
	return nil
}
