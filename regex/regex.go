package main

import (
	"regexp"
	"fmt"
)

const text = "My email is ccmouse@gmail.com.cn@abc.com"

/**

	. 代表任意一个字符
	+ 代表一个字符
 	* 表示任意字符

 */

func main()  {
	re := regexp.MustCompile(`[a-zA-Z0-9]+@[a-zA-Z0-9]+\.[a-zA-Z0-9]+`)
	match := re.FindString(text)
	fmt.Println(match)


}
