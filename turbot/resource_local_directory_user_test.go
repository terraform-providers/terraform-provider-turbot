package turbot

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/terraform-providers/terraform-provider-turbot/apiClient"
	"testing"
)

// test suites
func TestAccLocalDirectoryUser(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckLocalDirectoryUserDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLocalDirectoryUserConfig(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLocalDirectoryUserExists("turbot_local_directory_user.test_user"),
					resource.TestCheckResourceAttr(
						"turbot_local_directory_user.test_user", "title", "Kai Daguerre"),
					resource.TestCheckResourceAttr(
						"turbot_local_directory_user.test_user", "email", "kai@turbot.com"),
				),
			},
			{
				Config: testAccLocalDirectoryUserUpdateEmailConfig(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLocalDirectoryUserExists("turbot_local_directory_user.test_user"),
					resource.TestCheckResourceAttr(
						"turbot_local_directory_user.test_user", "title", "Kai Daguerre"),
					resource.TestCheckResourceAttr(
						"turbot_local_directory_user.test_user", "email", "kai2@turbot.com"),
				),
			},
			{
				Config: testAccLocalDirectoryUserUpdateTitleConfig(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLocalDirectoryUserExists("turbot_local_directory_user.test_user"),
					resource.TestCheckResourceAttr(
						"turbot_local_directory_user.test_user", "title", "Kai Daguerre2"),
					resource.TestCheckResourceAttr(
						"turbot_local_directory_user.test_user", "email", "kai@turbot.com"),
				),
			},
		},
	})
}

// configs
func testAccLocalDirectoryUserConfig() string {
	return `
resource "turbot_local_directory" "test_dir" {
	parent              = "tmod:@turbot/turbot#/"
	title               = "provider_test_directory"
	description         = "provider_test_directory"
	profile_id_template = "{{profile.email}}"
}

resource "turbot_local_directory_user" "test_user" {
	title        = "Kai Daguerre"
	email        = "kai@turbot.com"
	display_name = "Kai Daguerre"
	parent       = turbot_local_directory.test_dir.id
}
`
}

func testAccLocalDirectoryUserUpdateTitleConfig() string {
	return `
resource "turbot_local_directory" "test_dir" {
	parent              = "tmod:@turbot/turbot#/"
	title               = "provider_test_directory"
	description         = "provider_test_directory"
	profile_id_template = "{{profile.email}}"
}

resource "turbot_local_directory_user" "test_user" {
	title        = "Kai Daguerre2"
	email        = "kai@turbot.com"
	display_name = "Kai Daguerre"
	parent       = turbot_local_directory.test_dir.id
}`
}

func testAccLocalDirectoryUserUpdateEmailConfig() string {
	return `
resource "turbot_local_directory" "test_dir" {
	parent              = "tmod:@turbot/turbot#/"
	title               = "provider_test_directory"
	description         = "provider_test_directory"
	profile_id_template = "{{profile.email}}"
}

resource "turbot_local_directory_user" "test_user" {
	title        = "Kai Daguerre"
	email        = "kai2@turbot.com"
	display_name = "Kai Daguerre"
	parent       = turbot_local_directory.test_dir.id
}`
}

// helper functions
func testAccCheckLocalDirectoryUserExists(resource string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resource]
		if !ok {
			return fmt.Errorf("Not found: %s", resource)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No Record ID is set")
		}
		client := testAccProvider.Meta().(*apiClient.Client)
		_, err := client.ReadLocalDirectoryUser(rs.Primary.ID)
		if err != nil {
			return fmt.Errorf("error fetching item with resource %s. %s", resource, err)
		}
		return nil
	}
}

func testAccCheckLocalDirectoryUserDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*apiClient.Client)
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "turbot_local_directory_user" {
			_, err := client.ReadLocalDirectoryUser(rs.Primary.ID)
			if err == nil {
				return fmt.Errorf("Alert still exists")
			}
			if !apiClient.NotFoundError(err) {
				return fmt.Errorf("expected 'not found' error, got %s", err)
			}
		}
	}
	return nil
}
