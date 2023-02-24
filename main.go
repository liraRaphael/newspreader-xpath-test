package main

import (
	"embed"
	"html/template"
	"newspreader/models"
	collector "newspreader/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	fhtml "github.com/gofiber/template/html"
)

var embedDirStatic embed.FS

func main() {
	engine := fhtml.New("./views", ".html")
	engine.AddFunc(
		"render", func(s string) template.HTML {
			return template.HTML(s)
		},
	)

	logger := logger.New()

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Use(logger)
	app.Static(
		"/assets",
		"./views/assets",
	)

	app.Get("/api/paynot", func(c *fiber.Ctx) error {
		if c.Query("url") != "" {
			return c.JSON(collector.GetArticleInfo(c.Query("url")))
		}

		return c.Status(400).JSON(models.Error{Message: "Campo 'url' é obrigatório para a consulta."})
	})

	app.Get("/paynot", func(c *fiber.Ctx) error {
		if c.Query("url") != "" {
			article := collector.GetArticleInfo(c.Query("url"))

			return c.Render("artigo", fiber.Map{
				"Author":    article.Author,
				"Media":     article.Media,
				"MediaType": article.MediaType,
				"Title":     article.Title,
				"Text":      article.GetHtmlPreparedText(),
			})
		}

		return c.Status(400).JSON(models.Error{Message: "Campo 'url' é obrigatório para a consulta."})
	})

	app.Listen(":3000")
}
