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

	// Focus endpoints
	api.Post("/daily-focus", DailyFocusHandler(service))
	api.Get("/focus/recommendation", FocusRecommendationHandler(service))
	api.Post("/focus/session", StartFocusSessionHandler(service))
	api.Put("/focus/session/:id", CompleteFocusSessionHandler(service))
}

// DailyFocusHandler handles POST /daily-focus
// @Summary Get daily focus task recommendations
// @Description Returns a list of tasks optimized for focus sessions based on energy and available time
// @Tags focus
// @Accept json
// @Produce json
// @Param request body models.FocusRequest true "Focus request parameters"
// @Success 200 {array} models.FocusResponseItem "List of recommended tasks"
// @Failure 400 {object} models.APIError "Invalid request"
// @Failure 500 {object} models.APIError "Internal server error"
// @Router /daily-focus [post]
func DailyFocusHandler(service *Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req models.FocusRequest
		if err := c.BodyParser(&req); err != nil {
			log.Error("failed to parse focus request", "error", err)
			return c.Status(fiber.StatusBadRequest).JSON(models.APIError{
				Code:    fiber.StatusBadRequest,
				Message: "invalid JSON body",
			})
		}

		log.Debug("parsed focus request", "req", req)

		// Set defaults
		if req.Date == "" {
			req.Date = time.Now().Format("2006-01-02")
			log.Info("focus request using default date", "date", req.Date)
		}
		if req.Hours == 0 {
			req.Hours = 5
			log.Info("focus request using default hours", "hours", req.Hours)
		}

		log.Info("processing focus request", "date", req.Date, "hours", req.Hours)
		log.Debug("focus request context setup", "date", req.Date, "hours", req.Hours)

		ctx := context.Background()
		opts := models.FocusOptions{
			Date:       req.Date,
			MaxMinutes: int(req.Hours * 60), // Convert hours to minutes
			Energy:     "medium",            // Default energy - could be from query param
			Mode:       "deep",              // Default mode - could be from query param
			MaxTasks:   20,                  // Reasonable default
		}

		// Check for query parameters to override defaults
		if energy := c.Query("energy"); energy != "" {
			opts.Energy = energy
			log.Debug("overriding energy from query param", "energy", energy)
		}
		if mode := c.Query("mode"); mode != "" {
			opts.Mode = mode
			log.Debug("overriding mode from query param", "mode", mode)
		}

		log.Debug("calling service.GetFocusTasks", "opts", opts)
		items, err := service.GetFocusTasks(ctx, &opts)
		if err != nil {
			log.Error("focus service error", "date", req.Date, "hours", req.Hours, "error", err)
			return c.Status(fiber.StatusInternalServerError).JSON(models.APIError{
				Code:    fiber.StatusInternalServerError,
				Message: err.Error(),
			})
		}

		log.Debug("service.GetFocusTasks returned items", "count", len(items))

		// Convert to response format
		resp := make([]models.FocusResponseItem, 0, len(items))
		for _, item := range items {
			resp = append(resp, models.FocusResponseItem{
				T:   fmt.Sprintf("%d", item.RawTask.ID),
				P:   fmt.Sprintf("%d", item.RawTask.ProjectID),
				Est: float64(item.Metadata.Estimate) / 60.0, // Convert minutes to hours for API consistency
			})
		}

		log.Info("focus request completed", "date", req.Date, "items", len(items))
		log.Debug("focus response ready", "response", resp)
		return c.Status(fiber.StatusOK).JSON(resp)
	}
}

// FocusRecommendationHandler handles GET /focus/recommendation
// @Summary Get single best task recommendation
// @Description Returns the single most suitable task for the current focus session
// @Tags focus
// @Produce json
// @Param energy query string false "Energy level (low, medium, high, social)" default(medium)
// @Param mode query string false "Focus mode (deep, quick, admin)" default(deep)
// @Param minutes query int false "Available time in minutes" default(60)
// @Success 200 {object} models.SessionRecommendation "Task recommendation with session details"
// @Failure 500 {object} models.APIError "Internal server error"
// @Router /focus/recommendation [get]
func FocusRecommendationHandler(service *Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx := context.Background()

		// Parse query parameters
		opts := models.FocusOptions{
			Energy:     c.Query("energy", "medium"),
			Mode:       c.Query("mode", "deep"),
			MaxMinutes: c.QueryInt("minutes", 60),
			MaxTasks:   1, // Single recommendation
		}

		log.Debug("parsed recommendation options", "opts", opts)

		recommendation, err := service.GetTaskRecommendation(ctx, &opts)
		if err != nil {
			log.Error("failed to get task recommendation", "error", err)
			return c.Status(fiber.StatusInternalServerError).JSON(models.APIError{
				Code:    fiber.StatusInternalServerError,
				Message: err.Error(),
			})
		}

		if recommendation == nil {
			log.Debug("no suitable tasks found for recommendation", "opts", opts)
			return c.Status(fiber.StatusOK).JSON(models.APIError{
				Code:    fiber.StatusNoContent,
				Message: "no suitable tasks found for current criteria",
			})
		}

		log.Debug("recommendation found", "task_id", recommendation.RawTask.ID, "score", recommendation.FocusScore)

		response := models.SessionRecommendation{
			Task:      recommendation,
			CanExtend: recommendation.Metadata.Extend && opts.MaxMinutes >= recommendation.Metadata.Estimate,
			Reasoning: fmt.Sprintf("Selected based on %s energy, %s mode compatibility (score: %.1f)",
				opts.Energy, opts.Mode, recommendation.FocusScore),
		}

		log.Info("task recommendation generated",
			"task_id", recommendation.RawTask.ID,
			"score", recommendation.FocusScore)
		log.Debug("recommendation response ready", "response", response)

		return c.Status(fiber.StatusOK).JSON(response)
	}
}

// StartFocusSessionHandler handles POST /focus/session
// @Summary Start a new focus session
// @Description Creates a new focus session record for tracking
// @Tags focus
// @Accept json
// @Produce json
// @Param session body models.FocusSession true "Focus session details"
// @Success 201 {object} models.FocusSession "Created focus session"
// @Failure 400 {object} models.APIError "Invalid request"
// @Failure 500 {object} models.APIError "Internal server error"
// @Router /focus/session [post]
func StartFocusSessionHandler(service *Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var session models.FocusSession
		if err := c.BodyParser(&session); err != nil {
			log.Error("failed to parse focus session", "error", err)
			return c.Status(fiber.StatusBadRequest).JSON(models.APIError{
				Code:    fiber.StatusBadRequest,
				Message: "invalid JSON body",
			})
		}

		log.Debug("parsed focus session", "session", session)

		// Set session start time and generate ID
		session.ID = fmt.Sprintf("session_%d_%d", session.TaskID, time.Now().Unix())
		session.StartTime = time.Now().Format(time.RFC3339)
		session.Completed = false

		log.Debug("focus session initialized", "session_id", session.ID, "start_time", session.StartTime)

		// TODO: Store session in database/memory for tracking
		log.Info("focus session started",
			"session_id", session.ID,
			"task_id", session.TaskID,
			"planned_length", session.PlannedLength)

		return c.Status(fiber.StatusCreated).JSON(session)
	}
}

// CompleteFocusSessionHandler handles PUT /focus/session/:id
// @Summary Complete a focus session
// @Description Marks a focus session as completed and records performance data
// @Tags focus
// @Accept json
// @Produce json
// @Param id path string true "Session ID"
// @Param completion body models.FocusSession true "Session completion data"
// @Success 200 {object} models.FocusSession "Updated focus session"
// @Failure 400 {object} models.APIError "Invalid request"
// @Failure 404 {object} models.APIError "Session not found"
// @Failure 500 {object} models.APIError "Internal server error"
// @Router /focus/session/{id} [put]
func CompleteFocusSessionHandler(service *Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		sessionID := c.Params("id")

		var completion models.FocusSession
		if err := c.BodyParser(&completion); err != nil {
			log.Error("failed to parse session completion", "error", err)
			return c.Status(fiber.StatusBadRequest).JSON(models.APIError{
				Code:    fiber.StatusBadRequest,
				Message: "invalid JSON body",
			})
		}

		log.Debug("parsed session completion", "completion", completion)

		// Set completion data
		completion.ID = sessionID
		completion.EndTime = time.Now().Format(time.RFC3339)
		completion.Completed = true

		log.Debug("completion data set", "session_id", sessionID, "end_time", completion.EndTime)

		// Calculate actual length if not provided
		if completion.ActualLength == 0 && completion.StartTime != "" {
			startTime, err := time.Parse(time.RFC3339, completion.StartTime)
			if err == nil {
				completion.ActualLength = int(time.Since(startTime).Minutes())
				log.Debug("calculated actual session length", "actual_length", completion.ActualLength)
			} else {
				log.Debug("failed to parse start time for actual length calculation", "error", err)
			}
		}

		// TODO: Update session in database and update learning algorithms
		log.Info("focus session completed",
			"session_id", sessionID,
			"actual_length", completion.ActualLength,
			"effectiveness", completion.Effectiveness)
		log.Debug("completion response ready", "completion", completion)

		return c.Status(fiber.StatusOK).JSON(completion)
	}
}
