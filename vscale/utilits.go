package vscale

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

func saveConfig(data []byte) {

	var config ServerConfig
	err := json.Unmarshal(data, &config)
	if err != nil {
		log.Fatalln(err)
	}
	json_data, _ := json.MarshalIndent(config, "", "	")
	err = ioutil.WriteFile(fmt.Sprintf("configs/%s.json", config.Name), json_data, 0644)
}
