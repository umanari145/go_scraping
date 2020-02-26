package main

import (
	"github.com/go-easylog/el"
)

func main() {
	el.SetLogLevel(el.TRACE)
	el.Warn("--batch start--")

	el.Warn("--batch end--")

	//PHPでいうlist的な取り出し方
	x2, y2 := basic_fun("hoge", "foo")
	el.Warn("x2:" + x2)
	el.Warn("y2:" + y2)
}

/**
 * 引数をまとめた関数
 */
func basic_fun(x, y string) (string, string) {
	x2 := x + "aaa"
	y2 := y + "aaa"
	return x2, y2
}
