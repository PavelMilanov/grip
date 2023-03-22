package extentions

import (
	"text/template"
)


host_pipeline := `
[{{.Vendor}}]
{{.Server}} ansible_host={{.Ip}}
`
