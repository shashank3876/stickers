package middleware_test

import (
	"net/http"
	test "net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	middleware "shashank-gusain-backend-onboarding/stickers/delivery/middleware"
)

func TestCORS(t *testing.T) {
	e := echo.New()
	req := test.NewRequest(echo.GET, "/", nil)
	res := test.NewRecorder()
	c := e.NewContext(req, res)
	m := middleware.InitMiddleware()

	h := m.CORS(func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})

	err := h(c)
	require.NoError(t, err)
	assert.Equal(t, "*", res.Header().Get("Access-Control-Allow-Origin"))
}
