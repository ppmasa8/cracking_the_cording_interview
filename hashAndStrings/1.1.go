package main

import (
	"fmt"
	"os"
)

/*** 方針 ***/
// 文字列:strとおく。
// strを前から順に見ていき、{a:1,b:0,c:2,...}のようにハッシュテーブルに保持して
// valueが2以上のものが存在すれば真、しなければ偽というふうに実装する
// もし、新しいデータ構造を使えないという制約が入るなら、strをソートし（O(nlogn)）
// 前から順に②つずつ比較して、同じものがあるか見ていく。

// 制約なし //
//func main() {
//	str := "sakhgauisghanus"
//	hash := make(map[string]int)
//	for i := 0; i < len(str); i++ {
//		hash[string(str[i])]++
//	}
//
//	for _,v := range hash {
//		if v >= 2 {
//			fmt.Println("false")
//			os.Exit(1)
//		}
//	}
//	fmt.Println("true")
//	os.Exit(0)
//}
// 時間計算量 O(n)

// 制約ありver
func main() {
	str := "jdfhajhfagk"
	strArr := []string{}
	// string型を[]string型に変換
	for i := 0; i < len(str); i++ {
		strArr = append(strArr, str[i:i+1])
	}
	s := mergeSort(strArr)

	for i, v := range s {
		if i == len(s) {
			fmt.Println("true")
			os.Exit(0)
		}

		if v == s[i+1] {
			fmt.Println("false")
			os.Exit(1)
		}
	}
}
// 時間計算量O(nlogn)

func mergeSort(s []string) []string {
	var result []string
	if len(s) < 2 { return s }

	mid := len(s) / 2
	r := mergeSort(s[:mid])
	l := mergeSort(s[mid:])
	i, j := 0, 0

	for i < len(r) && j < len(l) {
		if r[i] > l[j] {
			result = append(result, l[j])
			j++
		} else {
			result = append(result, r[i])
			i++
		}
	}

	result = append(result, r[i:]...)
	result = append(result, l[j:]...)

	return result
}

