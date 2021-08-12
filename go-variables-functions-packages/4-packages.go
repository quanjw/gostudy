package main

import (
	"github.com/myuser/calculator"
	"rsc.io/quote"
)

func main() {

	//引用本地包（模块）
	total := calculator.Sum(3, 5)
	println(total)
	println("Version: ", calculator.Version)

	totalx := calculator.InternalSum(5)
	println(totalx)

	//引用外部（第三方）包
	println(quote.Hello())

}
