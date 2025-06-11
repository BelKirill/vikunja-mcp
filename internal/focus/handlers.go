package focus

import (
	"time"

	"github.com/charmbracelet/log"
	"github.com/gofiber/fiber/v2"
)

// --- Handler registration ----------------------------------------------------

func RegisterRoutes(app *fiber.App) {
	api := app.Group("/api/v1")
	api.Post("/daily-focus", focusHandler)
}

// --- Handlers ----------------------------------------------------------------

// @Summary      Create or update daily focus
// @Description  Set or update the daily focus for the user
// @Tags         focus
// @Accept       json
// @Produce      json
// @Param        request  body      FocusRequest  true  "Daily focus payload"
// @Success      200      {object}  FocusResponse
// @Failure      400      {object}  APIError
// @Failure      500      {object}  APIError
// @Router       /api/v1/daily-focus [post]
func focusHandler(c *fiber.Ctx) error {
	var req FocusRequest
	if err := c.BodyParser(&req); err != nil {
		log.Error("failed to parse focus request", "error", err)
		return c.Status(fiber.StatusBadRequest).JSON(APIError{Code: fiber.StatusBadRequest, Message: "invalid JSON body"})
	}

	// Set defaults: if date empty use tomorrow
	if req.Date == "" {
		req.Date = time.Now().Format("2006-01-02")
		log.Info("focus request using default date", "date", req.Date)
	}
	if req.Hours == 0 {
		req.Hours = 5
		log.Info("focus request using default hours", "hours", req.Hours)
	}

	log.Info("processing focus request", "date", req.Date, "hours", req.Hours)
	items, err := focusService.Focus(req.Date, req.Hours)
	if err != nil {
		log.Error("focus service error", "date", req.Date, "hours", req.Hours, "error", err)
		return c.Status(fiber.StatusInternalServerError).JSON(APIError{Code: fiber.StatusInternalServerError, Message: err.Error()})
	}

	// Build slim response
	resp := make([]FocusResponseItem, 0, len(items))
	for _, it := range items {
		resp = append(resp, FocusResponseItem{
			T:   it.TaskID,
			P:   it.Project,
			Est: it.Estimate,
		})
	}
	log.Info("focus request completed", "date", req.Date, "items", len(items))
	return c.Status(fiber.StatusOK).JSON(resp)
}
