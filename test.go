package main

import "fmt"
func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
			case c <- x:
				x, y = y, x+y
				fmt.Println("first layer")

			case z:= <-c:
				x, y = y, x+y
				fmt.Println("second layer")
				fmt.Println("z: ",z)

			case <-quit:
				fmt.Println("quit")
				return
}
	}
}
func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("shuchu",<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}