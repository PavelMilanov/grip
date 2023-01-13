package regru

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const RegruDir = "configs/regru"

var server ServerConfig

func ValidateAccount(token string, canal chan int) {
	url := "https://api.cloudvps.reg.ru/v1/account/info"
	client := http.Client{}
	request, err := http.NewRequest(http.MethodGet, url, nil)
	bearer := "Bearer " + token
	request.Header.Add("Authorization", bearer)

	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}

	canal <- response.StatusCode
}

func CreateServer(token string, template RegruServer, canal chan int) {
	url := "https://api.cloudvps.reg.ru/v1/reglets"
	data, _ := json.MarshalIndent(template, "", "	")
	client := http.Client{}
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(data))
	bearer := "Bearer " + token
	request.Header.Add("Authorization", bearer)
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	switch response.StatusCode {
	case 201:
		file, json_data := server.validateConfig(responseData)
		ioutil.WriteFile(file, json_data, 0644)
		canal <- response.StatusCode
	case 400:
		canal <- response.StatusCode
		panic(string(responseData))
	}
}

func GetServer() {
	files, err := ioutil.ReadDir(RegruDir)
	if err != nil {
		panic(err)
	}

	for _, config := range files {
		config := server.readConfig(config.Name())
		fmt.Println(config.Server.Name)
	}
}

func InspectServer(name string) {
	config := server.parceConfig(name + ".json")
	fmt.Printf("%s", config)
}

func RemoveServer(token string, name string, canal chan int) {
	config_file := fmt.Sprintf("%s.json", name)
	config := server.readConfig(config_file)
	url := fmt.Sprintf("https://api.cloudvps.reg.ru/v1/reglets/%d", config.Server.Id)
	client := http.Client{}
	request, err := http.NewRequest(http.MethodDelete, url, nil)
	bearer := "Bearer " + token
	request.Header.Add("Authorization", bearer)

	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}

	if response.StatusCode == 204 {
		os.Chdir(RegruDir)
		os.Remove(config_file)
		canal <- response.StatusCode
	} else {
		canal <- response.StatusCode
	}
}
