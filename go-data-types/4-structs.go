package main

import (
	"encoding/json"
	"fmt"
)

type Employee struct {
	ID        int
	FirstName string `json:"name"`
	LastName  string
	Address   string `json:"address,omitempty"`
}

type Manager struct {
	Employee
	ManagerID   int
}

func main() {
	//声明和初始化结构
	var john Employee
	fmt.Println(john)
	employee := Employee{1001, "John", "Doe", "Doe's Street"}
	fmt.Println(employee)
	employee2 := Employee{LastName: "Doe", FirstName: "John"}
	fmt.Println(employee2)

	employeeCopy := &employee
	employeeCopy.FirstName = "David"
	fmt.Println(employee)

	//结构嵌入
	manager := Manager{
		Employee: Employee{
			FirstName: "John",
		},
	}
	manager.LastName = "Doe"
	fmt.Println(manager.FirstName)

	//用 JSON 编码和解码结构
	data, _ := json.Marshal(manager)
	fmt.Printf("%s\n", data)

	var decoded []Manager
	json.Unmarshal(data, &decoded)
	fmt.Println( decoded)

}
