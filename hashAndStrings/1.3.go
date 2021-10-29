package main

import "fmt"

/*** 方針 ***/
// 素直に空白の個数をカウントし、%20に置き換えた分の文字数を追加してlengthをとっておき
// 後に、追加する。

func main() {
	str, length := []string{"s", "f", "d", " "}, 4

	for j := len(str) - 1; j >= 0 ; j-- {
		if str[j] == " " {
			str[j] = "%20"
		}
	}

	fmt.Println(str, length)
}
// 時間計算量O(n)
// なお、pythonなどと違い、配列の格納できる個数に制限がかかっている静的言語であり自動で
// 拡張してくれないためにしょうもない実装になっている。
