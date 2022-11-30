package vscale

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func ValidateAccount(token string) int {
	url := "https://api.vscale.io/v1/account"
	client := http.Client{}
	request, err := http.NewRequest(http.MethodGet, url, nil)
	request.Header.Add("X-Token", token)

	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}

	return response.StatusCode
}

func GetServers(token string) string {
	url := "https://api.vscale.io/v1/scalets"
	client := http.Client{}
	request, err := http.NewRequest(http.MethodGet, url, nil)
	request.Header.Add("X-Token", token)

	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	return string(responseData)
}

func CreateServer(token string, template VscaleServer) int {
	data, _ := json.MarshalIndent(template, "", "	")
	url := "https://api.vscale.io/v1/scalets"
	client := http.Client{}
	request, _ := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(data))
	request.Header.Add("X-Token", token)
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
		saveConfig(responseData)
	case 400:
		panic(string(responseData))
	}
	return response.StatusCode
}

func GetServer() {
	files, err := ioutil.ReadDir(VscaleDir)
	if err != nil {
		panic(err)
	}

	for _, config := range files {
		config := readConfig(config.Name())
		fmt.Println(config.Name)
	}
}

func InspectServer(name string) {
	config := parceConfig(name)
	fmt.Printf("%s", config)
}

func RemoveServer(token string, name string) int {
	config_file := fmt.Sprintf("%s.json", name)
	config := readConfig(config_file)
	url := fmt.Sprintf("https://api.vscale.io/v1/scalets/%d", config.Ctid)
	client := http.Client{}
	request, err := http.NewRequest(http.MethodDelete, url, nil)
	request.Header.Add("X-Token", token)

	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}

	if response.StatusCode == 200 {
		os.Remove(config_file)
		return response.StatusCode
	} else {
		return 404
	}
}
