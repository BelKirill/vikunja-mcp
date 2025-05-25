package review

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestReviewHandler_BadRequest(t *testing.T) {
	app := fiber.New()
	app.Post("/360-review", reviewHandler)

	// Invalid JSON
	req := httptest.NewRequest(http.MethodPost, "/360-review", nil)
	resp, _ := app.Test(req)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
}
