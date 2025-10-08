package main

import (
    "fmt"
    "time"
)

func main() {

    // 这是一个基本的 `switch` 语句。
    // Go 的 `switch` 不需要 `break`，匹配成功后会自动退出。
    i := 2
    fmt.Print("Write ", i, " as ")
    switch i {
    case 1:
        fmt.Println("one")
    case 2:
        fmt.Println("two")
    case 3:
        fmt.Println("three")
    }

    // 你可以在一个 `case` 中使用逗号分隔多个表达式。
    // `default` 分支是可选的。
    switch time.Now().Weekday() {
    case time.Saturday, time.Sunday:
        fmt.Println("It's the weekend")
    default:
        fmt.Println("It's a weekday")
    }

    // 不带表达式的 `switch` 是实现 `if/else` 逻辑的另一种方式。
    // `case` 表达式可以是任意布尔表达式。
    t := time.Now()
    switch {
    case t.Hour() < 12:
        fmt.Println("It's before noon")
    default:
        fmt.Println("It's after noon")
    }

    // 类型开关（Type Switch）用于判断一个接口变量的实际类型。
    whatAmI := func(i interface{}) {
        switch t := i.(type) {
        case bool:
            fmt.Println("I'm a bool")
        case int:
            fmt.Println("I'm an int")
        default:
            // `%T` 格式化占位符用于打印变量的类型。
            fmt.Printf("Don't know type %T\n", t)
        }
    }
    whatAmI(true)
    whatAmI(1)
    whatAmI("hey")
}