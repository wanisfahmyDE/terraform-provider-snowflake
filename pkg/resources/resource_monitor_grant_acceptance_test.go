package resources_test

import (
	"fmt"
	"testing"

	acc "github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/acceptance"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
)

func TestAcc_ResourceMonitor_defaults(t *testing.T) {
	wName := acc.TestClient().Ids.Alpha()
	roleName := acc.TestClient().Ids.Alpha()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: acc.TestAccProtoV6ProviderFactories,
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.RequireAbove(tfversion.Version1_5_0),
		},
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		CheckDestroy: nil,
		Steps: []resource.TestStep{
			{
				Config: resourceMonitorGrantConfig(wName, roleName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("snowflake_resource_monitor_grant.test", "monitor_name", wName),
					resource.TestCheckResourceAttr("snowflake_resource_monitor_grant.test", "privilege", "MONITOR"),
				),
			},
			// IMPORT
			{
				ResourceName:      "snowflake_resource_monitor_grant.test",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"enable_multiple_grants", // feature flag attribute not defined in Snowflake, can't be imported
				},
			},
		},
	})
}

func resourceMonitorGrantConfig(n, role string) string {
	return fmt.Sprintf(`

resource "snowflake_resource_monitor" "test" {
  name      = "%v"
}

resource "snowflake_role" "test" {
  name = "%v"
}

resource "snowflake_resource_monitor_grant" "test" {
  monitor_name = snowflake_resource_monitor.test.name
  roles          = [snowflake_role.test.name]
}
`, n, role)
}
