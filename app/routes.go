package app

import (
	"github.com/labstack/echo/v4"
	"go-boilerplate-project/app/types"
	"net/http"
)

func InitApi(e *echo.Echo, application *types.Application) {
	e.GET("/", func(context echo.Context) error {
		return context.String(http.StatusOK, "Hello world!")
	})
}
