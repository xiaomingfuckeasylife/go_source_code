package md5

import (
	"testing"
	"crypto/md5"
	"io"
	"fmt"
	"os"
)

// New MD5 checksum
func TestMd5New(t *testing.T){
	m := md5.New()
	_ , err := io.WriteString(m , "this is a test")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("%x",string(m.Sum(nil)))
	fmt.Printf("%v",string(m.Sum([]byte("this is a salt"))))
}

// new md5 check sum from file content
func TestMd5File(t *testing.T){
	f , err := os.Open("c:/Users/clark/go_source_code/src/source/crypto/md5/text.txt")
	if err != nil {
		t.Fatal(err.Error())
	}
	m := md5.New()
	_ , err = io.Copy(m , f)
	if err != nil {
		t.Fatal(err.Error())
	}
	fmt.Printf("%x",m.Sum(nil))
}

// sum return the checksum of the data
func TestCheckSum(t *testing.T){
	data := []byte("this is a test")
	buf := md5.Sum(data)
	fmt.Printf("%x",buf)
}