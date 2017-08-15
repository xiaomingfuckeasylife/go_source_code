package main

import (
	//"fmt"
)

func main(){

	type test struct {
		a string
	}
	t := test{
		a :"test",
	}
	println(t)
	//b := []byte(" i am ")
	//fmt.Printf("%s",b[:1])
	//println(3 << 2)
	//
	//const dict = `<?xml version="1.0"?>` + `<book>` + `<data>` +`<meta name="` + `" content="`
	//println(dict)
	//sc := bufio.NewScanner(os.Stdin)
	//
	//for sc.Scan() {
	//	if sc.Text() == "exit"{
	//		return
	//	}
	//	println(sc.Text())
	//}
	//
	//if err := sc.Err(); err != nil {
	//	println(err.Error())
	//}

	//println(math.MaxUint8 , " " , math.MaxInt8)
	//
	//b := make([]byte,100)
	//copy( b , "this is ")
	//fmt.Printf("%v",b)
	//c := make(chan string)
	//go func(){
	//	//defer func(){
	//	//	if rv := recover() ; rv != nil{
	//	//		println("recover :", )
	//	//	}
	//	//	println("q")
	//	//}()
	//	panic("this is a panic")
	//	c<-"s"
	//	close(c)
	//}()
	//time.Sleep(10 * time.Second)
	//println("q")
	//println(<-c)
	println("hello,world")
}
