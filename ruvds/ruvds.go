package ruvds

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/PavelMilanov/grip/text"
)

const RuVdsDir = "configs/ruvds"

var server ServerConfig

func ValidateAccount(template Account) (string, int) {
	/*
		Авторизует аккаунт на основе API-ключа. Возвращает session-ключ (токен API).
	*/
	urlPath := "https://ruvds.com/api/logon/"
	client := http.Client{}
	params := url.Values{
		"key":      {template.Key},
		"username": {template.Username},
		"password": {template.Password},
		"ednless":  {template.Endless},
	}
	encodedData := params.Encode()
	request, _ := http.NewRequest(http.MethodPost, urlPath, bytes.NewBuffer([]byte(encodedData)))
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	if response.StatusCode == 200 {
		var result map[string]interface{}
		err := json.Unmarshal(responseData, &result) // map[expire:19032023145445 rejectReason:0 sessionToken:a89f82521d66eeb4567bbe7d00b6675ed9a74fa37135468f30222facd74bc34f]
		if err != nil {
			panic(err)
		}
		token := fmt.Sprintf("%s", result["sessionToken"])
		return token, response.StatusCode
	}
	return "Error", response.StatusCode
}

func CreateServer(token string, template RuvdsServer) int {
	/*
		Делает POST-запрос к API на создание сервера, исходя из шаблона.
		Генерирует конфигурационный файл в json-формате.
	*/
	urlPath := "https://ruvds.com/api/server/create/"
	client := http.Client{}
	params := url.Values{
		"sessionToken":   {token},
		"datacenter":     {template.Datacenter},
		"tariff":         {template.Tariff},
		"os":             {template.Os},
		"cpu":            {template.Cpu},
		"ram":            {template.Ram},
		"drivesCount":    {template.DrivesCount},
		"drive0Tariff":   {template.Drive0Tariff},
		"drive0Capacity": {template.Drive0Сapacity},
		"drive0System":   {template.Drive0System},
		"ip":             {template.Ip},
		"ddosProtection": {template.DdosProtection},
		"paymentPeriod":  {template.PaymentPeriod},
		"promocode":      {template.Promocode},
		"computerName":   {template.ComputerName},
		"userComment":    {template.UserComment},
	}
	encodedData := params.Encode()
	request, _ := http.NewRequest(http.MethodPost, urlPath, bytes.NewBuffer([]byte(encodedData)))
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(responseData))
	switch response.StatusCode {
	case 200:
		file, json_data := server.validateConfig(responseData, template.ComputerName)
		ioutil.WriteFile(file, json_data, 0644)
	case 400:
		panic(string(responseData))
	}
	return response.StatusCode
}

func configServer(token string, name string) {
	/*
		Функция проходит по директории и ищет нужный файл по имени сервера, после
		делает запрос по API и редактирует файл
	*/
	files, err := ioutil.ReadDir(RuVdsDir)
	if err != nil {
		panic(err)
	}

	for _, item := range files {
		config := server.readConfig(item.Name())
		serverName := strings.Split(item.Name(), ".")
		if serverName[0] == name {
			urlPath := fmt.Sprintf("https://api.cloudvps.reg.ru/v1/reglets/%d", config.Id)
			client := http.Client{}
			params := url.Values{
				"sessionToken": {token},
			}
			encodedData := params.Encode()
			request, err := http.NewRequest(http.MethodGet, urlPath, bytes.NewBuffer([]byte(encodedData)))
			request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

			response, err := client.Do(request)
			if err != nil {
				panic(err)
			}
			defer response.Body.Close()
			responseData, err := ioutil.ReadAll(response.Body)
			if err != nil {
				panic(err)
			}
			switch response.StatusCode {
			case 200:
				file, json_data := server.validateConfig(responseData, serverName[0])
				file = fmt.Sprintf("%s.json", name)
				os.Chdir(RuVdsDir)
				err := ioutil.WriteFile(file, json_data, 0644) // перезаписывает конфиг. файл
				if err != nil {
					panic(err)
				}
			}
			break
		}
	}
}

func ShowServer() {
	/*
		Выводит список серверов по наличию конфигурационных файлов в директории.
	*/
	files, err := ioutil.ReadDir(RuVdsDir)
	if err != nil {
		panic(err)
	}

	for _, item := range files {
		// config := server.readConfig(item.Name())
		server := strings.Split(item.Name(), ".")
		fmt.Println(string(text.GREEN), server[0])
	}
}

func InspectServer(token string, name string) {
	/*
		Читает конфигурационный файл в директории по названию сервера и выводит его на печать.
	*/
	configServer(token, name)
	config := server.parceConfig(name + ".json")
	fmt.Printf("%s", config)
}

func ManageServer(token string, name string, command string) int {
	/*
		Делает различные запросы к API, в зависимости от команды.
	*/
	config_file := fmt.Sprintf("%s.json", name)
	config := server.readConfig(config_file)
	urlPath := fmt.Sprintf("https://ruvds.com/api/server/command")
	client := http.Client{}
	params := url.Values{
		"sessionToken": {token},
		"id":           {fmt.Sprintf("%d", config.Id)},
		"type":         {command},
	}
	encodedData := params.Encode()
	request, err := http.NewRequest(http.MethodPost, urlPath, bytes.NewBuffer([]byte(encodedData)))
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	if response.StatusCode == 200 && command == "remove" {
		os.Chdir(RuVdsDir)
		os.Remove(config_file)
	}
	return response.StatusCode
}

func GetServerParametrs(name string) ServerConfig {
	/*
		Возвращает структуру конфигурационного файла сервера.
	*/
	config_file := fmt.Sprintf("%s.json", name)
	os.Chdir(RuVdsDir)
	config := server.readConfig(config_file)
	return config
}
