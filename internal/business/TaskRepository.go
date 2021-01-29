package business

import "github.com/keuller/simple-api/internal/models"

type TaskRepository interface {
	Save(task models.Task) error

	Update(task models.Task) error

	FindByID(value string) (models.Task, error)

	FetchAllTasks() []models.Task

	Delete(value string) error

	Toggle(value string) error
}
