package main

import (
	"newspreader/models"
	collector "newspreader/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
	logger := logger.New()

	app.Use(logger)

	app.Get("/paynot", func(c *fiber.Ctx) error {
		if c.Query("url") != "" {
			return c.JSON(collector.SetCollector(c.Query("url")))
		}

		return c.Status(400).JSON(models.Error{Message: "Campo 'url' é obrigatório para a consulta."})
	})

	app.Listen(":3000")
}
