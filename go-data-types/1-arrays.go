package main

import "fmt"

func main() {
	//声明数组
	var arr[3]int
	arr[1] = 10
	fmt.Println(arr[0])
	fmt.Println(arr[1])
	fmt.Println(arr[len(arr)-1])

	//初始化数组
	cities := [5]string{"New York", "Paris", "Berlin", "Madrid"}
	fmt.Println("Cities:", cities)

	//数组中的省略号
	q := [...]int{1, 2, 3}
	citiess := [...]string{"New York", "Paris", "Berlin", "Madrid"}
	fmt.Println("q:", q)
	fmt.Println("Cities:", citiess)

	numbers := [...]int{99: -1}
	fmt.Println("First Position:", numbers[0])
	fmt.Println("Last Position:", numbers[99])
	fmt.Println("Length:", len(numbers))

	//多维数组
	var twoD [3][5]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 5; j++ {
			twoD[i][j] = (i + 1) * (j + 1)
		}
		fmt.Println("Row", i, twoD[i])
	}
	fmt.Println("\nAll at once:", twoD)

	var threeD [3][5][2]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 5; j++ {
			for k := 0; k < 2; k++ {
				threeD[i][j][k] = (i + 1) * (j + 1) * (k + 1)
			}
		}
	}
	fmt.Println("\nAll at once:", threeD)

}
