package bytes

import (
	"testing"
	"sort"
	"fmt"
	"unicode"
	"bytes"
	"os"
	"io"
	"encoding/base64"
)
// bytes package manipulate bytes . analogous to the facilities of the strings packages

// compare the first letter if they are the same then compare the second one .
func TestCompare(t *testing.T){
	b1 := []byte("bcadfafd")
	b2 := []byte("ab")
	println("1 " , bytes.Compare(b1,b2) )
	b2 = []byte("c")
	println("2 " , bytes.Compare(b1,b2))
	b2 = []byte("a")
	println("3 " , bytes.Compare(b1,b2))
}

// the idiot usage of compare search using binary search
func TestCompareSearch(t *testing.T){
	b1 := []byte("ab")
	b2 :=[][]byte{[]byte("ab"),[]byte("cd")}
	// has to has at least one byte larger than 0
	i := sort.Search(len(b2), func(i int) bool{
		return bytes.Compare(b2[i],b1) >= 0
	})
	if i < len(b2) && bytes.Equal(b1 , b2[i]){
		println("found one")
	}
}

// test a byte arr contains a another byte arr
func TestContains(t *testing.T){

	b1 := []byte("a")
	b2 := []byte("ab")

	println(bytes.Contains(b2,b1))
	b1 = nil
	println(bytes.Contains(b2,b1))
	b2 = nil
	println(bytes.Contains(b2,b1))

}

// test the count of byte that has the same int8 value
func TestCount(t *testing.T){
	b1 := []byte("abcde")
	b2 := []byte("a")
	println(bytes.Count(b1,b2))
	println(bytes.Count(b1,nil))
}

// case ignore equal
func TestEqualFolder(t *testing.T){
	println(bytes.EqualFold([]byte("Goa "),[]byte("goA")))
}

// separate by white space
func TestField(t *testing.T){
	fmt.Printf("%v",bytes.Fields([]byte("a b c d ")))
}

// separate by a function
func TestFieldFunc(t *testing.T){
	rb := bytes.FieldsFunc([]byte("i1 am 2gro4ot3"),func(r rune) bool {
		return !unicode.IsLetter(r) && unicode.IsNumber(r)
	})
	fmt.Printf("%q",rb)
}

// check if the bytes start or end with .
func TestHasPrefixAndHasSuffix(t *testing.T){
	b := "i am a genius"
	println(bytes.HasPrefix([]byte(b),[]byte("i")),
	bytes.HasSuffix([]byte(b),[]byte("us")))
}

func TestIndexAll(t *testing.T){
	b1 := []byte("i am a genius")
	b2 := []byte("a")
	println(bytes.Index( b1 ,b2 ))
	println(bytes.IndexAny(b1,"ddde"))
	println(bytes.IndexRune(b1,'k'))
	println("4 " , bytes.IndexFunc(b1,func(r rune) bool{
		return r > 'k'
	}))
}

// concatenates the value
func TestJoin(t *testing.T){
	s := [][]byte{[]byte{'a'},[]byte{'b'},[]byte{'c'}}
	fmt.Printf("%q",bytes.Join(s,[]byte("|")))
}

// filter alter the byte value
func TestMap(t *testing.T){
	println(bytes.Map(func (r rune) rune {
		if r >= 'a' && r <='z' {
			return r + 13
		}
		return r
	},[]byte("abc")))
}

func TestRepeat(t *testing.T){
	fmt.Printf("%q" , bytes.Repeat([]byte("a"),10))
}

func TestReplaces(t *testing.T){
	fmt.Printf("%q",bytes.Replace([]byte("abc"),[]byte("a"),[]byte("d"),1))
}

// test how to use ByteBuffer.
func TestByteBuffer(t *testing.T){
	var buf bytes.Buffer
	_ , err := buf.Write([]byte("hello,world"))
	if err != nil {
		t.Fatal(" this is a error")
	}
	// has to be a pointer.Write method has pointer receiver
	fmt.Fprintf(&buf,"this is %s" , "clark")
	buf.WriteTo(os.Stdout)
}

// byte buffer Reader stuff
func TestByteReader(t *testing.T){
	buf :=bytes.NewBufferString("R29waGVycyBydWxlIQ==")
	dec := base64.NewDecoder(base64.StdEncoding , buf)
	io.Copy(os.Stdout,dec)
}