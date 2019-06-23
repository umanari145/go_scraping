package fileutil

import (
	"testing"
)

func TestHttpAccess(t *testing.T) {
	GetHTTPData("https://paiza.jp/career/job_offers")
}
