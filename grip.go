package main

import (
	"log"
	"os"
)

func init() {
	log.SetPrefix("LOG: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	log.Println("init started")
}

func main() {
	switch os.Args[1] {
	case "init":
		cli_init()
	case "server":
		cli_server()
	}
}
