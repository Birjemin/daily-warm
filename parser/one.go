package parser

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
)

type One struct {
	name string
	url  string
}

func NewOne() *One {
	return &One{
		name: "one",
		url:  "http://wufazhuce.com/",
	}
}

// name
func (o *One) Name() string {
	return o.name
}

// route
func (o *One) Route() string {
	return o.url
}

// parse
func (o *One) Parse(body []byte) map[string]string {
	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(body))
	if err != nil {
		log.Fatal(err)
	}
	wrap := doc.Find(".fp-one .carousel .item.active")
	day := wrap.Find(".dom").Text()
	monthYear := wrap.Find(".may").Text()
	imgURL, _ := wrap.Find(".fp-one-imagen").Attr("src")
	return map[string]string{
		"ImgURL":   imgURL,
		"Date":     fmt.Sprintf("%s %s", day, monthYear),
		"Sentence": wrap.Find(".fp-one-cita a").Text(),
	}
}
