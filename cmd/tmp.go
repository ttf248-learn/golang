package main

import (
	"fmt"
	"os"
)

const (
	Unkow  = 0
	Female = 1
	Male   = 2
)

const (
	One = iota
	Two = iota
)

func main() {

	const conStr string = "const string"
	const conStr2 = "const string 2"

	fmt.Println(conStr, conStr2, Unkow, Female, Male, One, Two)

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
