package list

import (
	"testing"
	"container/list"
	"fmt"
)

func TestList(t *testing.T){

	l := list.New()
	l.PushFront("abc")
	l.PushBack(12)
	l.PushFront(true) // 12 abc true
	for  e := l.Front() ; e != nil ; e = e.Next() {
		fmt.Printf("%v\n",e.Value)
	}
}