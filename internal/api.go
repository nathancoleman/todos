package internal

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

var client = resty.New().
	SetBaseURL("https://dummyjson.com").
	SetHeader("Accept", "application/json")

type TODO struct {
	ID        int    `json:"id"`
	Title     string `json:"todo"`
	Completed bool   `json:"completed"`
	UserID    int    `json:"userId"`
}

func GetTODOs() ([]TODO, error) {
	var response struct {
		TODOs []TODO `json:"todos"`
	}

	_, err := client.R().
		SetResult(&response).
		Get("todos")

	return response.TODOs, err
}

func GetTODO(id int) (TODO, error) {
	var response TODO

	_, err := client.R().
		SetResult(&response).
		Get(fmt.Sprintf("todos/%d", id))

	return response, err
}

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Image     string `json:"image"`
	Username  string `json:"username"`
}

func GetUsers() ([]User, error) {
	var response struct {
		Users []User `json:"users"`
	}

	_, err := client.R().
		SetResult(&response).
		Get("users")

	return response.Users, err
}

func GetUser(id int) (User, error) {
	var response User

	_, err := client.R().
		SetResult(&response).
		Get(fmt.Sprintf("users/%d", id))

	return response, err
}
