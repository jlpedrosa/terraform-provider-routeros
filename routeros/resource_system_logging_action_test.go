package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

const testSystemLoggingEmailTask = "routeros_system_logging_action_email.action"

func TestAccSystemLoggingActions_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				Steps: []resource.TestStep{
					{
						Config: testAccSystemLoggingActionConfig(),
						Check: resource.ComposeTestCheckFunc(
							testAccCheckLoggingAction(testSystemLoggingEmailTask),
							resource.TestCheckResourceAttr(testSystemLoggingEmailTask, "name", "customemailaction"),
							resource.TestCheckResourceAttr(testSystemLoggingEmailTask, "email_to", "some@domain.com"),
							resource.TestCheckResourceAttr(testSystemLoggingEmailTask, "email_start_tls", "false"),
							resource.TestCheckResourceAttr(testSystemLoggingEmailTask, "target", "email"),
						),
					},
				},
			})

		})
	}
}

func testAccCheckLoggingAction(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("not found: %s", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no id is set")
		}

		return nil
	}
}

func testAccSystemLoggingActionConfig() string {
	return providerConfig + `
resource "routeros_system_logging_action_email" "action" {
    name = "customemailaction"
    email_to = "some@domain.com"
}
`
}
