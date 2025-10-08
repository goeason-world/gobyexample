package main

import (
    "fmt"
    "os"
)

// `defer` 语句用于安排一个函数调用，使其在执行 `defer` 的函数即将返回之前执行。
// 这个特性常用于执行清理操作，例如关闭文件。

func main() {
    // 创建文件
    f := createFile("/tmp/defer.txt")
    // 使用 `defer` 安排文件关闭操作。
    // 这将在 `main` 函数的末尾执行。
    defer closeFile(f)
    // 向文件写入数据。
    writeFile(f)
}

// `createFile` 创建一个文件并返回。
func createFile(p string) *os.File {
    fmt.Println("creating")
    f, err := os.Create(p)
    if err != nil {
        panic(err)
    }
    return f
}

// `writeFile` 向文件中写入数据。
func writeFile(f *os.File) {
    fmt.Println("writing")
    fmt.Fprintln(f, "data")
}

// `closeFile` 关闭文件，这个函数将被延迟执行。
func closeFile(f *os.File) {
    fmt.Println("closing")
    // 在关闭文件时检查并处理错误。
    err := f.Close()
    if err != nil {
        fmt.Fprintf(os.Stderr, "error: %v\n", err)
        os.Exit(1)
    }
}