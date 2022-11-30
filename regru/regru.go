package regru

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func ValidateAccount(token string) int {
	url := "https://api.cloudvps.reg.ru/v1/account/info"
	client := http.Client{}
	request, err := http.NewRequest(http.MethodGet, url, nil)
	bearer := "Bearer " + token
	request.Header.Add("Authorization", bearer)

	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}

	return response.StatusCode
}

func CreateServer(token string, template RegruServer) int {
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
		saveConfig(responseData)
	case 400:
		panic(string(responseData))
	}
	return response.StatusCode
}

func GetServer() {
	files, err := ioutil.ReadDir(RegruDir)
	if err != nil {
		panic(err)
	}

	for _, config := range files {
		config := readConfig(config.Name())
		fmt.Println(config.Server.Name)
	}
}
