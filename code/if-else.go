package main

import "fmt"

func main() {

    // 这是一个基本的 `if-else` 语句。
    // 条件表达式不需要用小括号 `()` 包围，但代码块必须用大括号 `{}` 包围。
    if 7%2 == 0 {
        fmt.Println("7 is even")
    } else {
        fmt.Println("7 is odd")
    }

    // 你可以只使用 `if` 语句，而没有 `else`。
    if 8%4 == 0 {
        fmt.Println("8 is divisible by 4")
    }

    // 在条件语句之前可以有一个初始化语句。
    // 在这个语句中声明的变量 `num` 在所有分支中都可用。
    if num := 9; num < 0 {
        fmt.Println(num, "is negative")
    } else if num < 10 {
        fmt.Println(num, "has 1 digit")
    } else {
        fmt.Println(num, "has multiple digits")
    }

    // 注意：Go 中没有三元运算符 `? :`，你需要使用完整的 `if-else` 语句。
}