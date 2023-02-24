package models

import "strings"

type Artigo struct {
	Author    string `json:"author"`
	Media     string `json:"media"`
	MediaType string `json:"mediaType"`
	Title     string `json:"title"`
	Text      string `json:"text"`
}

func (a *Artigo) GetHtmlPreparedText() string {
	clean_up := strings.Replace(a.Text, "\n\n", "\n", -1)

	return strings.Replace(clean_up, "\n", "</br></br>", -1)
}
