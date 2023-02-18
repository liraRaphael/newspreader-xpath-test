package models

type NewsSource struct {
	AuthorSelector string `json:"authorSelector"`
	ImageLink      string `json:"imageLink"`
	TitleSelector  string `json:"titleSelector"`
	TextSelector   string `json:"textSelector"`
}

type Selector struct {
	Source map[string]NewsSource
}

func (s *Selector) Scrape(newsSource string) NewsSource {
	news, ok := s.Source[newsSource]
	if !ok {
		// error
	}

	return news
}
