package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	go senddata(ch)
	time.Sleep(2e9)

	go getdata(ch)

	time.Sleep(1e9)

}

func senddata(ch chan string)  {
	ch <- "123"
	ch <- "3453"
	ch <- "523"
	ch <- "4562"
	ch <- "4572"
	ch <- "13434"
	ch <- "634"
}

func getdata(ch chan string )  {
	var input string
	for  {
		input = <- ch
		fmt.Println(input)
	}
}