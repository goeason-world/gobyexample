package main

import "fmt"

func main() {
    // 默认情况下，通道是无缓冲的，这意味着只有在有对应的接收者准备好接收时，
    // 发送操作才会成功。带缓冲的通道可以接受有限数量的值，而无需相应的接收者。

    // `make` 函数的第二个参数指定了通道的缓冲区大小。
    messages := make(chan string, 2)

    // 由于这个通道是带缓冲的，我们可以将这些值发送到通道中，而无需并发的接收操作。
    messages <- "buffered"
    messages <- "channel"

    // 稍后我们可以像往常一样接收这两个值。
    fmt.Println(<-messages)
    fmt.Println(<-messages)
}