package main

import "fmt"

// `mayPanic` 函数会引发一个 panic。
func mayPanic() {
    panic("a problem")
}

func main() {
    // `recover` 必须在 `defer` 函数中调用。
    // 当 `main` 函数发生 panic 时，`defer` 函数会被执行，
    // 其中的 `recover` 会捕获 `panic` 的值。
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered. Error:\n", r)
        }
    }()

    mayPanic()

    // 这行代码不会被执行，因为 `mayPanic` 发生 panic 后，
    // 执行会跳转到 `defer` 的函数，并在 `defer` 执行完毕后直接返回。
    fmt.Println("After mayPanic()")
}