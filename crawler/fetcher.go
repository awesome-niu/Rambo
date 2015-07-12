package crawler

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Fetcher struct {
	url    string
	header map[string]string
}

func NewFetcher(url string, header map[string]string) *Fetcher {
	return &Fetcher{url, header}
}

func (f *Fetcher) Get() string {
	client := &http.Client{}

	req, err := http.NewRequest("GET", f.url, nil)
	if err != nil {
		log.Fatal("Generate http request error.")
		os.Exit(-1)
	}

	for key, value := range f.header {
		req.Header.Add(key, value)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Fetch http response error.")
		os.Exit(-1)
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Fatal("Fetch http response error.")
		os.Exit(-1)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Read response body error.")
		os.Exit(-1)
	}

	return string(body)
}