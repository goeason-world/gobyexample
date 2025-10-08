package main

import (
    "fmt"
    "math"
)

// `const` 关键字用于声明一个常量。
// 这里我们声明一个字符串常量 `s`。
const s string = "constant"

func main() {
    fmt.Println(s)

    // `const` 可以在任何 `var` 语句可以出现的地方出现。
    // 数值常量没有类型，直到被赋予一个明确的类型，例如通过显式转换。
    const n = 500000000

    // 常量表达式支持任意精度的算术运算。
    // 这里的计算在编译时完成。
    const d = 3e20 / n
    fmt.Println(d)

    // 当需要时，数值常量可以被赋予一个类型。
    // 这里我们将 `d` 显式转换为 `int64` 类型。
    fmt.Println(int64(d))

    // 当在需要特定类型的上下文中使用时，数值常量会自动转换类型。
    // `math.Sin` 函数需要一个 `float64` 类型的参数，所以 `n` 会被自动转换为 `float64`。
    fmt.Println(math.Sin(n))
}