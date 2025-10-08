package main

import (
    "bytes"
    "fmt"
    "regexp"
)

func main() {
    // `regexp.MatchString` 用于判断一个模式是否存在于字符串中。
    match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
    fmt.Println(match)

    // 对于需要多次使用的正则表达式，更好的做法是先用 `regexp.Compile` 将其编译成一个 `Regexp` 对象。
    r, _ := regexp.Compile("p([a-z]+)ch")

    // `Regexp` 对象有很多方法。
    fmt.Println(r.MatchString("peach"))

    // `FindString` 查找模式的第一个匹配项。
    fmt.Println("FindString:", r.FindString("peach punch"))

    // `FindStringIndex` 查找第一个匹配项的开始和结束索引。
    fmt.Println("idx:", r.FindStringIndex("peach punch"))

    // `FindStringSubmatch` 查找第一个匹配项及其所有子匹配（捕获组）。
    fmt.Println(r.FindStringSubmatch("peach punch"))

    // `FindStringSubmatchIndex` 查找第一个匹配项及其子匹配的索引。
    fmt.Println(r.FindStringSubmatchIndex("peach punch"))

    // `FindAllString` 查找所有匹配项。
    fmt.Println(r.FindAllString("peach punch pinch", -1))

    // `FindAllStringSubmatchIndex` 查找所有匹配项及其子匹配的索引。
    fmt.Println("all:", r.FindAllStringSubmatchIndex(
        "peach punch pinch", -1))

    // `FindAllString` 的第二个参数可以限制匹配次数。
    fmt.Println(r.FindAllString("peach punch pinch", 2))

    // `Match` 方法可以用于字节切片。
    fmt.Println(r.Match([]byte("peach")))

    // `regexp.MustCompile` 在编译失败时会 panic，而不是返回一个错误。
    r = regexp.MustCompile("p([a-z]+)ch")
    fmt.Println("regexp:", r)

    // `ReplaceAllString` 替换所有匹配项为指定字符串。
    fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))

    // `ReplaceAllFunc` 使用自定义函数对每个匹配项进行替换。
    in := []byte("a peach")
    out := r.ReplaceAllFunc(in, bytes.ToUpper)
    fmt.Println(string(out))
}