package main

import "fmt"

func main() {
    jobs := make(chan int, 5)
    done := make(chan bool)

    // 在一个 worker goroutine 中处理任务。
    go func() {
        for {
            // `more` 的值在通道被关闭且没有更多值可接收时为 `false`。
            j, more := <-jobs
            if more {
                fmt.Println("received job", j)
            } else {
                fmt.Println("received all jobs")
                done <- true
                return
            }
        }
    }()

    // 发送 3 个任务到 `jobs` 通道，然后关闭它。
    for j := 1; j <= 3; j++ {
        jobs <- j
        fmt.Println("sent job", j)
    }
    close(jobs)
    fmt.Println("sent all jobs")

    // 等待 worker goroutine 完成。
    <-done

    // 从一个已关闭的通道接收值会立即成功，并返回该类型的零值。
    // 第二个返回值 `ok` 会是 `false`，表示通道已关闭。
    _, ok := <-jobs
    fmt.Println("received more jobs:", ok)
}