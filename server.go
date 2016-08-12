package main

import (
  "log"
  "os"
  "github.com/labstack/echo"
  "github.com/labstack/echo/engine/standard"
  "github.com/labstack/echo/middleware"
  "github.com/rafaeljesus/tracing-rest/api/healthz"
)

func main() {
  log := log.New(os.Stdout, "web ", log.LstdFlags)

  e := echo.New()
  e.Use(middleware.Logger())
  e.GET("/api/healthz", healthz.Index)
  e.Run(standard.New(":3000"))

  log.Println("Running tracing-rest alive!")
}
