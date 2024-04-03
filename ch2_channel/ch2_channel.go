package main

import (
	"fmt"
	"time"
)

func sendInt(ch chan int, input int) {
	ch <- input
}

func rangChannel() {
	ch := make(chan int, 2)

	for i := 0; i < 2; i++ {
		ch <- i
	}
	close(ch)

	for v := range ch {
		fmt.Println(v)
	}
}

func intChannelWorker(ch chan int, input int) {
	time.Sleep(1 * time.Millisecond)
	ch <- input
}

func strChannelWorker(ch chan string, input string) {
	time.Sleep(1 * time.Millisecond)
	ch <- input
}

func someWorker(ch chan bool) {
	fmt.Println("before some worker")
	ch <- true
	fmt.Println("after some worker")
}

func main() {
	fmt.Println("ch2 message")

	ch := make(chan int)

	go func() {
		ch <- 100
	}()

	go sendInt(ch, 200)

	fmt.Println(<-ch)
	fmt.Println(<-ch)

	rangChannel() // 如果用 go 启动当前函数，由于是异步，程序直接退出到了，会导致数据无法正常打印

	chInt := make(chan int)
	chStr := make(chan string)

	go intChannelWorker(chInt, 100)
	go strChannelWorker(chStr, "xiangtianlong")

selectLoop:
	for true {
		select {
		case <-chInt:
			fmt.Println("get int message")
		case <-chStr:
			fmt.Println("get str message")
		default:
			fmt.Println("break select loop")
			break selectLoop
		}
	}

	chBool := make(chan bool)
	go someWorker(chBool)

	<-chBool

	fmt.Println("demo end")
}
