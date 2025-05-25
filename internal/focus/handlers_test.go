package focus

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestFocusHandler_BadRequest(t *testing.T) {
	app := fiber.New()
	app.Post("/daily-focus", focusHandler)

	// Invalid JSON
	req := httptest.NewRequest(http.MethodPost, "/daily-focus", nil)
	resp, _ := app.Test(req)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
}

// Add more tests for valid requests and service errors as needed.
