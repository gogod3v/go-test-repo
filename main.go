package main

import (
	"net/http"
	
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		name := c.QueryParam("name")
		return c.String(http.StatusOK, "Hello, "+name+"!")
	})

	e.Logger.Fatal(e.Start(":1323"))
}

