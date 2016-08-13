package main

import (
  "github.com/labstack/echo"
  "github.com/labstack/echo/engine/standard"
  "github.com/labstack/echo/middleware"

  "github.com/rafaeljesus/tracing-rest/db"
  "github.com/rafaeljesus/tracing-rest/api/healthz"
  "github.com/rafaeljesus/tracing-rest/api/events"
)

func main() {
  db.Connect()

  e := echo.New()
  e.Use(middleware.Logger())

  r := e.Group("/v1")

  r.GET("/healthz", healthz.Index)
  r.GET("/events", events.Index)
  r.POST("/events", events.Create)

  e.Run(standard.New(":3000"))
}
