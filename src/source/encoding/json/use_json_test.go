package json

import (
	"testing"
	"encoding/json"
	"strings"
	"fmt"
	//"io"
	"io"
	"bytes"
	"os"
)

// make json data into struct
func TestJsonDecoding(t *testing.T){
	const jsonStr = `
	[
			{"name":"clark","age":10},
			{"name":"xiaoming","age":11}
	]
	`
	// create a json decoder
	dec := json.NewDecoder(strings.NewReader(jsonStr))

	var js []jsonStruct
	if err := dec.Decode(&js) ; err != nil {
		t.Fatal(err.Error())
	}
	for _ , c :=  range js {
		println(c.name,c.Age)
	}
}

type jsonStruct struct {
	name string
	Age int
	Id int
}

func (js jsonStruct) Stringer() string{
	return fmt.Sprintf("name is %s , age is %d",js.name,js.Age)
}

// parse json array stream
func TestJsonArrayStream(t *testing.T){
	const jsonStr = `
	[
			{"name":"clark","age":10},
			{"age":11,"id":1}
	]
	`
	dec :=json.NewDecoder(strings.NewReader(jsonStr))
	//read open bracket
	tt ,err := dec.Token()
	if err != nil {
		t.Fatal(err.Error())
	}
	fmt.Printf("1 %v",tt)
	// while the array contains values
	for dec.More() {
		var js jsonStruct
		if err = dec.Decode(&js) ; err != nil {
			t.Fatal(err.Error())
		}
		println(js.name,js.Age,js.Id)
	}
	// reading closing bracket
	ttt , err := dec.Token()
	if err != nil {
		t.Fatal(err.Error())
	}
	fmt.Printf(" 2 %v",ttt)
	//if _ , err = dec.Token(); err != nil {
	//	t.Fatal(err.Error())
	//}
}

func TestJsonToken(t *testing.T){
	const stream =`
		{"Message":"hello","array":[1,2,3],"Null":null,"Number":1.234}
	`
	dec := json.NewDecoder(strings.NewReader(stream))
	for {
		tt, err :=dec.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Fatal(err.Error())
		}
		fmt.Printf("%v\n",tt)
		if dec.More() {
			println("there is more ")
		}
	}
}

// indent the json data using prefix and suffix
func TestIndent(t *testing.T){
	type Road struct{
		Name string
		Number int
	}
	roads := []Road{
		{Name:"first streat",Number:1},
		{Name:"second streat",Number:2},
	}

	b , err := json.Marshal(roads)
	if err != nil {
		t.Fatal(err.Error())
	}
	var buf bytes.Buffer
	if err = json.Indent(&buf,b,"=","\t"); err != nil {
		t.Fatal(err.Error())
	}
	//io.Copy(os.Stdout,bytes.NewReader(b))
	buf.WriteTo(os.Stdout)
}

//  Marshal struct into json string
func TestMarshal(t *testing.T){
	jsArr := []jsonStruct{
		{name:"clark",Age:10,Id:1},
		{name:"clark1",Age:11,Id:2},
	}
	buf , err :=json.Marshal(jsArr)
	if err != nil {
		t.Fatal(err.Error())
	}
	os.Stdout.Write(buf)
}

// unmarshal json bytes to struct data . Decode is quite like unmarshal but more flexible
func TestUnmashal(t *testing.T){
	var jsStr jsonStruct
	err :=json.Unmarshal([]byte(`{"age":1,"Id":0}`),&jsStr)
	if err != nil{
		t.Fatal(err.Error())
	}
	println(jsStr.Age,jsStr.Id)
}

// RawMessage using to predefined some value ,and computation stuff
func TestRawMessage(t *testing.T){
	b := json.RawMessage(`{"Precomputed":true}`)
	c := struct {
		Header *json.RawMessage `json:"Header"`
		Body string 			`json:"body"`
	}{&b,"hello body"}
	buf , err := json.MarshalIndent(&c,"","\t")
	if err != nil {
		t.Fatal(err.Error())
	}
	os.Stdout.Write(buf)
}