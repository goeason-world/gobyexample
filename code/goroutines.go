package main

import (
    "fmt"
    "time"
)

// `f` 是一个我们将要在 goroutine 中运行的函数。
func f(from string) {
    for i := 0; i < 3; i++ {
        fmt.Println(from, ":", i)
    }
}

func main() {
    // 这是一个普通的同步函数调用。
    f("direct")

    // 在函数调用前加上 `go` 关键字，就会把调用放到一个新的 goroutine 中并立即返回。
    go f("goroutine")

    // 你也可以对匿名函数使用 goroutine。
    go func(msg string) {
        fmt.Println(msg)
    }("going")

    // `time.Sleep` 用于等待 goroutine 执行完成。
    // 在真实项目中，应使用 `sync.WaitGroup` 或 channel 来同步。
    time.Sleep(time.Second)
    fmt.Println("done")
}