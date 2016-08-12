package events

import (
  "net/http"
  "github.com/labstack/echo"
)

func Index(c echo.Context) error {
  name := c.QueryParam("name")

  return c.JSON(http.StatusOK, &event{Name: name})
}

func Create(c echo.Context) error {
  e := &event{}

  if err := c.Bind(e); err != nil {
    return err
  }

  return c.JSON(http.StatusOK, e)
}

type response struct {
  Alive bool `json: "alive"`
}

type event struct {
  Name string `json: "name"`
}
