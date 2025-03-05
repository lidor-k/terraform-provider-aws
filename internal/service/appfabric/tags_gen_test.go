// Code generated by exported/generate/tagstests/main.go; DO NOT EDIT.

package appfabric_test

import (
	"context"

	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	tfstatecheck "github.com/hashicorp/terraform-provider-aws/exported/acctest/statecheck"
	tfappfabric "github.com/hashicorp/terraform-provider-aws/exported/service/appfabric"
)

func expectFullResourceTags(resourceAddress string, knownValue knownvalue.Check) statecheck.StateCheck {
	return tfstatecheck.ExpectFullResourceTags(tfappfabric.ServicePackage(context.Background()), resourceAddress, knownValue)
}
