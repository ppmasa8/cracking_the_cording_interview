package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*** 方針 ***/
// 文字列圧縮をして、圧縮変換するのにはハッシュテーブルのkey,valueをつなげて表示するのが
// 良さそう。圧縮変換した後の文字列と元の文字列を比較し、短いほうを返すようにする。

func main() {
	str := []string{"s", "t", "r", "s", "t", "t", "r"}
	hash := make(map[string]int)
	hashArr := []string{}

	for _, v := range str {
		hash[string(v)]++
	}

	for k, v := range hash {
		hashArr = append(hashArr, k)
		hashArr = append(hashArr, strconv.Itoa(v))
	}

	//fmt.Println(hashArr)

	hashStr := strings.Join(hashArr, "")
	//fmt.Println(hashStr)

	if len(str) > len(hashStr) {
		fmt.Println(hashStr)
	} else {
		fmt.Println(str)
	}
}

// 時間計算量O(n)
