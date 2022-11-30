package main

import (
	"fmt"
	"os"
)

func main() {
	help_text := `
grip init	- add prodvider token. (vscale, regru)
grip vscale	- menu interaction of vscale-provider.
grip regru	- menu interaction of regru-provider. 
`

	// defer func() {
	// 	if r := recover(); r != nil {
	// 		fmt.Println(help_text)
	// 	}
	// }()

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
