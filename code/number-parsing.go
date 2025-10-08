package main

import (
    "fmt"
    "strconv"
)

func main() {
    // `strconv.ParseFloat` 用于将字符串解析为浮点数。
    f, _ := strconv.ParseFloat("1.234", 64)
    fmt.Println(f)

    // `strconv.ParseInt` 用于解析整数。
    i, _ := strconv.ParseInt("123", 0, 64)
    fmt.Println(i)

    // `ParseInt` 可以识别十六进制格式。
    d, _ := strconv.ParseInt("0x1c8", 0, 64)
    fmt.Println(d)

    // `strconv.ParseUint` 用于解析无符号整数。
    u, _ := strconv.ParseUint("789", 0, 64)
    fmt.Println(u)

    // `strconv.Atoi` 是一个方便的函数，用于将十进制字符串解析为 `int` 类型。
    k, _ := strconv.Atoi("135")
    fmt.Println(k)

    // 如果字符串无法被解析为有效的数字，解析函数会返回一个错误。
    _, e := strconv.Atoi("wat")
    fmt.Println(e)
}
