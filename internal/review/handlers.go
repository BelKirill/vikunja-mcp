package review

import (
	"github.com/gofiber/fiber/v2"
)

// RegisterRoutes attaches review endpoints.
func RegisterRoutes(app *fiber.App) {
	api := app.Group("/api/v1")
	api.Post("/360-review", reviewHandler)
}

// @Summary      360-review
// @Description  Fetch the nearest important task for each specified project
// @Tags         review
// @Accept       json
// @Produce      json
// @Param        request  body      ReviewRequest  true  "360-review request payload"
// @Success      200      {array}   ReviewResponseItem
// @Failure      400      {object}  fiber.Error
// @Failure      500      {object}  fiber.Error
// @Router       /api/v1/360-review [post]
func reviewHandler(c *fiber.Ctx) error {
	var req ReviewRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid JSON body")
	}
	// TODO: Allow running without project ID, brings all projects
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
