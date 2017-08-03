package tar

import (
	"testing"
	"bytes"
	"strings"
)

// TODO : (implements the test)
func Test1(t *testing.T){

	println(-1 << 56)

	print(4 | 7)
	print(16 >> 3) // 8 1 向右移动三个位置
	// 00000000 00000000 00000000 00000100
	// 00000000 00000000 00000000 00000111
	// 00000000 00000000 00000000 00000111
//println()
//	c:=1
//	d:=2
//	c+=d
//	print(c)

	println(string(bytes.Trim([]byte(" saf asdf  adsf  "), " \x00")))
	println(^uint(0) >> 62)
	// 11111111 11111111 11111111 11111111
	// 00000000 00000000 00000000 00000011

	println(strings.Trim("1212sfadfsf23fs12", "0123456789"))
}
