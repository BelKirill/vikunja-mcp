package peek

import (
	"github.com/charmbracelet/log"
	"github.com/gofiber/fiber/v2"
)

// RegisterRoutes attaches peek endpoints.
func RegisterRoutes(app *fiber.App) {
	api := app.Group("/api/v1")
	api.Post("/task-peek", taskPeekHandler)
	api.Post("/project-peek", projectPeekHandler)
}

// @Summary      Peek a task
// @Description  Fetch selected fields for a given task ID
// @Tags         peek
// @Accept       json
// @Produce      json
// @Param        request  body      TaskPeekRequest  true  "Task peek request payload"
// @Success      200      {object}  TaskPeekResponse
// @Failure      400      {object}  APIError
// @Failure      500      {object}  APIError
// @Router       /api/v1/task-peek [post]
func taskPeekHandler(c *fiber.Ctx) error {
	var req TaskPeekRequest
	if err := c.BodyParser(&req); err != nil {
		log.Error("failed to parse task peek request", "error", err)
		return c.Status(fiber.StatusBadRequest).JSON(APIError{Code: fiber.StatusBadRequest, Message: "invalid JSON body"})
	}
	if req.ID == "" {
		log.Error("task peek request missing ID")
		return c.Status(fiber.StatusBadRequest).JSON(APIError{Code: fiber.StatusBadRequest, Message: "id is required"})
	}
	if len(req.Fields) == 0 {
		req.Fields = []string{"name"}
		log.Info("task peek request using default fields", "task_id", req.ID)
	}

	log.Info("processing task peek request", "task_id", req.ID, "fields", req.Fields)
	data, err := taskService.Peek(req.ID, req.Fields)
	if err != nil {
		log.Error("task peek service error", "task_id", req.ID, "error", err)
		return c.Status(fiber.StatusInternalServerError).JSON(APIError{Code: fiber.StatusInternalServerError, Message: err.Error()})
	}

	resp := TaskPeekResponse{T: req.ID, F: data}
	log.Info("task peek request completed", "task_id", req.ID)
	return c.Status(fiber.StatusOK).JSON(resp)
}

// @Summary      Peek a project
// @Description  Fetch selected fields for a given project ID
// @Tags         peek
// @Accept       json
// @Produce      json
// @Param        request  body      ProjectPeekRequest  true  "Project peek request payload"
// @Success      200      {object}  ProjectPeekResponse
// @Failure      400      {object}  APIError
// @Failure      500      {object}  APIError
// @Router       /api/v1/project-peek [post]
func projectPeekHandler(c *fiber.Ctx) error {
	var req ProjectPeekRequest
	if err := c.BodyParser(&req); err != nil {
		log.Error("failed to parse project peek request", "error", err)
		return c.Status(fiber.StatusBadRequest).JSON(APIError{Code: fiber.StatusBadRequest, Message: "invalid JSON body"})
	}
	if req.ID == "" {
		log.Error("project peek request missing ID")
		return c.Status(fiber.StatusBadRequest).JSON(APIError{Code: fiber.StatusBadRequest, Message: "id is required"})
	}
	if len(req.Fields) == 0 {
		req.Fields = []string{"name"}
		log.Info("project peek request using default fields", "project_id", req.ID)
	}

	log.Info("processing project peek request", "project_id", req.ID, "fields", req.Fields)
	data, err := projectService.Peek(req.ID, req.Fields)
	if err != nil {
		log.Error("project peek service error", "project_id", req.ID, "error", err)
		return c.Status(fiber.StatusInternalServerError).JSON(APIError{Code: fiber.StatusInternalServerError, Message: err.Error()})
	}

	resp := ProjectPeekResponse{P: req.ID, F: data}
	log.Info("project peek request completed", "project_id", req.ID)
	return c.Status(fiber.StatusOK).JSON(resp)
}
