package persistence

import (
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/keuller/simple-api/internal/business"
	"github.com/keuller/simple-api/internal/models"
)

const (
	INSERT    = "INSERT INTO tasks (id, title, description, done, created_at) VALUES ($1, $2, $3, false, $4)"
	UPDATE    = "UPDATE tasks SET title = $1, description = $2 WHERE id = $3"
	DELETE    = "DELETE FROM tasks WHERE id = $1"
	TOGGLE    = "UPDATE tasks SET done = not done WHERE id = $1"
	FIND_ID   = "SELECT * FROM tasks WHERE id = $1"
	FETCH_ALL = "SELECT * FROM tasks ORDER BY created_at"
)

type TaskRepositoryImpl struct {
	db *sqlx.DB
}

func NewTaskRepository(db *sqlx.DB) business.TaskRepository {
	return TaskRepositoryImpl{db}
}

func (r TaskRepositoryImpl) Save(task models.Task) error {
	transaction := r.db.MustBegin()
	transaction.MustExec(INSERT, task.ID, task.Title, task.Description, task.CreatedAt)
	if err := transaction.Commit(); err != nil {
		return err
	}

	return nil
}

func (r TaskRepositoryImpl) Update(task models.Task) error {
	transaction := r.db.MustBegin()
	transaction.MustExec(UPDATE, task.Title, task.Description, task.ID)

	if err := transaction.Commit(); err != nil {
		return err
	}

	return nil
}

func (r TaskRepositoryImpl) FindByID(value string) (models.Task, error) {
	var entity models.Task

	if err := r.db.Get(&entity, FIND_ID, value); err != nil {
		return models.Task{}, err
	}

	return entity, nil
}

func (r TaskRepositoryImpl) FetchAllTasks() []models.Task {
	var tasks []models.Task

	if err := r.db.Select(&tasks, FETCH_ALL); err != nil {
		log.Printf("[WARN] %q", err)
		return make([]models.Task, 0)
	}

	return tasks
}

func (r TaskRepositoryImpl) Delete(value string) error {
	transaction := r.db.MustBegin()
	transaction.MustExec(DELETE, value)

	if err := transaction.Commit(); err != nil {
		return err
	}

	return nil
}

func (r TaskRepositoryImpl) Toggle(value string) error {
	transaction := r.db.MustBegin()
	transaction.MustExec(TOGGLE, value)

	if err := transaction.Commit(); err != nil {
		return err
	}

	return nil
}
