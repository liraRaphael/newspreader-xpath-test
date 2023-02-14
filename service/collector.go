package collector

import (
	"net/url"
	"newspreader/models"
	"strings"

	"github.com/gocolly/colly"
)

func SetCollector(u string) models.Artigo {
	c := colly.NewCollector()
	artigo := models.Artigo{}

	url_detail, _ := url.Parse(u)
	scheme := url_detail.Scheme + "://"
	host := url_detail.Host

	// Autor do artigo
	c.OnHTML(".hide-on-mobile .authors-names", func(e *colly.HTMLElement) {
		artigo.Author = e.Text
	})

	// Imagem do artigo
	c.OnHTML(".figure-image-container img", func(e *colly.HTMLElement) {
		img_link := e.Attr("src")

		if strings.Contains(img_link, "https://") {
			artigo.Image = img_link
		} else {
			artigo.Image = scheme + host + img_link
		}
	})

	// Texto do artigo
	c.OnHTML(".news-body p", func(e *colly.HTMLElement) {
		if artigo.Text != "" {
			artigo.Text += " "
		}

		artigo.Text += e.Text
	})

	c.Visit(u)

	return artigo
}
