package main

import (
	"fmt"
	"os"
)

const (
	Unknown = 0
	Female  = 1
	Male    = 2
)

const (
	One = iota
	Two = iota
)

func main() {

	const conStr string = "const string"
	const conStr2 = "const string 2"
	fmt.Println(conStr, conStr2, Unknown, Female, Male, One, Two)

	/* << 左移运算符，相当于乘以2的n次方 >> 右移运算符，相当于除以2的n次方 */

	var tmp1 int = 60
	fmt.Println(tmp1<<1, tmp1>>1)

	tmp2 := 30
	var tmp3 *int
	fmt.Println("tmp2:", tmp2, "tmp3:", tmp3)
	tmp3 = &tmp2
	fmt.Println("tmp2:", tmp2, "tmp3:", *tmp3)
	tmp2 = 60
	fmt.Println("tmp2:", tmp2, "tmp3:", *tmp3)

	_, b, c := multi_()
	fmt.Println(b, c)

	if len(os.Args) > 1 {
		fmt.Println("ch1 message", os.Args[1])
	} else {
		fmt.Println("ch1 message default")
	}

	if b == 2 {
		if c == "return message" {
			fmt.Println("b:", b, "c:", c)
		} else {
			fmt.Println("c is not equal return message")
		}
	} else {
		fmt.Println("b is not equal 2")
	}

	funcSwitchTest("A")
	funcSwitchTest("B")
	funcSwitchTest("")
}

func multi_() (int, int, string) {
	a, b, c := 1, 2, "return message 1"
	return a, b, c
}

func funcSwitchTest(grade string) {
	switch grade {
	case "A":
		fmt.Println("优秀")
	case "B":
		fmt.Println("良好")
	case "C":
		fmt.Println("及格")
	default:
		fmt.Println("未知")
	}
}
