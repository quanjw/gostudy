package main

import (
	"fmt"
	"math/rand"
	"regexp"
	"time"
)

func location(city string) (region string, continent string) {
	switch city {
	case "Delhi", "Hyderabad", "Mumbai", "Chennai", "Kochi":
		region, continent = "India", "Asia"
	case "Lafayette", "Louisville", "Boulder":
		region, continent = "Colorado", "USA"
	case "Irvine", "Los Angeles", "San Diego":
		region, continent = "California", "USA"
	default:
		region, continent = "Unknown", "Unknown"
	}
	return
}

func main() {
	//基本 switch 语法
	sec := time.Now().Unix()
	rand.Seed(sec)
	i := rand.Int31n(10)
	fmt.Println(i)
	switch i {
	case 0:
		fmt.Print("zero...")
	case 1:
		fmt.Print("one...")
	case 2:
		fmt.Print("two...")
	}

	fmt.Println("ok")

	//使用多个表达式
	region, continent := location("Irvine")
	fmt.Printf("John works in %s, %s\n", region, continent)

	//调用函数
	switch time.Now().Weekday().String() {
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Println("It's time to learn some Go.")
	default:
		fmt.Println("It's weekend, time to rest!")
	}

	fmt.Println(time.Now().Weekday().String())

	//switch 块没有表达式
	var email = regexp.MustCompile(`^[^@]+@[^@.]+\.[^@.]+`)
	var phone = regexp.MustCompile(`^[(]?[0-9][0-9][0-9][). \-]*[0-9][0-9][0-9][.\-]?[0-9][0-9][0-9][0-9]`)

	contact := "1341513456"

	switch {
	case email.MatchString(contact):
		fmt.Println(contact, "is an email")
	case phone.MatchString(contact):
		fmt.Println(contact, "is a phone number")
	default:
		fmt.Println(contact, "is not recognized")
	}

	//省略条件
	rand.Seed(time.Now().Unix())
	r := rand.Float32()
	switch {
	case r > 0.1:
		fmt.Println("Common case, 90% of the time")
	default:
		fmt.Println("10% of the time")
	}

	//使逻辑进入到下一个 case
	switch num := 15; {
	case num < 50:
		fmt.Printf("%d is less than 50\n", num)
		fallthrough
	case num > 100:
		fmt.Printf("%d is greater than 100\n", num)
		//fallthrough
	case num < 200:
		fmt.Printf("%d is less than 200", num)
	}
}
