package events

import (
	"github.com/labstack/echo"
	"net/http"

	"github.com/rafaeljesus/tracing-rest/models"
)

func Index(c echo.Context) error {
	Name := c.QueryParam("name")
	query := models.Query{Name}
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
