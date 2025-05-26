package focus

import (
	"time"

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
		return c.Status(fiber.StatusBadRequest).JSON(APIError{Code: fiber.StatusBadRequest, Message: "invalid JSON body"})
	}

	// Set defaults: if date empty use tomorrow
	if req.Date == "" {
		req.Date = time.Now().Format("2006-01-02")
	}
	if req.Hours == 0 {
		req.Hours = 5
	}

	items, err := focusService.Focus(req.Date, req.Hours)
	if err != nil {
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
	return c.Status(fiber.StatusOK).JSON(resp)
}
