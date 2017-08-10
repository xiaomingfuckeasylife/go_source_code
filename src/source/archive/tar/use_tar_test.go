package tar

import (
	"testing"
	"bytes"
	"archive/tar"
	"io"
	"os"
	"fmt"
)

type file struct{
	name string
	body string
}

// this is a example showed how to write a tar file. we are not actually create a disk file we just create the tar file in the memory .
// so we are not indicator the disk file name .like a.tar
func TestWriteReadTar(t *testing.T){

	// create a buffer to stole tar data
	buf :=new(bytes.Buffer)
	// files
	files := []file{
		{"1.txt","this is 1.txt"},
		{"2.txt","this is 2.txt"},
	}
	// new a tar file writer
	w := tar.NewWriter(buf)
	for _ , f := range files {
		err := w.WriteHeader(&tar.Header{
			Name:f.name,
			Mode:0600,
			Size:int64(len(f.body)),
		})
		if err != nil {
			t.Fatal(err.Error())
		}
		_ , err = w.Write([]byte(f.body))
		if err != nil {
			t.Fatal(err.Error())
		}
	}

	if err := w.Close() ; err != nil {
		t.Fatal(err.Error())
	}

	// read tar
	//r := tar.NewReader(buf)  // we can also do it like this but not easy to understand
	br  := bytes.NewReader(buf.Bytes())
	r  := tar.NewReader(br)
	for {
		// using a pointer like a iterator
		hdr , err := r.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Fatal(err.Error())
		}
		println(hdr.Name)
		if _ ,err = io.Copy(os.Stdout,r); err != nil {
			t.Fatal(err.Error())
		}
		fmt.Println()
	}
}

