package main

import (
    "fmt"
    "os"
    "os/signal"
    "syscall"
)

func main() {
    // 创建一个用于接收信号的 channel。
    sigs := make(chan os.Signal, 1)

    // `signal.Notify` 将指定的 channel 注册为给定信号的接收者。
    signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

    // 创建一个用于程序同步的 channel。
    done := make(chan bool, 1)

    // 启动一个 goroutine 来等待并处理信号。
    go func() {
        // 程序会在这里阻塞，直到接收到一个在上面注册过的信号。
        sig := <-sigs
        fmt.Println()
        fmt.Println(sig)

        // 当信号处理完成后，向 `done` channel 发送一个值。
        done <- true
    }()

    // 主程序继续执行，并等待完成信号。
    fmt.Println("awaiting signal")
    <-done
    fmt.Println("exiting")
}
