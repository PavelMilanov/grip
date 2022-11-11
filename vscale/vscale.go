package vscale

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type VscaleConfig struct {
	Make_from string `json:"make_from"`
	Rplan     string `json:"rplan"`
	Do_start  bool   `json:"do_start"`
	Name      string `json:"name"`
	Keys      string `json:"keys,omitempty"` // заменит keys пустым значением, если мы его не передаем
	Password  string `json:"password"`
	Location  string `json:"location"`
}

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

func CreateServer(token string, config VscaleConfig) (string, int) {
	data, _ := json.MarshalIndent(config, "", "	")
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

	return string(responseData), response.StatusCode
}

func GetServer(token string, id string) {

}

func RemoveServer(token string, id string) {

}
