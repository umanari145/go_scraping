package dateutil

import (
	"strconv"
	"strings"
	"time"
)

/**
 * 任意の日付フォーマットを与えられて整形して日付オブジェクトを返す
 * @type string dateStr 任意の日付けデータyyyyMMdd/yyyy-MM-dd/yyyy/MM/dd
 * @return time.Time 日付けオブジェクト
 */
func GetDateObject(dateStr string) time.Time {
	dateStr2 := strings.Replace(dateStr, "-", "", -1)
	dateStr3 := strings.Replace(dateStr2, "/", "", -1)
	yearStr := dateStr3[0:4]
	monthStr := dateStr3[4:6]
	dayStr := dateStr3[6:8]
	hourStr := dateStr3[8:10]
	minitusStr := dateStr3[10:12]
	secondStr := dateStr3[12:14]

	yearInt, _ := strconv.Atoi(yearStr)
	monthInt, _ := strconv.Atoi(monthStr)
	dayInt, _ := strconv.Atoi(dayStr)
	hourInt, _ := strconv.Atoi(hourStr)
	minitusInt, _ := strconv.Atoi(minitusStr)
	secondInt, _ := strconv.Atoi(secondStr)

	dateObj := time.Date(yearInt, time.Month(monthInt), dayInt, hourInt, minitusInt, secondInt, 0, time.Local)
	return dateObj
}
