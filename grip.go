package main

import (
	"fmt"
	"os"
)

func init() {
	os.Mkdir("configs", 0755)
}

func main() {
	help_text := `
grip init	- add prodvider token. (vscale)
grip vscale	- menu interaction of vscale-provider. 
`

	defer func() {
		if r := recover(); r != nil {
			fmt.Println(help_text)
		}
	}()

	switch os.Args[1] {
	case "init":
		cli_init()
	case "vscale":
		cli_vscale()
	case "regru":
		cli_regru()
	default:
		fmt.Println(help_text)
	}
}
