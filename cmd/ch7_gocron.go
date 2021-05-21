package main

import (
	"fmt"
	"time"

	"github.com/go-co-op/gocron"
)

func task(s string) {
	fmt.Println("I am running, about", s)
}

func main() {

	s := gocron.NewScheduler(time.Local)

	// every 类型的任务，除开星期类型的模式，其他类型的任务会立马执行一次，不用按照设置的时间下次再执行
	/*I am running, about 1s
	I am running, about 1min
	I am running, about 1hour
	I am running, about 1day
	I am running, about 1week
	I am running, about 1s
	I am running, about 1s*/

	// 每隔多久
	s.Every(1).Seconds().Do(task, "1s")
	s.Every(1).Minutes().Do(task, "1min")
	s.Every(1).Hours().Do(task, "1hour")
	s.Every(1).Days().Do(task, "1day")
	s.Every(1).Weeks().Do(task, "1week")

	// 每周几
	s.Every(1).Monday().Do(task, "monday")
	s.Every(1).Friday().Do(task, "friday")

	// 每天固定时间
	s.Every(1).Days().At("10:30").Do(task, "10:30")
	s.Every(1).Monday().At("17:30").Do(task, "17:30")
	s.Every(1).Friday().At("16:24").Do(task, "16:22")

	// 设置 crontab 字符串格式
	s.Cron("*/1 * * * *").Do(task, "crontab")

	// 启动任务，同步锁定当前线程
	s.StartBlocking()
}
