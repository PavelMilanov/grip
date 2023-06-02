package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/PavelMilanov/grip/extentions"
	"github.com/PavelMilanov/grip/text"
)

func cli_ansible() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf(text.ANSIBLE_MENU)
		}
	}()

	serverCommand := flag.NewFlagSet("ansible", flag.ExitOnError)
	serverCommand.Parse(os.Args[2:])

	switch os.Args[2] {
	case "build":
		mode := flag.String("f", "local", "local and custom ansible build file config.")
		flag.Parse()
		extentions.BuildAnsible(*mode)
	case "run":
		cmd := flag.String("cmd", "", "Command to run")
		flag.Parse()

		extentions.RunAnsible(*cmd)
	default:
		fmt.Printf(text.ANSIBLE_MENU)
	}

}
