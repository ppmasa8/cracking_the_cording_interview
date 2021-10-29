package main

import (
	"fmt"
	"strings"
)

/*** 方針 ***/
// 素直に置き換える

func main() {
	// input
	str, length := "sdfs sdfs sdf", 13

	fmt.Println(strings.Replace(str, " ", "%20", -1))

	fmt.Println(length)

}
