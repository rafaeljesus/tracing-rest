package healthz

import (
  "net/http"
  "net/http/httptest"
  "testing"
  "strings"
  "github.com/labstack/echo"
  "github.com/labstack/echo/engine/standard"
  "github.com/stretchr/testify/assert"

  "github.com/rafaeljesus/tracing-rest/api/healthz"
)

const (
  responseJSON = `{"Alive":true}`
)

func TestIndex(t *testing.T) {
  e := echo.New()
  req, err := http.NewRequest(echo.POST, "/v1/healthz", strings.NewReader(responseJSON))

  if assert.NoError(t, err) {
    req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
    rec := httptest.NewRecorder()
    ctx := e.NewContext(standard.NewRequest(req, e.Logger()), standard.NewResponse(rec, e.Logger()))

    if assert.NoError(t, healthz.Index(ctx)) {
      assert.Equal(t, http.StatusOK, rec.Code)
      assert.Equal(t, responseJSON, rec.Body.String())
    }
  }
}
