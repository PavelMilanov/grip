package main

import (
	"fmt"
	"os"
)

func init() {
	os.Mkdir("configs", 0755)
}

func main() {
	switch os.Args[1] {
	case "init":
		cli_init()
	case "server":
		cli_server()
	default:
		help_text := `
grip init	- add prodvider token. (vscale)
grip server	- menu interaction of server. 
`
		fmt.Println(help_text)
	}
}
