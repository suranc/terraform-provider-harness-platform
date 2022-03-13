package provider

import (
	"context"
	"fmt"

	"github.com/harness/harness-go-sdk/harness"
	"github.com/harness/harness-go-sdk/harness/helpers"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness-platform/internal/service"
	"github.com/harness/terraform-provider-harness-platform/internal/service/connector"
	"github.com/hashicorp/go-cleanhttp"
	"github.com/hashicorp/go-retryablehttp"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/logging"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	// Set descriptions to support markdown syntax, this will be used in document generation
	// and the language server.
	schema.DescriptionKind = schema.StringMarkdown

	// Customize the content of descriptions when output. For example you can add defaults on
	// to the exported descriptions if present.
	// schema.SchemaDescriptionBuilder = func(s *schema.Schema) string {
	// 	desc := s.Description
	// 	if s.Default != nil {
	// 		desc += fmt.Sprintf(" Defaults to `%v`.", s.Default)
	// 	}
	// 	return strings.TrimSpace(desc)
	// }
}

func Provider(version string) func() *schema.Provider {
	return func() *schema.Provider {
		p := &schema.Provider{
			Schema: map[string]*schema.Schema{
				"endpoint": {
					Description: fmt.Sprintf("The URL of the Harness API endpoint. The default is `https://app.harness.io/gateway`. This can also be set using the `%s` environment variable.", helpers.EnvVars.Endpoint.String()),
					Type:        schema.TypeString,
					Required:    true,
					DefaultFunc: schema.EnvDefaultFunc(helpers.EnvVars.Endpoint.String(), utils.BaseUrl),
				},
				"account_id": {
					Description: fmt.Sprintf("The Harness account id. This can also be set using the `%s` environment variable.", helpers.EnvVars.AccountId.String()),
					Type:        schema.TypeString,
					Required:    true,
					DefaultFunc: schema.EnvDefaultFunc(helpers.EnvVars.AccountId.String(), nil),
				},
				"api_key": {
					Description: fmt.Sprintf("The Harness API key. This can also be set using the `%s` environment variable.", helpers.EnvVars.ApiKey.String()),
					Type:        schema.TypeString,
					Optional:    true,
					DefaultFunc: schema.EnvDefaultFunc(helpers.EnvVars.ApiKey.String(), nil),
				},
			},
			DataSourcesMap: map[string]*schema.Resource{
				"harness_connector":    connector.DataSourceConnector(),
				"harness_current_user": service.DataSourceCurrentUser(),
				"harness_organization": service.DataSourceOrganization(),
				"harness_project":      service.DataSourceProject(),
				"harness_pipeline":     service.DataSourcePipeline(),
			},
			ResourcesMap: map[string]*schema.Resource{
				"harness_connector":    connector.ResourceConnector(),
				"harness_organization": service.ResourceOrganization(),
				"harness_project":      service.ResourceProject(),
				"harness_pipeline":     service.ResourcePipeline(),
			},
		}

		p.ConfigureContextFunc = configure(version, p)

		return p
	}
}

func getHttpClient() *retryablehttp.Client {
	httpClient := retryablehttp.NewClient()
	httpClient.HTTPClient.Transport = logging.NewTransport(harness.SDKName, cleanhttp.DefaultPooledClient().Transport)
	httpClient.RetryMax = 10
	return httpClient
}

// Setup the client for interacting with the Harness API
func configure(version string, p *schema.Provider) func(context.Context, *schema.ResourceData) (interface{}, diag.Diagnostics) {
	return func(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {

		client := nextgen.NewAPIClient(&nextgen.Configuration{
			AccountId: d.Get("account_id").(string),
			BasePath:  d.Get("endpoint").(string),
			DefaultHeader: map[string]string{
				helpers.HTTPHeaders.ApiKey.String(): d.Get("api_key").(string),
			},
			UserAgent:    fmt.Sprintf("terraform-provider-harness-platform-%s", version),
			HTTPClient:   getHttpClient(),
			DebugLogging: logging.IsDebugOrHigher(),
		})

		return client, nil
	}
}
