package main

import (
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
	}
}
