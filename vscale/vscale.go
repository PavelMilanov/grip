package vscale

import (
	"bytes"
	"encoding/json"
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

func GetServer(token string, id string) {

}

func RemoveServer(token string, id string) {

}

// func Test() {
// 	data := []byte(`{
// 	"ctid":15461047,
// 	"name":"cli-new",
// 	"status":"queued",
// 	"location":"msk0",
// 	"rplan":"small",
// 	"keys":[],
// 	"tags":[],
// 	"public_address":{},
// 	"private_address":{},
// 	"made_from":"debian_11_64_001_master",
// 	"hostname":"",
// 	"created":"11.11.2022 20:50:44",
// 	"active":true,
// 	"locked":true,
// 	"deleted":null,
// 	"block_reason":null,
// 	"block_reason_custom":null,
// 	"date_block":null
// }`)
// 	var config ServerConfig
// 	err := json.Unmarshal(data, &config)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	err = ioutil.WriteFile("data.json", data, 0644)
// 	// content, err := ioutil.ReadFile("data.json")
// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// }

// 	// fmt.Printf("File contents: %s", content)
// }
