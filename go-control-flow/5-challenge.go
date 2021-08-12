package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	//编写 FizzBuzz 程序
	if false {
		for i:=100;i>1;i-- {
			var number = FizzBuzz()
			print(number)
		}
	}


	//推测平方根
	//sqroot n+1 = sqroot n − (sqroot n * sqroot n − x) / (2 * sqroot n)
	if false {
		num := sqroot(5235)
		fmt.Printf(" square root is %f \n", num)
	}
	//要求用户输入一个数字，如果该数字为负数，则进入紧急状态
	val := 0
	for  {
		fmt.Print("Enter number: ")
		fmt.Scanf("%d", &val)
		fmt.Println("You entered:", val)
		switch  {
		case val ==  0:
			fmt.Println("0 is neither negative nor positive")
		case val > 0 :
			fmt.Printf("You entered: %d",val)
		case val < 0:
			panic("")
		}
	}

}

func FizzBuzz() ( number int)  {
	rand.Seed(time.Now().Unix())
	number = int(rand.Int31n(100))
	switch  {
	case number%3 ==0  && number%5 == 0:
		fmt.Println("Fizz")
	case number%3 == 0:
		fmt.Println("Fizz")
	case number%5 == 0:
		fmt.Println("Buzz")
	default:
		fmt.Println(number)
	}
	return
}

func sqroot(number float64) float64   {
	current := 1.0
	previous := 0.0
	for i:=0;i<10;i++ {
		previous = current
		current = previous-(previous*previous-number)/(2*previous)
		if current == previous {
			break
		}
		fmt.Printf("A guess for square root is %f \n",current)
	}
	return  current
}

