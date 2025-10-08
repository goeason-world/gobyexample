package main

import (
    "fmt"
    "sync"
    "time"
)

// `worker` 函数模拟一个耗时任务。
func worker(id int) {
    fmt.Printf("Worker %d starting\n", id)
    time.Sleep(time.Second)
    fmt.Printf("Worker %d done\n", id)
}

func main() {
    // `sync.WaitGroup` 用于等待一组 goroutine 完成。
    var wg sync.WaitGroup

    for i := 1; i <= 5; i++ {
        // `i := i` 是一个重要的步骤，它为每个 goroutine 创建了一个独立的变量 `i` 的副本。
        // 如果不这样做，所有的 goroutine 都会引用同一个 `i`，导致行为不符合预期。
        i := i

        // `wg.Go` (Go 1.22+) 是一个方便的方法，它会自动调用 `Add(1)` 和 `Done()`。
        // 它接收一个函数，并在一个新的 goroutine 中执行它。
        wg.Go(func() {
            worker(i)
        })

        /*
        // 在 Go 1.22 之前的版本中，你需要手动管理 WaitGroup 的计数器：
        wg.Add(1) // 增加计数器
        go func() {
            defer wg.Done() // 在 goroutine 完成时减少计数器
            worker(i)
        }()
        */
    }

    // `wg.Wait()` 会阻塞，直到 WaitGroup 的计数器变为 0。
    wg.Wait()
}