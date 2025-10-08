package main

import "fmt"

func main() {
    messages := make(chan string)
    signals := make(chan bool)

    // 非阻塞接收。
    // 如果 `messages` 通道有值，`select` 将会执行 `<-messages` 的 `case`。
    // 否则，它将立即执行 `default` 分支。
    select {
    case msg := <-messages:
        fmt.Println("received message", msg)
    default:
        fmt.Println("no message received")
    }

    // 非阻塞发送。
    // 如果 `messages` 通道有接收者，`select` 将会执行 `messages <- msg` 的 `case`。
    // 否则，它将立即执行 `default` 分支。
    msg := "hi"
    select {
    case messages <- msg:
        fmt.Println("sent message", msg)
    default:
        fmt.Println("no message sent")
    }

    // 我们可以在 `default` 分支之前使用多个 `case` 子句来实现多路非阻塞选择。
    select {
    case msg := <-messages:
        fmt.Println("received message", msg)
    case sig := <-signals:
        fmt.Println("received signal", sig)
    default:
        fmt.Println("no activity")
    }
}