package main

import "fmt"
import "rsc.io/quote"

func main() {
    //fmt.Println("Hello, World!")
    fmt.Println(quote.Go())

    var j = 4
    i := 3
    println(i)
    print(j)

    //常量
    const cons = 5
    print(cons)

    const (
        a = iota
        b = iota
        c = iota
    )

    print(a)
    print(b)
    print(c)
    //第一个 iota 等于 0，每当 iota 在新的一行被使用时，它的值都会自动加 1，并且没有赋值的常量默认会应用上一行的赋值表达式：



    //变量
    //var identifier [type] = value
    /*var a int = 15
    var i = 5
    var b bool = false
    var str string = "Go says hello to the world!"

    var a = 15
    var b = false
    var str = "Go says hello to the world!"*/


    //指针
    var i1 = 5
    fmt.Printf("An integer: %d, it's location in memory: %p\n", i1, &i1)

    var intP *int
    intP = &i1
    fmt.Printf("it's location in memory: %p, An intege: %d \n", intP, *intP)


}

