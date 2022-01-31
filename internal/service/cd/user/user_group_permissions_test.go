package user_test

import (
	"fmt"
	"testing"

	"github.com/harness-io/harness-go-sdk/harness/utils"
	"github.com/harness-io/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/require"
)

func TestAccResourceUserGroupPermissions_AccountPermissions(t *testing.T) {

	expectedName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(12))
	resourceName := "harness_user_group_permissions.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		// CheckDestroy:      testAccUserGroupDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceUserGroupPermissions_AccountPermissions(expectedName),
				Check: resource.ComposeTestCheckFunc(
					// resource.TestCheckResourceAttr(resourceName, "name", expectedName),
					testAccUserGroupPermissionCreation(t, resourceName, expectedName),
					// func (state *terraform.State) error {

					// }
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

// func TestAccResourceUserGroup_DeleteUnderlyingResource(t *testing.T) {

// 	expectedName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(12))
// 	resourceName := "harness_user_group.test"

// 	resource.UnitTest(t, resource.TestCase{
// 		PreCheck:          func() { acctest.TestAccPreCheck(t) },
// 		ProviderFactories: acctest.ProviderFactories,
// 		Steps: []resource.TestStep{
// 			{
// 				Config: testAccResourceUserGroupAccountPermissions(expectedName),
// 				Check: resource.ComposeTestCheckFunc(
// 					resource.TestCheckResourceAttr(resourceName, "name", expectedName),
// 				),
// 			},
// 			{
// 				PreConfig: func() {
// 					acctest.TestAccConfigureProvider()
// 					c := acctest.TestAccProvider.Meta().(*api.Client)

// 					grp, err := c.CDClient.UserClient.GetUserGroupByName(expectedName)
// 					require.NoError(t, err)
// 					require.NotNil(t, grp)

// 					err = c.CDClient.UserClient.DeleteUserGroup(grp.Id)
// 					require.NoError(t, err)
// 				},
// 				Config:             testAccResourceUserGroupAccountPermissions(expectedName),
// 				PlanOnly:           true,
// 				ExpectNonEmptyPlan: true,
// 			},
// 		},
// 	})
// }

// func TestAccResourceUserGroup_AppPermissions(t *testing.T) {

// 	expectedName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(12))
// 	resourceName := "harness_user_group.test"

// 	resource.UnitTest(t, resource.TestCase{
// 		PreCheck:          func() { acctest.TestAccPreCheck(t) },
// 		ProviderFactories: acctest.ProviderFactories,
// 		CheckDestroy:      testAccUserGroupDestroy(resourceName),
// 		Steps: []resource.TestStep{
// 			{
// 				Config: testAccResourceUserGroupAppPermissions(expectedName),
// 				Check: resource.ComposeTestCheckFunc(
// 					resource.TestCheckResourceAttr(resourceName, "name", expectedName),
// 					testAccUserGroupPermissionCreation(t, resourceName, expectedName),
// 				),
// 			},
// 			{
// 				ResourceName:      resourceName,
// 				ImportState:       true,
// 				ImportStateVerify: true,
// 			},
// 		},
// 	})
// }

func testAccUserGroupPermissionCreation(t *testing.T, resourceName string, name string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		group, err := testAccGetUserGroup(resourceName, state)
		require.NoError(t, err)
		require.NotNil(t, group)
		// require.Equal(t, name, group.Name)

		return nil
	}
}

// func testAccGetUserGroup(resourceName string, state *terraform.State) (*graphql.UserGroup, error) {
// 	r := acctest.TestAccGetResource(resourceName, state)
// 	c := acctest.TestAccGetApiClientFromProvider()
// 	id := r.Primary.ID

// 	return c.CDClient.UserClient.GetUserGroupById(id)
// }

// func testAccUserGroupDestroy(resourceName string) resource.TestCheckFunc {
// 	return func(state *terraform.State) error {
// 		app, _ := testAccGetUserGroup(resourceName, state)
// 		if app != nil {
// 			return fmt.Errorf("Found user group: %s", app.Id)
// 		}

// 		return nil
// 	}
// }

// func testAccResourceUserGroupAppPermissions(name string) string {
// 	return fmt.Sprintf(`
// 		resource "harness_user_group" "test" {
// 			name = "%s"
// 			description = "my description"

// 			permissions {
// 				app_permissions {

// 					all {
// 						app_ids = []
// 						actions = ["CREATE", "READ", "UPDATE", "DELETE"]
// 					}

// 					deployment {
// 						actions = ["READ", "ROLLBACK_WORKFLOW", "EXECUTE_PIPELINE", "EXECUTE_WORKFLOW"]
// 						filters = ["NON_PRODUCTION_ENVIRONMENTS"]
// 					}

// 					deployment {
// 						actions = ["READ"]
// 						filters = ["PRODUCTION_ENVIRONMENTS"]
// 					}

// 					environment {
// 						actions = ["CREATE", "READ", "UPDATE", "DELETE"]
// 						filters = ["NON_PRODUCTION_ENVIRONMENTS"]
// 					}

// 					environment {
// 						actions = ["READ"]
// 						filters = ["PRODUCTION_ENVIRONMENTS"]
// 					}

// 					pipeline {
// 						actions = ["CREATE", "READ", "UPDATE", "DELETE"]
// 						filters = ["NON_PRODUCTION_PIPELINES"]
// 					}

// 					pipeline {
// 						actions = ["READ"]
// 						filters = ["PRODUCTION_PIPELINES"]
// 					}

// 					provisioner {
// 						actions = ["UPDATE", "DELETE"]
// 					}

// 					provisioner {
// 						actions = ["CREATE", "READ"]
// 					}

// 					service {
// 						actions = ["UPDATE", "DELETE"]
// 					}

// 					service {
// 						actions = ["UPDATE", "DELETE"]
// 					}

// 					template {
// 						actions = ["CREATE", "READ", "UPDATE", "DELETE"]
// 					}

// 					workflow {
// 						actions = ["UPDATE", "DELETE"]
// 						filters = ["NON_PRODUCTION_WORKFLOWS",]
// 					}

// 					workflow {
// 						actions = ["CREATE", "READ"]
// 						filters = ["PRODUCTION_WORKFLOWS", "WORKFLOW_TEMPLATES"]
// 					}

// 				}
// 			}
// 		}
// `, name)
// }

func testAccResourceUserGroupPermissions_AccountPermissions(name string) string {
	return fmt.Sprintf(`
		resource "harness_user_group" "test" {
			name = "%s"
		}
		
		resource "harness_user_group_permissions" "test" {
			user_group_id = harness_user_group.test.id

			account_permissions = ["ADMINISTER_OTHER_ACCOUNT_FUNCTIONS", "MANAGE_API_KEYS"]
		}
`, name)
}