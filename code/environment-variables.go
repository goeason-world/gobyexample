package main

import (
    "fmt"
    "os"
    "strings"
)

func main() {
    // `os.Setenv` 用于设置一个环境变量。
    os.Setenv("FOO", "1")
    // `os.Getenv` 用于获取一个环境变量的值。
    fmt.Println("FOO:", os.Getenv("FOO"))
    // 如果环境变量未被设置，`os.Getenv` 会返回一个空字符串。
    fmt.Println("BAR:", os.Getenv("BAR"))

    fmt.Println()
    // `os.Environ` 返回一个包含了所有环境变量的字符串切片。
    for _, e := range os.Environ() {
        pair := strings.SplitN(e, "=", 2)
        fmt.Println(pair[0])
    }
}
