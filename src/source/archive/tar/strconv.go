package tar

import "bytes"

// check if str contains any byte that is over 127 which is the boundary of ASCII(0 - 127)
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



