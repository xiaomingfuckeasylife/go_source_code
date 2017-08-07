package tar

import (
	"bytes"
	"strconv"
	"strings"
	"fmt"
	"time"
)

// check if str contains any byte that is over 127 which is the boundary of ASCII[-128,127]
// the name for string type I normally use str . but in go everything should be as simple as possible .
func isASCII(s string) bool{
	for _ , v := range s {  // this place the source code using alia c instead of v . c stand for current value . they bose cool.
		if v >= 0x80 {
			return false
		}
	}
	return true
}


// get the ascii out of a string .
func toASCII(s string) string{

	if isASCII(s) {
		return s
	}

	//var b []byte
	//
	//for _ , c := range s {
	//	if c >= 0x80 {
	//		continue;
	//	}
	//	b = append(b, byte(c))
	//}

	//return string(b)

	// the source code using byteBuffer and it's method WriteBuffer() .

	var b bytes.Buffer

	for _ , c := range s {
		if c < 0x80 {
			b.WriteByte(byte(c))
		}
	}
	return b.String()
}

// parser parse []byte data into some other type
// the err indicator the current parser success or not
type parser struct {
	err error
}

// formatter the opposite of parser format other type to []byte
// the same as parser
type formatter struct {
	err error
}

// parse []byte data into string . it stops when the byte is null value .
// which means 0. which indicator the stream is end.
// like in c we always add one byte to the end of string indicator the end
func (parse *parser) parseString(b []byte) string {

	//for i:= 0 ;i < len(b); i++ {
	//	if b[i] == 0 {
	//		return string(b[0:i])
	//	}
	//}
	//return string(b)

	// compare to the code i write the source code is much better of course . usually we use n alia to stand for the current times ,
	// count down , number

	n := 0

	for n < len(b) && b[n] != 0 {
		n++
	}

	return string(b[0:n])
}

// get ascii out of s and put it int b . and remember that byte end indicator is 0
func (format *formatter) formatString(b []byte, s string){

	if len(b) < len(s) {
		format.err = ErrWriteTooLong
		return
	}
	s = toASCII(s)

	copy(b , s)

	if len(b) > len(s) {
		b[len(s)] = 0
	}
}

// check if x can be encoded in n bytes .
// base256 does not need to put a extra null byte to the end indicator the end of string like octal does .
// so all the n byte can be used as output.
// 256 = 2^8 which means it has eight bytes which is 64 bits.if n > 8 then the int64 data must be able to encoded in .
// If operating in binary mode, this assumes strict GNU binary mode; which means
// that the first byte can only be either 0x80 or 0xff. Thus, the first byte is
// equivalent to the sign bit in two's complement form.
func fitsInBase256(n int , x int64) bool {
	binBits := uint((n-1)) * 8 // how many binary column it has
	// include the left contains and not right
	return n >= 9 || (x >= -1 << binBits && x < 1 << binBits)
}

// parsing numeric using either base256 or base8
// may return negative value
// if the integer is overflow . the a error will be set
func (p *parser) parseNumeric(b []byte) int64{
	// Check for base-256 (binary) format first.
	// If the first bit is set, then all following bits constitute a two's
	// complement encoded number in big-endian byte order.
	if len(b) > 0 && b[0]&0x80 != 0 {
		p.parseBase256(b)
	}
	return p.parseOctal(b)
}

// TODO this part not quite understand .
// but i think the code bellow is how the base256 rule works so there is that .
func (p *parser) parseBase256(b []byte) int64{
	// Handling negative numbers relies on the following identity:
	//	-a-1 == ^a
	//
	// If the number is negative, we use an inversion mask to invert the
	// data bytes and treat the value as an unsigned number.
	var inv byte // 0x00 if positive or zero, 0xff if negative
	if b[0]&0x40 != 0 {
		inv = 0xff
	}

	var x uint64
	for i, c := range b {
		c ^= inv // Inverts c only if inv is 0xff, otherwise does nothing
		if i == 0 {
			c &= 0x7f // Ignore signal bit in first byte
		}
		if (x >> 56) > 0 {
			p.err = ErrHeader // Integer overflow
			return 0
		}
		x = x<<8 | uint64(c)
	}
	if (x >> 63) > 0 {
		p.err = ErrHeader // Integer overflow
		return 0
	}
	if inv == 0xff {
		return ^int64(x)
	}
	return int64(x)
}

// parse byte into int64 using base8
func (p *parser) parseOctal(b []byte) int64{
	// trim left pad and tail pad of space and null value
	b = bytes.Trim(b," \x00")
	if len(b)  == 0{
		return 0
	}
	u ,uErr := strconv.ParseUint(p.parseString(b),8,64)
	if uErr != nil {
		p.err = ErrHeader
	}
	return int64(u)
}

func (f *formatter) formatNumeric(b []byte, x int64){
	if fitsInBase256(len(b),x) {
		f.formatBase256(b,x)
	}
	//return f.formatOctal(b)
}

// This is the reverse algorithm of  parseBase256 so we can understand bose by each other
func (f *formatter) formatBase256(b []byte , x int64){
	for i:=len(b)-1 ; i >= 0 ;i-- {
		b[i] = byte(x)  // when it is out of scope (-127,127) then it will circulate . for example if the x is 269 the byte(269) is actually 13
		x >>= 8			// left move 8 bits.
	}
	b[0] |= 0x80     // indicator the binary format
}

func (f *formatter) formatOctal(b []byte , x int64) {
	s := strconv.FormatInt(x , 8)
	if n := len(b) - len(s) - 1 ; n > 0 {
		s = strings.Repeat("0",n) + s
	}
	f.formatString(b,s)
}

// TODO: the code bellow is absolutely copy by hand .
// TODO:just typinng we can still learn a lot remember. do not take anything for granted .

// parsePAXTime takes a string of the form %d.%d as described in the PAX
// specification .Note that this implementation allows for negative timestamps,
// which is allowed for by the PAX specification , but not always portable.
func parsePAXTime(s string) (time.Time,error) {
	const maxNanoSecondDigits = 9

	// Split string into seconds and sub-seconds parts .
	ss , sn := s , ""
	// by the way it seems to me the strings.IndexByte and strings.Index does not has much difference .
	if pos := strings.IndexByte(ss ,'.');pos >= 0 {
		ss  , sn = ss[:pos] , ss[pos+1:]
	}

	// Parse the seconds
	secs , err := strconv.ParseInt(ss,10,64)

	if err != nil {
		return time.Time{} , ErrHeader
	}

	if len(sn) == 0 {
		return time.Unix(secs,0) ,  nil
	}

	if strings.Trim(sn,"0123456789") != "" {
		return time.Time{} ,ErrHeader
	}

	if len(sn) < maxNanoSecondDigits {
		sn += strings.Repeat("0" , maxNanoSecondDigits - len(sn)) // Right pad
	}else {
		sn = sn[:maxNanoSecondDigits]								 // Right truncate
	}

	nsecs , _ := strconv.ParseInt(sn , 10 , 64) // Must succeed
	if len(ss) > 0 && ss[0] == '-' {
		return time.Unix(secs , -1 * nsecs),nil
	}
	return time.Unix(secs,nsecs),nil
}

// parsePAXRecord parses the input PAX record stirng into a key-value pair.
// If parsing is successful . it will slice off the currently read record and
// return the remainder as r.
//
// A PAX record is of the following form .
// `%d = %s=%s\n`%(size , key , value)
func parsePAXRecords(s string) (k , v , r string , err error){
	sp := strings.IndexByte(s , ' ')
	if sp == -1 {
		return "", "", s , ErrHeader
	}

	n , perr := strconv.ParseInt(s[:sp], 10 , 0)

	if perr != nil || n < 5 || int64(len(s)) < n {
		return "" , "" , s , ErrHeader
	}

	rec , nl , rem := s[sp+1:n-1] , s[n-1:n] , s[n:]
	if nl != "\n" {
		return "","",s , ErrHeader
	}

	eq := strings.IndexByte(rec , '=')

	if eq == -1 {
		return "" , "" , s , ErrHeader
	}
	return rec[:eq] , rec[eq+1:],rem , nil
}

// formatPAXRecord formats a single PAX record , prefixing it with the appropriate length
func formatPAXRecord(k,v string) string {
	const padding = 3 // Extra padding for ' ' , '=' , and '\n'
	size := len(k) + len(v) + padding
	size += len(strconv.Itoa(size))		// integer size
	record := fmt.Sprintf("%d %s=%s\n",size,k , v)

	// Final adjestment if adding size field increased the record size.
	if len(record) != size {
		size = len(record)
		record = fmt.Sprintf("%d %s=%s\n",size,k , v)
	}
	return record
}