package main

import (
	"fmt"
	"regexp"
)

func main()  {
	//var str = "abc123def"
	//var re = "[0-9]+"
	//regexp := regexp.MustCompile(re)
	//match := regexp.FindStringSubmatch(str)
	//fmt.Println(match)


	var str = "abc1223abc777"
	//var patt1 = "^[0-9]+"
	var patt2 = "^[^0-9]+$"

	re := regexp.MustCompile(patt2)
	match := re.FindStringSubmatch(str)
	fmt.Println(match)



}
