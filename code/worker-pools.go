package main

import (
    "fmt"
    "time"
)

// `worker` 函数定义了一个工作池中的工作单元。
// 它接收一个 `jobs` 通道来获取任务，并将结果发送到 `results` 通道。
func worker(id int, jobs <-chan int, results chan<- int) {
    // `for range` 会在 `jobs` 通道被关闭后自动退出循环。
    for j := range jobs {
        fmt.Println("worker", id, "started  job", j)
        time.Sleep(time.Second) // 模拟耗时操作
        fmt.Println("worker", id, "finished job", j)
        results <- j * 2
    }
}

func main() {
    const numJobs = 5
    jobs := make(chan int, numJobs)
    results := make(chan int, numJobs)

    // 启动 3 个 worker，它们最初会因为没有任务而阻塞。
    for w := 1; w <= 3; w++ {
        go worker(w, jobs, results)
    }

    // 发送 5 个任务到 `jobs` 通道。
    for j := 1; j <= numJobs; j++ {
        jobs <- j
    }
    // 关闭 `jobs` 通道，表示所有任务都已发送。
    close(jobs)

    // 收集所有任务的结果。
    // 这也确保了 worker goroutine 在程序退出前已经完成。
    for a := 1; a <= numJobs; a++ {
        <-results
    }
}