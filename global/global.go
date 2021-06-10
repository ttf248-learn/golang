package global

import (
	"fmt"
	"os"
)

func GetConsulInfo() string {
	var ret string

	if len(os.Args) > 1 && os.Args[1] == "local" {
		ret = "localhost:8500"
	} else {
		ret = "134.175.47.100:8500"
	}
	fmt.Println("GetConsulInfo:", ret)
	return ret
}
