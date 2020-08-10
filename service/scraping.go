package service

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/go-easylog/el"
	strip "github.com/grokify/html-strip-tags-go"
	"github.com/saintfish/chardet"
	"golang.org/x/net/html/charset"
)

//Comments はコメント
var Comments []Comment

//Comment は１つ１つのコメント
type Comment struct {
	//No はコメント番号
	No string
	//Contents はコメント番号
	Contents string
}

//GetAllHTMLData は全ページのデータ取得
func GetAllHTMLData(URL string) {

	el.SetLogLevel(el.TRACE)
	el.SetRotateLog("./%Y/%M/%D.log")

	el.Info("--page番号の取得--")
	targetURL := fmt.Sprintf("%s/a/1", URL)
	reader := GetStream(targetURL)
	allPageNo := getPageNo(reader)
	el.Info(`--全page数 ` + allPageNo + `--`)

	//allPageNoInt, _ := strconv.Atoi(allPageNo)
	//	for i := 0; i <= allPageNoInt; i++ {
	for i := 0; i <= 2; i++ {
		GetCommentData(URL, i)
	}
	fmt.Println(Comments)

}

//GetCommentData はコメントデータの取得
func GetCommentData(URL string, pageNo int) {
	targetURL := fmt.Sprintf("%s/%d", URL, pageNo)
	reader := GetStream(targetURL)
	parseComment(reader)
}

//GetStream はストリームデータの取得
func GetStream(URL string) io.Reader {

	res, _ := http.Get(URL)

	// 読み取り
	buf, _ := ioutil.ReadAll(res.Body)

	// 文字コード判定
	det := chardet.NewTextDetector()
	detRslt, _ := det.DetectBest(buf)

	// 文字コード変換
	bReader := bytes.NewReader(buf)
	reader, _ := charset.NewReaderLabel(detRslt.Charset, bReader)
	return reader
}

//getPageNo はページ番号の取得
func getPageNo(reader io.Reader) string {
	var pageNo string
	doc, _ := goquery.NewDocumentFromReader(reader)
	doc.Find("#jump_bt").Each(func(_ int, s *goquery.Selection) {
		pageNo = s.Find("span.total").Text()
	})
	return pageNo
}

/**
 * parseComment はページごとのコメントデータを取得
 * @type io.Reader reader
 * @return string HTMLデータ
 */
func parseComment(reader io.Reader) {

	doc, _ := goquery.NewDocumentFromReader(reader)
	doc.Find("#thread_c article").Each(func(_ int, s *goquery.Selection) {
		CommentData := Comment{}
		commentNo, _ := s.Find(".fancybox_com").Attr("title")
		rawData := s.Find(".res").Text()
		strippedData := strip.StripTags(rawData)

		CommentData.No = commentNo
		CommentData.Contents = strippedData
		Comments = append(Comments, CommentData)
	})
}
