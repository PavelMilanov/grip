package ruvds

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestValidateConfig(t *testing.T) {
	serverName := "test"
	json1, _ := ioutil.ReadFile("test.json")
	json2 := []byte(`{"rejectReason": 0,"id": 38420,"cost": 279}`)
	file1, json_data1 := server.validateConfig(json1, serverName)
	file2, json_data2 := server.validateConfig(json2, serverName)
	fmt.Println(file1, file2)
	fmt.Printf("%s", json_data1)
	fmt.Printf("%s", json_data2)
}
