package main

import (
    "bufio"
    "fmt"
    "net/http"
)

func main() {
    // `http.Get` 是发送一个 HTTP GET 请求最简单的方式。
    resp, err := http.Get("https://gobyexample.com")
    if err != nil {
        panic(err)
    }
    // `resp.Body` 是一个 `io.ReadCloser`，在使用完毕后必须被关闭。
    defer resp.Body.Close()

    // `resp.Status` 包含了响应的状态码和文本描述。
    fmt.Println("Response status:", resp.Status)

    // `resp.Body` 是一个 `io.Reader`，可以从中读取响应的正文内容。
    scanner := bufio.NewScanner(resp.Body)
    for i := 0; scanner.Scan() && i < 5; i++ {
        fmt.Println(scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        panic(err)
    }
}
