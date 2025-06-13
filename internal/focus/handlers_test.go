package focus

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestFocusHandler_BadRequest(t *testing.T) {
	app := fiber.New()
	service, err := NewService()
	if err != nil {
		t.Fatalf("failed to create service: %v", err)
	}
	app.Post("/daily-focus", FocusHandler(service))

	// Invalid JSON
	req := httptest.NewRequest(http.MethodPost, "/daily-focus", nil)
	resp, _ := app.Test(req)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
}

func TestFocusHandler_ValidRequest(t *testing.T) {
	app := fiber.New()
	s, err := NewService()
	if err != nil {
		t.Fatalf("failed to create service: %v", err)
	}
	app.Post("/daily-focus", FocusHandler(s))

	body := `{"date":"2025-06-12","hours":2}`
	req := httptest.NewRequest(http.MethodPost, "/daily-focus", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
}

func TestFocusHandler_Defaults(t *testing.T) {
	app := fiber.New()
	s, err := NewService()
	if err != nil {
		t.Fatalf("failed to create service: %v", err)
	}
	app.Post("/daily-focus", FocusHandler(s))

	body := `{}`
	req := httptest.NewRequest(http.MethodPost, "/daily-focus", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
}

func TestFocusHandler_ServiceError(t *testing.T) {
	app := fiber.New()
	// Use a nil Vikunja to force error
	s := &Service{Vikunja: nil}
	app.Post("/daily-focus", FocusHandler(s))

	body := `{"date":"2025-06-12","hours":2}`
	req := httptest.NewRequest(http.MethodPost, "/daily-focus", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
}

// Add more tests for other scenarios as needed.
