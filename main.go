package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"net/http"
)

func main() {
	doc, err := goquery.NewDocument("http://weekly.manong.io/issues/77")
	if err != nil {
		fmt.Println(err)
	}
	doc.Find("h4").Each(func(num int, s *goquery.Selection) {
		fmt.Println(s.Contents().Text())
	})
	
}

func httpGetMaNong() {
	response, err := http.Get("http://weekly.manong.io/issues/77")
	if err != nil {
		fmt.Println(err)
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(body))
}
