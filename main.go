package main

import (
	"github.com/go-easylog/el"
)

func main() {
	el.SetLogLevel(el.TRACE)
	el.Warn("--batch start--")

	el.Warn("--batch end--")
}
