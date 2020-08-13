package service

import (
	"bytes"
	"fmt"
	"go_scraping/dbutil"
	"io"
	"io/ioutil"
	"net/http"
	"sort"
	"strconv"

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
	//ThreadNo はコメント番号
	ThreadNo string
	//Contents はコメント本体
	Contents string
	//CommentDate はコメントを書き込んだ日
	CommentDate string
}

//GetAllHTMLData は全ページのデータ取得
func GetAllHTMLData(URL string) {

	el.Info("--page番号の取得--")
	targetURL := fmt.Sprintf("%s/a/1", URL)
	reader := GetStream(targetURL)
	allPageNo := getPageNo(reader)
	el.Info(`--全page数 ` + allPageNo + `--`)

	//allPageNoInt, _ := strconv.Atoi(allPageNo)
	//for i := 1; i <= allPageNoInt; i++ {
	for i := 1; i <= 2; i++ {
		strInt := strconv.Itoa(i)
		el.Info(`--page ` + strInt + `--`)
		GetCommentData(URL, i)
	}

	sort.SliceStable(Comments, func(i, j int) bool { return Comments[i].ThreadNo < Comments[j].ThreadNo })
	fmt.Println(Comments)

	dbConn, err := dbutil.Connect()
	if err != nil {
		el.Fatal("Cannnot Connect DB")
	}

	for _, eachComment := range Comments {
		dbConn.Create(&eachComment)
	}
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
		threadNo, _ := s.Find(".fancybox_com").Attr("title")
		rawData := s.Find(".res").Text()
		commentDate := s.Find(".date").Text()
		strippedData := strip.StripTags(rawData)

		CommentData.ThreadNo = threadNo
		CommentData.Contents = strippedData
		CommentData.CommentDate = commentDate
		Comments = append(Comments, CommentData)
	})
}
