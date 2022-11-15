package vscale

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func saveConfig(data []byte) {
	var config ServerConfig
	err := json.Unmarshal(data, &config)
	if err != nil {
		log.Fatalln(err)
	}
	json_data, _ := json.MarshalIndent(config, "", "	")
	ioutil.WriteFile(fmt.Sprintf("configs/%s.json", config.Name), json_data, 0644)
}

func readConfig(file string) {
	content := parceConfig(file)
	var config ServerConfig
	json.Unmarshal(content, &config)
	fmt.Println(config.Name)

}

func parceConfig(file string) []byte {
	os.Chdir("configs")
	filename := fmt.Sprintf("%s", file)
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return content
}
