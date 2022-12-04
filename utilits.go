package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/PavelMilanov/grip/regru"
	"github.com/PavelMilanov/grip/vscale"
	"github.com/joho/godotenv"
)

func init() {
	os.Mkdir(vscale.VscaleDir, 0755)
	os.Mkdir(regru.RegruDir, 0755)
}

func env(key string) string {
	err := godotenv.Load()

	if err != nil {
		panic(err)
	}
	return os.Getenv(key)
}
func save_token(token string, vendor string) {
	new_variable := fmt.Sprintf("%s_TOKEN=%s\n", strings.ToUpper(vendor), token)
	if check_environment(new_variable) {
		new_env := validate_token(token, vendor)
		update_environment_variables(new_env)
	} else {
		file, err := os.OpenFile(".env", os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			fmt.Println("Unable to create file:", err)
			panic(err)
		}
		defer file.Close()
		file.WriteString(new_variable)
		fmt.Println("Token initialized successful!")
	}
}

func validate_token(token string, vendor string) *os.File {
	file_vendor := fmt.Sprintf("%s_TOKEN", strings.ToUpper(vendor))

	env, _ := os.OpenFile(".env", os.O_RDONLY, 0666)
	defer env.Close()

	scanner := bufio.NewScanner(env)
	tmp_file, _ := ioutil.TempFile(".", "tmp.txt") // создаем временный файл для сохранения новых значений и перезаписи основного потом
	for scanner.Scan() {
		file_string := strings.Split(scanner.Text(), "=")             // [VSCALE_TOKEN e299d3f826c5051ecef365fcbb7dceaf00b2cf88daac95e77c6a083ef38ed947]
		if file_vendor == file_string[0] && token == file_string[1] { // если названия переменной равны, токены равны
			fmt.Println("Token already exists")
		} else if file_vendor == file_string[0] && token != file_string[1] { // если названия переменной равны, токены не равны
			var flag string
			fmt.Printf("%s already exists, continue? (Y/N)", file_vendor)
			fmt.Scan(&flag)
			if strings.HasPrefix(strings.ToLower(flag), "y") {
				tmp_file.WriteString(file_vendor + "=" + token + "\n")
			} else {
				os.Remove(tmp_file.Name())
			}
		} else {
			tmp_file.WriteString(scanner.Text() + "\n")
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return tmp_file
}

func update_environment_variables(tmp *os.File) {
	env, _ := os.OpenFile(".env", os.O_WRONLY, 0666)
	defer env.Close()
	env.Truncate(0)
	new_data, _ := ioutil.ReadFile(tmp.Name())
	env.Write(new_data)
	os.Remove(tmp.Name())
}

func check_environment(pattern string) bool {
	env, _ := os.OpenFile(".env", os.O_RDONLY, 0666)
	defer env.Close()
	var check bool
	scanner := bufio.NewScanner(env)
	for scanner.Scan() {
		if pattern == scanner.Text() {
			check = true
			return check
		}
	}
	return check
}
