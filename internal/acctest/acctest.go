package acctest

import (
	"context"
	"fmt"
	"sync"
	"testing"

	"github.com/harness/harness-go-sdk/harness/helpers"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness-platform/internal/provider"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const (
	TestAccSecretFileId      = "2WnPVgLGSZW6KbApZuxeaw"
	TestAccDefaultUsageScope = `
	usage_scope {
		environment_filter_type = "NON_PRODUCTION_ENVIRONMENTS"
	}
`
)

func TestAccConfigureProvider() {
	TestAccProviderConfigure.Do(func() {
		TestAccProvider = provider.Provider("dev")()

		config := map[string]interface{}{
			"endpoint":   helpers.EnvVars.Endpoint.GetWithDefault(utils.BaseUrl),
			"account_id": helpers.EnvVars.AccountId.Get(),
			"api_key":    helpers.EnvVars.ApiKey.Get(),
			"ng_api_key": helpers.EnvVars.NGApiKey.Get(),
		}

		TestAccProvider.Configure(context.Background(), terraform.NewResourceConfigRaw(config))
	})
}

func TestAccPreCheck(t *testing.T) {
	TestAccConfigureProvider()
}

var TestAccProvider *schema.Provider
var TestAccProviderConfigure sync.Once

func TestAccGetResource(resourceName string, state *terraform.State) *terraform.ResourceState {
	rm := state.RootModule()
	return rm.Resources[resourceName]
}

func TestAccGetApiClientFromProvider() *nextgen.APIClient {
	return TestAccProvider.Meta().(*nextgen.APIClient)
}

// providerFactories are used to instantiate a provider during acceptance testing.
// The factory function will be invoked for every Terraform CLI command executed
// to create a provider server to which the CLI can reattach.
var ProviderFactories = map[string]func() (*schema.Provider, error){
	"harness": func() (*schema.Provider, error) {
		return provider.Provider("dev")(), nil
	},
}

func TestAccResourceAwsCloudProvider(name string) string {
	return fmt.Sprintf(`
	data "harness_secret_manager" "default" {
		default = true
	}

	resource "harness_encrypted_text" "aws_access_key" {
		name = "%[1]s_access_key"
		value = "%[2]s"
		secret_manager_id = data.harness_secret_manager.default.id
	}

	resource "harness_encrypted_text" "aws_secret_key" {
		name = "%[1]s_secret_key"
		value = "%[3]s"
		secret_manager_id = data.harness_secret_manager.default.id
	}
	
	resource "harness_cloudprovider_aws" "test" {
		name = "%[1]s"
		access_key_id_secret_name = harness_encrypted_text.aws_access_key.name
		secret_access_key_secret_name = harness_encrypted_text.aws_secret_key.name
	}	
`, name, helpers.TestEnvVars.AwsAccessKeyId.Get(), helpers.TestEnvVars.AwsSecretAccessKey.Get())
}

func TestAccResourceInfraDefEnvironment(name string) string {
	return fmt.Sprintf(`
		resource "harness_application" "test" {
			name = "%[1]s"
		}

		resource "harness_environment" "test" {
			name = "%[1]s"
			app_id = harness_application.test.id
			type = "NON_PROD"
		}
`, name)
}

func TestAccResourceGitConnector_default(name string) string {

	return fmt.Sprintf(`
		data "harness_secret_manager" "test" {
			default = true
		}

		resource "harness_encrypted_text" "test" {
			name 							= "%[1]s"
			value 					  = "foo"
			secret_manager_id = data.harness_secret_manager.test.id
		}

		resource "harness_git_connector" "test" {
			name = "%[1]s"
			url = "https://github.com/micahlmartin/harness-demo"
			branch = "master"
			generate_webhook_url = true
			password_secret_id = harness_encrypted_text.test.id
			url_type = "REPO"
			username = "someuser"
		}	
`, name)
}
