package tools

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/provider"
	"github.com/hashicorp/terraform-provider-aws/internal/service/ec2"
)

func ReadResource(resource_type string, id string) any {
	resource := ec2.ResourceInstance()
	data := resource.Data(nil)
	data.SetId(id)

	p, err := provider.New(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	meta := p.Meta().(*conns.AWSClient) // .EC2Client(context.Background())

	resource.ReadWithoutTimeout(context.Background(), data, meta)

	return data.State().Attributes
}
