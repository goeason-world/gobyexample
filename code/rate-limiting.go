package main

import (
    "fmt"
    "time"
)

func main() {
    // ---- 基础限流：一个请求必须等一个令牌 ----
    requests := make(chan int, 5)
    for i := 1; i <= 5; i++ {
        requests <- i
    }
    close(requests)

    // `time.Tick` 返回一个通道，该通道会按照指定的时间间隔发送事件。
    limiter := time.Tick(200 * time.Millisecond)

    for req := range requests {
        // 在处理每个请求之前，我们都会阻塞等待，直到从 `limiter` 通道接收到一个值。
        <-limiter
        fmt.Println("request", req, time.Now())
    }

    // ---- 突发限流：允许短暂的高峰 ----
    // 创建一个带缓冲的通道作为令牌桶。
    burstyLimiter := make(chan time.Time, 3)

    // 预先填充令牌桶，以允许突发请求。
    for i := 0; i < 3; i++ {
        burstyLimiter <- time.Now()
    }

    // 每 200 毫秒向令牌桶中添加一个新的令牌，直到达到容量 3。
    go func() {
        for t := range time.Tick(200 * time.Millisecond) {
            burstyLimiter <- t
        }
    }()

    burstyRequests := make(chan int, 5)
    for i := 1; i <= 5; i++ {
        burstyRequests <- i
    }
    close(burstyRequests)

    for req := range burstyRequests {
        // 从令牌桶中获取一个令牌。如果令牌桶为空，将会阻塞等待。
        <-burstyLimiter
        fmt.Println("request", req, time.Now())
    }
}