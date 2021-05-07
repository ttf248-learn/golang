package main

import (
	"fmt"
	"os"
)

func main() {

	_, b, c := multi_()

	fmt.Println(b, c)

	if len(os.Args) > 1 {
		fmt.Println("ch1 message", os.Args[1])
	}
}

func multi_() (int, int, string) {
	a, b, c := 1, 2, "return message"
	return a, b, c
}
