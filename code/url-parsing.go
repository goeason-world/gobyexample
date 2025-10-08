package main

import (
    "fmt"
    "net"
    "net/url"
)

func main() {
    s := "postgres://user:pass@host.com:5432/path?k=v#f"

    // `url.Parse` 是解析 URL 字符串的核心函数。
    u, err := url.Parse(s)
    if err != nil {
        panic(err)
    }

    // 协议
    fmt.Println(u.Scheme)

    // 用户认证信息
    fmt.Println(u.User)
    fmt.Println(u.User.Username())
    p, _ := u.User.Password()
    fmt.Println(p)

    // 主机和端口
    fmt.Println(u.Host)
    host, port, _ := net.SplitHostPort(u.Host)
    fmt.Println(host)
    fmt.Println(port)

    // 路径和片段
    fmt.Println(u.Path)
    fmt.Println(u.Fragment)

    // 查询参数
    fmt.Println(u.RawQuery)
    m, _ := url.ParseQuery(u.RawQuery)
    fmt.Println(m)
    fmt.Println(m["k"][0])
}
