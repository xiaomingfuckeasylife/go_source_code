package bufio

import (
	"testing"
	"bufio"
	"bytes"
	"strconv"
	"fmt"
	"os"
)

// bufio mainly read textual files . read a bunch of bytes into a buffer .

// using ScanWord function to filter input string
func TestScanWord(t *testing.T){

	const input = "1234 4566 123 1332444121314212"

	scan := bufio.NewScanner(bytes.NewReader([]byte(input)))

	split:= func(data []byte, atEOF bool) (advance int, token []byte, err error){
		advance , token , err = bufio.ScanWords(data,atEOF)
		if token != nil && err == nil {
			_ , err = strconv.ParseInt(string(token),10,32)
		}
		return
	}

	scan.Split(split)

	for scan.Scan()  {
		fmt.Println(scan.Text())
	}

	if err := scan.Err(); err != nil {
		t.Fatal(err.Error())
	}
}

// a empty string as the final token
func TestErrFinalToken(t *testing.T){

	s := "1,2,3,4, "

	sc := bufio.NewScanner(bytes.NewReader([]byte(s)))

	split:=func(data []byte, atEOF bool) (advance int, token []byte, err error){
		for i:=0 ; i < len(data);i++ {
			if data[i] == ',' {
				return i+1 , data[:i],nil
			}
		}
		// tell the scaner there is no more token left . not trigger a err in scan .
		return 0 ,nil ,bufio.ErrFinalToken
	}

	sc.Split(split)

	for sc.Scan() {
		fmt.Printf("%v",sc.Text())
	}

	if err :=sc.Err();err != nil  {
		fmt.Fprintf(os.Stderr,"err %v",err.Error())
	}

}

// the test has to be test in a Main package line separator by stdio
func TestLineScan(t *testing.T){

	sc := bufio.NewScanner(os.Stdin)

	for sc.Scan() {
		println(sc.Text())
	}

	if err := sc.Err(); err != nil {
		println(err.Error())
	}
}

// test words count scanner
func TestWordsCount(t *testing.T){

	s := "I am a genius"
	sc := bufio.NewScanner(bytes.NewReader([]byte(s)))
	sc.Split(bufio.ScanWords)
	count:=0
	for sc.Scan() {count++}
	if err := sc.Err();err != nil {
		println(err.Error())
	}
	println("count : ", count)

}