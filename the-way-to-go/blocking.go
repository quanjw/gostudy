package main

import (
	"fmt"
	"time"
)

func f11(in chan int)  {
	fmt.Println(<-in)
}

func main() {
	//out := make(chan int)
	out := make(chan int,100)
	out <- 2
	go f11(out)
	//out <- 2
	time.Sleep(1e9)
}




