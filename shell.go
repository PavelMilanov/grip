package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/PavelMilanov/grip/vscale"
	"github.com/joho/godotenv"
)

func env(key string) string {
	err := godotenv.Load()

	if err != nil {
		log.Fatalln(err)
	}
	return os.Getenv(key)
}

func cli_init() {
	initCommand := flag.NewFlagSet("init", flag.ExitOnError)
	vendorProvider := initCommand.String("provider", "vscale", "vendor provider")
	vendorToken := initCommand.String("token", "", "vendor token")

	initCommand.Parse(os.Args[2:])

	switch statusCode := vscale.ValidateAccount(*vendorToken); statusCode {
	case 200:
		file, err := os.OpenFile(".env", os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println("Unable to create file:", err)
			log.Panicln(err)
		}
		defer file.Close()
		file.WriteString(fmt.Sprintf("%s_TOKEN=%s", strings.ToUpper(*vendorProvider), *vendorToken))
		fmt.Println("Token initialized successful!")
	case 403:
		fmt.Println("Token invalid!")
	}
}

func cli_server() {
	serverCommand := flag.NewFlagSet("server", flag.ExitOnError)
	serverImage := serverCommand.String("image", "debian_11_64_001_master", "OS image to server. Default: debian_11")
	serverPlan := serverCommand.String("plan", "small", "Plan to server. Default: small")
	serverState := serverCommand.Bool("start", false, "Server start status. Default: false")
	serverName := serverCommand.String("name", "", "Server name")
	serverPassword := serverCommand.String("pwd", "", "Server password")
	serverLocation := serverCommand.String("loc", "msk0", "Server location")

	serverCommand.Parse(os.Args[3:])

	token := env("VSCALE_TOKEN")
	switch os.Args[2] {
	case "ls":
		info := vscale.GetServers(token)
		fmt.Println(info)
	case "create":
		config := vscale.VscaleServer{
			Image:    *serverImage,
			Rplan:    *serverPlan,
			Do_start: *serverState,
			Name:     *serverName,
			Password: *serverPassword,
			Location: *serverLocation,
		}
		status := vscale.CreateServer(token, config)
		switch status {
		case 201:
			fmt.Println("Server successfully created")
		case 400:
			fmt.Println("Invalid data")
		}
	case "inspect":
	case "delete":
	}
}
