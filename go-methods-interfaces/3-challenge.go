package main

import (
	"errors"
	"fmt"
)

type Account  struct {
	FirstName  string
	LastName   string
}

func (a *Account) ChangeName( s string)  {
	a.FirstName = s
}

type Employee  struct {
	Account
	Credits  float64
}

func (e Employee) String() string {
	return fmt.Sprintf("Name: %s %s\nCredits: %.2f\n", e.FirstName, e.LastName, e.Credits)
}

func CreateEmployee(firstName string , lastName string, credits float64) (*Employee, error) {
	return &Employee{Account{firstName, lastName}, credits}, nil
}

func (e *Employee) AddCredits(amount float64) (float64, error) {
	if amount > 0.0 {
		e.Credits += amount
		return e.Credits, nil
	}
	return 0.0, errors.New("Invalid credit amount.")
}

func (e *Employee) RemoveCredits(amount float64) (float64, error) {
	if amount > 0.0 {
		if amount <= e.Credits {
			e.Credits -= amount
			return e.Credits, nil
		}
		return 0.0, errors.New("You can't remove more credits than the account has.")
	}
	return 0.0, errors.New("You can't remove negative numbers.")
}

func (e *Employee) CheckCredits() float64 {
	return e.Credits
}

func main() {
	//1.创建一个名为 Account 的自定义类型，此类型包含帐户所有者的名字和姓氏。 此类型还必须加入 ChangeName 的功能
	quanjw := Account{"cuan","jw"}
	quanjw.ChangeName("quan")
	fmt.Println(quanjw)

	//2.创建另一个名为 Employee 的自定义类型，此类型包含用于将贷方数额存储为类型 float64 并嵌入 Account 对象的变量。
	//类型还必须包含 AddCredits、RemoveCredits 和 CheckCredits 的功能。 你需要展示你可以通过 Employee 对象更改帐户名称。
	bruce, _ := CreateEmployee("Bruce", "Lee", 500)
	fmt.Println(bruce.CheckCredits())
	credits, err := bruce.AddCredits(250)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("New Credits Balance = ", credits)
	}

	_, err = bruce.RemoveCredits(2500)
	if err != nil {
		fmt.Println("Can't withdraw or overdrawn!", err)
	}

	bruce.ChangeName("Mark")

	//3.将字符串方法写入 Account 对象，以便按包含名字和姓氏的格式打印 Employee 名称。
	fmt.Println(bruce)

}
