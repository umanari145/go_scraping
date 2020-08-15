package service

import (
	"bytes"
	"fmt"
	"go_scraping/dbutil"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"sort"
	"strconv"

	"github.com/PuerkitoBio/goquery"
	"github.com/go-easylog/el"
	strip "github.com/grokify/html-strip-tags-go"
	"github.com/saintfish/chardet"
	"golang.org/x/net/html/charset"
)

//Comments はCommentの複数形
type Comments []Comment

//Comment は１つ１つのコメント
type Comment struct {
	//ThreadNo はコメント番号
	ThreadNo string
	//Contents はコメント本体
	Contents string
	//CommentDate はコメントを書き込んだ日
	CommentDate string
}

//Thread はスレッド
type Thread struct {
	//ThreadKey はスレッドの文字列 yyyy-mm-dd hh:mm:ss
	ThreadKey string
	//Title はスレッド名
	Title string
	//isClose は1000まで到達しているか
	isClose bool
}

//GetAllHTMLData は全ページのデータ取得
func GetAllHTMLData(URL string) {
	//InsertThread(URL)

	InsertComment(URL)
}

//InsertThread はスレッドの挿入
func InsertThread(URL string) {
	/*targetURL := fmt.Sprintf("%s/%d", URL, 1)
	reader := GetStream(targetURL)
	Thread, err := parseThread(reader)
	if err != nil {
		return
	}

	dbConn, err := dbutil.Connect()
	if err != nil {
		el.Fatal("Cannnot Connect DB %s", err.Error)
		return
	}

	for _, eachComment := range Comments {
		dbConn.Create(&eachComment)
	}
	*/
}

//parseThread はスレッドの解析
func parseThread(reader io.Reader) Thread {
	doc, _ := goquery.NewDocumentFromReader(reader)
	Thread := Thread{}
	doc.Find("#thread_top").Each(func(_ int, s *goquery.Selection) {
		threadTitle := s.Find("h1").Text()
		Thread.Title = threadTitle
	})
	return Thread
}

//InsertComment はコメントのinsert
func InsertComment(URL string) {

	el.Info("--page番号の取得--")
	targetURL := fmt.Sprintf("%s/a/1", URL)
	reader := GetStream(targetURL)
	allPageNo := getPageNo(reader)
	el.Info(`--全page数 ` + allPageNo + `--`)

	var maxLoop int

	if os.Getenv("TEST_LOOP") != "" {
		maxLoop, _ = strconv.Atoi(os.Getenv("DO_TEST"))
	} else {
		maxLoop, _ = strconv.Atoi(allPageNo)
	}

	comments := Comments{}
	for i := 1; i <= maxLoop; i++ {
		strInt := strconv.Itoa(i)
		el.Info(`--page ` + strInt + `--`)
		comments = GetCommentData(comments, URL, i)
	}

	sort.SliceStable(comments, func(i, j int) bool { return comments[i].ThreadNo < comments[j].ThreadNo })

	dbConn, err := dbutil.Connect()
	if err != nil {
		el.Fatal("Cannnot Connect DB")
	}
	defer dbConn.Close()

	for _, eachComment := range comments {
		dbConn.Create(&eachComment)
	}
}

//GetCommentData はコメントデータの取得
func GetCommentData(comments Comments, URL string, pageNo int) Comments {
	targetURL := fmt.Sprintf("%s/%d", URL, pageNo)
	reader := GetStream(targetURL)
	comments = parseComment(reader, comments)
	return comments
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

//getPangeNo はページ番号を取得
func getPageNo(reader io.Reader) string {
	var pageNo string
	doc, _ := goquery.NewDocumentFromReader(reader)
	doc.Find("#jump_bt").Each(func(_ int, s *goquery.Selection) {
		pageNo = s.Find("span.total").Text()
	})
	return pageNo
}

//parseComment はページごとのコメントデータを取得
func parseComment(reader io.Reader, comments Comments) Comments {

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
		comments = append(comments, CommentData)
	})

	return comments
}
