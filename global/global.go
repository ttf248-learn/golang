package global

import (
	"fmt"
	"os"
)

func GetConsulInfo() string {
	var ret string

	if os.Args[1] == "remote" {
		ret = "134.175.47.100:8500"
	} else {
		ret = "localhost:8500"
	}
	fmt.Println("GetConsulInfo:", ret)
	return ret
}
