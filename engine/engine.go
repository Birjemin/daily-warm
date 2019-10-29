package engine

import (
	"github.com/barryyan/daily-warm/fetcher"
	"github.com/barryyan/daily-warm/parser"
)

func Run(seeds []parser.IParser) map[string]interface{} {
	ch := make(chan map[string]interface{}, len(seeds))
	defer close(ch)

	for _, p := range seeds {
		go Fetch(p, ch)
	}

	data := map[string]interface{}{}
	for range seeds {
		temp := <-ch
		for k, v := range temp {
			data[k] = v
		}
	}

	return data
}

func Fetch(p parser.IParser, ch chan map[string]interface{}) {
	body, err := fetcher.Fetch(p.Route())
	if err == nil {
		ch <- map[string]interface{}{p.Name(): p.Parse(body)}
	} else {
		ch <- map[string]interface{}{p.Name(): map[string]string{}}
	}
}
