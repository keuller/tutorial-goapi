package factory

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/keuller/simple-api/internal/application/persistence"
	"github.com/keuller/simple-api/internal/business"
	"github.com/keuller/simple-api/internal/config"

	_ "github.com/lib/pq"
)

var repository business.TaskRepository

func init() {
	db, err := sqlx.Connect("postgres", connectionString())

	if err != nil {
		panic(err)
	}

	repository = persistence.NewTaskRepository(db)
}

func connectionString() string {
	config := config.Get()

	host := config.GetString("PG_HOST")
	port := config.GetString("PG_PORT")
	pgdb := config.GetString("PG_DATABASE")
	schema := config.GetString("PG_SCHEMA")
	user := config.GetString("PG_USER")
	pass := config.GetString("PG_PASS")

	return fmt.Sprintf("host=%s port=%s dbname=%s search_path=%s user=%s password=%s sslmode=disable", host, port, pgdb, schema, user, pass)
}

func GetTaskService() business.TaskService {
	return business.NewTaskService(repository)
}
