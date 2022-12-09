package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	userModel "github.com/therasuldev/model"
)

const URL = "https://jsonplaceholder.typicode.com/users"

var client *http.Client

func GetJson(url string, data interface{}) error {

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(data)
}

func getUsers() {
	var users userModel.Users

	err := GetJson(URL, &users.User)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	} else {
		for i := 0; i < len(users.User); i++ {
			fmt.Println(users.User[i].Address.Geo.Lat)
		}
	}
}

func main() {
	client = &http.Client{Timeout: time.Second * 10}
	getUsers()
}
