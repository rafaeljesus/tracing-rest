package healthz

import (
  "net/http"
  "github.com/labstack/echo"
)

func Index(c echo.Context) error {
  return c.JSON(http.StatusOK, &response{Alive: true})
}

type response struct {
  Alive bool `json: "alive"`
}
