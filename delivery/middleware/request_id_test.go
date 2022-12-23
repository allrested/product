package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	appMiddleware "github.com/allrested/product/delivery/middleware"
	"github.com/allrested/product/entity"
	"github.com/allrested/product/mocks"
)

func TestGenerateCID(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	handler := func(c echo.Context) error {
		return c.String(http.StatusOK, "test")
	}

	mockLogger := new(mocks.Logger)
	mockJWTSvc := new(mocks.JWTService)
	cid := appMiddleware.NewMiddleware(mockJWTSvc, mockLogger).RequestID()
	h := cid(handler)
	err := h(c)

	require.NoError(t, err)
	assert.NotNil(t, rec.Header().Get(entity.RequestIDHeader))
}
