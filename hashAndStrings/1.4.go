package main

import (
	"fmt"
	"os"
)

/*** 方針 ***/
// 回分の順列かどうかの判定をするには、ハッシュテーブルのvalueに複数個奇数が出てこなければOK

func main() {
	str := []string{"s", "t", "d", "t", "s"}
	hash := make(map[string]int)

	for _, v := range str {
		hash[string(v)]++
	}

	cnt := 0
	for _, v := range hash {
		if cnt > 1 {
			fmt.Println("false")
			os.Exit(1)
		}
		if v % 2 != 0 {
			cnt++
		}
	}

	if cnt > 1 {
		fmt.Println("false")
		os.Exit(1)
	} else {
		fmt.Println("true")
		os.Exit(0)
	}
}
