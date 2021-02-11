package web

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/keuller/simple-api/internal/application/factory"
	"github.com/keuller/simple-api/internal/models"
)

var service = factory.GetTaskService()

// CreateHandler - HTTP handler that creates a task
func createHandler(w http.ResponseWriter, r *http.Request) {
	var data models.AddTask

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := service.CreateNewTask(data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	render.JSON(w, r, map[string]string{
		"status":  "OK",
		"message": "Task was added.",
	})
}

func readByIDHandler(w http.ResponseWriter, r *http.Request) {
	resource, err := service.FindTaskById(chi.URLParam(r, "id"))

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	render.JSON(w, r, resource)
}

func updateHandler(w http.ResponseWriter, r *http.Request) {
	var data models.UpdateTask

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	data.ID = chi.URLParam(r, "id")

	if err := service.UpdateTask(data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	render.JSON(w, r, map[string]string{
		"message": "Task has been updated.",
	})
}

// removes an specific task by ID
func removeHandler(w http.ResponseWriter, r *http.Request) {
	if err := service.RemoveTask(chi.URLParam(r, "id")); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	render.JSON(w, r, map[string]string{})
}

// returns task list
func listTaskHandler(w http.ResponseWriter, r *http.Request) {
	tasks := service.ListTasks()
	render.JSON(w, r, tasks)
}

// changes task status
func toggleTaskHandler(w http.ResponseWriter, r *http.Request) {
	if err := service.ToggleTask(chi.URLParam(r, "id")); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	render.NoContent(w, r)
}
