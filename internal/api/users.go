package api

import "fmt"

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
