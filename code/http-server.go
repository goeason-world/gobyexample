package main

import (
    "fmt"
    "net/http"
)

// `hello` 是一个简单的处理器函数。
func hello(w http.ResponseWriter, req *http.Request) {
    // `http.ResponseWriter` 用于填充 HTTP 响应。
    fmt.Fprintf(w, "hello\n")
}

// `headers` 处理器会读取并打印出 HTTP 请求中的所有头信息。
func headers(w http.ResponseWriter, req *http.Request) {
    for name, headers := range req.Header {
        for _, h := range headers {
            fmt.Fprintf(w, "%v: %v\n", name, h)
        }
    }
}

func main() {
    // `http.HandleFunc` 用于将一个处理器函数注册到一个特定的 URL 路径上。
    http.HandleFunc("/hello", hello)
    http.HandleFunc("/headers", headers)

    // `http.ListenAndServe` 用于启动服务器并开始在指定的端口上监听请求。
    http.ListenAndServe(":8090", nil)
}
