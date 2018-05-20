package parser

import (
	"github.com/ericfengchao/crawler/Crawler/engine/model"
	"log"
	"regexp"
)

const cityPageRegexp = `<td><a href="(http://album.zhenai.com/u/[\d]+)"[^>]+>([^<]+)</a></td>`

var cityPageRe = regexp.MustCompile(cityPageRegexp)

func ParseCity(contents []byte, pageType string) model.ParseResult {
	matches := cityPageRe.FindAllSubmatch(contents, -1)

	result := model.ParseResult{}
	for _, m := range matches {
		result.Requests = append(result.Requests, model.Request{
			Url:        string(m[1]),
			PageTitle:  string(m[2]),
			ParserFunc: ParseProfile,
		})
		//result.Items = append(result.Items, "User "+string(m[2]))
		log.Printf("Page %s Got User %s\n", pageType, string(m[2]))
	}
	return result
}
