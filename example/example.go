package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/vdparikh/wrap"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())

	e.GET("/hello", func(c echo.Context) error {
		return c.JSON(200, "HELLO")
	})

	server := wrap.Route(e)

	lambda.Start(server)
}
