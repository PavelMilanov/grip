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
	/*
		Генерирует файл hosts для ansible и генерирует group_vars.
	*/
	var hosts = "{{range .}}[{{.Vendor}}]\n{{range .Alias}}{{.Server}} ansible_host={{.Ip}}\n{{end}}{{end}}"
	exit := make(chan bool)
	var vendors []string
	for _, vendor := range model {
		vendors = append(vendors, vendor.Vendor)
	}
	go generateGroupVarsFiles(vendors, exit)

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
	<-exit
}

func generateGroupVarsFiles(vendors []string, exit chan bool) {
	/*
		Создает файлы по шаблону в директории ansible/group_vars.
	*/
	os.Chdir("ansible/group_vars")
	data := []byte("ansible_user: root\nansible_ssh_private_key_file: ./.ssh/id_rsa")
	for _, fileName := range vendors {
		os.Create(fileName)
		os.WriteFile(fileName, data, 0600)
	}
	exit <- true
}
