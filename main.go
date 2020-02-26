package main

import (
	"fmt"

	"github.com/go-easylog/el"
)

var i, j int = 1, 2

func main() {
	el.SetLogLevel(el.TRACE)
	el.Warn("--batch start--")

	el.Warn("--batch end--")

	//PHPでいうlist的な取り出し方
	x2, y2 := multiArgs("hoge", "foo")
	el.Warn("x2:" + x2)
	el.Warn("y2:" + y2)

	el.Warn("i+j:")
	el.Warn(i + j)

	//自動での型変換
	el.Warn("自動での型変更")
	c, python, java := true, false, "no!"
	//boolean
	el.Warn(c)
	//boolean
	el.Warn(python)
	//string
	el.Warn(java)
	fmt.Println("--下記のように変数をカンマつなぎで表記可能--")
	fmt.Println(c, python, java)

	fmt.Println("定数の表記＋sprintfのような表記")
	var i int
	var f float64
	var b bool
	var s string
	//phpのsprintfに近い表記
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}

/**
 * 引数をまとめた関数
 */
func multiArgs(x, y string) (string, string) {
	x2 := x + "aaa"
	y2 := y + "aaa"
	return x2, y2
}
