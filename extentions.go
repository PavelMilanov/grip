package main

import (
	"flag"
	"fmt"
	"os"
)

func cli_ansible() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("ansible menu")
		}
	}()

	serverCommand := flag.NewFlagSet("ansible", flag.ExitOnError)
	serverCommand.Parse(os.Args[2:])
}
