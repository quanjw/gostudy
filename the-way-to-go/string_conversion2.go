package main

import (
	"fmt"
	"strconv"
)

func main() {
	var orig string = "ABC"
	// var an int
	var newS string
	// var err error

	fmt.Printf("The size of ints is: %d\n", strconv.IntSize)
	// anInt, err = strconv.Atoi(origStr)
	an, err := strconv.Atoi(orig)
	fmt.Print(an)
	fmt.Print(err)
	if err != nil {
		fmt.Printf("orig %s is not an integer - exiting with error\n", orig)
		return
	}
	fmt.Printf("The integer is %d\n", an)
	an = an + 5
	newS = strconv.Itoa(an)
	fmt.Printf("The new string is: %s\n", newS)


	/*if err := file.Chmod(0664); err != nil {
		fmt.Println(err)
		return err
	}*/


	/*if value, ok := readData(); ok {
		…
	}*/
	switch 0 {
	case 0:
		fallthrough
	case 1:
		main()

	}
}