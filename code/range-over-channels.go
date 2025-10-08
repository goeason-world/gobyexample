package main

import "fmt"

func main() {
    // 创建一个带缓冲的通道。
    queue := make(chan string, 2)
    queue <- "one"
    queue <- "two"
    // 关闭通道。这告诉接收者不会再有新的值被发送。
    close(queue)

    // `for range` 可以用于遍历通道中的值。
    // 循环会自动在通道被关闭且没有更多值可接收时终止。
    for elem := range queue {
        fmt.Println(elem)
    }
}