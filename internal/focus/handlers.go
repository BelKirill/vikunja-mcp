package focus

import (
	"context"
	"fmt"
	"time"

	"github.com/BelKirill/vikunja-mcp/models"
	"github.com/charmbracelet/log"
	"github.com/gofiber/fiber/v2"
)

// --- Handler registration ----------------------------------------------------

// RegisterRoutes registers REST API routes for focus endpoints.
func RegisterRoutes(app *fiber.App) {
	api := app.Group("/api/v1")

	service, err := NewService()
	if err != nil {
		log.Error("failed to create focus service", "error", err)
		return
	}

	api.Post("/daily-focus", FocusHandler(service))
}

// Handler for POST /daily-focus
func FocusHandler(service *Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req models.FocusRequest
		if err := c.BodyParser(&req); err != nil {
			log.Error("failed to parse focus request", "error", err)
			return c.Status(fiber.StatusBadRequest).JSON(models.APIError{Code: fiber.StatusBadRequest, Message: "invalid JSON body"})
		}

		if req.Date == "" {
			req.Date = time.Now().Format("2006-01-02")
			log.Info("focus request using default date", "date", req.Date)
		}
		if req.Hours == 0 {
			req.Hours = 5
			log.Info("focus request using default hours", "hours", req.Hours)
		}

		log.Info("processing focus request", "date", req.Date, "hours", req.Hours)

		ctx := context.Background()
		opts := models.FocusOptions{
			Date:     req.Date,
			Hours:    float32(req.Hours),
			Energy:   "medium",
			Mode:     "",
			MaxItems: 20,
		}
		items, err := service.GetFocusTasks(ctx, opts)
		if err != nil {
			log.Error("focus service error", "date", req.Date, "hours", req.Hours, "error", err)
			return c.Status(fiber.StatusInternalServerError).JSON(models.APIError{Code: fiber.StatusInternalServerError, Message: err.Error()})
		}

		resp := make([]models.FocusResponseItem, 0, len(items))
		for _, it := range items {
			resp = append(resp, models.FocusResponseItem{
				T:   fmt.Sprintf("%d", it.TaskID),
				P:   fmt.Sprintf("%d", it.Project),
				Est: 0, // TODO: map estimate if available
			})
		}
		log.Info("focus request completed", "date", req.Date, "items", len(items))
		return c.Status(fiber.StatusOK).JSON(resp)
	}
}
