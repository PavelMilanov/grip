package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/PavelMilanov/grip/vscale"
	"github.com/joho/godotenv"
)

func env(key string) string {
	err := godotenv.Load()

	if err != nil {
		panic("Error loading .env file")
	}
	return os.Getenv(key)
}

func main() {

	initCommand := flag.NewFlagSet("init", flag.ExitOnError)
	vendorProvider := initCommand.String("provider", "vscale", "vendor provider")
	vendorToken := initCommand.String("token", "", "vendor token")

	serverCommand := flag.NewFlagSet("server", flag.ExitOnError)

	switch os.Args[1] {
	case "init":
		initCommand.Parse(os.Args[2:])
		switch statusCode := vscale.ValidateAccount(*vendorToken); statusCode {
		case 200:
			file, err := os.OpenFile(".env", os.O_WRONLY|os.O_CREATE, 0666)
			if err != nil {
				fmt.Println("Unable to create file:", err)
				os.Exit(1)
			}
			file.WriteString(fmt.Sprintf("%s_TOKEN=%s", strings.ToUpper(*vendorProvider), *vendorToken))
			defer file.Close()
			fmt.Println("Token initialized successful!")
		case 403:
			fmt.Println("Token invalid!")
		}
	case "server":
		serverCommand.Parse(os.Args[2:])
		token := env("VSCALE_TOKEN")
		switch os.Args[2] {
		case "ls":
			vscale.GetServers(token)
		}
	}
}
