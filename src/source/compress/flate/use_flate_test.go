package flate

import (
	"testing"
	"bytes"
	"compress/flate"
	"io"
	"strings"
	"fmt"
	"os"
	"sync"
	"log"
)

// test the usage of compress flate read and write using a predefined dictionary
func TestFlate(t *testing.T){

	const dict = `<?xml version="1.0"?>` + `<book>` + `<data>` +`<meta name="` + `" content="`

	const data = `<?xml version="1.0"?>
<book>
	<meta name="title" content="The Go Programming Language"/>
	<meta name="authors" content="Alan Donovan and Brian Kernighan"/>
	<meta name="published" content="2015-10-26"/>
	<meta name="isbn" content="978-0134190440"/>
	<data>...</data>
</book>
`
	// using a buffer to write data into
	var buf bytes.Buffer
	w , err := flate.NewWriterDict(&buf,flate.DefaultCompression,[]byte(dict))
	if err != nil {
		t.Fatal(err.Error())
	}

	// copy data from reader to buffer
	if _ , err = io.Copy(w,strings.NewReader(data)) ; err != nil {
		t.Fatal(err.Error())
	}

	if err = w.Close() ; err != nil {
		t.Fatal(err.Error())
	}

	fmt.Println("decompressed output using the dictionary :")
	rc := flate.NewReaderDict(bytes.NewReader(buf.Bytes()),[]byte(dict))
	if _ , err = io.Copy(os.Stdout,rc); err != nil {
		t.Fatal(err.Error())
	}
	if err = rc.Close() ; err != nil {
		t.Fatal(err.Error())
	}
	fmt.Println()

	fmt.Println("Substrings matched by the dictionary marked with #:")
	hashDict := []byte(dict)
	for i := range hashDict {
		hashDict[i] = '#'
	}
	rc = flate.NewReaderDict(&buf,[]byte(hashDict))
	if _ ,err = io.Copy(os.Stdout,rc); err != nil {
		t.Fatal(err.Error())
	}
	if err = rc.Close(); err != nil {
		t.Fatal(err.Error())
	}
}

//In performance critical applications, Reset can be used to discard the current compressor or decompressor state and
// reinitialize them quickly by taking advantage of previously allocated memory.
func TestFlateReset(t *testing.T){

	proverbs := []string{
		"Don't communicate by sharing memory, share memory by communicating.\n",
		"Concurrency is not parallelism.\n",
		"The bigger the interface, the weaker the abstraction.\n",
		"Documentation is for users.\n",
	}
	// reader to read string str
	var r strings.Reader
	// buffer to zw and zr
	var b bytes.Buffer
	// a io.copyBuffer buffer
	buf := make([]byte, 32 << 10 )
	// create a compress writer
	zw , err := flate.NewWriter(nil,flate.DefaultCompression)
	if err != nil {
		t.Fatal(err.Error())
	}
	// create a compress reader
	zr := flate.NewReader(nil)

	for _ ,s := range proverbs {
		// reset string reader
		r.Reset(s)
		// reset buffer
		b.Reset()
		// reset compress writer
		zw.Reset(&b)
		if _ , err := io.CopyBuffer(zw,&r,buf) ; err != nil {
			t.Fatal(err.Error())
		}
		if err = zw.Close(); err != nil {
			t.Fatal(err.Error())
		}
		// reset compress reader
		if err = zr.(flate.Resetter).Reset(&b,nil); err != nil {
			t.Fatal(err.Error())
		}
		if _ , err := io.CopyBuffer(os.Stdout,zr,buf); err != nil {
			t.Fatal(err.Error())
		}
		if err = zr.Close() ; err != nil {
			t.Fatal(err.Error())
		}
	}
}

// test synchronize compress
func TestSyncrhonization(t *testing.T){
	// add two thread using pipe to simulate the network .
	var wg sync.WaitGroup
	// wait all the task finish
	defer wg.Wait()
	// network pipeline
	rp , wp := io.Pipe()
	// add sender task
	wg.Add(1)
	go func(){
		// finish sending
		defer wg.Done()
		zw , err := flate.NewWriter(wp,flate.BestSpeed) ;
		if err != nil {
			t.Fatal(err.Error())
		}
		// declare a fix byte size
		b := make([]byte,256)
		for _ , m := range strings.Fields("I am genius") {
			// the first byte store the length of bytes
			b[0] = uint8(copy(b[1:],m))
			// has to 1 + len(m) for the first byte is using to store the length of the str
			if zw.Write(b[:1+len(m)]);err != nil {
				t.Fatal(err.Error())
			}
			// flush data into pipe
			if err := zw.Flush();err != nil {
				log.Fatal(err.Error())
			}
		}
		// after using WriterCloser remember to close
		if err = zw.Close(); err != nil {
			t.Fatal(err.Error())
		}
	}()

	// add receiver task . receive data from pipeline
	wg.Add(1)
	go func(){
		// finish task
		defer wg.Done()
		zr := flate.NewReader(rp)
		b := make([]byte,256)
		for {
			// read the first byte to get the number of bytes in the pipe line should read
			if _, err := io.ReadFull(zr,b[:1]);err != nil{
				if err == io.EOF {
					break // The transmitter closed the stream
				}
				t.Fatal(err.Error())
			}
			n := int(b[0])
			// after reading the first byte from the pipeline then there will no such byte in the pipeline.
			if _ , err := io.ReadFull(zr,b[:n]); err != nil {
				log.Fatal(err)
			}
			fmt.Printf("Received %d bytes : %s\n",n , b[:n])
		}
		fmt.Println()
		if err := zr.Close() ; err != nil {
			t.Fatal(err)
		}
	}()
}

