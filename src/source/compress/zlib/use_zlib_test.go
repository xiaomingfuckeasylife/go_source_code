package zlib

import (
	"testing"
	"bytes"
	"compress/zlib"
	"io"
	"os"
	"fmt"
)

// read as gzip
func TestRead(t *testing.T){
	buf := []byte{120, 156, 202, 72, 205, 201, 201, 215, 81, 40, 207,
		47, 202, 73, 225, 2, 4, 0, 0, 255, 255, 33, 231, 4, 147}
	r := bytes.NewReader(buf)
	rc , err := zlib.NewReader(r)
	if err != nil {
		t.Fatal(err.Error())
	}
	io.Copy(os.Stdout,rc)
	if err = rc.Close() ; err != nil {
		t.Fatal(err.Error())
	}
}

// write as gzip
func TestWriter(t *testing.T){
	var buf bytes.Buffer
	zw := zlib.NewWriter(&buf)
	_ , err := zw.Write([]byte("hello,world"))
	if err != nil {
		t.Fatal(err.Error())
	}
	if err := zw.Close(); err != nil {
		t.Fatal(err.Error())
	}
	fmt.Printf("%s",buf.Bytes())
}