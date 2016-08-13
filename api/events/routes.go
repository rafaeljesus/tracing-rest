package events

import (
  "net/http"
  "github.com/labstack/echo"

  "github.com/rafaeljesus/tracing-rest/db"
  "github.com/rafaeljesus/tracing-rest/models"
)

func Index(c echo.Context) error {
  return c.JSON(http.StatusOK, &response{Alive: true})
}

func Create(c echo.Context) error {
  event := &models.Event{}

  if err := c.Bind(event); err != nil {
    return err
  }

  db.Repo.Create(&event)

  return c.JSON(http.StatusOK, &event)
}

type response struct {
  Alive bool `json: "alive"`
}
