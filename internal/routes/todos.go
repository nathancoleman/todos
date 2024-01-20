package routes

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"github.com/nathancoleman/todos/internal/api"
)

func getTODOs(w http.ResponseWriter, _ *http.Request) {
	todos, err := api.GetTODOs()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tData := struct {
		TODOs []api.TODO
	}{
		TODOs: todos,
	}

	err = templates.ExecuteTemplate(w, "todos.htmx", tData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
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

	err = templates.ExecuteTemplate(w, "todo.htmx", todo)
	if err != nil {
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
