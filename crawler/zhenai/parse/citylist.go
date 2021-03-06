package parse

import (
	"spider/crawler/engine"
	"regexp"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

func ParseCityList(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityListRe)
	matchs := re.FindAllStringSubmatch(string(contents), -1)

	result := engine.ParseResult{}
	for _, m := range matchs {
		result.Items = append(result.Items, "City " + m[2]) //城市
		result.Requests = append(result.Requests, engine.Request{
			Url:       m[1],
			ParseFunc: ParseCity,
		})
	}

	return result
}
