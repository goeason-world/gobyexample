package main

import (
    "fmt"
    "strings"
)

// 为 `fmt.Println` 创建一个别名，方便使用。
var p = fmt.Println

func main() {
    // `strings` 包提供了大量用于处理字符串的实用函数。

    p("Contains:  ", strings.Contains("test", "es"))
    p("Count:     ", strings.Count("test", "t"))
    p("HasPrefix: ", strings.HasPrefix("test", "te"))
    p("HasSuffix: ", strings.HasSuffix("test", "st"))
    p("Index:     ", strings.Index("test", "e"))
    p("Join:      ", strings.Join([]string{"a", "b"}, "-"))
    p("Repeat:    ", strings.Repeat("a", 5))
    p("Replace:   ", strings.Replace("foo", "o", "0", -1)) // -1 表示替换所有
    p("Replace:   ", strings.Replace("foo", "o", "0", 1))  // 1 表示只替换 1 次
    p("Split:     ", strings.Split("a-b-c-d-e", "-"))
    p("ToLower:   ", strings.ToLower("TEST"))
    p("ToUpper:   ", strings.ToUpper("test"))
}