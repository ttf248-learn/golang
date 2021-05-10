package main

import "fmt"

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
}
