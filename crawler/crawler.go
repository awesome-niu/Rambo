package crawler

import "fmt"

type Crawler struct {
	jobs    chan Fetcher
	results chan int
}

func NewCrawler(size int) *Crawler {
	return &Crawler{make(chan Fetcher, size), make(chan int, size)}
}

func (c Crawler) AddJob(job Fetcher) {
	c.jobs <- job
}

func(c Crawler) Run() {
	for job := range c.jobs {
		fmt.Println("--------------------------------")
		fmt.Println(job.Get())
	}
}

func Test() {
	c := NewCrawler(10)
	c.AddJob(NewFetcher("http://www.baidu.com", nil))
	c.AddJob(NewFetcher("http://www.v2ex.com", nil))
	c.AddJob(NewFetcher("http://www.newsmth.net", nil))

	c.Run()
}

