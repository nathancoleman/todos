package api

import (
	"fmt"
)

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

func DeleteTODO(id int) error {
	_, err := client.R().
		Delete(fmt.Sprintf("todos/%d", id))

	return err
}
