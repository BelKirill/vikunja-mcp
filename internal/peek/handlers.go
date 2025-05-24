package peek

import "github.com/gofiber/fiber/v2"

// RegisterRoutes attaches peek endpoints.
func RegisterRoutes(app *fiber.App) {
    api := app.Group("/api/v1")
    api.Post("/task-peek", taskPeekHandler)
    api.Post("/project-peek", projectPeekHandler)
}

// --- Handlers ----------------------------------------------------------------

func taskPeekHandler(c *fiber.Ctx) error {
    var req TaskPeekRequest
    if err := c.BodyParser(&req); err != nil {
        return fiber.NewError(fiber.StatusBadRequest, "invalid JSON body")
    }
    if req.ID == "" {
        return fiber.NewError(fiber.StatusBadRequest, "id is required")
    }
    if len(req.Fields) == 0 {
        req.Fields = []string{"title"}
    }

    // Delegate to the service
    data, err := taskService.Peek(req.ID, req.Fields)
    if err != nil {
        return fiber.NewError(fiber.StatusInternalServerError, err.Error())
    }

    resp := TaskPeekResponse{T: req.ID, F: data}
    return c.Status(fiber.StatusOK).JSON(resp)
}

func projectPeekHandler(c *fiber.Ctx) error {
    var req ProjectPeekRequest
    if err := c.BodyParser(&req); err != nil {
        return fiber.NewError(fiber.StatusBadRequest, "invalid JSON body")
    }
    if req.ID == "" {
        return fiber.NewError(fiber.StatusBadRequest, "id is required")
    }
    if len(req.Fields) == 0 {
        req.Fields = []string{"name"}
    }

    // Delegate to the service
    data, err := projectService.Peek(req.ID, req.Fields)
    if err != nil {
        return fiber.NewError(fiber.StatusInternalServerError, err.Error())
    }

    resp := ProjectPeekResponse{P: req.ID, F: data}
    return c.Status(fiber.StatusOK).JSON(resp)
}