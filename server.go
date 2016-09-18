package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/fasthttp"
	"github.com/labstack/echo/middleware"

	"github.com/rafaeljesus/tracing-rest/api/events"
	"github.com/rafaeljesus/tracing-rest/api/healthz"
	"github.com/rafaeljesus/tracing-rest/db"
	"os"
)

func main() {
	db.Connect()

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())
	e.Use(middleware.Secure())
	e.Use(middleware.Gzip())

	r := e.Group("/v1")

	r.GET("/healthz", healthz.Index)
	r.GET("/events", events.Index)
	r.POST("/events", events.Create)

	e.Run(fasthttp.New(":" + os.Getenv("TRACKING_REST_PORT")))
}
