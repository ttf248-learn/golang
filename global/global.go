package global

import (
	"fmt"
)

func GetConsulInfo() string {
	ret := "localhost:8500"
	fmt.Println("GetConsulInfo:", ret)
	return ret
}
