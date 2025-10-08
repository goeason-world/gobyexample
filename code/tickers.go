package main

import (
    "fmt"
    "time"
)

func main() {
    // `time.NewTicker` 用于按照固定间隔重复触发事件。
    ticker := time.NewTicker(500 * time.Millisecond)
    done := make(chan bool)

    go func() {
        for {
            select {
            case <-done:
                return
            case t := <-ticker.C:
                fmt.Println("Tick at", t)
            }
        }
    }()

    // 等待一段时间，让 ticker 触发几次。
    time.Sleep(1600 * time.Millisecond)
    // 停止 ticker。一旦停止，就不会再有新的事件发送到 `ticker.C`。
    ticker.Stop()
    // 通知 goroutine 退出。
    done <- true
    fmt.Println("Ticker stopped")
}