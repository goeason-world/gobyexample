package main

import "fmt"

func main() {

    // `var` 关键字用于声明一个变量。
    // Go 会自动推断已初始化变量的类型。
    var a = "initial"
    fmt.Println(a)

    // 你可以一次性声明多个变量。
    // 这里我们显式指定了变量的类型为 `int`。
    var b, c int = 1, 2
    fmt.Println(b, c)

    // Go 会自动推断出这是一个布尔类型（bool）。
    var d = true
    fmt.Println(d)

    // 未经初始化的变量会被赋予其类型的“零值”。
    // 对于 `int` 类型，零值是 `0`。
    var e int
    fmt.Println(e)

    // `:=` 语法是声明并初始化变量的简写形式。
    // 它等价于 `var f string = "apple"`。
    // `:=` 语法只能在函数内部使用。
    f := "apple"
    fmt.Println(f)
}