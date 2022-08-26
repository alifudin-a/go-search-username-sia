package routes

import (
	restsearch "github.com/alifudin-a/go-search-username-sia/pkg/http/rest/search-username-sia"

	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	logger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html"
)

func Routes() {
	engine := html.New("./pkg/static/views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(cors.New())

	app.Get("/", func(c *fiber.Ctx) error {
		// Render index template
		return c.Render("index", fiber.Map{
			"Title": "Hello, World!",
		})
	})

	sia := app.Group("/sia")

	sia.Post("/user/search", restsearch.NewSearchUsernameSIAHandler().SearchUsernameSIAHandler)

	app.Listen(":" + os.Getenv("APP_PORT"))
}
