package extentions

import (
	"fmt"
	"os"
	"os/exec"
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

	file, err := os.OpenFile("ansible/hosts", os.O_WRONLY, 0600)
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
	data := []byte("ansible_user: root\nansible_ssh_private_key_file: ~/.ssh/id_rsa")
	for _, fileName := range vendors {
		fileNamePath := fmt.Sprintf("ansible/group_vars/%s", fileName)
		os.WriteFile(fileNamePath, data, 0600)
	}
	exit <- true
}

func RunAnsible(command string) {
	entrypoint := fmt.Sprintf("--entrypoint=%s", command)
	cmd := exec.Command("docker", "run", "--rm", "--name=ansible-playbook", "--network=host", entrypoint, "-it", "ansible")
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Run()
}

func buildAnsibleImage() {
	cmd := exec.Command("docker", "images", "|", "awk", "'{print $1}'", "|", "tail", "-1")
	out, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	fmt.Println(out)
}
