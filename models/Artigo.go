package models

type Artigo struct {
	Author    string `json:"author"`
	Media     string `json:"media"`
	MediaType string `json:"mediaType"`
	Title     string `json:"title"`
	Text      string `json:"text"`
}
