package resources_test

import (
	"fmt"
	"testing"

	acc "github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/acceptance"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
)

func testRolesAndShares(t *testing.T, path string, roles []string) func(*terraform.State) error {
	t.Helper()
	return func(state *terraform.State) error {
		is := state.RootModule().Resources[path].Primary

		if c, ok := is.Attributes["roles.#"]; !ok || mustParseInt(t, c) != int64(len(roles)) {
			return fmt.Errorf("expected roles.# to equal %d but got %s", len(roles), c)
		}
		r, err := extractList(is.Attributes, "roles")
		if err != nil {
			return err
		}

		// TODO case sensitive?
		if !listSetEqual(roles, r) {
			return fmt.Errorf("expected roles %#v but got %#v", roles, r)
		}

		return nil
	}
}

func TestAcc_DatabaseGrant(t *testing.T) {
	roleName := acc.TestClient().Ids.Alpha()
	shareName := acc.TestClient().Ids.Alpha()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: acc.TestAccProtoV6ProviderFactories,
		PreCheck:                 func() { acc.TestAccPreCheck(t) },
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.RequireAbove(tfversion.Version1_5_0),
		},
		CheckDestroy: nil,
		Steps: []resource.TestStep{
			{
				Config: databaseGrantConfig(roleName, shareName, acc.TestDatabaseName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("snowflake_database_grant.test", "database_name", acc.TestDatabaseName),
					resource.TestCheckResourceAttr("snowflake_database_grant.test", "privilege", "USAGE"),
					resource.TestCheckResourceAttr("snowflake_database_grant.test", "roles.#", "1"),
					resource.TestCheckResourceAttr("snowflake_database_grant.test", "shares.#", "1"),
					resource.TestCheckResourceAttr("snowflake_database_grant.test", "shares.#", "1"),
					testRolesAndShares(t, "snowflake_database_grant.test", []string{roleName}),
				),
			},
			// IMPORT
			{
				PreConfig:         func() { fmt.Println("[DEBUG] IMPORT") },
				ResourceName:      "snowflake_database_grant.test",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"enable_multiple_grants", // feature flag attribute not defined in Snowflake, can't be imported
				},
			},
		},
	})
}

// TODO(el): fix this test
// func TestAccDatabaseGrant_dbNotExists(t *testing.T) {
// 	dbName := acc.TestClient().Ids.Alpha()
// 	roleName := acc.TestClient().Ids.Alpha()

// 	resource.Test(t, resource.TestCase{
// 		ProtoV6ProviderFactories: acc.TestAccProtoV6ProviderFactories,
//		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
//			tfversion.RequireAbove(tfversion.Version1_5_0),
//		},
// 		Steps: []resource.TestStep{
// 			{
// 				// Note the DB we're trying to grant to doesn't exist
// 				// This tests we don't error out, but do delete remote state
// 				Config: fmt.Sprintf(`
// resource "snowflake_database_grant" "test" {
// 	database_name = "%v"
//   roles         = ["%v"]
// }`, dbName, roleName),
// 				ResourceName: "snowflake_database_grant.test",
// 				ImportStateId: ,
// 				Check: resource.ComposeTestCheckFunc(
// 					func(state *terraform.State) error {
// 						id := state.RootModule().Resources["snowflake_database_grant.test"].Primary.ID
// 						if id != "" {
// 							return errors.Errorf("Expected empty ID but got %s", id)
// 						}
// 						return nil
// 					},
// 				),
// 			},
// 		},
// 	})
// }

func databaseGrantConfig(role, share, databaseName string) string {
	return fmt.Sprintf(`
resource "snowflake_role" "test" {
  name = "%v"
}

resource "snowflake_share" "test" {
  name = "%v"
}

resource "snowflake_database_grant" "test" {
  database_name = "%s"
  roles         = [snowflake_role.test.name]
  shares        = [snowflake_share.test.name]
}
`, role, share, databaseName)
}
