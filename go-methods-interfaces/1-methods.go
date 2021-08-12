package main

import (
	"fmt"
	"strings"
)

type triangle struct {
	size int
}

type coloredTriangle struct {
	triangle
	color string
}

type square struct {
	size int
}

type upperstring string

func (s upperstring) Upper() string {
	return strings.ToUpper(string(s))
}


func (t triangle) perimeter() int {
	return t.size * 3
}
func (s square) perimeter() int {
	return s.size * 4
}

func (t coloredTriangle) perimeter() int {
	return t.size * 3 * 2
}

func (t *triangle) doubleSize() {
	t.size *= 2
}

func main() {
	//声明方法
	t := triangle{3}
	s := square{4}
	fmt.Println("Perimeter (triangle):", t.perimeter())
	fmt.Println("Perimeter (square):", s.perimeter())

	//方法中的指针\\
	//pptr := &t
	(&t).doubleSize()
	t.doubleSize()
	fmt.Println("Size:", t.size)
	fmt.Println("Perimeter:", t.perimeter())

	//声明其他类型的方法
	ss := upperstring("Learning Go!")
	fmt.Println(ss)
	fmt.Println(ss.Upper())


	//嵌入方法
	tt := coloredTriangle{triangle{3}, "blue"}
	fmt.Println("Size:", tt.size)
	fmt.Println("Perimeter", tt.perimeter())
	fmt.Println("Perimeter (colored)", tt.perimeter())
	fmt.Println("Perimeter (normal)", tt.triangle.perimeter())

	//方法中的封装
	//在 Go 中，只需使用大写标识符，即可公开方法，使用非大写的标识符将方法设为私有方法。
}
