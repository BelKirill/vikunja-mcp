package main

import (
	"context"
	"os"
	"time"

	"github.com/BelKirill/vikunja-mcp/pkg/vikunja/client"
	"github.com/charmbracelet/log"
)

func main() {
	log.Info("Connecting to MY actual Vikunja instance...")

	// This is YOUR system, YOUR data
	c, err := client.NewClient()
	if err != nil {
		log.Error("Connection failed", "err", err)
		log.Warn("Check your VIKUNJA_URL and VIKUNJA_TOKEN env vars")
		os.Exit(1)
	}

	log.Info("Client created, making real API call...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Touch the real system
	tasks, err := c.GetAllTasks(ctx)
	if err != nil {
		log.Error("API call failed", "err", err)
		log.Warn("Let me see what the actual response looks like...")
		os.Exit(1)
	}

	log.Info("SUCCESS! Found real tasks from your K3s instance", "count", len(tasks))

	if len(tasks) > 0 {
		log.Info("Here's your first actual task:")
		log.Info("ID", "id", tasks[0].ID)
		log.Info("Title", "title", tasks[0].Title)
		log.Info("Description", "description", tasks[0].Description)
		log.Info("See? Your brain still works. The magic is still there.")
	}
}
