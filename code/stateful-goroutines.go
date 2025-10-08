package main

import (
    "fmt"
    "math/rand"
    "sync/atomic"
    "time"
)

// `readOp` 和 `writeOp` 结构体封装了读和写操作，
// 并且都包含一个用于响应的通道 `resp`。
type readOp struct {
    key  int
    resp chan int
}
type writeOp struct {
    key  int
    val  int
    resp chan bool
}

func main() {
    // `readOps` 和 `writeOps` 将记录我们执行的总操作数。
    var readOps uint64
    var writeOps uint64

    // `reads` 和 `writes` 通道分别用于发起读和写请求。
    reads := make(chan readOp)
    writes := make(chan writeOp)

    // 这是拥有状态的 goroutine，它独占 `state` map。
    // 这个 goroutine 反复从 `reads` 和 `writes` 通道中选择请求并执行。
    go func() {
        var state = make(map[int]int)
        for {
            select {
            case read := <-reads:
                read.resp <- state[read.key]
            case write := <-writes:
                state[write.key] = write.val
                write.resp <- true
            }
        }
    }()

    // 启动 100 个 goroutine 来模拟并发读取。
    for r := 0; r < 100; r++ {
        go func() {
            for {
                read := readOp{
                    key:  rand.Intn(5),
                    resp: make(chan int),
                }
                reads <- read
                <-read.resp
                atomic.AddUint64(&readOps, 1)
                time.Sleep(time.Millisecond)
            }
        }()
    }

    // 同样，我们启动 10 个 goroutine 来模拟并发写入。
    for w := 0; w < 10; w++ {
        go func() {
            for {
                write := writeOp{
                    key:  rand.Intn(5),
                    val:  rand.Intn(100),
                    resp: make(chan bool),
                }
                writes <- write
                <-write.resp
                atomic.AddUint64(&writeOps, 1)
                time.Sleep(time.Millisecond)
            }
        }()
    }

    // 让 goroutine 运行一秒钟。
    time.Sleep(time.Second)

    // 最后，获取并报告总操作数。
    readOpsFinal := atomic.LoadUint64(&readOps)
    fmt.Println("readOps:", readOpsFinal)
    writeOpsFinal := atomic.LoadUint64(&writeOps)
    fmt.Println("writeOps:", writeOpsFinal)
}