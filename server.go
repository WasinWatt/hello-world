package main

import (
	"github.com/WasinWatt/hello-world/app"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e = app.WithTransport(e)

	e.Logger.Fatal(e.Start(":3000"))
}
