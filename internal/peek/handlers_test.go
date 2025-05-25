package peek

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestTaskPeekHandler_BadRequest(t *testing.T) {
	app := fiber.New()
	app.Post("/task-peek", taskPeekHandler)

	// Invalid JSON
	req := httptest.NewRequest(http.MethodPost, "/task-peek", nil)
	resp, _ := app.Test(req)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
}

func TestProjectPeekHandler_BadRequest(t *testing.T) {
	app := fiber.New()
	app.Post("/project-peek", projectPeekHandler)

	// Invalid JSON
	req := httptest.NewRequest(http.MethodPost, "/project-peek", nil)
	resp, _ := app.Test(req)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
}
