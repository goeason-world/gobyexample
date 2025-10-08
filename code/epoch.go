package main

import (
    "fmt"
    "time"
)

func main() {
    // 使用 `time.Now()` 获取当前时间。
    now := time.Now()

    // `Unix()` 或 `UnixNano()` 获取自 Unix 纪元以来的秒数或纳秒数。
    fmt.Println(now.Unix())
    fmt.Println(now.UnixMilli())
    fmt.Println(now.UnixNano())

    // 你也可以将自纪元以来的秒数或纳秒数转换回 `time.Time`。
    fmt.Println(time.Unix(now.Unix(), 0))
    fmt.Println(time.Unix(0, now.UnixNano()))
}
