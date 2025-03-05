// Code generated by exported/generate/tagstests/main.go; DO NOT EDIT.

package appflow_test

import (
	"context"

	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	tfstatecheck "github.com/hashicorp/terraform-provider-aws/exported/acctest/statecheck"
	tfappflow "github.com/hashicorp/terraform-provider-aws/exported/service/appflow"
)

func expectFullResourceTags(resourceAddress string, knownValue knownvalue.Check) statecheck.StateCheck {
	return tfstatecheck.ExpectFullResourceTags(tfappflow.ServicePackage(context.Background()), resourceAddress, knownValue)
}
