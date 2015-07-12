package crawler
import "time"

type Cron struct {
	c Crawler
}

func NewCron(c Crawler) *Cron {
	return &Cron{c}
}

func (cron *Cron) Run() {
	cron.c.AddJob(NewFetcher("http://www.163.com", nil))
	cron.c.AddJob(NewFetcher("http://www.baidu.com", nil))
	cron.c.AddJob(NewFetcher("http://www.ifeng.com", nil))
	time.Sleep(2*time.Second)
	cron.c.AddJob(NewFetcher("http://www.cnbeta.com", nil))
	time.Sleep(3*time.Second)
	cron.c.AddJob(NewFetcher("http://www.feng.com", nil))
}