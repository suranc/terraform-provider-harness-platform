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
			"identifier": {
				Description: "Unique identifier of the pipeline.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"org_id": {
				Description: "Unique identifier of the organization.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"project_id": {
				Description: "Unique identifier of the project.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"name": {
				Description: "Name of the pipeline.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"pipeline_yaml": {
				Description: "YAML of the pipeline.",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}
}

func dataSourcePipelineRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*nextgen.APIClient)

	resp, _, err := c.PipelinesApi.GetPipeline(ctx,
		c.AccountId,
		d.Get("org_id").(string),
		d.Get("project_id").(string),
		d.Id(),
		&nextgen.PipelinesApiGetPipelineOpts{},
	)

	if err != nil {
		e := err.(nextgen.GenericSwaggerError)
		if e.Code() == nextgen.ErrorCodes.ResourceNotFound {
			d.SetId("")
			d.MarkNewResource()
			return nil
		}
		return diag.Errorf(err.(nextgen.GenericSwaggerError).Error())
	}

	readPipeline(d, resp.Data)

	return nil
}
