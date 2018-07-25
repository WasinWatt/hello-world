package app

import (
	"net/http"

	"github.com/labstack/echo"
)

// WithTransport returns new eco handler
func WithTransport(e *echo.Echo) *echo.Echo {
	e.GET("/hello", func(ctx echo.Context) error {
		return ctx.JSON(http.StatusOK, struct {
			Message string
			Sender  string
		}{
			"Hello",
			"Pun",
		})
	})

	e.GET("/", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "Status OK :)")
	})

	return e
}
