package comutil

import (
	"fmt"
	"math"
)

/**
 * 小数点の切り捨て
 * @type float64 f 少数
 * @type int places 桁数
 * @return float64 切り捨て後の少数
 */
func Round(f float64, places int) float64 {
	shift := math.Pow(10, float64(places))
	return math.Floor(f*shift+.5) / shift
}

/**
 * mapのデバッグ
 *
 * @return map mapdata マップデータ
 */
func DebugMap(mapdata map[string]string) {
	for k, v := range mapdata {
		fmt.Println(fmt.Sprintf("key:%s  value:%s", k, v))
	}
}
