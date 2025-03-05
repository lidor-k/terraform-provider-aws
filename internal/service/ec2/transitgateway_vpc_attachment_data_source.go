// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-aws/exported/conns"
	"github.com/hashicorp/terraform-provider-aws/exported/errs/sdkdiag"
	tftags "github.com/hashicorp/terraform-provider-aws/exported/tags"
	"github.com/hashicorp/terraform-provider-aws/exported/tfresource"
	"github.com/hashicorp/terraform-provider-aws/names"
)

// @SDKDataSource("aws_ec2_transit_gateway_vpc_attachment", name="Transit Gateway VPC Attachment")
// @Tags
// @Testing(tagsTest=false)
func dataSourceTransitGatewayVPCAttachment() *schema.Resource {
	return &schema.Resource{
		ReadWithoutTimeout: dataSourceTransitGatewayVPCAttachmentRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"appliance_mode_support": {
				Type:     schema.TypeString,
				Computed: true,
			},
			names.AttrARN: {
				Type:     schema.TypeString,
				Computed: true,
			},
			"dns_support": {
				Type:     schema.TypeString,
				Computed: true,
			},
			names.AttrFilter: customFiltersSchema(),
			names.AttrID: {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv6_support": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"security_group_referencing_support": {
				Type:     schema.TypeString,
				Computed: true,
			},
			names.AttrSubnetIDs: {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			names.AttrTransitGatewayID: {
				Type:     schema.TypeString,
				Computed: true,
			},
			names.AttrTags: tftags.TagsSchemaComputed(),
			names.AttrVPCID: {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vpc_owner_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceTransitGatewayVPCAttachmentRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	conn := meta.(*conns.AWSClient).EC2Client(ctx)

	input := &ec2.DescribeTransitGatewayVpcAttachmentsInput{}

	input.Filters = append(input.Filters, newCustomFilterList(
		d.Get(names.AttrFilter).(*schema.Set),
	)...)

	if len(input.Filters) == 0 {
		// Don't send an empty filters list; the EC2 API won't accept it.
		input.Filters = nil
	}

	if v, ok := d.GetOk(names.AttrID); ok {
		input.TransitGatewayAttachmentIds = []string{v.(string)}
	}

	transitGatewayVPCAttachment, err := findTransitGatewayVPCAttachment(ctx, conn, input)

	if err != nil {
		return sdkdiag.AppendFromErr(diags, tfresource.SingularDataSourceFindError("EC2 Transit Gateway VPC Attachment", err))
	}

	d.SetId(aws.ToString(transitGatewayVPCAttachment.TransitGatewayAttachmentId))
	d.Set("appliance_mode_support", transitGatewayVPCAttachment.Options.ApplianceModeSupport)
	vpcOwnerID := aws.ToString(transitGatewayVPCAttachment.VpcOwnerId)
	arn := arn.ARN{
		Partition: meta.(*conns.AWSClient).Partition(ctx),
		Service:   names.EC2,
		Region:    meta.(*conns.AWSClient).Region(ctx),
		AccountID: vpcOwnerID,
		Resource:  fmt.Sprintf("transit-gateway-attachment/%s", d.Id()),
	}.String()
	d.Set(names.AttrARN, arn)
	d.Set("dns_support", transitGatewayVPCAttachment.Options.DnsSupport)
	d.Set("ipv6_support", transitGatewayVPCAttachment.Options.Ipv6Support)
	d.Set("security_group_referencing_support", transitGatewayVPCAttachment.Options.SecurityGroupReferencingSupport)
	d.Set(names.AttrSubnetIDs, transitGatewayVPCAttachment.SubnetIds)
	d.Set(names.AttrTransitGatewayID, transitGatewayVPCAttachment.TransitGatewayId)
	d.Set(names.AttrVPCID, transitGatewayVPCAttachment.VpcId)
	d.Set("vpc_owner_id", vpcOwnerID)

	setTagsOut(ctx, transitGatewayVPCAttachment.Tags)

	return diags
}
