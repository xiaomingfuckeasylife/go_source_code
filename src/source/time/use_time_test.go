package time

import (
	"testing"
	"fmt"
	"time"
)

// using like a timer wait a specific time then execution some action
func TestAfter(t *testing.T){
	c := make(chan int , 1)
	//c <- 0
	//out :for {
		select {
			case v :=<-c:
				fmt.Println(v)
			case <-time.After(time.Second * 10):
				fmt.Println("time out ")
				//break out
		}
	//}
}

// Sleep pauses the current goroutine sleep a few times. if the wait time is 0 or negative time the time return immediately
func TestSleep(t *testing.T){

	time.Sleep(10  * time.Second)
	println("now i am freshing up")

}

// this is simulate a timer like a timer in java
func TestTicker(t *testing.T){

	tm := time.Tick(time.Second * 10) // chan time
	for now :=range tm {
		fmt.Printf("%v\n",now)
	}

}

// duration is just a type for int64 for time is probably vary big so we using a int64 and a alia for that type
func  TestDuration(t *testing.T){

	d := 10
	println(time.Duration(d) * time.Second)

}

// declare a date
func TestDate(t *testing.T){

	date := time.Date(2017,time.September,10,10,10,10,0,time.UTC)
	println(date.String())

}

// test using month
func TestMonth(t *testing.T){

	y , m , d := time.Now().Date()
	fmt.Printf("this is year = %d , month=%d , day=%d",y,m,d)
	if time.November == m && d == 10 {
		println("Happy Go day")
	}

}

func TestParse(t *testing.T){
	const testDate  = "2017-9-09"
	ti , err :=time.Parse("2017-9-09",testDate)
	if err != nil {
		t.Fatal(err.Error())
	}
	y , m , d :=ti.Date()
	println(y , m , d )
}













