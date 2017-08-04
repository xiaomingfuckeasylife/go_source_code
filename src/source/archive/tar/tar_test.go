package tar

import (
	"testing"
	"bytes"
	"strings"
	"fmt"
	"strconv"
)

// TODO : (implements the test)

func TestRandom(t *testing.T){



	var x int64 = 1000		// 1 * 8^3 + 7 * 8^2 + 5 * 8
	println(strconv.FormatInt(x , 8),"----------")

	i := int64(269) // 2^8 + 2^2 + 1  => 0001 0000 0101 if it is out of boundary.
	//i >>= 4		// 					  0000 0001 0000 0101
	fmt.Printf("%v ***" , byte(i))

	// x raised to the power of the column from the right -1) * (the number found in the column)
	// base2 		   0
	// base8 		 000
	// base16		0000
	// base32	   00000	// every five bits indicator a letter

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

	// 1000 0000 0x80
	// 0xff 15 * 16 + 15 = 256
	// 1 0000 0000
	// 0x40 4 * 16 = 64
	// 0100 0000

}

func TestASCII(t *testing.T){

	s := "htis is a 你好word"

	println(isASCII(s))

	println(toASCII(s))
}


func TestParse(t *testing.T) {

	b := []byte("htis is a \x00 你好word")

	p := &parser{}

	println(p.parseString(b))

}
