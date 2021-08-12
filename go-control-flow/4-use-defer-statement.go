package main

import (
	"fmt"
	"io"
	"os"
)

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic("Panic in g() (major)")
	}
	defer fmt.Println("Defer in g()", i)
	fmt.Println("Printing in g()", i)
	g(i + 1)
}

func main() {
	//defer 函数
	for i := 1; i <= 4; i++ {
		//defer fmt.Println("deferred", -i)
		fmt.Println("regular", i)
	}

	f, err := os.Create("notes.txt")
	if err != nil {
		return
	}
	defer f.Close()

	if _, err = io.WriteString(f, "Learning Go!"); err != nil {
		return
	}

	_ = f.Sync()

	//panic 函数
	//g(0)
	//fmt.Println("Program finished successfully!")


	//recover 函数
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()
	g(0)
	fmt.Println("Program finished successfully!")

}
