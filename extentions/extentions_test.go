package extentions

import (
	"testing"
)

func TestGenerateAnsibleEnvirons(t *testing.T) {
	t.Run("Generate", func(t *testing.T) {
		vendor1 := AnsibleHost{Vendor: "local", Alias: []AnsibleAlias{{"test1", "ip1"}, {"test2", "ip2"}}}
		vendor2 := AnsibleHost{Vendor: "dev", Alias: []AnsibleAlias{{"work", "127.0.0.1"}, {"dev", "127.0.0.2"}, {"preprod", "127.0.0.3"}}}
		vendor3 := AnsibleHost{Vendor: "prod", Alias: []AnsibleAlias{{"product", "10.0.0.1"}}}
		model := []AnsibleHost{vendor1, vendor2, vendor3}
		GenerateAnsibleHostsFile(model)
	})
}

func TestBuildAnsible(t *testing.T) {
	t.Run("Docker", func(t *testing.T) {
		image := buildAnsibleImage()
		if image {
			t.Log("Image find")
		} else {
			t.Log("No image find")
		}
	})
}
