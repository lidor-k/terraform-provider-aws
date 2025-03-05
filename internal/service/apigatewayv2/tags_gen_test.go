// Code generated by exported/generate/tagstests/main.go; DO NOT EDIT.

package apigatewayv2_test

import (
	"context"

	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	tfstatecheck "github.com/hashicorp/terraform-provider-aws/exported/acctest/statecheck"
	tfapigatewayv2 "github.com/hashicorp/terraform-provider-aws/exported/service/apigatewayv2"
)

func expectFullResourceTags(resourceAddress string, knownValue knownvalue.Check) statecheck.StateCheck {
	return tfstatecheck.ExpectFullResourceTags(tfapigatewayv2.ServicePackage(context.Background()), resourceAddress, knownValue)
}

func expectFullDataSourceTags(resourceAddress string, knownValue knownvalue.Check) statecheck.StateCheck {
	return tfstatecheck.ExpectFullDataSourceTags(tfapigatewayv2.ServicePackage(context.Background()), resourceAddress, knownValue)
}
