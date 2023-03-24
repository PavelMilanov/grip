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

func GenerateAnsibleHostsFile(model []AnsibleHost) {
	var hosts = "{{range .}}[{{.Vendor}}]\n{{range .Alias}}{{.Server}} ansible_host={{.Ip}}\n{{end}}{{end}}"

	tmpl, err := template.New("ansible_hosts").Parse(hosts)
	if err != nil {
		panic(err)
	}
	file, err := os.OpenFile("hosts", os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(file, model)
	if err != nil {
		panic(err)
	}
	defer file.Close()
}
