package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//基本 for 循环语法
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println("sum of 1..100 is", sum)

	///空的预处理语句和后处理语句
	var num int64
	rand.Seed(time.Now().Unix())
	fmt.Println(num)
	for num != 5 {
		num = rand.Int63n(10)
		fmt.Println(num)
	}

	//无限循环和 break 语句
	var numx int32
	sec := time.Now().Unix()
	rand.Seed(sec)

	for {
		fmt.Print("Writting inside the loop...")
		if numx = rand.Int31n(10); numx == 5 {
			fmt.Println("finish!")
			break
		}
		fmt.Println(numx)
	}

	//Continue 语句
	sumxx := 0
	for num := 1; num <= 100; num++ {
		if num%5 == 0 {
			continue
		}
		sumxx += num
	}
	fmt.Println("The sum of 1 to 100, but excluding numbers divisible by 5, is", sumxx)

}
