package vscale

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type VscaleServer struct {
	Image    string `json:"make_from"`
	Size     string `json:"rplan"`
	Do_start bool   `json:"do_start"`
	Name     string `json:"name"`
	Keys     []int  `json:"keys,omitempty"` // заменит keys пустым значением, если мы его не передаем
	Password string `json:"password"`
	Location string `json:"location"`
}

type ServerConfig struct {
	Ctid        int      `json:"ctid"`
	Name        string   `json:"name"`
	Status      string   `json:"status"`
	Location    string   `json:"location"`
	Size        string   `json:"rplan"`
	Keys        []int    `json:"keys,omitempty"`
	Tags        []string `json:"tags,omitempty"`
	PublicAddr  Network  `json:"public_address,omitempty"`
	PrivateAddr Network  `json:"private_address,omitempty"`
	Image       string   `json:"made_from,omitempty"`
	CreateTime  string   `json:"created,omitempty"`
	Active      bool     `json:"active"`
	Loced       bool     `json:"loced"`
	Deleted     bool     `json:"deleted,omitempty"`
}

type Network struct {
	Ip   string `json:"address,omitempty"`
	Mask string `json:"netmask,omitempty"`
	Gw   string `json:"gateway,omitempty"`
}

func (s ServerConfig) parceResponce(data []byte) int {
	err := json.Unmarshal(data, &s)
	if err != nil {
		panic(err)
	}
	return s.Ctid
}

func (s ServerConfig) validateConfig(data []byte) (string, []byte) {
	err := json.Unmarshal(data, &s)
	if err != nil {
		panic(err)
	}
	json_data, _ := json.MarshalIndent(s, "", "	")
	file := fmt.Sprintf("%s/%s.json", VscaleDir, s.Name)
	return file, json_data
}

func (s ServerConfig) readConfig(file string) ServerConfig {
	content := s.parceConfig(file)
	json.Unmarshal(content, &s)
	return s
}

func (s ServerConfig) parceConfig(file string) []byte {
	os.Chdir(VscaleDir)
	content, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	return content
}
