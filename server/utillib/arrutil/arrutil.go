package arrutil

import (
	"strings"
)

/**
 * ハッシュをグループ化させる
 * @type slice arr対象の配列
 * @type slice keyArr グループキー
 * @type string splitKey キー同士をつなぐ文字列
 * @return map グループ化された配列
 */
func GroupByMultiKey(arr []map[string]string, keyArr []string, splitKey string) map[string][]map[string]string {
	groupedHash := map[string][]map[string]string{}
	for _, v := range arr {
		var hashKeyArr []string
		var hashKey string
		hashKey = ""
		for _, k := range keyArr {
			hashKeyArr = append(hashKeyArr, v[k])
		}
		hashKey = strings.Join(hashKeyArr, splitKey)
		groupedHash[hashKey] = append(groupedHash[hashKey], v)
	}
	return groupedHash
}

/**
 * 対象要素が配列の中にあるか
 * @type string val 対象要素
 * @type slice array 配列
 * @return bool true(存在する)/false(存在しない)
 */
func InArray(val string, array []string) (isInArray bool) {
	for _, v := range array {
		if val == v {
			return true
		}
	}
	return false
}

/**
 * 配列のキーを特定コードなどに置き換える
 * @type string listKey 特定コード
 * @type map arr 対象の配列
 * @type map キーを特定の文字列にした配列
 */
func MakeArrToHash(listKey string, arr []map[string]string) map[string][]map[string]string {
	hashArr := map[string][]map[string]string{}
	for _, v := range arr {
		hashArr[v[listKey]] = append(hashArr[v[listKey]], v)
	}
	return hashArr
}

/**
 * マップからキーと値を取得
 * @type map map
 * @return []string キーの配列
 * @return []string 値の配列
 */
func GetmapEle(singleMap map[string]string) ([]string, []string) {
	var keylist []string
	var vallist []string
	for key, values := range singleMap {
		keylist = append(keylist, key)
		vallist = append(vallist, values)
	}
	return keylist, vallist
}

/**
 *  arr1 < arr2 の前提でarr1にあってarr2にない集合を出力
 * @type []string arr1 配列1
 * @type []string arr2 配列2
 * @return []string 値の配列
 */
func GetArrDiff(arr1 []string, arr2 []string) []string {
	var arr3 []string
	for _, val := range arr2 {
		if InArray(val, arr1) == false {
			arr3 = append(arr3, val)
		}
	}
	return arr3
}
