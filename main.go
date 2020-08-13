package main

import (
	"github.com/go-easylog/el"
)

func main() {
	el.SetLogLevel(el.TRACE)
	el.SetRotateLog("./logs/%Y/%M/%D.log")
	el.Info("--scraping start--")

	el.Info("--scraping end--")
}
