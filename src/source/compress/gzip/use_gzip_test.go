package gzip

import (
	"testing"
	"compress/gzip"
	"bytes"
	"time"
	"fmt"
	"io"
	"os"
)

// reading and writing gzip format compress file
func TestGzipReadWrite(t *testing.T){

	var buf bytes.Buffer
	gw := gzip.NewWriter(&buf)
	gw.Name = "gzip.txt"
	gw.Comment = "this is a test"
	gw.ModTime = time.Now()
	if _ , err := gw.Write([]byte("hello world")) ; err != nil {
		t.Fatal(err.Error())
	}
	if err := gw.Flush(); err != nil {
		t.Fatal(err.Error())
	}
	if err := gw.Close(); err != nil {
		t.Fatal(err.Error())
	}

	gr , err := gzip.NewReader(&buf)
	if err != nil {
		t.Fatal(err.Error())
	}
	fmt.Printf("%s %s %s \n",gr.Name , gr.Comment , gr.ModTime )
	io.Copy(os.Stdout,gr)

}