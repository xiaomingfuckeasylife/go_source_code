package goByExamples

import (
	"testing"
	"os"
	"fmt"
	"flag"
	"os/exec"
	"os/signal"
	"syscall"
)

// test args usage
func TestArgs(t *testing.T){
	argsWithProg := os.Args

	fmt.Printf("\n  %v \n %v %v \n",argsWithProg , argsWithProg[1] , len(argsWithProg))
}

// test flag usage
func TestFlag(t *testing.T){

	wordPtr := flag.String("word","foo","a string")
	numPtr  := flag.Int("numb",42,"an int")
	boolPtr := flag.Bool("fork",false,"a bool")

	var svar string
	flag.StringVar(&svar,"svar","bar","a string var")

	flag.Parse()

	fmt.Printf("%v %v %v %v %v" , *wordPtr , *numPtr , *boolPtr , svar , flag.Args())
}

// test env usage
func TestEnv(t *testing.T){
	os.Setenv("Foo","2")
	println(os.Getenv("Foo"))
	for _ , c :=range os.Environ() {
		println(c)
	}
}

// test exec usage
func TestExec(t *testing.T){
	dateCmd := exec.Command("ipconfig")

	dateOut, err := dateCmd.Output()

	if err != nil {
		panic(err)
	}
	fmt.Println("> ipconfig")
	fmt.Println(string(dateOut))

}


// test signal usage
func TestSignal(t *testing.T){

	sigs := make(chan os.Signal,1)
	done := make(chan bool,1)

	signal.Notify(sigs,syscall.SIGINT,syscall.SIGTERM)

	go func (){
		s := <-sigs
		println()
		fmt.Printf("%v",s)
		done <- true
	}()

	println("wait for signal")
	<-done
	println("exit")
}

func TestExit(t *testing.T){
	// defer out
	defer println("defer out")
	os.Exit(3)
	// echo $?  // 3
}