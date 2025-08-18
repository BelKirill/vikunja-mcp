package focus

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/BelKirill/vikunja-mcp/models"
	"github.com/charmbracelet/log"
)

const ExportVersion = "1.0"

// ExportProjectToJSON exports a single project's tasks to a local JSON file
func (s *Service) ExportProjectToJSON(ctx context.Context, projectID int64, options models.ExportOptions) (*models.ExportResult, error) {
	log.Info("ExportProjectToJSON called", "project_id", projectID, "output_path", options.OutputPath)

	// Validate output path
	if options.OutputPath == "" {
		return nil, fmt.Errorf("output_path is required")
	}

	// Ensure output directory exists
	if err := ensureDirectoryExists(options.OutputPath); err != nil {
		log.Error("Failed to ensure output directory exists", "error", err)
		return nil, fmt.Errorf("failed to create output directory: %w", err)
	}

	// Fetch project metadata
	project, err := s.Vikunja.GetProject(ctx, projectID)
	if err != nil {
		log.Error("Failed to fetch project", "project_id", projectID, "error", err)
		return nil, fmt.Errorf("failed to fetch project: %w", err)
	}

	// Build filter for project tasks
	filter := fmt.Sprintf("project_id = %d", projectID)
	if !options.IncludeCompleted {
		filter += " && done = false"
	}

	// Apply custom filter if provided
	if options.CustomFilter != "" {
		filter = fmt.Sprintf("(%s) && (%s)", filter, options.CustomFilter)
	}

	log.Debug("Using filter for export", "filter", filter)

	// Fetch and process tasks
	rawTasks, err := s.Vikunja.GetFilteredTasks(ctx, &filter)
	if err != nil {
		log.Error("Failed to fetch filtered tasks", "filter", filter, "error", err)
		return nil, fmt.Errorf("failed to fetch tasks: %w", err)
	}

	log.Info("Fetched tasks for export", "count", len(rawTasks))

	// Enrich tasks with metadata if requested
	var enrichedTasks []models.Task
	if options.IncludeMetadata {
		enrichedTasks = s.EnrichTasksParallel(ctx, rawTasks)
	} else {
		// Convert raw tasks to Task format without enrichment
		for _, rawTask := range rawTasks {
			enrichedTasks = append(enrichedTasks, models.Task{
				Identifier:       rawTask.Identifier,
				RawTask:          &rawTask,
				CleanDescription: rawTask.Description,
				Metadata:         nil,
			})
		}
	}

	// Collect comments if requested
	comments := make(map[int64][]models.Comment)
	if options.IncludeComments {
		log.Debug("Fetching comments for tasks", "task_count", len(enrichedTasks))
		for _, task := range enrichedTasks {
			taskComments, err := s.Vikunja.GetTaskCommentsByID(ctx, task.RawTask.ID)
			if err != nil {
				log.Warn("Failed to fetch comments for task", "task_id", task.RawTask.ID, "error", err)
				continue
			}
			if len(taskComments) > 0 {
				comments[task.RawTask.ID] = taskComments
			}
		}
		log.Debug("Collected comments", "total_comments", len(comments))
	}

	// Collect unique labels from tasks
	labelMap := make(map[int]models.PartialLabel)
	for _, task := range enrichedTasks {
		for _, label := range task.RawTask.Labels {
			labelMap[label.ID] = label
		}
	}

	// Convert label map to slice
	var labels []models.PartialLabel
	for _, label := range labelMap {
		labels = append(labels, label)
	}

	// Build export structure
	export := &models.ProjectExport{
		Project:  project,
		Tasks:    enrichedTasks,
		Labels:   labels,
		Comments: comments,
		ExportMeta: models.ExportMetadata{
			ExportedAt:    time.Now(),
			TaskCount:     len(enrichedTasks),
			IncludesDone:  options.IncludeCompleted,
			FilterUsed:    filter,
			ExportVersion: ExportVersion,
			ProjectTitle:  project.Title,
		},
	}

	// Write to file
	fileSize, err := writeExportToFile(export, options.OutputPath, options.PrettyPrint)
	if err != nil {
		log.Error("Failed to write export to file", "error", err)
		return nil, fmt.Errorf("failed to write export file: %w", err)
	}

	result := &models.ExportResult{
		Success:      true,
		FilePath:     options.OutputPath,
		TaskCount:    len(enrichedTasks),
		ProjectCount: 1,
		FileSize:     fileSize,
	}

	log.Info("Export completed successfully",
		"file_path", options.OutputPath,
		"task_count", len(enrichedTasks),
		"file_size", fileSize)

	return result, nil
}

// ExportFilteredTasksToJSON exports tasks based on a custom filter
func (s *Service) ExportFilteredTasksToJSON(ctx context.Context, filter string, options models.ExportOptions) (*models.ExportResult, error) {
	log.Info("ExportFilteredTasksToJSON called", "filter", filter, "output_path", options.OutputPath)

	// Validate output path
	if options.OutputPath == "" {
		return nil, fmt.Errorf("output_path is required")
	}

	// Ensure output directory exists
	if err := ensureDirectoryExists(options.OutputPath); err != nil {
		log.Error("Failed to ensure output directory exists", "error", err)
		return nil, fmt.Errorf("failed to create output directory: %w", err)
	}

	// Build final filter
	finalFilter := filter
	if !options.IncludeCompleted && !strings.Contains(strings.ToLower(filter), "done") {
		finalFilter = fmt.Sprintf("(%s) && done = false", filter)
	}

	// Apply custom filter if provided
	if options.CustomFilter != "" {
		finalFilter = fmt.Sprintf("(%s) && (%s)", finalFilter, options.CustomFilter)
	}

	log.Debug("Using filter for export", "filter", finalFilter)

	// Fetch and process tasks
	rawTasks, err := s.Vikunja.GetFilteredTasks(ctx, &finalFilter)
	if err != nil {
		log.Error("Failed to fetch filtered tasks", "filter", finalFilter, "error", err)
		return nil, fmt.Errorf("failed to fetch tasks: %w", err)
	}

	log.Info("Fetched tasks for export", "count", len(rawTasks))

	// Enrich tasks with metadata if requested
	var enrichedTasks []models.Task
	if options.IncludeMetadata {
		enrichedTasks = s.EnrichTasksParallel(ctx, rawTasks)
	} else {
		// Convert raw tasks to Task format without enrichment
		for _, rawTask := range rawTasks {
			enrichedTasks = append(enrichedTasks, models.Task{
				Identifier:       rawTask.Identifier,
				RawTask:          &rawTask,
				CleanDescription: rawTask.Description,
				Metadata:         nil,
			})
		}
	}

	// Collect comments if requested
	comments := make(map[int64][]models.Comment)
	if options.IncludeComments {
		log.Debug("Fetching comments for tasks", "task_count", len(enrichedTasks))
		for _, task := range enrichedTasks {
			taskComments, err := s.Vikunja.GetTaskCommentsByID(ctx, task.RawTask.ID)
			if err != nil {
				log.Warn("Failed to fetch comments for task", "task_id", task.RawTask.ID, "error", err)
				continue
			}
			if len(taskComments) > 0 {
				comments[task.RawTask.ID] = taskComments
			}
		}
		log.Debug("Collected comments", "total_comments", len(comments))
	}

	// Collect unique labels from tasks
	labelMap := make(map[int]models.PartialLabel)
	for _, task := range enrichedTasks {
		for _, label := range task.RawTask.Labels {
			labelMap[label.ID] = label
		}
	}

	// Convert label map to slice
	var labels []models.PartialLabel
	for _, label := range labelMap {
		labels = append(labels, label)
	}

	// Build export structure (no specific project for filtered export)
	export := &models.ProjectExport{
		Project:  nil, // No specific project for filtered export
		Tasks:    enrichedTasks,
		Labels:   labels,
		Comments: comments,
		ExportMeta: models.ExportMetadata{
			ExportedAt:    time.Now(),
			TaskCount:     len(enrichedTasks),
			IncludesDone:  options.IncludeCompleted,
			FilterUsed:    finalFilter,
			ExportVersion: ExportVersion,
			ProjectTitle:  "Filtered Export",
		},
	}

	// Write to file
	fileSize, err := writeExportToFile(export, options.OutputPath, options.PrettyPrint)
	if err != nil {
		log.Error("Failed to write export to file", "error", err)
		return nil, fmt.Errorf("failed to write export file: %w", err)
	}

	result := &models.ExportResult{
		Success:      true,
		FilePath:     options.OutputPath,
		TaskCount:    len(enrichedTasks),
		ProjectCount: 0, // Filtered export doesn't target specific projects
		FileSize:     fileSize,
	}

	log.Info("Filtered export completed successfully",
		"file_path", options.OutputPath,
		"task_count", len(enrichedTasks),
		"file_size", fileSize)

	return result, nil
}

// ExportAllProjectsToJSON exports all user projects to a single JSON file
func (s *Service) ExportAllProjectsToJSON(ctx context.Context, options models.ExportOptions) (*models.ExportResult, error) {
	log.Info("ExportAllProjectsToJSON called", "output_path", options.OutputPath)

	// Validate output path
	if options.OutputPath == "" {
		return nil, fmt.Errorf("output_path is required")
	}

	// Ensure output directory exists
	if err := ensureDirectoryExists(options.OutputPath); err != nil {
		log.Error("Failed to ensure output directory exists", "error", err)
		return nil, fmt.Errorf("failed to create output directory: %w", err)
	}

	// Fetch all projects
	projects, err := s.Vikunja.GetAllProjects(ctx)
	if err != nil {
		log.Error("Failed to fetch all projects", "error", err)
		return nil, fmt.Errorf("failed to fetch projects: %w", err)
	}

	log.Info("Fetched projects for export", "count", len(projects))

	var projectExports []models.ProjectExport
	totalTasks := 0

	// Export each project
	for _, project := range projects {
		log.Debug("Exporting project", "project_id", project.ID, "title", project.Title)

		// Create temporary options for individual project export
		tempOptions := options
		tempOptions.OutputPath = "" // Don't write individual files

		// Use ExportProjectToJSON logic but don't write to file
		result, err := s.buildProjectExport(ctx, int64(project.ID), tempOptions)
		if err != nil {
			log.Warn("Failed to export project", "project_id", project.ID, "error", err)
			continue
		}

		projectExports = append(projectExports, *result)
		totalTasks += result.ExportMeta.TaskCount
	}

	// Build multi-project export structure
	multiExport := &models.MultiProjectExport{
		Projects: projectExports,
		ExportMeta: models.ExportMetadata{
			ExportedAt:    time.Now(),
			TaskCount:     totalTasks,
			IncludesDone:  options.IncludeCompleted,
			ExportVersion: ExportVersion,
			ProjectTitle:  "All Projects Export",
		},
	}

	// Write to file
	fileSize, err := writeMultiExportToFile(multiExport, options.OutputPath, options.PrettyPrint)
	if err != nil {
		log.Error("Failed to write multi-export to file", "error", err)
		return nil, fmt.Errorf("failed to write export file: %w", err)
	}

	result := &models.ExportResult{
		Success:      true,
		FilePath:     options.OutputPath,
		TaskCount:    totalTasks,
		ProjectCount: len(projectExports),
		FileSize:     fileSize,
	}

	log.Info("Multi-project export completed successfully",
		"file_path", options.OutputPath,
		"project_count", len(projectExports),
		"task_count", totalTasks,
		"file_size", fileSize)

	return result, nil
}

// buildProjectExport builds a project export without writing to file
func (s *Service) buildProjectExport(ctx context.Context, projectID int64, options models.ExportOptions) (*models.ProjectExport, error) {
	// Fetch project metadata
	project, err := s.Vikunja.GetProject(ctx, projectID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch project: %w", err)
	}

	// Build filter for project tasks
	filter := fmt.Sprintf("project_id = %d", projectID)
	if !options.IncludeCompleted {
		filter += " && done = false"
	}

	// Apply custom filter if provided
	if options.CustomFilter != "" {
		filter = fmt.Sprintf("(%s) && (%s)", filter, options.CustomFilter)
	}

	// Fetch and process tasks
	rawTasks, err := s.Vikunja.GetFilteredTasks(ctx, &filter)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch tasks: %w", err)
	}

	// Enrich tasks with metadata if requested
	var enrichedTasks []models.Task
	if options.IncludeMetadata {
		enrichedTasks = s.EnrichTasksParallel(ctx, rawTasks)
	} else {
		// Convert raw tasks to Task format without enrichment
		for _, rawTask := range rawTasks {
			enrichedTasks = append(enrichedTasks, models.Task{
				Identifier:       rawTask.Identifier,
				RawTask:          &rawTask,
				CleanDescription: rawTask.Description,
				Metadata:         nil,
			})
		}
	}

	// Collect comments if requested
	comments := make(map[int64][]models.Comment)
	if options.IncludeComments {
		for _, task := range enrichedTasks {
			taskComments, err := s.Vikunja.GetTaskCommentsByID(ctx, task.RawTask.ID)
			if err != nil {
				log.Warn("Failed to fetch comments for task", "task_id", task.RawTask.ID, "error", err)
				continue
			}
			if len(taskComments) > 0 {
				comments[task.RawTask.ID] = taskComments
			}
		}
	}

	// Collect unique labels from tasks
	labelMap := make(map[int]models.PartialLabel)
	for _, task := range enrichedTasks {
		for _, label := range task.RawTask.Labels {
			labelMap[label.ID] = label
		}
	}

	// Convert label map to slice
	var labels []models.PartialLabel
	for _, label := range labelMap {
		labels = append(labels, label)
	}

	// Build export structure
	export := &models.ProjectExport{
		Project:  project,
		Tasks:    enrichedTasks,
		Labels:   labels,
		Comments: comments,
		ExportMeta: models.ExportMetadata{
			ExportedAt:    time.Now(),
			TaskCount:     len(enrichedTasks),
			IncludesDone:  options.IncludeCompleted,
			FilterUsed:    filter,
			ExportVersion: ExportVersion,
			ProjectTitle:  project.Title,
		},
	}

	return export, nil
}

// writeExportToFile writes a ProjectExport to a JSON file
func writeExportToFile(export *models.ProjectExport, outputPath string, prettyPrint bool) (int64, error) {
	var jsonData []byte
	var err error

	if prettyPrint {
		jsonData, err = json.MarshalIndent(export, "", "  ")
	} else {
		jsonData, err = json.Marshal(export)
	}

	if err != nil {
		return 0, fmt.Errorf("failed to marshal export data: %w", err)
	}

	err = os.WriteFile(outputPath, jsonData, 0644)
	if err != nil {
		return 0, fmt.Errorf("failed to write file: %w", err)
	}

	return int64(len(jsonData)), nil
}

// writeMultiExportToFile writes a MultiProjectExport to a JSON file
func writeMultiExportToFile(export *models.MultiProjectExport, outputPath string, prettyPrint bool) (int64, error) {
	var jsonData []byte
	var err error

	if prettyPrint {
		jsonData, err = json.MarshalIndent(export, "", "  ")
	} else {
		jsonData, err = json.Marshal(export)
	}

	if err != nil {
		return 0, fmt.Errorf("failed to marshal export data: %w", err)
	}

	err = os.WriteFile(outputPath, jsonData, 0644)
	if err != nil {
		return 0, fmt.Errorf("failed to write file: %w", err)
	}

	return int64(len(jsonData)), nil
}

// ensureDirectoryExists creates the directory for the output path if it doesn't exist
func ensureDirectoryExists(outputPath string) error {
	dir := filepath.Dir(outputPath)
	if dir == "." {
		return nil // Current directory
	}

	return os.MkdirAll(dir, 0755)
}
