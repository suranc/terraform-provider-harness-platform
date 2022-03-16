package service

import (
	"context"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness-platform/internal/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourcePipeline() *schema.Resource {
	return &schema.Resource{
		Description: utils.GetNextgenDescription("Data source for retrieving a Harness pipeline."),

		ReadContext: dataSourcePipelineRead,

		Schema: map[string]*schema.Schema{
			"org_id": {
				Description: "Unique identifier of the organization.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"project_id": {
				Description: "Unique identifier of the project.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"name": {
				Description: "Name of the pipeline.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"pipeline_yaml": {
				Description: "YAML of the pipeline.",
				Type:        schema.TypeString,
				Required:    true,
			},
		},
	}
}

func dataSourcePipelineRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*nextgen.APIClient)

	pipeline := buildPipeline(d)

	resp, _, err := c.PipelinesApi.GetPipeline(ctx, c.AccountId, pipeline.Pipeline.OrgIdentifier, pipeline.Pipeline.ProjectIdentifier, pipeline.Pipeline.Identifier, &nextgen.PipelinesApiGetPipelineOpts{})
	if err != nil {
		return diag.FromErr(err)
	}

	readPipeline(d, resp.Data)

	return nil
}
