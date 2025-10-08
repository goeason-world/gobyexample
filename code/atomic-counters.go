package main

import (
    "fmt"
    "sync"
    "sync/atomic"
)

func main() {
    // `sync/atomic` 包提供了低级别的原子内存操作，适合实现无锁的并发计数器。
    // `atomic.Uint64` 是一个无符号 64 位整数的原子类型。
    var ops atomic.Uint64

    // 使用 `sync.WaitGroup` 等待所有 goroutine 完成。
    var wg sync.WaitGroup

    // 启动 50 个 goroutine，每个执行 1000 次递增。
    for i := 0; i < 50; i++ {
        wg.Add(1)

        go func() {
            defer wg.Done()
            for c := 0; c < 1000; c++ {
                // `ops.Add(1)` 原子地增加计数器。
                ops.Add(1)
            }
        }()
    }

    // 等待所有 goroutine 完成。
    wg.Wait()

    // `ops.Load()` 原子地读取最终计数值。
    fmt.Println("ops:", ops.Load())
}