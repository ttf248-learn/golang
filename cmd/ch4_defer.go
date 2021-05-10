package main

import (
	"fmt"
	"time"
)

func main() {

	// defer 本质内部是个链表，先进后出进行遍历
	{
		fmt.Println("enter first code block")
		defer fmt.Println("1")
		fmt.Println("leave first code block")
	}
	defer fmt.Println("2")

	// defer 和普通的 go 函数一样，都是通过值传递，在执行到 defer 的时候已经完成了计算，内部的数据结构保存了计算结果
	tmpTime := time.Now()
	defer fmt.Println("consume time:", time.Since(tmpTime))
	time.Sleep(time.Second * 1)

	tmpTime = time.Now()
	defer func() { fmt.Println("consume time:", time.Since(tmpTime)) }()
	time.Sleep(time.Second * 1)
}
