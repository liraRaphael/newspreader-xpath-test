package collector

import (
	"net/url"
	"newspreader/models"
	"strings"

	"github.com/gocolly/colly"
)

func GetArticleInfo(u string) models.Artigo {
	news := map[string]models.NewsSource{
		"www.estadao.com.br": {
			AuthorSelector: ".hide-on-mobile .authors-names",
			ImageLink:      ".figure-image-container img",
			TitleSelector:  "h1 .cover-titulo",
			TextSelector:   ".news-body p",
		},
		"www.folha.uol.com.br": {
			AuthorSelector: ".c-signature .c-signature__author a",
			ImageLink:      ".widget-image figure .c-image-aspect-ratio noscript img .img-responsive",
			TitleSelector:  "h1 .c-content-head__title",
			TextSelector:   ".c-news__content .c-news__body p",
		},
		"www1.folha.uol.com.br": {
			AuthorSelector: ".c-signature .c-signature__author a",
			ImageLink:      ".widget-image figure .c-image-aspect-ratio noscript img .img-responsive",
			TitleSelector:  "h1 .c-content-head__title",
			TextSelector:   ".c-news__content .c-news__body p",
		},
		"www.correiobraziliense.com.br": {
			AuthorSelector: ".article .autor .item .name",
			ImageLink:      ".responsive-img picture .cb-article-destaque",
			TitleSelector:  ".materia-title h1",
			TextSelector:   ".article p.texto",
		},
		"www.cnnbrasil.com.br": {
			AuthorSelector: ".author__content address .author__image .author__info .author__name .author__group span a",
			ImageLink:      "article .post__header .img__destaque img",
			TitleSelector:  "article .post__header h1.post__title",
			TextSelector:   "article .post__content p",
		},
	}

	url_detail, _ := url.Parse(u)
	scheme := url_detail.Scheme + "://"
	host := url_detail.Host

	selector := models.Selector{Source: news}
	bs := selector.Scrape(host)

	return ScrapeContent(u, host, scheme, bs)
}

func ScrapeContent(
	article_url string,
	host string,
	scheme string,
	bs models.NewsSource) models.Artigo {
	c := colly.NewCollector()
	artigo := models.Artigo{}

	// Autor do artigo
	c.OnHTML(bs.AuthorSelector, func(e *colly.HTMLElement) {
		artigo.Author = e.Text
	})

	// Imagem do artigo
	c.OnHTML(bs.ImageLink, func(e *colly.HTMLElement) {
		img_link := e.Attr("src")

		if strings.Contains(img_link, "https://") {
			artigo.Image = img_link
		} else {
			artigo.Image = scheme + host + img_link
		}
	})

	// Título do artigo
	c.OnHTML(bs.TitleSelector, func(e *colly.HTMLElement) {
		artigo.Title = e.Text
	})

	// Texto do artigo
	c.OnHTML(bs.TextSelector, func(e *colly.HTMLElement) {
		if artigo.Text != "" {
			artigo.Text += " "
		}

		artigo.Text += e.Text
	})

	c.Visit(article_url)

	return artigo
}
