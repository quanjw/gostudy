package main

import "fmt"

func givemeanumber() int {
	return -1
}
func somenumber() int {
	return -7
}

func main() {
	//if 语句的语法
	x := 27
	if x%2 == 1 {
		fmt.Println(x, "is even")
	}
	//复合 if 语句
	if num := givemeanumber(); num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has only one digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

	if numx := somenumber(); numx < 0 {
		fmt.Println(numx, "is negative")
	} else if numx < 10 {
		fmt.Println(numx, "has 1 digit")
	} else {
		fmt.Println(numx, "has multiple digits")
	}
	//fmt.Println(numx)
}
