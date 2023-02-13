package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/PavelMilanov/grip/regru"
	"github.com/PavelMilanov/grip/vscale"
)

func cli_init() {
	help_text := `
grip init -provider=<provider> -token=<provider token>
`
	messages := make(chan int)

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
		go vscale.ValidateAccount(*vendorToken, messages)
		statusCode := <-messages
		fmt.Println("Chek token...")
		switch statusCode {
		case 200:
			save_token(*vendorToken, *vendorProvider)
		case 403:
			fmt.Printf("%s token invalid!", *vendorProvider)
		default:
			fmt.Println(help_text)
		}

	case "regru":
		go regru.ValidateAccount(*vendorToken, messages)
		statusCode := <-messages
		fmt.Println("Chek token...")
		switch statusCode {
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
grip vscale stop	- stop server.
grip vscale start	- start server.
grip vscale restart	- restart server.
`
	messages := make(chan int)

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
		vscale.ShowServer()
	case "create":
		createCommand := flag.NewFlagSet("create", flag.ExitOnError)
		createImage := createCommand.String("image", "debian_11_64_001_master", "OS image to server. Default: debian_11")
		createPlan := createCommand.String("plan", "small", "Plan to server. Default: small")
		createState := createCommand.Bool("start", false, "Server start status. Default: false")
		createName := createCommand.String("name", "", "Server name")
		createPassword := createCommand.String("pwd", "", "Server password")
		createLocation := createCommand.String("loc", "msk0", "Server location")

		createCommand.Parse(os.Args[3:])

		data := vscale.VscaleServer{
			Image:    *createImage,
			Size:     *createPlan,
			Do_start: *createState,
			Name:     *createName,
			Password: *createPassword,
			Location: *createLocation,
		}
		go vscale.CreateServer(token, data, messages)
		fmt.Println("Server creating...")
		status := <-messages
		switch status {
		case 201:
			fmt.Println("Server successfully created")
		case 400:
			fmt.Println("Invalid data")
		}
	case "inspect":
		vscale.InspectServer(token, os.Args[3])
	case "rm":
		go vscale.RemoveServer(token, os.Args[3], messages)
		fmt.Println("Server removing...")
		status := <-messages
		switch status {
		case 200:
			fmt.Println("Server successfully removed")
		case 404:
			fmt.Println("Server don't removed. Error")
		}
	case "stop":
		go vscale.ManageServer(token, os.Args[3], "stop", messages)
		fmt.Println("Server stopping...")
		status := <-messages
		switch status {
		case 200:
			fmt.Println("Server successfully stopped")
		case 404:
			fmt.Printf("Server don't stopped. Error")
		}
	case "start":
		go vscale.ManageServer(token, os.Args[3], "start", messages)
		fmt.Println("Server started...")
		status := <-messages
		switch status {
		case 200:
			fmt.Println("Server successfully started")
		case 404:
			fmt.Printf("Server don't started. Error")
		}
	case "restart":
		go vscale.ManageServer(token, os.Args[3], "restart", messages)
		fmt.Println("Server restarted...")
		status := <-messages
		switch status {
		case 200:
			fmt.Println("Server successfully restarted")
		case 404:
			fmt.Printf("Server don't restarted. Error")
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

	messages := make(chan int)

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

		data := regru.RegruServer{
			Image:    *createImage,
			Size:     *createPlan,
			Name:     *createName,
			Backups:  *createBackup,
			Location: *createLocation,
		}
		go regru.CreateServer(token, data, messages)
		fmt.Println("Server creating...")
		status := <-messages
		switch status {
		case 201:
			fmt.Println("Server successfully created")
		case 400:
			fmt.Println("Invalid data")
		}
	case "inspect":
		regru.InspectServer(os.Args[3])
	case "rm":
		go regru.RemoveServer(token, os.Args[3], messages)
		status := <-messages
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
