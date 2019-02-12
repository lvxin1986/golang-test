package main

import (
	"fmt"
	"time"
)

func main() {
	var c1, c2, c3 chan int
	var i1, i2 int
	c1 = make(chan int, 1)
	c2 = make(chan int, 1)
	c3 = make(chan int, 1)
	go func() {
		num := 1
		for {
			fmt.Println(num)
			time.Sleep(1e9)
			num++
			if num%3== 0 {
				c1 <- 1
			} else if num%3== 1{
				c2 <- 2
			} else {
				c3 <- 3
			}
		}
	}()
	for {
		select {
		case i1 = <-c1:
			fmt.Printf("received ", i1, " from c1\n")
		case i2 = <-c2:
			fmt.Printf("received ", i2, "  from  c2\n")
		case i3, ok := (<-c3): // same as: i3, ok := <-c3
			if ok {
				fmt.Printf("received ", i3, " from c3\n")
			} else {
				fmt.Printf("c3 is closed\n")
			}
		//default:
		//	fmt.Printf("no communication\n")
		}
	}
}