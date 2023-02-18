package models

type Artigo struct {
	Author string `json:"author"`
	Image  string `json:"image"`
	Title  string `json:"title"`
	Text   string `json:"text"`
}
