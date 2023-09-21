package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testIPFirewallFilterAddress = "routeros_firewall_filter.rule"

func TestAccIPFirewallFilterTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/ip/firewall/filter", "routeros_firewall_filter"),
				Steps: []resource.TestStep{
					{
						Config: testAccIPFirewallFilterConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testIPFirewallFilterAddress),
							resource.TestCheckResourceAttr(testIPFirewallFilterAddress, "action", "accept"),
						),
					},
				},
			})

		})
	}
}

func testAccIPFirewallFilterConfig() string {
	return providerConfig + `
resource "routeros_firewall_filter" "rule" {
	action 		= "accept"
	chain   	= "forward"
	src_address = "10.0.0.1"
	dst_address = "10.0.1.1"
	dst_port 	= "443"
	protocol 	= "tcp"
}

resource "routeros_ip_firewall_filter" "testepeg" {
	action = "add-dst-to-address-list"
	address_list_timeout = "00:00:10"
	protocol = "tcp"
	tls_host = "globo"
	address_list = "teste"
	chain = "forward"
	src_address_list = "LAN"
}
`
}
