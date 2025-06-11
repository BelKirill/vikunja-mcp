package review

import (
	"github.com/charmbracelet/log"
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
// @Failure      400      {object}  APIError
// @Failure      500      {object}  APIError
// @Router       /api/v1/360-review [post]
func reviewHandler(c *fiber.Ctx) error {
	var req ReviewRequest
	if err := c.BodyParser(&req); err != nil {
		log.Error("failed to parse review request", "error", err)
		return c.Status(fiber.StatusBadRequest).JSON(APIError{Code: fiber.StatusBadRequest, Message: "invalid JSON body"})
	}
	// TODO: Allow running without project ID, brings all projects
	if len(req.ProjectIDs) == 0 {
		log.Error("review request missing project IDs")
		return c.Status(fiber.StatusBadRequest).JSON(APIError{Code: fiber.StatusBadRequest, Message: "project_ids is required"})
	}

	log.Info("processing review request", "project_count", len(req.ProjectIDs))
	items, err := reviewService.Review(req.ProjectIDs)
	if err != nil {
		log.Error("review service error", "project_count", len(req.ProjectIDs), "error", err)
		return c.Status(fiber.StatusInternalServerError).JSON(APIError{Code: fiber.StatusInternalServerError, Message: err.Error()})
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
	log.Info("review request completed", "project_count", len(req.ProjectIDs), "items", len(items))
	return c.Status(fiber.StatusOK).JSON(resp)
}
