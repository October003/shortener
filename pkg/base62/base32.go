package base62

import (
	"math"
	"strings"
)

// 62进制转换的模块

// 0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKMNOPQRSTUVWXYZ

// 0-9: 0-9
// a-z: 10-35
// A-Z: 36-61

// 10进制数    转换    62进制数
//    0                 0
//    1					1
//	  10				a
//    11				b
//    61				Z
//    62				10
//    63				11
//    6347				1En

// 如何实现62进制转换

const base62Str = `0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ`
// 为了避免被人恶意请求，我们可以将上面的字符串打乱


// Int2String 十进制数转为62进制字符串
func Int2String(seq uint64) string {
	if seq == 0 {
		return string(base62Str[0])
	}
	bl := []byte{} // 23 40 1
	for seq > 0 {
		mod := seq % 62
		div := seq / 62
		bl = append(bl, base62Str[mod])
		seq = div
	}
	// 最后把得到数反转一下
	return string(reverse(bl))
}

func reverse(s []byte) []byte {
	for i, j := 0, len(s)-1; i < len(s)/2; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

// String2Int 62进制字符串转换为十进制数
func String2Int(s string) (seq uint64) {
	bl := []byte(s)
	bl = reverse(bl)
	// 从右往左遍历
	for idx, b := range bl {
		base := math.Pow(62, float64(idx))
		seq += uint64(strings.Index(base62Str, string(b))) * uint64(base)
	}
	return seq
}
