package main

import (
	"fmt"
	"os"

	"github.com/PavelMilanov/grip/text"
)

func main() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf(text.GRIP_MENU)
		}
	}()
	switch os.Args[1] {
	case "init":
		cli_init()
	case "vscale":
		cli_vscale()
	case "regru":
		cli_regru()
	case "ruvds":
		cli_ruvds()
	case "ansible":
		cli_ansible()
	default:
		fmt.Printf(text.GRIP_MENU)
	}
}
