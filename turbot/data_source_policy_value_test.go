package turbot

import (
	"github.com/hashicorp/terraform/helper/resource"
	"testing"
)

// todo test more policy formats: array, templated, calculated (e.g. stack source)

func TestAccPolicyValueDataSource_Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckPolicyValueConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.turbot_policy_value.test_policy", "value", "alpha.turbot.com"),
					resource.TestCheckResourceAttr(
						"data.turbot_policy_value.test_policy", "precedence", "must"),
				),
			},
		},
	})

}
func testAccCheckPolicyValueConfig() string {
	return `
data "turbot_policy_value" "test_policy" {
  resource = "tmod:@turbot/turbot#/"
  policy_type = "tmod:@turbot/turbot#/policy/types/domainName"
}
`
}
