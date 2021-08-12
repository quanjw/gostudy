package main

import "fmt"

func main() {
	//声明和初始化切片
	months := []string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}
	quarter1 := months[0:3]
	quarter2 := months[3:6]
	quarter3 := months[6:9]
	quarter4 := months[9:12]

	fmt.Println(months)
	fmt.Println("Length:", len(months))
	fmt.Println("Capacity:", cap(months))

	fmt.Println(quarter1, len(quarter1), cap(quarter1))
	fmt.Println(quarter2, len(quarter2), cap(quarter2))
	fmt.Println(quarter3, len(quarter3), cap(quarter3))
	fmt.Println(quarter4, len(quarter4), cap(quarter4))

	//切片项
	quarter22 := months[3:6]
	quarter22Extended := quarter2[:4]
	fmt.Println(quarter2, len(quarter22), cap(quarter22))
	fmt.Println(quarter22Extended, len(quarter22Extended), cap(quarter22Extended))

	//追加项
	var numbers []int
	for i := 0; i < 10; i++ {
		numbers = append(numbers, i)
		fmt.Printf("%d\tcap=%d\t%v\n", i, cap(numbers), numbers)
	}

	//删除项
	letters := []string{"A", "B", "C", "D", "E"}
	remove := 2

	if remove < len(letters) {

		fmt.Println("Before", letters, "Remove ", letters[remove])

		letters = append(letters[:remove], letters[remove+1:]...)

		fmt.Println("After", letters)
	}

	//创建切片的副本
	letterss := []string{"A", "B", "C", "D", "E"}
	fmt.Println("Before", letterss)

	slice1 := letterss[0:2]
	slice2 := letterss[1:4]

	slice1[1] = "Z"

	fmt.Println("After", letterss)
	fmt.Println("Slice2", slice2)

	//--------------------------------------
	lettersss := []string{"A", "B", "C", "D", "E"}
	fmt.Println("Before", lettersss)

	slice11 := lettersss[0:2]
	fmt.Println(slice11, len(slice11), cap(slice11))


	slice22 := make([]string,3)
	copy(slice22, lettersss[1:4])

	slice11[1] = "Z"

	fmt.Println("After", lettersss)
	fmt.Println("Slice2", slice22)


}
