package parse

import (
	"testing"
	"spider/crawler/fetcher"
	"fmt"
)

func TestParseProfile(t *testing.T) {
	contents , err := fetcher.Fetch("http://album.zhenai.com/u/106385696")
	if err != nil{
		t.Fatal(err)
	}

	result := ParseProfile(contents,"小朵")
	fmt.Println("item count : " ,len(result.Items))
	fmt.Println("--------------------------------------------")
	fmt.Println(result.Items)


}
