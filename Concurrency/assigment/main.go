// package main

// func sumsq(start int, end int, c chan int) {
// 	var s int
// 	for i := start; i < end; i++ {
// 		s += i * i
// 	}
// 	c <- s
// }

// func main() {
// 	runtime.GOMAXPROCS(1)
// 	c1 := make(chan int)
// 	c2 := make(chan int)
// 	c3 := make(chan int)
// 	c4 := make(chan int)
// 	c5 := make(chan int)

// 	go sumsq(1, 200, c1)

// 	go sumsq(200, 400, c2)

// 	go sumsq(400, 600, c3)

// 	go sumsq(600, 800, c4)

// 	go sumsq(800, 1000, c5)

// 	sum := <-c1 + <-c2 + <-c3 + <-c4 + <-c5
// 	fmt.Println(sum)
// }

package main

import (
	"fmt"
	"sync"
)

func main() {
	num := make(chan int, 1000)
	var mx sync.Mutex
	var wg sync.WaitGroup
	var c int
	wg.Add(5)
	for i := 1; i <= 5; i++ {
		go func(num <-chan int, mx *sync.Mutex, i int) {
			for j := range num {
				fmt.Printf("The Value of %d is being processed from %d \n", j, i)
				mx.Lock()
				c += j * j
				mx.Unlock()
			}
			wg.Done()
		}(num, &mx, i)
	}

	for count := 0; count < 1000; count++ {
		num <- count
	}
	close(num)

	wg.Wait()
	fmt.Println(c)
}
