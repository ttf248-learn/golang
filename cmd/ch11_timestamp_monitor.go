package main

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/atotto/clipboard"
)

func main() {
	// 创建一个通道来接收退出信号
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// 创建一个通道来传递剪贴板的内容
	clipboardChan := make(chan string)

	// 开始剪贴板监控
	go monitorClipboard(clipboardChan)

	// 启动一个 goroutine 来处理剪贴板内容
	go processClipboard(clipboardChan)

	// 等待退出信号
	<-quit
	fmt.Println("Exiting...")
}

func monitorClipboard(clipboardChan chan<- string) {
	previousClipboard := ""
	for {
		clipboardContent, err := clipboard.ReadAll()
		if err != nil {
			fmt.Println("Error reading clipboard:", err)
		} else if clipboardContent != previousClipboard {
			clipboardChan <- clipboardContent
			previousClipboard = clipboardContent
		}
		time.Sleep(1 * time.Second)
	}
}

func processClipboard(clipboardChan <-chan string) {
	for {
		select {
		case clipboardContent := <-clipboardChan:
			if isValidTimestamp(clipboardContent) {
				timestamp, err := parseTimestamp(clipboardContent)
				if err != nil {
					fmt.Println("Error parsing timestamp:", err)
				} else {
					fmt.Println("Timestamp:", timestamp)
				}
			}
		}
	}
}

func isValidTimestamp(timestamp string) bool {
	return len(timestamp) == 10 || len(timestamp) == 13
}

func parseTimestamp(timestampStr string) (time.Time, error) {
	if len(timestampStr) == 10 || len(timestampStr) == 13 {
		// 将时间戳字符串转换为整数
		timestampInt, err := strconv.ParseInt(timestampStr, 10, 64)
		if err != nil {
			return time.Time{}, err
		}

		// 将时间戳转换为 time.Time 类型
		timestamp := time.Unix(timestampInt, 0)
		return timestamp, nil
	}

	return time.Time{}, fmt.Errorf("Invalid timestamp length: %d", len(timestampStr))
}
