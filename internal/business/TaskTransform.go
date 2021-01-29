package business

import (
	"time"

	"github.com/google/uuid"
	"github.com/keuller/simple-api/internal/models"
)

func currentTime() time.Time {
	location, _ := time.LoadLocation("America/Sao_Paulo")
	return time.Now().In(location)
}

func AddTaskToEntity(data models.AddTask) models.Task {
	newId := uuid.Must(uuid.NewRandom())
	return models.Task{
		ID:          newId.String(),
		Title:       data.Title,
		Description: data.Description,
		Done:        false,
		CreatedAt:   currentTime(),
	}
}

func UpdateTaskToEntity(data models.UpdateTask) models.Task {
	return models.Task{
		ID:          data.ID,
		Title:       data.Title,
		Description: data.Description,
		Done:        false,
		CreatedAt:   time.Now(),
	}
}
