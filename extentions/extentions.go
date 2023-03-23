package extentions

import (
	"os"
	"text/template"
)

type AnsibleHost struct {
	Vendor string
	Alias  []AnsibleAlias
}

type AnsibleAlias struct {
	Server, Ip string
}

func test(model []AnsibleHost) {
	var hosts = `
{{range .}}
[{{.Vendor}}]
{{range .Alias}}
{{.Server}} ansible_host={{.Ip}}
{{end}}
{{end}}
`

	tmpl, err := template.New("ansible_hosts").Parse(hosts)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, model)
	if err != nil {
		panic(err)
	}
}
