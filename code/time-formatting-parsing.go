package main

import (
    "fmt"
    "time"
)

func main() {
    p := fmt.Println

    // `time.Format` 使用一个布局字符串来格式化时间。
    // `time.RFC3339` 是一个预定义的布局常量。
    t := time.Now()
    p(t.Format(time.RFC3339))

    // `time.Parse` 使用相同的布局字符串来解析时间字符串。
    t1, _ := time.Parse(
        time.RFC3339,
        "2012-11-01T22:08:41+00:00")
    p(t1)

    // 你可以创建任何你需要的自定义布局。
    p(t.Format("3:04PM"))
    p(t.Format("Mon Jan _2 15:04:05 2006"))
    p(t.Format("2006-01-02T15:04:05.999999-07:00"))
    form := "3 04 PM"
    t2, _ := time.Parse(form, "8 41 PM")
    p(t2)

    // 对于纯数字的时间表示，你也可以使用 `fmt.Printf`。
    fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-07:00\n",
        t.Year(), t.Month(), t.Day(),
        t.Hour(), t.Minute(), t.Second())

    // 如果输入的时间字符串与布局不匹配，`time.Parse` 会返回一个错误。
    ansic := "Mon Jan _2 15:04:05 2006"
    _, e := time.Parse(ansic, "8:41PM")
    p(e)
}
