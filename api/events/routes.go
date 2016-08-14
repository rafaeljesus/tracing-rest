package events

import (
  "net/http"
  "github.com/labstack/echo"
  "github.com/rafaeljesus/tracing-rest/models"
)

func Index(c echo.Context) error {
  Name := c.QueryParam("name")
  Status := c.QueryParam("status")
  query := models.Query{Name, Status}
  events := models.Search(query)
  return c.JSON(http.StatusOK, &events)
}

func Create(c echo.Context) error {
  event := &models.Event{}
  if err := c.Bind(event); err != nil {
    return err
  }
  event.Save()
  return c.JSON(http.StatusOK, &event)
}
