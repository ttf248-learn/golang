package main

import (
	"fmt"
	"time"
)

func timestamp_to_time(timestamp_in int64) {
	// 假设你有一个Unix时间戳
	timestamp := int64(timestamp_in) // 替换为你自己的时间戳

	// 将时间戳转换为time.Time类型
	t := time.Unix(timestamp, 0)

	// 格式化时间为年-月-日 时:分:秒
	formattedTime := t.Format("2006-01-02 15:04:05")

	// 输出结果
	fmt.Println("格式化时间:", formattedTime)
}

func main() {
	timestamp_to_time(-594806400)
	timestamp_to_time(-621355968)
}
