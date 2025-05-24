package focus

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

// --- Handler registration ----------------------------------------------------

func RegisterReviewRoutes(app *fiber.App) {
    api := app.Group("/api/v1")
    api.Post("/daily-focus", focusHandler)
}


// --- Handlers ----------------------------------------------------------------

func focusHandler(c *fiber.Ctx) error {
    var req FocusRequest
    if err := c.BodyParser(&req); err != nil {
        return fiber.NewError(fiber.StatusBadRequest, "invalid JSON body")
    }

    // Set defaults: if date empty use tomorrow
    if req.Date == "" {
        req.Date = time.Now().Add(24 * time.Hour).Format("2006-01-02")
    }
    // if hours zero, service may derive from calendar

    items, err := focusService.Focus(req.Date, req.Hours)
    if err != nil {
        return fiber.NewError(fiber.StatusInternalServerError, err.Error())
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