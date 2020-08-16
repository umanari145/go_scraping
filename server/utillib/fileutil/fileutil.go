package fileutil

import (
	"io/ioutil"
	"net/http"
)

/**
 * GetHTTPData HTTPデータの取得
 * @type string url URL
 * @return string HTMLデータ
 */
func GetHTTPData(url string) string {
	resp, _ := http.Get(url)
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)

	return string(byteArray)
}
