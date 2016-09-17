package events

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/rafaeljesus/tracing-rest/api/events"
	"github.com/rafaeljesus/tracing-rest/db"
)

func init() {
	db.Connect()
}

func TestIndex(t *testing.T) {
	db.Connect()
	e := echo.New()

	q := make(url.Values)
	q.Set("name", "order_created")

	req, err := http.NewRequest(echo.GET, "/v1/events/?"+q.Encode(), nil)

	if assert.NoError(t, err) {
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		ctx := e.NewContext(standard.NewRequest(req, e.Logger()), standard.NewResponse(rec, e.Logger()))

		if assert.NoError(t, events.Index(ctx)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	}
}
