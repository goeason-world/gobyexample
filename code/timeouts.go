package main

import (
    "fmt"
    "time"
)

func main() {
    // 创建一个带缓冲的通道，以避免 goroutine 泄漏。
    c1 := make(chan string, 1)
    go func() {
        // 模拟一个耗时 2 秒的操作。
        time.Sleep(2 * time.Second)
        c1 <- "result 1"
    }()

    // 使用 `select` 实现超时控制。
    // `time.After` 会在指定时间后向其返回的通道发送一个值。
    select {
    case res := <-c1:
        fmt.Println(res)
    case <-time.After(1 * time.Second): // 1 秒超时
        fmt.Println("timeout 1")
    }

    // 第二个例子，超时时间为 3 秒。
    c2 := make(chan string, 1)
    go func() {
        time.Sleep(2 * time.Second)
        c2 <- "result 2"
    }()

    select {
    case res := <-c2:
        fmt.Println(res)
    case <-time.After(3 * time.Second): // 3 秒超时
        fmt.Println("timeout 2")
    }
}