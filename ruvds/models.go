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
	Server []ServerConfigItems `json:"items"`
	// Premium      bool                `json:"premiun,omitempty"`
	// Tarif        TarifItems          `json:"tariff"`
}

type ServerConfigItems struct {
	Id            int              `json:"id"`
	Cpu           int              `json:"cpu"`
	Ram           int              `json:"ram"`
	Vram          int              `json:"vram"`
	Disk          []map[string]int `json:"drive"`
	Ip            IpItems          `json:"ip"`
	Tarif         TarifItems       `json:"tariff"`
	PaymentPeriod int              `json:"paymentPeriod"`
	AdminPassword string           `json:"defaultAdminPassword"`
	Running       bool             `json:"running"`
}

type TarifItems struct {
	Premium bool           `json:"premium"`
	Id      int            `json:"id"`
	Cpu     int            `json:"cpu"`
	Ram     int            `json:"ram"`
	Vram    float32        `json:"vram"`
	Drive   map[string]int `json:"drive"`
	Ip      int            `json:"ip"`
	Active  bool           `json:"active"`
}

type IpItems struct {
	Count int      `json:"count"`
	Ip    []string `json:"assigned"`
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
