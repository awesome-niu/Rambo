package main

import (
  "net/url"
  "fmt"
  "github.com/PuerkitoBio/goquery"
  "io/ioutil"
  "net/http"
  "strings"
)

func main() {
  doc, err := goquery.NewDocument("http://weekly.manong.io/issues/77")
  if err != nil {
    fmt.Println(err)
  }
  doc.Find("h4").Each(func(num int, s *goquery.Selection) {

    if strings.Contains(strings.ToLower(s.Contents().Text()), "android") ||
      strings.Contains(strings.ToLower(s.Contents().Text()), "go") {

      fmt.Println(s.Contents().Text())
      content, isExist := s.Find("a").Attr("href")

      if !isExist {
        return
      }

      path, _ := url.Parse(content)

      fmt.Printf("url parse %p\n", path)

      content = strings.Replace(content, "%3A", ":", -1)
      content = strings.Replace(content, "%2F", "/", -1)
      content = strings.Replace(content, "&amp", " ", -1)

      //      fmt.Println(content)
      start := strings.Index(content, "url=")
      end := strings.Index(content, "&aid")
      fmt.Println(content[start+4 : end])
    }

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
