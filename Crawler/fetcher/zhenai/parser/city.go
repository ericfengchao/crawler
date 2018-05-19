package parser

import (
	"regexp"

	"github.com/ericfengchao/crawler/Crawler/engine/model"
)

const cityListPageRegex = `<th><a href="(http://album.zhenai.com/u/[0-9]+)[^>]*>([^>]+)</a></th>`

func ParseCity(contents []byte) *model.ParseResult {
	re := regexp.MustCompile(cityListPageRegex)
	matches := re.FindAllSubmatch(contents, -1)

	result := model.ParseResult{}
	for _, m := range matches {
		result.Requests = append(result.Requests, model.Request{
			Url:        string(m[1]),
			ParserFunc: ParseProfile,
		})
		//result.Items = append(result.Items, "User "+string(m[2]))
	}
	return &result
}
