package model

type Request struct {
	Url        string
	PageTitle  string
	ParserFunc func([]byte, string) ParseResult
}

type ParseResult struct {
	Requests []Request
	Items    []interface{}
}
