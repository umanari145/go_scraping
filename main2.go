package main

import (
	"fmt"

	"github.com/go-easylog/el"
)

//インスタンス変数
var i, j int = 1, 2

//インスタンス定数
const TEL = "0123456789"

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

	const EMAIL = "umanari145@gmail.com"
	//当然代入するとエラー
	//EMAIL = "aaaa"
	fmt.Println("EMAIL 定数" + EMAIL)
	fmt.Println("TEL (インスタンス)定数" + TEL)

	fmt.Println("defer")

	deferMethod()

	deferMethod2()
}

/**
 * 引数をまとめた関数
 */
func multiArgs(x, y string) (string, string) {
	x2 := x + "aaa"
	y2 := y + "aaa"
	return x2, y2
}

func deferMethod() {
	defer fmt.Println("順番は最初だが関数終了時点で発動")
	fmt.Println("通常の関数")
}

func deferMethod2() {
	defer func() {
		fmt.Println("クロージャ的な使い方でのdefer")
	}()
	fmt.Println("通常の関数")
}
