package service

import (
	"context"
	"strings"

	"gopkg.in/yaml.v2"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness-platform/internal/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type PipelineYAML struct {
	Pipeline struct {
		Name              string `yaml:"name"`
		Identifier        string `yaml:"identifier"`
		ProjectIdentifier string `yaml:"projectIdentifier"`
		OrgIdentifier     string `yaml:"orgIdentifier"`
	}
}

func ResourcePipeline() *schema.Resource {
	return &schema.Resource{
		Description: utils.GetNextgenDescription("Resource for creating a Harness pipeline."),

		ReadContext:   resourcePipelineRead,
		UpdateContext: resourcePipelineUpdate,
		DeleteContext: resourcePipelineDelete,
		CreateContext: resourcePipelineCreate,
		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
				// <org_id>/<project_id>/<pipeline_id>
				parts := strings.Split(d.Id(), "/")
				d.Set("org_id", parts[0])
				d.Set("project_id", parts[1])
				d.SetId(parts[2])

				return []*schema.ResourceData{d}, nil
			},
		},

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

// Build PipelineYAML object from stored pipeline yaml
func buildPipeline(d *schema.ResourceData) *PipelineYAML {
	var Pipeline PipelineYAML
	yaml.Unmarshal([]byte(d.Get("pipeline_yaml").(string)), &Pipeline)

	return &Pipeline
}

// Read response from API out to the stored identifiers
func readPipeline(d *schema.ResourceData, PmsPipelineResponse *nextgen.PmsPipelineResponse) {
	d.Set("pipeline_yaml", PmsPipelineResponse.YamlPipeline)
	var Pipeline PipelineYAML

	yaml.Unmarshal([]byte(PmsPipelineResponse.YamlPipeline), &Pipeline)

	d.SetId(Pipeline.Pipeline.Identifier)
	d.Set("org_id", Pipeline.Pipeline.OrgIdentifier)
	d.Set("project_id", Pipeline.Pipeline.ProjectIdentifier)
	d.Set("name", Pipeline.Pipeline.Name)
}

func resourcePipelineRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*nextgen.APIClient)

	pipeline := buildPipeline(d)

	resp, _, err := c.PipelinesApi.GetPipeline(ctx, c.AccountId, pipeline.Pipeline.OrgIdentifier, pipeline.Pipeline.ProjectIdentifier, pipeline.Pipeline.Identifier, &nextgen.PipelinesApiGetPipelineOpts{})
	if err != nil {
		return diag.Errorf(err.(nextgen.GenericSwaggerError).Error())
	}

	readPipeline(d, resp.Data)

	return nil
}

func resourcePipelineCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*nextgen.APIClient)

	pipeline := buildPipeline(d)

	_, _, err := c.PipelinesApi.PostPipeline(ctx, d.Get("pipeline_yaml").(string), c.AccountId, pipeline.Pipeline.OrgIdentifier, pipeline.Pipeline.ProjectIdentifier, &nextgen.PipelinesApiPostPipelineOpts{})
	if err != nil {
		return diag.Errorf(err.(nextgen.GenericSwaggerError).Error())
	}

	resp, _, err := c.PipelinesApi.GetPipeline(ctx, c.AccountId, pipeline.Pipeline.OrgIdentifier, pipeline.Pipeline.ProjectIdentifier, pipeline.Pipeline.Identifier, &nextgen.PipelinesApiGetPipelineOpts{})
	if err != nil {
		return diag.Errorf(err.(nextgen.GenericSwaggerError).Error())
	}

	readPipeline(d, resp.Data)

	return nil
}

func resourcePipelineUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*nextgen.APIClient)

	pipeline := buildPipeline(d)

	_, _, err := c.PipelinesApi.UpdatePipeline(ctx, d.Get("pipeline_yaml").(string), c.AccountId, pipeline.Pipeline.OrgIdentifier, pipeline.Pipeline.ProjectIdentifier, pipeline.Pipeline.Identifier, &nextgen.PipelinesApiUpdatePipelineOpts{})
	if err != nil {
		return diag.Errorf(err.(nextgen.GenericSwaggerError).Error())
	}

	resp, _, err := c.PipelinesApi.GetPipeline(ctx, c.AccountId, pipeline.Pipeline.OrgIdentifier, pipeline.Pipeline.ProjectIdentifier, pipeline.Pipeline.Identifier, &nextgen.PipelinesApiGetPipelineOpts{})
	if err != nil {
		return diag.Errorf(err.(nextgen.GenericSwaggerError).Error())
	}

	readPipeline(d, resp.Data)

	return nil
}

func resourcePipelineDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*nextgen.APIClient)

	pipeline := buildPipeline(d)

	_, _, err := c.PipelinesApi.DeletePipeline(ctx, c.AccountId, pipeline.Pipeline.OrgIdentifier, pipeline.Pipeline.ProjectIdentifier, pipeline.Pipeline.Identifier, &nextgen.PipelinesApiDeletePipelineOpts{})
	if err != nil {
		return diag.Errorf(err.(nextgen.GenericSwaggerError).Error())
	}

	d.SetId("")

	return nil
}
