package parser

import (
	"log"
	"regexp"

	"github.com/ericfengchao/crawler/Crawler/engine/model"
)

const cityPageRegexp = `<td><a href="(http://album.zhenai.com/u/[\d]+)"[^>]+>([^<]+)</a></td>`

// `<th><a href="(http://album.zhenai.com/u/[0-9]+)[^>]*>([^>]+)</a></th>`

func ParseCity(contents []byte, pageType string) model.ParseResult {
	re := regexp.MustCompile(cityPageRegexp)
	matches := re.FindAllSubmatch(contents, -1)

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
