package parser

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"log"
	"strings"
)

type Trivia struct {
	name string
	url  string
}

func NewTrivia() *Trivia {
	return &Trivia{
		name: "trivia",
		url:  "http://www.lengdou.net/random",
	}
}

// name
func (t *Trivia) Name() string {
	return t.name
}

// route
func (t *Trivia) Route() string {
	return t.url
}

// parse
func (t *Trivia) Parse(body []byte) map[string]string {
	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(body))
	if err != nil {
		log.Fatal(err)
	}
	wrap := doc.Find(".container .media .media-body")
	imgURL, _ := wrap.Find(".topic-img img").Attr("src")

	return map[string]string{
		"ImgURL":      strings.Trim(imgURL, " "),
		"Description": strings.Split(wrap.Find(".topic-content").Text(), "#")[0],
	}
}
