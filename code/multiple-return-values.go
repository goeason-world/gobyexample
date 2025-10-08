package main

import "fmt"

// Go 内置支持多返回值功能。
// 函数签名中的 `(int, int)` 表示该函数返回两个 `int` 值。
func vals() (int, int) {
    return 3, 7
}

func main() {
    // 使用多重赋值从函数调用中接收多个返回值。
    a, b := vals()
    fmt.Println(a)
    fmt.Println(b)

    // 如果只需要部分返回值，可以使用空白标识符 `_` 忽略不需要的值。
    _, c := vals()
    fmt.Println(c)
}