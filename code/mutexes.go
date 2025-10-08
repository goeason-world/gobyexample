package main

import (
    "fmt"
    "sync"
)

// `Container` 包含一个我们希望并发访问的 map。
// 我们在结构体中加入一个 `sync.Mutex` 来同步对 map 的访问。
type Container struct {
    mu       sync.Mutex
    counters map[string]int
}

// `inc` 方法在访问 `counters` 之前，先锁定互斥锁，
// 并在函数结束时使用 `defer` 语句来解锁。
func (c *Container) inc(name string) {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.counters[name]++
}

func main() {
    // 注意，互斥锁的零值可以直接使用，所以这里不需要初始化。
    c := Container{
        counters: map[string]int{"a": 0, "b": 0},
    }

    var wg sync.WaitGroup

    // `doIncrement` 函数在一个循环中增加一个指定名称的计数器。
    doIncrement := func(name string, n int) {
        for i := 0; i < n; i++ {
            c.inc(name)
        }
        wg.Done()
    }

    // 我们启动多个 goroutine 来并发地增加计数器。
    wg.Add(3)
    go doIncrement("a", 10000)
    go doIncrement("b", 10000)
    go doIncrement("a", 10000)

    // 等待所有 goroutine 完成。
    wg.Wait()
    fmt.Println(c.counters)
}