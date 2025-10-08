package main

import "os"

func main() {
    // `panic` 是 Go 语言中一个内置函数，用于产生一个意外的、无法正常处理的错误。
    // 当程序遇到一个无法恢复的错误时，可以调用 `panic` 来中断程序的正常执行流程。
    panic("a problem")

    // `panic` 的一个常见用法是，如果一个函数返回了一个我们不知道如何处理（或不想处理）的错误值，就中止程序。
    // 这是一个在创建文件时发生意外错误的例子。
    _, err := os.Create("/tmp/file")
    if err != nil {
        panic(err)
    }
}