package main

import (
	"os"
	"strconv"
)

func main() {
	//number1, _ := strconv.Atoi(os.Args[1])
	//number2, _ := strconv.Atoi(os.Args[2])
	//println("Sum:", number1+number2)
	sum := sum(os.Args[1], os.Args[2])
	println("Sum:", sum)

	sum, mul := calc(os.Args[1], os.Args[2])
	println("Sum:", sum)
	println("Mul:", mul)

	sums, _ := calc(os.Args[1], os.Args[2])
	println("Sum:", sums)

	//更改函数参数值（指针）
	firstName := "John"
	updateName(&firstName)
	println(firstName)
}

/*func sum(number1 string, number2 string) int {
	int1, _ := strconv.Atoi(number1)
	int2, _ := strconv.Atoi(number2)
	return int1 + int2
}*/

func sum(number1 string, number2 string) (result int) {
	int1, _ := strconv.Atoi(number1)
	int2, _ := strconv.Atoi(number2)
	result = int1 + int2
	return
}

//返回多个值
func calc(number1 string, number2 string) (sum int, mul int) {
	int1, _ := strconv.Atoi(number1)
	int2, _ := strconv.Atoi(number2)
	sum = int1 + int2
	mul = int1 * int2
	return
}

func updateName(name *string) {
	*name = "David"
}