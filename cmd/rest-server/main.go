// @title        Vikunja MCP Server API
// @version      1.0
// @description  MCP server exposing micro-tools for Vikunja integration
// @host         localhost:8000
// @BasePath     /
package main

import (
	"os"

	"github.com/BelKirill/vikunja-mcp/docs"
	"github.com/BelKirill/vikunja-mcp/internal/focus"
	"github.com/charmbracelet/log"
	fiber "github.com/gofiber/fiber/v2"
	fiberSwagger "github.com/gofiber/swagger"
)

func main() {
	app := fiber.New()

	app.Get("/openapi.json", func(c *fiber.Ctx) error {
		c.Type("json")
		return c.SendString(docs.SwaggerInfo.ReadDoc())
	})

	app.Get("/docs/*", fiberSwagger.HandlerDefault)

	// tool routes
	focus.RegisterRoutes(app)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	log.Info("starting MCP server", "port", port)
	if err := app.Listen(":" + port); err != nil {
		log.Fatal(err)
	}
}
