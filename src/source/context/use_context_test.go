package context

import (
	"testing"
	"context"
	"fmt"
	"time"
)

// context defines the Context type , which carries deadlines , cancelation singals , and other request-scoped values across
// API boundaries and between processes.
// go is so beautiful
func TestWithCancel(t *testing.T){
	// gen generates integers in a separate goroutine and sends them to the returned channeled.
	// The callers of gen need to cancel the context once they are done consuming generated integers not to leak the internal gorontine stared by gen
	gen := func(ctx context.Context) <- chan int{
		dst := make(chan int)
		n := 1
		go func(){
			for {
				select {
					// when ctx is done stop go routine . reclaim the resource.
					case <-ctx.Done():
						return // returning not to leak the goroutine
					case dst <- n :
						n++
				}
			}
		}()
		return dst
	}

	// creating a context with a cancel function to cancel the context .
	ctx , cancel := context.WithCancel(context.Background())
	// when the main goroutine is done cancel it .
	defer cancel()
	for c :=range gen(ctx) {
		fmt.Printf("%v", c)
		if c == 5 {
			break
		}
	}
}


// set a dead line for the context . // almost the same as WithTimeout
func TestDeadLine(t *testing.T){
	d := time.Now().Add(50 * time.Microsecond)
	ctx , cancel := context.WithDeadline(context.Background(),d)
	// add this line in case of the goroutine not canceled
	defer cancel()
	select {
		case <-time.After(1 * time.Second):
		fmt.Printf("sleeping here")
		case <- ctx.Done():
		fmt.Printf("done ")
	}
}

// get the contents of context . we should better make the key different type . incase of different package collapse problem
func TestWithValue(t *testing.T){
	type myStr string
	k := myStr("name")
	ctx := context.WithValue(context.Background(), k ,"clark")
	has := func(ctx context.Context ,  s myStr ){
		if v := ctx.Value(s) ; v != nil {
			println( s ," is found")
		}else{
			println( s , "can not found")
		}
	}
	has(ctx,k)
	has(ctx,myStr("age"))
}