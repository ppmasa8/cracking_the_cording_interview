package main

import (
	"fmt"
	"os"
)

/*** 方針 ***/
// ①番目のstrをハッシュテーブルに展開し、2番目のstrをそこからマイナスする。
// valueをかぞえ、１または-1が以下のようでなければfalseを返す
// 1, -1 → 0, 0
// 1, -1 → 1, 0
// 1, -1 → 0, 1
// 1, -1 → 1, 1

func main() {
	str1, str2 := []string{"s", "r", "f"}, []string{"s", "r", "f", "t"}
	hash := make(map[string]int)
	plus, minus := 0, 0

	for _, v := range str1 {
		hash[string(v)]++
	}

	for _, v := range str2 {
		hash[string(v)]--
	}

	for _, v := range hash {
		if plus > 1 || minus > 1 {
			fmt.Println("false")
			os.Exit(1)
		} else if v > 1 || v < -1 {
			fmt.Println("false")
			os.Exit(1)
		}

		if v == -1 {
			minus++
		} else if v == 1 {
			plus++
		}
	}

	fmt.Println(hash)

	if plus > 1 || minus > 1 {
		fmt.Println("false")
		os.Exit(1)
	} else {
		fmt.Println("true")
		os.Exit(0)
	}

}
// 時間計算量O(n)