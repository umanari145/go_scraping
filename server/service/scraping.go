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
	//ThreadID はthread番号
	ThreadID int
	//ThreadNo はコメント番号
	ThreadNo string
	//Contents はコメント本体
	Contents string
	//CommentDate はコメントを書き込んだ日
	CommentDate string
}

//Thread はスレッド
type Thread struct {
	//URL スレッドのURL
	URL string
	//ThreadKey はスレッドの文字列 yyyy-mm-dd hh:mm:ss
	ThreadKey string
	//Title はスレッド名
	Title string
	//isClose は1000まで到達しているか
	isClose bool
}

var currentMaxComment int
var currentDonePageNo int

//InsertComment はコメントのinsert
func InsertComment(threadID int) {

	URL := getThreadURL(threadID)

	if URL == "" {
		el.Info("処理対象のスレッドは存在しません")
		return
	}

	el.Info("--page番号の取得--")
	targetURL := fmt.Sprintf("%s/a/1", URL)
	el.Infof("targetURL %s", targetURL)
	reader := getStream(targetURL)
	allPageNo := getPageNo(reader)
	el.Infof("--全page数 %s--", allPageNo)

	getCurrentMaxComment(threadID)

	var maxLoop int

	fmt.Println(allPageNo)

	if os.Getenv("TEST_LOOP") != "" {
		maxLoop, _ = strconv.Atoi(os.Getenv("DO_TEST"))
	} else {
		tmpLoop, _ := strconv.Atoi(allPageNo)
		maxLoop = tmpLoop - currentDonePageNo
		if maxLoop == 0 {
			el.Info("取得すべきページはありません。")
		}
	}
	fmt.Println(maxLoop)
	comments := Comments{}
	for i := 1; i <= maxLoop; i++ {
		strInt := strconv.Itoa(i)
		el.Info(`--page ` + strInt + `--`)
		comments = getCommentData(comments, URL, i)
	}

	sort.SliceStable(comments, func(i, j int) bool { return comments[i].ThreadNo < comments[j].ThreadNo })

	dbConn, err := dbutil.Connect()
	if err != nil {
		el.Fatal("Cannnot Connect DB")
	}
	defer dbConn.Close()

	for _, eachComment := range comments {
		eachComment.ThreadID = threadID
		dbConn.Create(&eachComment)
	}

}

//getThreadURL はスレッドのURLを取得
func getThreadURL(threadID int) string {
	dbConn, err := dbutil.Connect()
	if err != nil {
		el.Fatal("Cannnot Connect DB")
		return ""
	}
	defer dbConn.Close()

	thread := Thread{}
	dbConn.Where("id = ? and is_close = ?", threadID, false).First(&thread)

	return thread.URL
}

//CurrentThread はSQLの結果
type CurrentThread struct {
	ThreadNo int
}

//PerComment 1ページあたりのページ番号
const PerComment = 10

//getCurrentMaxComment は現在のコメントの取得
func getCurrentMaxComment(threadID int) {

	dbConn, err := dbutil.Connect()
	if err != nil {
		el.Fatal("Cannnot Connect DB")
		return
	}
	defer dbConn.Close()

	var currentThread CurrentThread
	dbConn.Raw("SELECT MAX(thread_no) AS thread_no FROM comments WHERE thread_id = ?", threadID).Scan(&currentThread)

	if currentThread.ThreadNo != 0 {
		currentMaxComment = currentThread.ThreadNo
		currentDonePageNo = currentThread.ThreadNo / PerComment
	} else {
		currentMaxComment = 0
		currentDonePageNo = 0
	}
	el.Infof("current CommentNo %d", currentThread.ThreadNo)
	el.Infof("current PageNo (has done) %d", currentDonePageNo)

}

//GetCommentData はコメントデータの取得
func getCommentData(comments Comments, URL string, pageNo int) Comments {
	targetURL := fmt.Sprintf("%s/%d", URL, pageNo)
	reader := getStream(targetURL)
	comments = parseComment(reader, comments)
	return comments
}

//GetStream はストリームデータの取得
func getStream(URL string) io.Reader {

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

		threadNoInt, _ := strconv.Atoi(threadNo)

		if threadNoInt > currentMaxComment {

			rawData := s.Find(".res").Text()
			commentDate := s.Find(".date").Text()
			strippedData := strip.StripTags(rawData)

			CommentData.ThreadNo = threadNo
			CommentData.Contents = strippedData
			CommentData.CommentDate = commentDate
			comments = append(comments, CommentData)
		}
	})

	return comments
}
