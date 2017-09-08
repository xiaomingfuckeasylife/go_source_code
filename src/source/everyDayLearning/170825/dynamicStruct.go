package main

// how to use reflect in go to dynamic create a struct type .
import (
"fmt"
"reflect"
)

var typeRegistry = make(map[string]reflect.Type)

func registerType(elem interface{}) {
	// get underly typed element
	t := reflect.TypeOf(elem).Elem()
	typeRegistry[t.Name()] = t
}

func newStruct(name string) (interface{}, bool) {
	elem, ok := typeRegistry[name]
	if !ok {
		return nil, false
	}
	// create a type interface by the type element .
	return reflect.New(elem).Elem().Interface(), true
}

func init() {
	registerType((*test)(nil))
}

type test struct {
	Name string
	Sex  int
}

// the efficiency is not very good .
func main() {
	structName := "test"

	s, ok := newStruct(structName)
	if !ok {
		return
	}

	fmt.Println(s, reflect.TypeOf(s))

	t, ok := s.(test)
	if !ok {
		return
	}
	t.Name = "i am test"
	fmt.Println(t, reflect.TypeOf(t))
}
