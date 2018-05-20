package parser

import (
	"log"
	"regexp"

	"github.com/ericfengchao/crawler/Crawler/engine/model"
)

const cityListPageRegexp = `<a href="(http://city.zhenai.com/[[:alnum:]]+)"[^>]+>([^<]+)</a>`

var cityListPageRe = regexp.MustCompile(cityListPageRegexp)

func ParseCityList(contents []byte, pageType string) model.ParseResult {
	matches := cityListPageRe.FindAllSubmatch(contents, -1)

	result := model.ParseResult{}
	for _, m := range matches {
		m_ := m
		result.Requests = append(result.Requests, model.Request{
			Url:        string(m_[1]),
			PageTitle:  string(m_[2]),
			ParserFunc: ParseCity,
		})
		//result.Items = append(result.Items, "City "+string(m[2]))
		log.Printf("Page %s Got City %s\n", pageType, m_[2])
	}
	return result
}

func NilParser([]byte) *model.ParseResult {
	return nil
}
