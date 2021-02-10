package business

import (
	"fmt"
	"log"

	m "github.com/keuller/simple-api/internal/models"
)

type TaskService struct {
	repository TaskRepository
}

// constructor function
func NewTaskService(repo TaskRepository) TaskService {
	return TaskService{repo}
}

func (s TaskService) CreateNewTask(data m.AddTask) error {
	log.Println("TaskService::Create new task")

	task := AddTaskToEntity(data)
	log.Println(fmt.Sprintf("Task: %v", task))
	if err := s.repository.Save(task); err != nil {
		return err
	}

	return nil
}

func (s TaskService) UpdateTask(data m.UpdateTask) error {
	log.Println("TaskService::Update")

	_, err := s.repository.FindByID(data.ID)
	if err != nil {
		return err
	}

	task := UpdateTaskToEntity(data)

	if err := s.repository.Update(task); err != nil {
		return err
	}

	return nil
}

func (s TaskService) FindTaskById(value string) (m.TaskResource, error) {
	entity, err := s.repository.FindByID(value)

	if err != nil {
		return m.TaskResource{}, err
	}

	return m.TaskResource{
		ID:          entity.ID,
		Title:       entity.Title,
		Description: entity.Description,
		Done:        entity.Done,
	}, nil
}

func (s TaskService) ListTasks() []m.TaskResource {
	tasks := s.repository.FetchAllTasks()
	list := make([]m.TaskResource, 0)

	for _, task := range tasks {
		list = append(list, m.TaskResource{ID: task.ID, Title: task.Title, Description: task.Description, Done: task.Done})
	}

	return list
}

func (s TaskService) RemoveTask(value string) error {
	return s.repository.Delete(value)
}

func (s TaskService) ToggleTask(value string) error {
	return s.repository.Toggle(value)
}
