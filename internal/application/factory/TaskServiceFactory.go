package factory

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/keuller/simple-api/internal/application/persistence"
	"github.com/keuller/simple-api/internal/business"

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
	if err := godotenv.Load(); err != nil {
		if err = godotenv.Load("../../../.env"); err != nil { // only for unit test purpose
			panic(err.Error())
		}
	}

	host := os.Getenv("PG_HOST")
	port := os.Getenv("PG_PORT")
	pgdb := os.Getenv("PG_DATABASE")
	schema := os.Getenv("PG_SCHEMA")
	user := os.Getenv("PG_USER")
	pass := os.Getenv("PG_PASS")

	return fmt.Sprintf("host=%s port=%s dbname=%s search_path=%s user=%s password=%s sslmode=disable", host, port, pgdb, schema, user, pass)
}

func GetTaskService() business.TaskService {
	return business.NewTaskService(repository)
}
