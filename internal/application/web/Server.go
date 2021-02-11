package web

import (
	"net/http"

	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Server() *echo.Echo {
	app := echo.New()

	app.Use(middleware.Recover())
	app.Use(middleware.Gzip())

	app.GET("/", func(ctx echo.Context) error {
		return ctx.JSON(http.StatusOK, map[string]string{
			"message": "It works!",
		})
	})

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
