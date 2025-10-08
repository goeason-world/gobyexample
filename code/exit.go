package main

import (
    "fmt"
    "os"
)

func main() {
    // `defer` 语句注册的函数在 `os.Exit` 调用时不会被执行。
    defer fmt.Println("!")

    // `os.Exit` 会立即以给定的状态码终止当前程序。
    os.Exit(3)
}
