package zip

import (
	"testing"
	"archive/zip"
	"io"
	"os"
	"bytes"
	"compress/flate"
	"fmt"
)

// read a zip file
func TestZipReader(t *testing.T){
	// open zip
	rc , err := zip.OpenReader("C:\\Users\\clark\\Desktop\\timeline\\important\\mysqlBack.zip")
	if err != nil {
		t.Fatal(err.Error())
	}
	// remember to close it when we are done with zip when we use ReaderClose
	defer  rc.Close()
	for _ , f := range rc.File {
		frc , err := f.Open()
		if err != nil {
			t.Fatal(err.Error())
		}
		io.CopyN(os.Stdout,frc ,100)
		frc.Close()
		println()
	}
}

// write a zip files . quite like tar
func TestZipWriter(t *testing.T){
	buf := new(bytes.Buffer)
	zw := zip.NewWriter(buf)
	var files = []struct {
		name string
		body string
	}{
		{name:"1.txt",body: "hello world"},
		{name:"2.txt",body: "hello world"},
	}
	for _ , f := range files {
		w , err := zw.Create(f.name)
		if err != nil {
			t.Fatal(err.Error())
		}
		_ , err = w.Write([]byte(f.body))
		if err != nil {
			t.Fatal(err.Error())
		}
	}
	err := zw.Close()
	if err != nil{
		t.Fatal(err.Error())
	}
}

// write customer compress zip files
func TestZipCstmCps(t *testing.T){
	buf := new(bytes.Buffer)
	zw := zip.NewWriter(buf)
	zw.RegisterCompressor(zip.Deflate, func(w io.Writer) (io.WriteCloser, error) {
		return flate.NewWriter(w,flate.BestCompression)
	})
	// now we can compress zip  file  using customer compressor
}
