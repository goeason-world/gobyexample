package main

import "fmt"

// 这是一个接收两个 `int` 参数并返回它们的和的函数。
func plus(a int, b int) int {
    // Go 要求显式使用 `return` 语句返回值，不会自动返回最后一个表达式的值。
    return a + b
}

// 当有多个连续的相同类型参数时，可以省略前面参数的类型声明，
// 只在最后一个参数处声明类型。
func plusPlus(a, b, c int) int {
    return a + b + c
}

func main() {
    // 使用 `函数名(参数)` 的方式调用函数。
    res := plus(1, 2)
    fmt.Println("1+2 =", res)

    res = plusPlus(1, 2, 3)
    fmt.Println("1+2+3 =", res)
}