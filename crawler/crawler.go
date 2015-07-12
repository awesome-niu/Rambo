package crawler

import (
	"sync"
	"io/ioutil"
	"os"
	"fmt"
	"time"
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

func Test() {
	c := NewCrawler(10)
	c.AddJob(NewFetcher("http://www.163.com", nil))
	c.AddJob(NewFetcher("http://www.baidu.com", nil))
	c.AddJob(NewFetcher("http://www.ifeng.com", nil))
	c.AddJob(NewFetcher("http://www.cnbeta.com", nil))
	go func() {
		time.Sleep(5*time.Second)
		c.AddJob(NewFetcher("http://www.feng.com", nil))
	}()

	c.Run()

	//	tasks := make(chan func(), 64)
	//	tasks <- func() {
	//		fmt.Println(1)
	//	}
	//	tasks <- func() {
	//		fmt.Println(2)
	//	}
	//	tasks <- func() {
	//		fmt.Println(3)
	//	}
	//	tasks <- func() {
	//		fmt.Println(4)
	//	}
	//
	//	go func() {
	//		time.Sleep(3*time.Second)
	//		tasks <- func() {
	//			fmt.Println(5)
	//		}
	//		time.Sleep(1*time.Second)
	//		tasks <- func() {
	//			fmt.Println(6)
	//		}
	//		time.Sleep(2*time.Second)
	//		tasks <- func() {
	//			fmt.Println(7)
	//		}
	//	}()
	//
	//	ticker := time.NewTicker(1*time.Second)
	//	quit := make(chan bool)
	//	for {
	//		select {
	//		case <-ticker.C:
	//			fmt.Println(0)
	//			var wg sync.WaitGroup
	//			for i := 0; i<2; i++ {
	//				wg.Add(1)
	//				go func() {
	//					for task := range tasks {
	//						task()
	//					}
	//					wg.Done()
	//				}()
	//			}
	//			wg.Wait()
	//		case <-quit:
	//			ticker.Stop()
	//			return
	//		}
	//	}
}



