package main

import (
    "fmt"
    "time"
)

func main() {
    // 创建两个通道
    c1 := make(chan string)
    c2 := make(chan string)

    // 启动两个 goroutine，分别在 1 秒和 2 秒后向通道发送消息
    go func() {
        time.Sleep(1 * time.Second)
        c1 <- "one"
    }()
    go func() {
        time.Sleep(2 * time.Second)
        c2 <- "two"
    }()

    // 使用 `select` 等待两个通道的消息
    // `select` 会阻塞，直到其中一个 case 可以运行
    for i := 0; i < 2; i++ {
        select {
        case msg1 := <-c1:
            fmt.Println("received", msg1)
        case msg2 := <-c2:
            fmt.Println("received", msg2)
        }
    }
}