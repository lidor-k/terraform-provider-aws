package tools

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-aws/internal/service/ec2"
	"github.com/hashicorp/terraform-provider-aws/internal/service/rds"
	"github.com/hashicorp/terraform-provider-aws/internal/service/s3"
)

var resources = map[string]func() *schema.Resource{
	"instance":          ec2.ResourceInstance,
	"vpc":               ec2.ResourceVPC,
	"subnet":            ec2.ResourceSubnet,
	"network-interface": ec2.ResourceNetworkInterface,
	"internet-gateway":  ec2.ResourceInternetGateway,
	"security-group":    ec2.ResourceSecurityGroup,
	"volume":            ec2.ResourceEBSVolume,
	"snapshot":          ec2.ResourceEBSVolumeSnapshot,
	"dhcp-options":      ec2.ResourceVPCDHCPOptions,
	"db":                rds.ResourceInstance,
	"db-snapshot":       rds.ResourceSnapshot,
	"cluster":           rds.ResourceCluster,
	"cluster-snapshot":  rds.ResourceClusterSnapshot,
	"bucket":            s3.ResourceBucket,
}
