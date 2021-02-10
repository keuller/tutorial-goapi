package web

import (
	"net/http"
	"os"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func Server() *gin.Engine {
	mode := os.Getenv("MODE")
	if mode != "dev" {
		gin.SetMode(gin.ReleaseMode)
	}

	app := gin.Default()
	app.RedirectTrailingSlash = true

	app.Use(gin.Recovery())
	app.Use(gzip.Gzip(gzip.DefaultCompression))

	app.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "It works!",
		})
	})

	app.GET("/async", AsyncHandler)

	crud := app.Group("/v1/tasks")
	{
		crud.POST("", CreateHandler)
		crud.GET("", ListTaskHandler)
		crud.GET("/:id", ReadByIdHandler)
		crud.PUT("/:id", UpdateHandler)
		crud.PUT("/:id/toggle", ToggleTaskHandler)
		crud.DELETE("/:id", RemoveHandler)
	}

	return app
}
