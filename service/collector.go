package collector

import (
	"net/url"
	"newspreader/models"
	"regexp"
	"strings"

	"github.com/gocolly/colly"
)

func GetArticleInfo(u string) models.Artigo {
	news := map[string]models.NewsSource{
		"www.estadao.com.br": {
			AuthorSelector: ".hide-on-mobile .authors-names",
			MediaLink:      "#fusion-app div .news-header",
			MediaAttr:      "style",
			TitleSelector:  ".news-header h1.cover-titulo, .columnist-header .columnist-container .columnist-title, .content-header-entrevista .left-container-header-entrevista .left-content h1",
			TextSelector:   ".news-body p",
			Source:         "Estadao",
		},
		"www.folha.uol.com.br": {
			AuthorSelector: ".c-signature .c-signature__author a, .c-top-columnist__wrapper .c-top-columnist__content .c-top-columnist__name a",
			MediaLink:      ".widget-image figure .c-image-aspect-ratio.c-image-aspect-ratio--3x2 img",
			MediaAttr:      "data-src",
			TitleSelector:  ".c-content-head .c-content-head__wrap h1.c-content-head__title",
			TextSelector:   ".c-news__content .c-news__body p",
			Source:         "Folha",
		},
		"www1.folha.uol.com.br": {
			AuthorSelector: ".c-signature .c-signature__author a, .c-top-columnist__wrapper .c-top-columnist__content .c-top-columnist__name a",
			MediaLink:      ".widget-image figure .c-image-aspect-ratio.c-image-aspect-ratio--3x2 img",
			MediaAttr:      "data-src",
			TitleSelector:  ".c-content-head .c-content-head__wrap h1.c-content-head__title",
			TextSelector:   ".c-news__content .c-news__body p",
			Source:         "Folha",
		},
		"www.correiobraziliense.com.br": {
			AuthorSelector: ".article .autor .item .name",
			MediaLink:      ".responsive-img picture img.cb-article-destaque",
			MediaAttr:      "data-src",
			TitleSelector:  ".materia-title h1",
			TextSelector:   ".article p.texto",
			Source:         "Correio Braziliense",
		},
		"www.cnnbrasil.com.br": {
			AuthorSelector: ".author__content address .author__image .author__info .author__name .author__group span a",
			MediaLink:      "article .post__header .img__destaque img, article .post__header .post__video .overlay-wrapper .video-button img",
			MediaAttr:      "src",
			TitleSelector:  "article .post__header h1.post__title",
			TextSelector:   "article .post__content p",
			Source:         "CNN Brasil",
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
	c.OnHTML(bs.MediaLink, func(e *colly.HTMLElement) {
		media_link := e.Attr(bs.MediaAttr)

		if bs.Source == "Estadao" {
			re := regexp.MustCompile(`url\("(\S+)"\)`)
			match := re.FindStringSubmatch(media_link)

			media_link = scheme + host + match[1]
		}

		is_youtube_link := strings.Contains(media_link, "youtube")

		if bs.Source == "CNN Brasil" && is_youtube_link {
			re := regexp.MustCompile(`vi/(\S+)/sddefault`)
			match := re.FindStringSubmatch(media_link)

			media_link = "https://www.youtube.com/embed/" + match[1]
		}

		artigo.Media = media_link
		artigo.MediaType = "Image"

		if is_youtube_link {
			artigo.MediaType = "Video"
		}
	})

	// Título do artigo
	c.OnHTML(bs.TitleSelector, func(e *colly.HTMLElement) {
		artigo.Title = e.Text
	})

	// Texto do artigo
	c.OnHTML(bs.TextSelector, func(e *colly.HTMLElement) {
		if artigo.Text != "" {
			artigo.Text += "\n"
		}

		artigo.Text += e.Text
	})

	c.Visit(article_url)

	return artigo
}
