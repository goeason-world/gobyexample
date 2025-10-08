package main

import (
    "fmt"
    "os"
)

func main() {
    // `os.Args` 是一个字符串切片，它包含了程序启动时的所有命令行参数。
    argsWithProg := os.Args
    // `os.Args[1:]` 只包含传递给程序的实际参数。
    argsWithoutProg := os.Args[1:]

    // 你可以通过索引访问单个参数。
    // 注意：如果索引超出范围，程序会 panic。
    if len(os.Args) > 3 {
        arg := os.Args[3]
        fmt.Println(arg)
    }

    fmt.Println(argsWithProg)
    fmt.Println(argsWithoutProg)
}
