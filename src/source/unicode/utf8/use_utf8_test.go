package utf8

import (
	"testing"
	"unicode/utf8"
	"fmt"
)

// this package is used to judge if the character
// utf8 used more often than utf16 and any other . because utf8 indicate usually one Character

// test using decode last rune method
func TestDecodeLastRune(t *testing.T){
	b := []byte("hello,世界")
	for {
		r , n := utf8.DecodeLastRune(b)
		if n == 0 {
			break
		}
		fmt.Printf("%q \n",r)
		b = b[:len(b)-n]
	}
}

// test decode rune one at a time
func TestDecodeRune(t *testing.T){
	b := []byte("hello,世界")
	for {
		r , n := utf8.DecodeRune(b)
		if n == 0 {
			break
		}
		fmt.Printf("%q\n",r)
		b = b[n:]
	}
}

// encode a rune using utf8
func TestEncodeRune(t *testing.T){
	r := '世'
	buf :=make([]byte,3)
	n := utf8.EncodeRune(buf,r)
	fmt.Printf("%d \n",buf) //[228 184 150]
	println(n)
}

// test some byte if they are a full rune .
func TestFullRune(t *testing.T){
	b := []byte{228,184,150}
	println(utf8.FullRune(b))
	println(utf8.FullRune(b[:len(b)-1]))
}

//check how many count there are
func TestRuneCount(t *testing.T){

	b := []byte("hello,世界")
	println(utf8.RuneCount(b))

}

// how long is the rune is .s
func TestRuneLen(t *testing.T){

	r :=rune('世')
	println(utf8.RuneLen(rune(r)))
	r = rune('h')
	println(utf8.RuneLen(r))

}

// check the byte is a starter of rune .
func TestRuneStart(t *testing.T){
	b := []byte("a世")
	println(utf8.RuneStart(b[0]))
	println(utf8.RuneStart(b[1]))
	println(utf8.RuneStart(b[2]))
}

// check if the character is a valid utf8 character
func TestValid(t *testing.T){
	valid := []byte("hello")
	inValid :=[]byte{ 0xff , 0xfe,0xfd}
	println(utf8.Valid(valid),utf8.Valid(inValid))
}

