package main

import (
	"fmt"
	"os"
)

func init() {
	os.Mkdir("configs", 0755)
}

func help_text() {
	help_text := `
grip init	- add prodvider token. (vscale)
grip vscale	- menu interaction of vscale-provider. 
`
	fmt.Println(help_text)
}

func main() {
	switch os.Args[1] {
	case "init":
		cli_init()
	case "vscale":
		cli_vscale()
	default:
		help_text()
	}
}
