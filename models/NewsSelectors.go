package models

type NewsSource struct {
	AuthorSelector string `json:"authorSelector"`
	MediaLink      string `json:"mediaLink"`
	MediaAttr      string `json:"mediaAttr"`
	TitleSelector  string `json:"titleSelector"`
	TextSelector   string `json:"textSelector"`
	Source         string `json:"source"`
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
