package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/keuller/simple-api/internal/application/factory"
	"github.com/keuller/simple-api/internal/models"
)

var service = factory.GetTaskService()

// HTTP handler that creates a task
func CreateHandler(ctx *gin.Context) {
	var data models.AddTask

	if err := ctx.Bind(&data); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := service.CreateNewTask(data); err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "OK",
		"message": "Task was added.",
	})
}

func ReadByIdHandler(ctx *gin.Context) {
	resource, err := service.FindTaskById(ctx.Param("id"))

	if err != nil {
		ctx.AbortWithError(http.StatusNotFound, err)
		return
	}

	ctx.JSON(http.StatusOK, resource)
}

func UpdateHandler(ctx *gin.Context) {
	var data models.UpdateTask

	if err := ctx.Bind(&data); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	data.ID = ctx.Param("id")

	if err := service.UpdateTask(data); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Task has been updated.",
	})
}

func RemoveHandler(ctx *gin.Context) {
	if err := service.RemoveTask(ctx.Param("id")); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{})
}

func ListTaskHandler(ctx *gin.Context) {
	tasks := service.ListTasks()
	ctx.JSON(http.StatusOK, tasks)
}

func ToggleTaskHandler(ctx *gin.Context) {
	if err := service.ToggleTask(ctx.Param("id")); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{})
}
