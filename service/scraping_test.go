package service

import (
	"testing"
)

func TestGetAllHTMLData(t *testing.T) {
	url := "https://kanto.hostlove.com/fuzoku_chiba/20090316232116"
	GetAllHTMLData(url)
}

func TestGetCommentData(t *testing.T) {
	url := "https://kanto.hostlove.com/fuzoku_chiba/20200614174142"
	GetCommentData(url, 1)
}
