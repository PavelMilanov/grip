package regru

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
