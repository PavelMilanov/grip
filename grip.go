package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	log.SetPrefix("LOG: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	log.Println("init started")
}

func env(key string) string {
	err := godotenv.Load()

	if err != nil {
		log.Fatalln(err)
	}
	return os.Getenv(key)
}

func main() {
	switch os.Args[1] {
	case "init":
		cli_init()
	case "server":
		cli_server()
	}
}
