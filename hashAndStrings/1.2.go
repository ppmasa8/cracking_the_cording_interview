package main

import (
	"fmt"
	"os"
)

/*** 方針 ***/
//　２つの文字列が並び替えかどうか調べるには、要素の個数が一致していれば良い。
// したがって、文字列１のハッシュテーブルを作り、文字列2を前から順にみてハッシュテーブルに
// あれば１つデクリメントし、存在しない文字であれば即false返し。hashテーブルのvalueがすっかり
// すべて0になればtrue,それ以外はfalseを返す。

func main() {
	str1 := "ohayou"
	str2 := "ohatou"

	if len(str1) != len(str2) {
		fmt.Println("false")
		os.Exit(1)
	}

	hash := make(map[string]int)
	for i := 0; i < len(str1); i++ {
		hash[string(str1[i])]++
	}

	for j := 0; j < len(str2); j++ {
		hash[string(str2[j])]--
	}

	for _,v := range hash {
		if v == 0 {
			fmt.Println("false")
			os.Exit(1)
		}
	}
	fmt.Println("true")
	os.Exit(0)
}
// 時間計算量O(n)