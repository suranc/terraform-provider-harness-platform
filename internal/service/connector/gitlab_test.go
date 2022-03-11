package connector_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness-platform/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceConnector_gitlab_Http_token(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_connector.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccConnectorDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceConnector_gitlab_http_token(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "gitlab.0.url", "https://gitlab.com/account"),
					resource.TestCheckResourceAttr(resourceName, "gitlab.0.connection_type", "Account"),
					resource.TestCheckResourceAttr(resourceName, "gitlab.0.validation_repo", "some_repo"),
					resource.TestCheckResourceAttr(resourceName, "gitlab.0.delegate_selectors.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "gitlab.0.credentials.0.http.0.username", "admin"),
					resource.TestCheckResourceAttr(resourceName, "gitlab.0.credentials.0.http.0.token_ref", "account.TEST_aws_secret_key"),
				),
			},
			{
				Config: testAccResourceConnector_gitlab_http_token(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "gitlab.0.credentials.0.http.0.username", "admin"),
					resource.TestCheckResourceAttr(resourceName, "gitlab.0.credentials.0.http.0.token_ref", "account.TEST_aws_secret_key"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccResourceConnector_gitlab_Ssh(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	resourceName := "harness_connector.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccConnectorDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceConnector_gitlab_ssh(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "gitlab.0.url", "https://gitlab.com/account"),
					resource.TestCheckResourceAttr(resourceName, "gitlab.0.connection_type", "Account"),
					resource.TestCheckResourceAttr(resourceName, "gitlab.0.validation_repo", "some_repo"),
					resource.TestCheckResourceAttr(resourceName, "gitlab.0.delegate_selectors.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "gitlab.0.credentials.0.ssh.0.ssh_key_ref", "account.test"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccResourceConnector_gitlab_token(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_connector.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccConnectorDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceConnector_gitlab_token(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "gitlab.0.url", "https://gitlab.com/account"),
					resource.TestCheckResourceAttr(resourceName, "gitlab.0.connection_type", "Account"),
					resource.TestCheckResourceAttr(resourceName, "gitlab.0.validation_repo", "some_repo"),
					resource.TestCheckResourceAttr(resourceName, "gitlab.0.delegate_selectors.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "gitlab.0.credentials.0.http.0.username", "admin"),
					resource.TestCheckResourceAttr(resourceName, "gitlab.0.api_authentication.0.token_ref", "account.TEST_aws_secret_key"),
				),
			},
			{
				Config: testAccResourceConnector_gitlab_token(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "gitlab.0.credentials.0.http.0.username", "admin"),
					resource.TestCheckResourceAttr(resourceName, "gitlab.0.credentials.0.http.0.token_ref", "account.TEST_aws_secret_key"),
					resource.TestCheckResourceAttr(resourceName, "gitlab.0.api_authentication.0.token_ref", "account.TEST_aws_secret_key"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccResourceConnector_gitlab_http_token(id string, name string) string {
	return fmt.Sprintf(`
		resource "harness_connector" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			description = "test"
			tags = ["foo:bar"]

			gitlab {
				url = "https://gitlab.com/account"
				connection_type = "Account"
				validation_repo = "some_repo"
				delegate_selectors = ["harness-delegate"]
				credentials {
					http {
						username = "admin"
						token_ref = "account.TEST_aws_secret_key"
					}
				}
			}
		}
`, id, name)
}

func testAccResourceConnector_gitlab_token(id string, name string) string {
	return fmt.Sprintf(`
		resource "harness_connector" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			description = "test"
			tags = ["foo:bar"]

			gitlab {
				url = "https://gitlab.com/account"
				connection_type = "Account"
				validation_repo = "some_repo"
				delegate_selectors = ["harness-delegate"]
				credentials {
					http {
						username = "admin"
						token_ref = "account.TEST_aws_secret_key"
					}
				}
				api_authentication {
					token_ref = "account.TEST_aws_secret_key"
				}
			}
		}
`, id, name)
}

func testAccResourceConnector_gitlab_ssh(id string, name string) string {
	return fmt.Sprintf(`
		resource "harness_connector" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			description = "test"
			tags = ["foo:bar"]

			gitlab {
				url = "https://gitlab.com/account"
				connection_type = "Account"
				validation_repo = "some_repo"
				delegate_selectors = ["harness-delegate"]
				credentials {
					ssh {
						ssh_key_ref = "account.test"
					}
				}
			}
		}
`, id, name)
}
