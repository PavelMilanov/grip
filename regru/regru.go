package regru

import (
	"net/http"
)

func ValidateAccount(token string) int {
	url := "https://api.cloudvps.reg.ru/v1/account/info"
	client := http.Client{}
	request, err := http.NewRequest(http.MethodGet, url, nil)
	bearer := "Bearer " + token
	request.Header.Add("Authorization", bearer)

	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}

	return response.StatusCode
}
