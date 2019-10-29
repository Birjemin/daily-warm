package parser

type IParser interface {
	Name() string
	Route() string
	Parse(body []byte) map[string]string
}