package vscale

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type server struct {
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

func GetServers(token string) {
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

	fmt.Println(string(responseData))
}

func CreateServer(token string) {

	data, _ := json.MarshalIndent(server{
		Make_from: "debian_11_64_001_master",
		Rplan:     "small",
		Do_start:  false,
		Name:      "Old-Test2",
		Password:  "P@ssw0rd7",
		Location:  "msk0",
	}, "", "	")

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

	fmt.Println(string(responseData))
}
