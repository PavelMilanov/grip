package regru

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

const RegruDir = "configs/regru"

func saveConfig(data []byte) {
	var config ServerConfig
	err := json.Unmarshal(data, &config)
	if err != nil {
		panic(err)
	}
	json_data, _ := json.MarshalIndent(config, "", "	")
	ioutil.WriteFile(fmt.Sprintf("%s/%s.json", RegruDir, config.Server.Name), json_data, 0644)
}

func readConfig(file string) ServerConfig {
	content := parceConfig(file)
	var config ServerConfig
	json.Unmarshal(content, &config)
	return config
}

func parceConfig(file string) []byte {
	os.Chdir(RegruDir)
	content, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	return content
}
