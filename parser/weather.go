package parser

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"log"
	"strings"
)

type Weather struct {
	name string
	url  string
}

func NewWeather(local string) *Weather {
	return &Weather{
		name: "weather" + local,
		url:  "https://tianqi.moji.com/weather/china/" + local,
	}
}

// name
func (w *Weather) Name() string {
	return w.name
}

// route
func (w *Weather) Route() string {
	return w.url
}

// parse
func (w *Weather) Parse(body []byte) map[string]string {
	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(body))
	if err != nil {
		log.Fatal(err)
	}
	wrap := doc.Find(".wea_info .left")
	humidityDesc := strings.Split(wrap.Find(".wea_about span").Text(), " ")
	humidity := "未知"
	if len(humidityDesc) >= 2 {
		humidity = humidityDesc[1]
	}
	limit := wrap.Find(".wea_about b").Text()

	return map[string]string{
		"City":     doc.Find("#search .search_default em").Text(),
		"Temp":     wrap.Find(".wea_weather em").Text() + "°",
		"Weather":  wrap.Find(".wea_weather b").Text(),
		"Air":      wrap.Find(" .wea_alert em").Text(),
		"Humidity": humidity,
		"Wind":     wrap.Find(".wea_about em").Text(),
		"Limit":    limit,
		"Note":     strings.ReplaceAll(wrap.Find(".wea_tips em").Text(), "。", ""),
	}
}
