package main

import "fmt"

func readStr(str *string , c chan int , c1 chan int){
	<-c
	println("the current str value is ", *str)
	c1<-1
}

func writeStr(str *string , c chan int){

	*str = "change str2"
	c<-1

}


func changeSlice(slice []string) {
	slice[0] = "1212"
}

func main(){

	slice := []string{"1","2","3","4"}
	changeSlice(slice)
	fmt.Print(slice[0])

	//
	//var arr [3]string
	//arr1 := [3]string{"1","2","3"}
	//arr = arr1
	//arr[1] = "a"
	//fmt.Println(arr[1],arr1[1])



	//fmt.Printf("the markt is going down right now %s" , "d" )
	//c := make(chan int)
	//c1 := make( chan int )
	//str := "hello,world"
	//go readStr(&str,c ,c1)
	//go writeStr(&str,c)
	//str = "change str1"
	//<-c1
	//for {
	//	<-time.After(10 * time.Second)
	//	println("ten second passed")
	//}

	//wordPtr := flag.String("word","foo","a string")
	//numPtr  := flag.Int("numb",42,"an int")
	//boolPtr := flag.Bool("fork",false,"a bool")
	//
	//var svar string
	//flag.StringVar(&svar,"svar","bar","a string var")
	//
	//flag.Parse()
	//
	//fmt.Printf("%v %v %v %v %v" , *wordPtr , *numPtr , *boolPtr , svar , flag.Args())

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
}
