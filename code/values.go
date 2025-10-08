package main

import "fmt"

func main() {

    // Go 支持多种基本数据类型，包括字符串、整数、浮点数、布尔值等。

    // 字符串可以通过 `+` 运算符进行拼接。
    fmt.Println("go" + "lang")

    // 整数支持常见的算术运算。
    fmt.Println("1+1 =", 1+1)

    // 浮点数支持精确的小数运算。
    // 注意：整数除法 `7/3` 的结果是 `2`，会舍弃小数部分。
    fmt.Println("7.0/3.0 =", 7.0/3.0)

    // 布尔值支持逻辑运算。
    // `&&` 逻辑与：两者都为真时才为真。
    fmt.Println(true && false)
    // `||` 逻辑或：至少一个为真时就为真。
    fmt.Println(true || false)
    // `!` 逻辑非：取反。
    fmt.Println(!true)
}