package main

import (
    "fmt"
    "net/http"
    "time"
)

// `hello` 是一个模拟耗时操作的 HTTP 处理器。
func hello(w http.ResponseWriter, req *http.Request) {
    // `req.Context()` 用于获取当前请求关联的 `Context`。
    ctx := req.Context()
    fmt.Println("server: hello handler started")
    defer fmt.Println("server: hello handler ended")

    // 使用 `select` 等待耗时操作或 `Context` 取消。
    select {
    case <-time.After(10 * time.Second):
        // 操作完成，正常返回响应。
        fmt.Fprintf(w, "hello\n")
    case <-ctx.Done():
        // `Context` 被取消了（例如，客户端断开连接）。
        err := ctx.Err()
        fmt.Println("server:", err)
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func main() {
    http.HandleFunc("/hello", hello)
    http.ListenAndServe(":8090", nil)
}
