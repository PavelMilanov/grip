package regru

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/PavelMilanov/grip/text"
)

const RegruDir = "configs/regru"

var server ServerConfig

func ValidateAccount(token string, canal chan int) {
	/*
		Получает по API список серверов.
	*/
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
	/*
		Делает POST-запрос к API на создание сервера, исходя из шаблона.
		Генерирует конфигурационный файл в json-формате.
	*/
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

func configServer(token string, name string) {
	/*
		Функция проходит по директории и ищет нужный файл по имени сервера, после
		делает запрос по API и редактирует файл
	*/
	files, err := ioutil.ReadDir(RegruDir)
	if err != nil {
		panic(err)
	}

	for _, config := range files {
		config := server.readConfig(config.Name())
		if config.Server.Name == name {
			url := fmt.Sprintf("https://api.cloudvps.reg.ru/v1/reglets/%d", config.Server.Id)
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

			switch response.StatusCode {
			case 200:
				file, json_data := server.validateConfig(responseData)
				file = fmt.Sprintf("%s.json", name)
				os.Chdir(RegruDir)
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
	files, err := ioutil.ReadDir(RegruDir)
	if err != nil {
		panic(err)
	}

	for _, config := range files {
		config := server.readConfig(config.Name())
		fmt.Println(string(text.GREEN), config.Server.Name)
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

func ManageServer(token string, name string, command string, canal chan int) {
	/*
		Делает POST-запрос исходя из переданного параметра. (start/stop/restart).
	*/
	config_file := fmt.Sprintf("%s.json", name)
	config := server.readConfig(config_file)
	param := map[string]string{
		"type": fmt.Sprintf("%s", command),
	}
	data, _ := json.Marshal(param)
	// data_json, _ := json.MarshalIndent(command, "", "	")
	url := fmt.Sprintf("https://api.vscale.io/v1/scalets/%d/actions", config.Server.Id)
	client := http.Client{}
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(data))
	bearer := "Bearer " + token
	request.Header.Add("Authorization", bearer)
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}

	if response.StatusCode == 200 {
		canal <- response.StatusCode
	} else {
		canal <- response.StatusCode
	}
}
