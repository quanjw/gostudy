package main

import "fmt"

func main() {
	//声明和初始化映射
	studentsAge := map[string]int{
		"john": 32,
		"bob":  31,
	}
	fmt.Println(studentsAge)

	studentsAge2:= make(map[string]int)
	fmt.Println(studentsAge2)

	//添加项
	studentsAge["john"] = 32
	studentsAge["bob"] = 31
	fmt.Println(studentsAge)

	//访问项
	fmt.Println("Bob's age is", studentsAge["bob"])

	age, exist := studentsAge["christy"]
	if exist {
		fmt.Println("Christy's age is", age)
	} else {
		fmt.Println("Christy's age couldn't be found")
	}

	//删除项
	delete(studentsAge, "christy")
	fmt.Println(studentsAge)

	//映射中的循环
	for name, age := range studentsAge {
		fmt.Printf("%s\t%d\n", name, age)
	}
	for _, age := range studentsAge {
		fmt.Printf("Ages %d\n", age)
	}
	for name := range studentsAge {
		fmt.Printf("Names %s\n", name)
	}
}
