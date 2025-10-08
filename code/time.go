package main

import (
    "fmt"
    "time"
)

func main() {
    p := fmt.Println

    // 获取并打印当前时间。
    now := time.Now()
    p(now)

    // 通过提供年月日等信息创建一个 `time.Time` 对象。
    // 时间总是与一个 `Location`（时区）相关联。
    then := time.Date(
        2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
    p(then)

    // 你可以提取时间的各个组成部分。
    p(then.Year())
    p(then.Month())
    p(then.Day())
    p(then.Hour())
    p(then.Minute())
    p(then.Second())
    p(then.Nanosecond())
    p(then.Location())

    // 输出是星期几。
    p(then.Weekday())

    // 比较两个时间，测试一个是否在另一个之前、之后或相等。
    p(then.Before(now))
    p(then.After(now))
    p(then.Equal(now))

    // `Sub` 方法返回一个 `Duration`，表示两个时间之间的时间差。
    diff := now.Sub(then)
    p(diff)

    // 我们可以用各种单位来表示 `Duration` 的长度。
    p(diff.Hours())
    p(diff.Minutes())
    p(diff.Seconds())
    p(diff.Nanoseconds())

    // 可以使用 `Add` 将一个 `Duration` 添加到一个时间点上，得到一个新的时间点。
    p(then.Add(diff))
    // `Add` 也可以接收负数，表示从一个时间点减去一个 `Duration`。
    p(then.Add(-diff))
}
