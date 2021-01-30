package web

import (
	"net/http"

	"github.com/keuller/simple-api/internal/application/factory"
	"github.com/keuller/simple-api/internal/models"
	echo "github.com/labstack/echo/v4"
)

var service = factory.GetTaskService()

// HTTP handler that creates a task
func CreateHandler(ctx echo.Context) error {
	var data models.AddTask

	if err := ctx.Bind(&data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := service.CreateNewTask(data); err != nil {
		return echo.NewHTTPError(http.StatusBadGateway, err.Error())
	}

	return ctx.JSON(http.StatusOK, map[string]string{
		"status":  "OK",
		"message": "Task was added.",
	})
}

func ReadByIdHandler(ctx echo.Context) error {
	resource, err := service.FindTaskById(ctx.Param("id"))

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return ctx.JSON(http.StatusOK, resource)
}

func UpdateHandler(ctx echo.Context) error {
	var data models.UpdateTask

	if err := ctx.Bind(&data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	data.ID = ctx.Param("id")

	if err := service.UpdateTask(data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	return ctx.JSON(http.StatusOK, map[string]string{
		"message": "Task has been updated.",
	})
}

func RemoveHandler(ctx echo.Context) error {
	if err := service.RemoveTask(ctx.Param("id")); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return ctx.JSON(http.StatusNoContent, map[string]string{})
}

func ListTaskHandler(ctx echo.Context) error {
	tasks := service.ListTasks()
	return ctx.JSON(http.StatusOK, tasks)
}

func ToggleTaskHandler(ctx echo.Context) error {
	if err := service.ToggleTask(ctx.Param("id")); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return ctx.JSON(http.StatusNoContent, map[string]string{})
}
