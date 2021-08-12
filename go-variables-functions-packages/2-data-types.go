package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {

	//整数数字
	var integer8 int8 = 127
	var integer16 int16 = 32767
	var integer32 int32 = 2147483647
	var integer64 int64 = 9223372036854775807
	fmt.Println(integer8, integer16, integer32, integer64)

	rune := 'G'
	fmt.Println(rune)
	fmt.Printf("%c \n",rune)

	var int32x int64= 2147483648
	fmt.Println(int32x)

	var integer int = -10
	fmt.Println(integer)

	//浮点数字
	var float float32 = 214
	var float32x float32 = 2147483647
	var float64x float64 = 9223372036854775807
	println(float,float32x, float64x)

	println(math.MaxFloat32, math.MaxFloat64)

	const e = 2.71828
	const Avogadro = 6.02214129e23
	const Planck = 6.62606957e-34

	//布尔型
	var featureFlag bool = true
	println(featureFlag)

	//字符串
	var firstName string = "John"
	lastName := "Doe"
	println(firstName, lastName)

	fullName := "John Doe \t(alias \"Foo\")\n"
	println(fullName)

	//默认值
	var defaultInt int
	var defaultFloat32 float32 = 111
	var defaultFloat64 float64
	var defaultBool bool
	var defaultString string
	println(defaultInt, defaultBool, defaultFloat32, defaultFloat64, defaultString)

	//类型转换
	var integer16xx int16 = 127
	var integer32xx int32 = 32767
	println(int32(integer16xx) + integer32xx)

	i, _ := strconv.Atoi("-42")
	s := strconv.Itoa(-42)
	println(i, s)
}