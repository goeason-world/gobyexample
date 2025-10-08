package main

import "fmt"

func main() {
    // 通道（channel）是 goroutine 之间通信的管道。
    // 使用 `make(chan val-type)` 创建一个新的通道。
    messages := make(chan string)

    // 使用 `channel <-` 语法将一个值发送到通道中。
    // 这里我们在一个新的 goroutine 中发送 "ping" 到 `messages` 通道。
    go func() { messages <- "ping" }()

    // 使用 `<-channel` 语法从通道中接收一个值。
    // 默认情况下，发送和接收都会阻塞，直到双方都准备好。
    // 这个特性使得我们可以在没有其他同步机制的情况下等待 goroutine 的完成。
    msg := <-messages
    fmt.Println(msg)
}