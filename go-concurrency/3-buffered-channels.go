package main

import (
	"fmt"
	"net/http"
	"time"
)

func send(ch chan string, message string) {
	ch <- message
}
//----------------------------

func checkAPI3(api string, ch chan string) {
	_, err := http.Get(api)
	if err != nil {
		ch <- fmt.Sprintf("ERROR: %s is down!\n", api)
		return
	}

	ch <- fmt.Sprintf("SUCCESS: %s is up and running!\n", api)
}

func send2(ch chan<- string, message string) {
	fmt.Printf("Sending: %#v\n", message)
	ch <- message
}

func read2(ch <-chan string) {
	fmt.Printf("Receiving: %#v\n", <-ch)
}
//-----------------------
func process(ch chan string) {
	time.Sleep(3 * time.Second)
	ch <- "Done processing!"
}

func replicate(ch chan string) {
	time.Sleep(1 * time.Second)
	ch <- "Done replicating!"
}

func main() {
	/*size := 4
	ch := make(chan string, size)
	send(ch, "one")
	send(ch, "two")
	go send(ch, "three")
	go send(ch, "four")
	fmt.Println("All data sent to the channel ...")

	for i := 0; i < size; i++ {
		fmt.Println(<-ch)
	}

	fmt.Println("Done!")*/
	//----------------------------------------
	//有缓冲 channel
	start := time.Now()

	apis := []string{
		"https://management.azure.com",
		"https://dev.azure.com",
		"https://api.github.com",
		"https://outlook.office.com/",
		"https://api.somewhereintheinternet.com/",
		"https://graph.microsoft.com",
	}

	ch := make(chan string, 10)

	for _, api := range apis {
		go checkAPI3(api, ch)
	}

	for i := 0; i < len(apis); i++ {
		fmt.Print(<-ch)
	}

	elapsed := time.Since(start)
	fmt.Printf("Done! It took %v seconds!\n", elapsed.Seconds())

	//Channel 方向
	ch2 := make(chan string, 1)
	send2(ch2, "Hello World!")
	read2(ch2)

	//多路复用
	ch3 := make(chan string)
	ch4 := make(chan string)
	go process(ch3)
	go replicate(ch4)

	for i := 0; i < 2; i++ {
		select {
		case process := <-ch3:
			fmt.Println(process)
		case replicate := <-ch4:
			fmt.Println(replicate)
		}
	}
}
