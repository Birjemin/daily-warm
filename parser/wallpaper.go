package parser

import (
	"encoding/json"
	"log"
)

// PoemRes response data
type WallPaperRes struct {
	Images []struct {
		Url   string `json:"url"`
		Title string `json:"copyright"`
	} `json:"images"`
}

type WallPaper struct {
	name string
	url  string
}

func NewWallpaper() *WallPaper {
	return &WallPaper{
		name: "wallpaper",
		url:  "https://cn.bing.com/HPImageArchive.aspx?format=js&idx=0&n=1&nc=1477557984674&pid=hp&video=1",
	}
}

// name
func (w *WallPaper) Name() string {
	return w.name
}

// route
func (w *WallPaper) Route() string {
	return w.url
}

// parse
func (w *WallPaper) Parse(body []byte) map[string]string {
	// Load the HTML document
	var resJSON WallPaperRes
	err := json.Unmarshal(body, &resJSON)
	if err != nil {
		log.Fatalf("Fetch json from %s error: %s", w.Route(), err)
	}
	origin := resJSON.Images
	return map[string]string{
		"Title":  origin[0].Title,
		"ImgURL": "https://cn.bing.com/" + origin[0].Url,
	}
}
