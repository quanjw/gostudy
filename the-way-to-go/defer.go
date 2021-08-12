package main

import (
	"fmt"
	"log"
	//"strings"
)

var where = log.Print

func main() {

	//a()
	//f()



	b()
	//strings.IndexFunc()

}

//关键字 defer 的用法类似于面向对象编程语言 Java 和 C# 的 finally 语句块，它一般用于释放某些已分配的资源。
/*func a() {
	i := 0
	defer fmt.Println(i)
	i++
	return
}*/

//当有多个 defer 行为被注册时，它们会以逆序执行（类似栈，即后进先出）：
func f() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}
}

func trace(s string) string {
	fmt.Println("entering:", s)
	return s
}

func un(s string) {
	fmt.Println("leaving:", s)
}

func a() {
	defer un(trace("a"))
	fmt.Println("in a")
}

func b() {
	defer un(trace("b"))
	where()
	fmt.Println("in b")
	where()
	a()
	where()
}

