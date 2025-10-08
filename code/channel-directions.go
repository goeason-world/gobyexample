package main

import "fmt"

// `ping` 函数只接受一个用于发送值的通道。
// `chan<- string` 表示这是一个只写通道。
// 尝试从这个通道接收值会导致编译错误。
func ping(pings chan<- string, msg string) {
    pings <- msg
}

// `pong` 函数接受一个用于接收值的通道 (`pings`) 和一个用于发送值的通道 (`pongs`)。
// `<-chan string` 表示这是一个只读通道。
func pong(pings <-chan string, pongs chan<- string) {
    msg := <-pings
    pongs <- msg
}

func main() {
    pings := make(chan string, 1)
    pongs := make(chan string, 1)

    // 发送消息到 pings 通道
    ping(pings, "passed message")

    // 从 pings 通道接收消息，并发送到 pongs 通道
    pong(pings, pongs)

    // 从 pongs 通道接收并打印消息
    fmt.Println(<-pongs)
}