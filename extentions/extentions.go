package extentions

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"text/template"

	"github.com/PavelMilanov/grip/text"
)

type AnsibleHost struct {
	Vendor string
	Alias  []AnsibleAlias
}

type AnsibleAlias struct {
	Server, Ip string
}

func generateAnsibleHostsFile(model []AnsibleHost) {
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

	file, err := os.OpenFile("extentions/ansible/hosts", os.O_WRONLY, 0600)
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
	data := []byte("ansible_user: root\nansible_ssh_private_key_file: /.ssh/id_rsa")
	for _, fileName := range vendors {
		fileNamePath := fmt.Sprintf("extentions/ansible/group_vars/%s", fileName)
		os.WriteFile(fileNamePath, data, 0600)
	}
	exit <- true
}

func BuildAnsible(mode string) {
	if mode == "local" {
		vendors := []AnsibleHost{}
		files, err := ioutil.ReadDir("configs")
		if err != nil {
			panic(err)
		}
		for _, item := range files {
			if item.IsDir() {
				vendors = append(vendors, AnsibleHost{Vendor: fmt.Sprintf("%s", item.Name()), Alias: []AnsibleAlias{{"server", "ip"}}})
			}
		}
		generateAnsibleHostsFile(vendors)
	} else if mode == "custom" {
		return // добавить логику
	}
	if buildAnsibleImage() {
		cmd := exec.Command("docker", "build", "extentions/ansible", "-t", "ansible:grip")
		cmd.Stdout = os.Stdout
		cmd.Stdin = os.Stdin
		cmd.Stderr = os.Stderr
		cmd.Run()
	} else {
		fmt.Println(string(text.RED), "Image already exist")
	}
}

func RunAnsible(command string) {
	entrypoint := fmt.Sprintf("--entrypoint=%s", command)
	cmd := exec.Command("docker", "run", "--rm", "--name=ansible-playbook", "--network=host", entrypoint, "-it", "ansible:grip")
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Run()
}

func buildAnsibleImage() bool {
	cmd, err := exec.Command("docker", "images").Output()
	if err != nil {
		panic(err)
	}
	if strings.Contains(string(cmd), "ansible") {
		return false
	} else {
		return true
	}
}
