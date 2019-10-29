package parser

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"log"
)

type English struct {
	name string
	url  string
}

func NewEnglish() *English {
	return &English{
		name: "english",
		url:  "http://dict.eudic.net/home/dailysentence",
	}
}

// name
func (e *English) Name() string {
	return e.name
}

// route
func (e *English) Route() string {
	return e.url
}

// parse
func (e *English) Parse(body []byte) map[string]string {
	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(body))
	if err != nil {
		log.Fatal(err)
	}
	wrap := doc.Find(".containter .head-img")
	imgURL, _ := wrap.Find(".himg").Attr("src")
	return map[string]string{
		"ImgURL":   imgURL,
		"Sentence": wrap.Find(".sentence .sect_en").Text(),
	}
}
