package main

import (
    "os"

		"github.com/charmbracelet/log"
    "github.com/gofiber/fiber/v2"
    "github.com/BelKirill/vikunja-mcp/internal/peek"
)

func main() {
    app := fiber.New()

    // peek routes
    peek.RegisterRoutes(app)

    port := os.Getenv("PORT")
    if port == "" {
        port = "8000"
    }

    log.Info("starting MCP server", "port", port)
    if err := app.Listen(":" + port); err != nil {
        log.Fatal(err)
    }
}