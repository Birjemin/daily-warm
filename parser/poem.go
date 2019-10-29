package parser

import (
	"encoding/json"
	"log"
	"strings"
)

// PoemRes response data
type PoemRes struct {
	Status string `json:"status"`
	Data   struct {
		Origin struct {
			Title   string   `json:"title"`
			Dynasty string   `json:"dynasty"`
			Author  string   `json:"author"`
			Content []string `json:"content"`
		} `json:"origin"`
	} `json:"data"`
}

type Poem struct {
	name string
	url  string
}

func NewPoem() *Poem {
	return &Poem{
		name: "poem",
		url:  "https://v2.jinrishici.com/one.json",
	}
}

// name
func (p *Poem) Name() string {
	return p.name
}

// route
func (p *Poem) Route() string {
	return p.url
}

// parse
func (p *Poem) Parse(body []byte) map[string]string {
	var resJSON PoemRes
	err := json.Unmarshal(body, &resJSON)
	if err != nil {
		log.Fatalf("Fetch json from %s error: %s", p.Route(), err)
	}
	status := resJSON.Status
	if status != "success" {
		log.Fatalf("Get poem status %s, res: %s", status, resJSON)
	}
	origin := resJSON.Data.Origin
	return map[string]string{
		"Title":   origin.Title,
		"Dynasty": origin.Dynasty,
		"Author":  origin.Author,
		"Content": strings.Join(origin.Content, "</br> "),
	}
}
