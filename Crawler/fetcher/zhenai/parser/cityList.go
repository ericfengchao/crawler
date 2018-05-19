package parser

import (
	"regexp"

	"github.com/ericfengchao/crawler/Crawler/engine/model"
)

const cityPageRegex = `<a href="(http://www.zhenai.com/zhenghun/[[:alpha:]]+)"[^>]+>([^<]+)</a>`

func ParseCityList(contents []byte) *model.ParseResult {
	re := regexp.MustCompile(cityPageRegex)
	matches := re.FindAllSubmatch(contents, -1)

	result := model.ParseResult{}
	for _, m := range matches {
		result.Requests = append(result.Requests, model.Request{
			Url:        string(m[1]),
			ParserFunc: ParseCity,
		})
		//result.Items = append(result.Items, "City "+string(m[2]))
	}
	return &result
}

func NilParser([]byte) *model.ParseResult {
	return nil
}
