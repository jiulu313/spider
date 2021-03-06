package fetcher

import (
	"io"
	"bufio"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"net/http"
	"fmt"
	"golang.org/x/text/transform"
	"io/ioutil"
	"golang.org/x/text/encoding/unicode"
	"log"
)

//获取网页的内容
func Fetch(url string) ([]byte,error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil,err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error:status code:", resp.StatusCode)
		return nil,fmt.Errorf("wrong status code: %d",resp.StatusCode)
	}

	e := determineEncoding(resp.Body)
	utf8Reader := transform.NewReader(resp.Body, e.NewDecoder())
	return ioutil.ReadAll(utf8Reader)
}


//返回网页的编码
func determineEncoding(r io.Reader) encoding.Encoding {
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		log.Printf("Fetcher error: %v",err)
		return unicode.UTF8
	}

	e, _, _ := charset.DetermineEncoding(bytes, "")

	return e
}