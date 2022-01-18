package main

import (
	"fmt"
	"sync"
	"time"
)

var count int

func generator(nums ...int) <-chan int {
	out := make(chan int)
	time.Sleep(time.Millisecond * 1000)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func square(in <-chan int, wg *sync.WaitGroup) {
	// out := make(chan int)s
	var mx sync.Mutex
	go func() {
		for n := range in {
			fmt.Printf("The Value of %d is being processed \n", n)
			mx.Lock()
			count += n * n
			mx.Unlock()
		}
		defer wg.Done()
	}()
}

func main() {
	in := generator(1, 2, 3, 4, 5, 6, 7, 8, 9)
	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 5; i++ {
		square(in, &wg)
	}
	wg.Wait()
	fmt.Println(count)
}
