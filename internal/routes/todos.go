package routes

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"

	"github.com/nathancoleman/todos/internal/api"
)

func getTODOs(w http.ResponseWriter, r *http.Request) {
	todos, err := api.GetTODOs()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if r.Header.Get("HX-Request") == "true" {

	} else {
		tData := struct {
			TODOs []api.TODO
		}{
			TODOs: todos,
		}

		err = templates.ExecuteTemplate(w, "todos.html", tData)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func getTODO(w http.ResponseWriter, r *http.Request) {
	todoID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	todo, err := api.GetTODO(todoID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = templates.ExecuteTemplate(w, "todo_list_item.htmx", todo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func putTODO(w http.ResponseWriter, r *http.Request) {
	todoID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	todo, err := api.GetTODO(todoID)
	if err != nil {
		http.Error(w, "", http.StatusNotFound)
		return
	}

	body := struct {
		Completed string `json:"completed"`
	}{}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	todo.Completed = strings.ToLower(body.Completed) == "on"

	if err = templates.ExecuteTemplate(w, "todo_list_item.htmx", todo); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func deleteTODO(w http.ResponseWriter, r *http.Request) {
	todoID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := api.DeleteTODO(todoID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
