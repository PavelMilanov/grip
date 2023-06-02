package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/PavelMilanov/grip/regru"
	"github.com/PavelMilanov/grip/ruvds"
	"github.com/PavelMilanov/grip/text"
	"github.com/PavelMilanov/grip/vscale"
)

func cli_init() {
	/*
		Команды для инициализации API-вендроров.
	*/

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf(text.INIT_MENU)
		}
	}()

	initCommand := flag.NewFlagSet("init", flag.ExitOnError)
	vendorProvider := initCommand.String("provider", "vscale", "vendor provider")
	vendorToken := initCommand.String("token", "", "vendor token")
	ruvdsUserName := initCommand.String("username", "", "username for ruvds account")
	ruvdsUserPassword := initCommand.String("password", "", "password for ruvds account")
	initCommand.Parse(os.Args[2:])

	switch *vendorProvider {
	case "vscale":
		statusCode := vscale.ValidateAccount(*vendorToken)
		fmt.Println("Chek token...")
		switch statusCode {
		case 200:
			save_token(*vendorToken, *vendorProvider)
		case 403:
			fmt.Printf("%s token invalid!", *vendorProvider)
		default:
			fmt.Printf(text.INIT_MENU)
		}
	case "regru":
		statusCode := regru.ValidateAccount(*vendorToken)
		fmt.Println("Chek token...")
		switch statusCode {
		case 200:
			save_token(*vendorToken, *vendorProvider)
		case 403:
			fmt.Printf("%s token invalid!", *vendorProvider)
		default:
			fmt.Printf(text.INIT_MENU)
		}
	case "ruvds":
		data := ruvds.Account{
			Key:      *vendorToken,
			Username: *ruvdsUserName,
			Password: *ruvdsUserPassword,
			Endless:  "1",
		}
		token, statusCode := ruvds.ValidateAccount(data)
		fmt.Println("Chek token...")
		switch statusCode {
		case 200:
			save_token(token, *vendorProvider)
		case 403:
			fmt.Printf("%s token invalid!", *vendorProvider)
		default:
			fmt.Printf(text.INIT_MENU)
		}

	default:
		fmt.Printf("%s is not supported provider.", *vendorProvider)
	}
}

func cli_vscale() {
	/*
		Команды для управления инфраструктурой через API vscale.
	*/
	messages := make(chan int)

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf(text.VSCALE_MENU)
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
			fmt.Println(string(text.CYAN), "Server successfully created")
		case 400:
			fmt.Println(string(text.RED), "Server don't created. Eror")
		}
	case "inspect":
		vscale.InspectServer(token, os.Args[3])
	case "rm":
		go vscale.RemoveServer(token, os.Args[3], messages)
		fmt.Println("Server removing...")
		status := <-messages
		switch status {
		case 200:
			fmt.Println(string(text.CYAN), "Server successfully removed")
		case 404:
			fmt.Println(string(text.RED), "Server don't removed. Error")
		}
	case "stop":
		go vscale.ManageServer(token, os.Args[3], "stop", messages)
		fmt.Println("Server stopping...")
		status := <-messages
		switch status {
		case 200:
			fmt.Println(string(text.CYAN), "Server successfully stopped")
		case 404:
			fmt.Println(string(text.RED), "Server don't stopped. Error")
		}
	case "start":
		go vscale.ManageServer(token, os.Args[3], "start", messages)
		fmt.Println("Server started...")
		status := <-messages
		switch status {
		case 200:
			fmt.Println(string(text.CYAN), "Server successfully started")
		case 404:
			fmt.Println(string(text.RED), "Server don't started. Error")
		}
	case "restart":
		go vscale.ManageServer(token, os.Args[3], "restart", messages)
		fmt.Println("Server restarted...")
		status := <-messages
		switch status {
		case 200:
			fmt.Println(string(text.CYAN), "Server successfully restarted")
		case 404:
			fmt.Println(string(text.RED), "Server don't restarted. Error")
		}
	case "ssh":
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("Server name error.\nWrite: 'grip vscale ls' for get servers")
			}
		}()
		server := os.Args[3]
		server_alias := vscale.GetServerParametrs(server)
		ssh_connection(server_alias.PublicAddr.Ip)
	default:
		fmt.Printf(text.VSCALE_MENU)
	}
}

func cli_regru() {
	/*
		Команды для управления инфраструктурой через API reg.ru.
	*/

	messages := make(chan int)

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf(text.REGRU_MENU)
		}
	}()
	serverCommand := flag.NewFlagSet("regru", flag.ExitOnError)
	serverCommand.Parse(os.Args[2:])

	token := env("REGRU_TOKEN")
	switch os.Args[2] {
	case "ls":
		regru.ShowServer()
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
			fmt.Println(string(text.CYAN), "Server successfully created")
		case 400:
			fmt.Println(string(text.RED), "Server don't created. Error")
		}
	case "inspect":
		regru.InspectServer(token, os.Args[3])
	case "rm":
		go regru.RemoveServer(token, os.Args[3], messages)
		status := <-messages
		switch status {
		case 204:
			fmt.Println(string(text.CYAN), "Server successfully removed")
		case 404:
			fmt.Println(string(text.RED), "Server don't removed. Error")
		}
	case "stop":
		go regru.ManageServer(token, os.Args[3], "stop", messages)
		fmt.Println("Server stopping")
		status := <-messages
		switch status {
		case 200:
			fmt.Println(string(text.CYAN), "Server successfully stopped")
		case 404:
			fmt.Println(string(text.RED), "Server don't stopped. Error")
		}
	case "start":
		go regru.ManageServer(token, os.Args[3], "start", messages)
		fmt.Println("Server started")
		status := <-messages
		switch status {
		case 200:
			fmt.Println(string(text.CYAN), "Server successfully started")
		case 404:
			fmt.Println(string(text.RED), "Server don't started. Error")
		}
	case "restart":
		go regru.ManageServer(token, os.Args[3], "reboot", messages)
		fmt.Println("Server restarted")
		status := <-messages
		switch status {
		case 200:
			fmt.Println(string(text.CYAN), "Server successfully restarted")
		case 404:
			fmt.Println(string(text.RED), "Server don't restarted. Error")
		}
	case "ssh":
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("Server name error.\nWrite: 'grip regru ls' for get servers")
			}
		}()
		server := os.Args[3]
		server_alias := regru.GetServerParametrs(server)
		ssh_connection(server_alias.Server.Ip)
	default:
		fmt.Printf(text.REGRU_MENU)
	}
}

func cli_ruvds() {
	/*
		Команды для управления инфраструктурой через API ruvds.
	*/
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf(text.RUVDS_MENU)
		}
	}()
	serverCommand := flag.NewFlagSet("ruvds", flag.ExitOnError)
	serverCommand.Parse(os.Args[2:])

	token := env("RUVDS_TOKEN")
	switch os.Args[2] {
	case "ls":
		ruvds.ShowServer()
	case "create":
		createCommand := flag.NewFlagSet("create", flag.ExitOnError)
		createDatacenter := createCommand.String("localtion", "1", "Datacenter. Default: 1 (Moscow)")
		createTariff := createCommand.String("tarif", "14", "Tarif of server. Default: 14 (Regular)")
		createOs := createCommand.String("os", "45", "OS of server. Default: Debian 11")
		createCpu := createCommand.String("cpu", "1", "CPU of server. Default: 1GB")
		createRam := createCommand.String("ram", "0.5", "RAM of server. Default: 1GB")
		createDrivesCount := createCommand.String("disks", "1", "Count disks of server. Default: 1")
		createDrive0Tariff := createCommand.String("type", "1", "Tarif of disk. HDD or SDD. Default: 1") // 1: HDD, 3: SSD
		createDrive0Сapacity := createCommand.String("capacity", "10", "Capacity of disk. Default: 15 GB")
		createDrive0System := createCommand.String("boot", "true", "System OS. Default: True")
		createIp := createCommand.String("ip", "1", "Count of ip by server. Default: 1")
		createDdosProtection := createCommand.String("ddos", "0", "DDOS protection of server. Default: 0")
		createPaymentPeriod := createCommand.String("payment", "2", "lease of payment. Default: 2 (monthy)")
		createPromocode := createCommand.String("promo", "", "Ruvds promocode. Default: <empty>")
		createComputerName := createCommand.String("name", "ruvds-server", "Name of server. Default: ruvds-server")
		createUserComment := createCommand.String("comment", "", "User comment. Default: <empty>")

		createCommand.Parse(os.Args[3:])

		data := ruvds.RuvdsServer{
			Datacenter:     *createDatacenter,
			Tariff:         *createTariff,
			Os:             *createOs,
			Cpu:            *createCpu,
			Ram:            *createRam,
			DrivesCount:    *createDrivesCount,
			Drive0Tariff:   *createDrive0Tariff,
			Drive0Сapacity: *createDrive0Сapacity,
			Drive0System:   *createDrive0System,
			Ip:             *createIp,
			DdosProtection: *createDdosProtection,
			PaymentPeriod:  *createPaymentPeriod,
			Promocode:      *createPromocode,
			ComputerName:   *createComputerName,
			UserComment:    *createUserComment,
		}
		status := ruvds.CreateServer(token, data)
		fmt.Println("Server creating...")
		switch status {
		case 201:
			fmt.Println(string(text.CYAN), "Server successfully created")
		case 400:
			fmt.Println(string(text.RED), "Server don't created. Error")
		}
	case "inspect":
		ruvds.InspectServer(token, os.Args[3])
	case "rm":
		status := ruvds.ManageServer(token, os.Args[3], "remove")
		switch status {
		case 204:
			fmt.Println(string(text.CYAN), "Server successfully removed")
		case 404:
			fmt.Println(string(text.RED), "Server don't removed. Error")
		}
	case "stop":
		status := ruvds.ManageServer(token, os.Args[3], "stop")
		fmt.Println("Server stopping")
		switch status {
		case 200:
			fmt.Println(string(text.CYAN), "Server successfully stopped")
		case 404:
			fmt.Println(string(text.RED), "Server don't stopped. Error")
		}
	case "start":
		status := ruvds.ManageServer(token, os.Args[3], "start")
		fmt.Println("Server started")
		switch status {
		case 200:
			fmt.Println(string(text.CYAN), "Server successfully started")
		case 404:
			fmt.Println(string(text.RED), "Server don't started. Error")
		}
	case "restart":
		status := ruvds.ManageServer(token, os.Args[3], "reset")
		fmt.Println("Server restarted")
		switch status {
		case 200:
			fmt.Println(string(text.CYAN), "Server successfully restarted")
		case 404:
			fmt.Println(string(text.RED), "Server don't restarted. Error")
		}
	case "ssh":
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("Server name error.\nWrite: 'grip ruvds ls' for get servers")
			}
		}()
		// server := os.Args[3]
		// server_alias := ruvds.GetServerParametrs(server)
		// ssh_connection(server_alias.Server.Ip)
	default:
		fmt.Printf(text.RUVDS_MENU)
	}

}
