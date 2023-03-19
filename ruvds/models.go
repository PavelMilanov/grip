package ruvds

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Account struct {
	Key      string
	Username string
	Password string
	Endless  string
}

type RuvdsServer struct {
	Datacenter     string
	Tariff         string
	Os             string
	Cpu            string
	Ram            string
	DrivesCount    string
	Drive0Tariff   string
	Drive0Сapacity string
	Drive0System   string
	Ip             string
	DdosProtection string
	PaymentPeriod  string
	Promocode      string
	ComputerName   string
	UserComment    string
}

type ServerConfig struct {
	Id int `json:"id"`
}

// func (s ServerConfig) parceResponce(data []byte) int {
// 	err := json.Unmarshal(data, &s)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return s.Id
// }

func (s ServerConfig) validateConfig(data []byte, server string) (string, []byte) {
	err := json.Unmarshal(data, &s)
	if err != nil {
		panic(err)
	}
	json_data, _ := json.MarshalIndent(s, "", "	")
	file := fmt.Sprintf("%s/%s.json", RuVdsDir, server)
	return file, json_data
}

func (s ServerConfig) readConfig(file string) ServerConfig {
	content := s.parceConfig(file)
	json.Unmarshal(content, &s)
	return s
}

func (s ServerConfig) parceConfig(file string) []byte {
	os.Chdir(RuVdsDir)
	content, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	return content
}
