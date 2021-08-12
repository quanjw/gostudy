package main

import "fmt"

func main() {
	ch := make(chan string)
	go senddata1(ch)
	getdata2(ch)
}

func senddata1(ch chan string) {
	ch <- "Washington"
	ch <- "Tripoli"
	ch <- "London"
	ch <- "Beijing"
	ch <- "Tokio"
	close(ch)
}

func getdata2(ch chan string)  {
	/*for{
		input,open :=<-ch
		if !open{
			break
		}
		fmt.Println(input)
	}*/

	for input := range ch {
		//process(input)
		fmt.Println(input)
	}
}