package main

import (
	"fmt"
	"net/http"
	"time"
)

func checkAPI2(api string, ch chan string) {
	_, err := http.Get(api)
	if err != nil {
		ch <- fmt.Sprintf("ERROR: %s is down!\n", api)
		return
	}

	ch <- fmt.Sprintf("SUCCESS: %s is up and running!\n", api)
}

func main() {
	//无缓冲 channel
	apis := []string{
		"https://management.azure.com",
		"https://dev.azure.com",
		"https://api.github.com",
		"https://outlook.office.com/",
		"https://api.somewhereintheinternet.com/",
		"https://graph.microsoft.com",
	}


	start2 := time.Now()

	ch := make(chan string)

	for _, api := range apis {
		go checkAPI2(api, ch)
	}

	//fmt.Print(<-ch)
	for i := 0; i < len(apis); i++ {
		fmt.Print(<-ch)
	}
	elapsed2 := time.Since(start2)
	fmt.Printf("Done! It took %v seconds!\n", elapsed2.Seconds())
}
