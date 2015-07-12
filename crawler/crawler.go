package crawler

import (
	"sync"
	"io/ioutil"
	"os"
	"fmt"
)

type Crawler struct {
	jobs chan *Fetcher
}

func NewCrawler(size int) *Crawler {
	return &Crawler{make(chan *Fetcher, size)}
}

func (c Crawler) AddJob(job *Fetcher) {
	c.jobs <- job
}

func (c Crawler) Run() {
	var wg sync.WaitGroup

	for i := 0; i<2; i++ {
		wg.Add(1)
		go func() {
			for job := range c.jobs {
				fn := "D:\\" + job.url[7:] + ".txt"
				s := []byte(job.Get())
				err := ioutil.WriteFile(fn, s, os.ModeAppend)
				if err != nil {
					fmt.Println(err.Error())
				}
			}
			wg.Done()
		}()
	}

	//	close(c.jobs)

	wg.Wait()
	fmt.Println("over")
}



