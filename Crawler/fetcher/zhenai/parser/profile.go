package parser

import (
	"log"
	"regexp"
	"strconv"

	engine_model "github.com/ericfengchao/crawler/Crawler/engine/model"
	"github.com/ericfengchao/crawler/Crawler/model"
)

var ageRe = regexp.MustCompile(`<td><span class="label">年龄：</span>([1-9][0-9]*)岁</td>`)

var heightRe = regexp.MustCompile(`<td><span class="label">身高：</span>([0-9]+)CM</td>`)

var IncomeRe = regexp.MustCompile(` <td><span class="label">月收入：</span>(.+元)</td>`)

var nameRe = regexp.MustCompile(`<a class="name fs24">(.+)</a>`)

func ParseProfile(contents []byte) *engine_model.ParseResult {
	profile := model.Profile{}
	if name := extractStr(contents, nameRe); name != "" {
		profile.Name = name
	} else {
		// user name is the basic required field or discard this user
		return nil
	}

	if ageMatch, err := strconv.Atoi(extractStr(contents, ageRe)); err == nil {
		profile.Age = ageMatch
	}

	if heightMatch, err := strconv.Atoi(extractStr(contents, heightRe)); err == nil {
		profile.Height = heightMatch
	}

	if incomeMatch := extractStr(contents, heightRe); incomeMatch != "" {
		profile.Income = incomeMatch
	}
	log.Printf("Profile Crawled: %v\n", profile)
	result := engine_model.ParseResult{}
	result.Items = []interface{}{profile}
	return &result
}

func extractStr(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)

	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}
