package main

import (
    "fmt"
    "time"
)

func main() {
    // `time.NewTimer` 用于在未来某个时刻触发一次事件。
    timer1 := time.NewTimer(2 * time.Second)

    // `<-timer1.C` 会阻塞，直到定时器到期。
    <-timer1.C
    fmt.Println("Timer 1 fired")

    // 如果你只是想等待一段时间，可以使用 `time.Sleep`。
    // 定时器之所以有用，是因为你可以在定时器触发之前取消它。
    timer2 := time.NewTimer(time.Second)
    go func() {
        <-timer2.C
        fmt.Println("Timer 2 fired")
    }()
    // `Stop` 方法可以取消定时器。
    stop2 := timer2.Stop()
    if stop2 {
        fmt.Println("Timer 2 stopped")
    }

    // 等待足够长的时间，以确认 `timer2` 确实没有触发。
    time.Sleep(2 * time.Second)
}