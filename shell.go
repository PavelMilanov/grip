package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/PavelMilanov/grip/regru"
	"github.com/PavelMilanov/grip/vscale"
	"github.com/joho/godotenv"
)

func init() {
	os.Mkdir(vscale.VscaleDir, 0755)
	os.Mkdir(regru.RegruDir, 0755)
}

func env(key string) string {
	err := godotenv.Load()

	if err != nil {
		panic(err)
	}
	return os.Getenv(key)
}

func save_token(token string, vendor string) {
	file, err := os.OpenFile(".env", os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Unable to create file:", err)
		panic(err)
	}
	defer file.Close()
	file.WriteString(fmt.Sprintf("%s_TOKEN=%s\n", strings.ToUpper(vendor), token))
	fmt.Println("Token initialized successful!")
}

func cli_init() {
	help_text := `
grip init -provider=<provider> -token=<provider token>
`
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(help_text)
		}
	}()

	initCommand := flag.NewFlagSet("init", flag.ExitOnError)
	vendorProvider := initCommand.String("provider", "vscale", "vendor provider")
	vendorToken := initCommand.String("token", "", "vendor token")

	initCommand.Parse(os.Args[2:])

	switch *vendorProvider {
	case "vscale":
		switch statusCode := vscale.ValidateAccount(*vendorToken); statusCode {
		case 200:
			save_token(*vendorToken, *vendorProvider)
		case 403:
			fmt.Printf("%s token invalid!", *vendorProvider)
		default:
			fmt.Println(help_text)
		}

	case "regru":
		switch statusCode := regru.ValidateAccount(*vendorToken); statusCode {
		case 200:
			save_token(*vendorToken, *vendorProvider)
		case 403:
			fmt.Printf("%s token invalid!", *vendorProvider)
		default:
			fmt.Println(help_text)
		}
	default:
		fmt.Printf("%s is not supported provider.", *vendorProvider)
	}
}

func cli_vscale() {
	help_text := `
grip vscale ls		- view servers.
grip vscale create	- create new server.
grip vscale inspect	- inspect server config by name.
grip vscale rm		- remove server by name.
`
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(help_text)
			fmt.Println(r)
		}
	}()
	serverCommand := flag.NewFlagSet("vscale", flag.ExitOnError)
	serverCommand.Parse(os.Args[2:])

	token := env("VSCALE_TOKEN")
	switch os.Args[2] {
	case "ls":
		vscale.GetServer()
	case "create":
		createCommand := flag.NewFlagSet("create", flag.ExitOnError)
		createImage := createCommand.String("image", "debian_11_64_001_master", "OS image to server. Default: debian_11")
		createPlan := createCommand.String("plan", "small", "Plan to server. Default: small")
		createState := createCommand.Bool("start", false, "Server start status. Default: false")
		createName := createCommand.String("name", "", "Server name")
		createPassword := createCommand.String("pwd", "", "Server password")
		createLocation := createCommand.String("loc", "msk0", "Server location")

		createCommand.Parse(os.Args[3:])

		config := vscale.VscaleServer{
			Image:    *createImage,
			Size:     *createPlan,
			Do_start: *createState,
			Name:     *createName,
			Password: *createPassword,
			Location: *createLocation,
		}
		status := vscale.CreateServer(token, config)
		switch status {
		case 201:
			fmt.Println("Server successfully created")
		case 400:
			fmt.Println("Invalid data")
		}
	case "inspect":
		vscale.InspectServer(os.Args[3])
	case "rm":
		status := vscale.RemoveServer(token, os.Args[3])
		switch status {
		case 200:
			fmt.Println("Server successfully removed")
		case 404:
			fmt.Println("Server don't removed. Error")
		}
	default:
		fmt.Println(help_text)
	}
}

func cli_regru() {
	help_text := `
grip regru ls		- view servers.
grip regru create	- create new server.
grip regru inspect	- inspect server config by name.
grip regru rm		- remove server by name.
`
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(help_text)
			fmt.Println(r)
		}
	}()
	serverCommand := flag.NewFlagSet("regru", flag.ExitOnError)
	serverCommand.Parse(os.Args[2:])

	token := env("REGRU_TOKEN")
	switch os.Args[2] {
	case "ls":
		regru.GetServer()
	case "create":
		createCommand := flag.NewFlagSet("create", flag.ExitOnError)
		createImage := createCommand.String("image", "debian-11-amd64", "OS image to server. Default: debian_11")
		createPlan := createCommand.String("plan", "base-1", "Plan to server. Default: small")
		createName := createCommand.String("name", "regru-vps", "Server name")
		createBackup := createCommand.Bool("bkp", false, "Backuping to server")
		createLocation := createCommand.String("loc", "msk1", "Server location")

		createCommand.Parse(os.Args[3:])

		config := regru.RegruServer{
			Image:    *createImage,
			Size:     *createPlan,
			Name:     *createName,
			Backups:  *createBackup,
			Location: *createLocation,
		}
		regru.CreateServer(token, config)
	case "inspect":
		regru.InspectServer(os.Args[3])
	case "rm":
		status := regru.RemoveServer(token, os.Args[3])
		switch status {
		case 204:
			fmt.Println("Server successfully removed")
		case 404:
			fmt.Println("Server don't removed. Error")
		}
	default:
		fmt.Println(help_text)
	}
}
