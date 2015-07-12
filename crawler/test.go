package crawler
import (
	"time"
	"fmt"
)

func TestFetcher() {
	f := NewFetcher("http://www.baidu.com", nil)
	fmt.Println(f.Get())
}

func TestCrawler() {
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
}

func TestCron() {
	c := NewCrawler(10)

	cron := NewCron(*c)
	go func() {
		cron.Run()
	}()

	c.Run()
}
