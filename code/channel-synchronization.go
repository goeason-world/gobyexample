package main

import (
    "fmt"
    "time"
)

// `worker` 函数将在一个 goroutine 中运行。
// `done` 通道将用于通知另一个 goroutine 这个函数已经完成了工作。
func worker(done chan bool) {
    fmt.Print("working...")
    time.Sleep(time.Second) // 模拟耗时任务
    fmt.Println("done")

    // 发送一个值来通知我们已经完成了。
    done <- true
}

func main() {
    // 创建一个 `done` 通道，我们将用它来同步。
    done := make(chan bool, 1)
    // 启动一个 worker goroutine，并给它 `done` 通道。
    go worker(done)

    // 在 `main` goroutine 中阻塞，直到我们从 `done` 通道接收到一个通知。
    // 如果我们移除了这行代码，程序将在 `worker` 启动之前就退出了。
    <-done
}