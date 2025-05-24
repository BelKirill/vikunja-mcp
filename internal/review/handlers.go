package review

import (
	"github.com/gofiber/fiber/v2"
)

// --- Handler registration ----------------------------------------------------

func RegisterReviewRoutes(app *fiber.App) {
    api := app.Group("/api/v1")
    api.Post("/360-review", reviewHandler)
}


// --- Handlers ----------------------------------------------------------------

func reviewHandler(c *fiber.Ctx) error {
    var req ReviewRequest
    if err := c.BodyParser(&req); err != nil {
        return fiber.NewError(fiber.StatusBadRequest, "invalid JSON body")
    }
    if len(req.ProjectIDs) == 0 {
        return fiber.NewError(fiber.StatusBadRequest, "project_ids is required")
    }

    items, err := reviewService.Review(req.ProjectIDs)
    if err != nil {
        return fiber.NewError(fiber.StatusInternalServerError, err.Error())
    }

    // Build slim response
    resp := make([]ReviewResponseItem, 0, len(items))
    for _, it := range items {
        resp = append(resp, ReviewResponseItem{
            P: it.ProjectID,
            T: it.TaskID,
            D: it.DueDate.Format("2006-01-02"),
        })
    }
    return c.Status(fiber.StatusOK).JSON(resp)
}