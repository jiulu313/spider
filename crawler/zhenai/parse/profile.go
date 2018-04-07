package parse

import (
	"spider/crawler/engine"
	"regexp"
	"strconv"
	"spider/crawler/model"
)

var ageRe = regexp.MustCompile(`<td><span class="label">年龄：</span>([\D]+)岁</td>`)
var marriageRe = regexp.MustCompile(`<td><span class="label">婚况：</span>([^<]+)</td>`)
var heightRe = regexp.MustCompile(`<td><span class="label">身高：</span>(\d+)CM</td>`)
var incomeRe = regexp.MustCompile(`<td><span class="label">月收入：</span>([^<]+)</td>`)
var weightRe = regexp.MustCompile(`<td><span class="label">体重：</span><span field="">(\d+)KG</span></td>`)
var genderRe = regexp.MustCompile(`<td><span class="label">性别：</span><span field="">([^<]+)</span></td>`)
var xinzhuoRe = regexp.MustCompile(`<td><span class="label">星座：</span><span field="">([^<]+)</span></td>`)
var educationRe = regexp.MustCompile(`<td><span class="label">学历：</span>([^<]+)</td>`)
var occupationRe = regexp.MustCompile(`<td><span class="label">职业： </span>([^<]+)</td>`)
var hokouRe = regexp.MustCompile(`<td><span class="label">籍贯：</span>([^<]+)</td>`)
var houseRe = regexp.MustCompile(`<td><span class="label">住房条件：</span><span field="">([^<]+)</span></td>`)
var carRe = regexp.MustCompile(`<td><span class="label">是否购车：</span><span field="">([^<]+)</span></td>`)



func ParseProfile(contents []byte,name string) engine.ParseResult {
	profile := model.Profile{}

	profile.Name = name;

	//age
	profile.Age, _ = strconv.Atoi(string(extractString(contents, ageRe)))

	//性别
	profile.Gender = extractString(contents,genderRe)

	//身高
	profile.Height, _ = strconv.Atoi(extractString(contents,heightRe))

	//体重
	profile.Weight,_ = strconv.Atoi(extractString(contents,weightRe))

	//收入
	profile.Income = extractString(contents,incomeRe)

	//marriage
	profile.Marriage = extractString(contents, marriageRe)

	//教育
	profile.Education = extractString(contents,educationRe)

	//职业
	profile.Occupation = extractString(contents,occupationRe)

	//户口
	profile.Hokou = extractString(contents,hokouRe)

	//星座
	profile.Xinzuo = extractString(contents,xinzhuoRe)

	//房子
	profile.House = extractString(contents,hokouRe)

	//车子
	profile.Car = extractString(contents,carRe)


	result := engine.ParseResult{
		Items:[]interface{}{profile},
	}
	return result
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)

	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}
