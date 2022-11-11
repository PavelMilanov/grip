package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/PavelMilanov/grip/vscale"
)

func main() {

	initCommand := flag.NewFlagSet("init", flag.ExitOnError)
	vendorProvider := initCommand.String("provider", "vscale", "vendor provider")
	vendorToken := initCommand.String("token", "", "vendor token")

	if os.Args[1] == "init" {
		initCommand.Parse(os.Args[2:])
		switch statusCode := vscale.ValidateAccount(*vendorToken); statusCode {
		case 200:
			vendor := fmt.Sprintf("%s_TOKEN", *vendorProvider)
			os.Setenv(vendor, *vendorToken)
			fmt.Println("Token initialized successful!")
		case 403:
			fmt.Println("Token invalid!")
		}

	}
}
