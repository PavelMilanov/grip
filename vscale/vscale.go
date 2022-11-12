package vscale

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type VscaleServer struct {
	Image    string `json:"make_from"`
	Rplan    string `json:"rplan"`
	Do_start bool   `json:"do_start"`
	Name     string `json:"name"`
	Keys     []int  `json:"keys,omitempty"` // заменит keys пустым значением, если мы его не передаем
	Password string `json:"password"`
	Location string `json:"location"`
}

type ServerConfig struct {
	Ctid        int               `json:"ctid"`
	Name        string            `json:"name"`
	Status      string            `json:"status"`
	Location    string            `json:"location"`
	Rplan       string            `json:"rplan"`
	Keys        []int             `json:"keys,omitempty"`
	Tags        []string          `json:"tags,omitempty"`
	PublicAddr  map[string]string `json:"public_address,omitempty"`
	PrivateAddr map[string]string `json:"private_address,omitempty"`
	Image       string            `json:"made_from,omitempty"`
	CreateTime  string            `json:"created,omitempty"`
	Active      bool              `json:"active"`
	Loced       bool              `json:"loced"`
	Deleted     bool              `json:"deleted,omitempty"`
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
		log.Panicln(string(responseData))
	}
	return response.StatusCode
}

func GetServer() {
	files, err := ioutil.ReadDir("configs/")
	if err != nil {
		panic(err)
	}

	for _, config := range files {
		readConfig(config.Name())
	}
}

func InspectServer(name string) {
	config := parceConfig(name)
	fmt.Printf("%s", config)
}

func RemoveServer(token string, id string) {

}
