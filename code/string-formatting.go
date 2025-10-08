package main

import (
    "fmt"
    "os"
)

type point struct {
    x, y int
}

func main() {
    p := point{1, 2}

    // `%v` 打印结构体的值
    fmt.Printf("struct1: %v\n", p)

    // `%+v` 打印结构体字段名
    fmt.Printf("struct2: %+v\n", p)

    // `%#v` 打印 Go 语法表示
    fmt.Printf("struct3: %#v\n", p)

    // `%T` 打印类型
    fmt.Printf("type: %T\n", p)

    // `%t` 打印布尔值
    fmt.Printf("bool: %t\n", true)

    // `%d` 打印整数
    fmt.Printf("int: %d\n", 123)

    // `%b` 打印二进制
    fmt.Printf("bin: %b\n", 14)

    // `%c` 打印对应的 Unicode 字符
    fmt.Printf("char: %c\n", 33)

    // `%x` 打印十六进制表示
    fmt.Printf("hex: %x\n", 456)

    // `%f` 打印浮点数
    fmt.Printf("float1: %f\n", 78.9)

    // `%e` 和 `%E` 打印科学计数法
    fmt.Printf("float2: %e\n", 123400000.0)
    fmt.Printf("float3: %E\n", 123400000.0)

    // `%s` 打印字符串
    fmt.Printf("str1: %s\n", "\"string\"")

    // `%q` 为字符串加上双引号
    fmt.Printf("str2: %q\n", "\"string\"")

    // `%x` 将字符串转换为十六进制
    fmt.Printf("str3: %x\n", "hex this")

    // `%p` 打印指针
    fmt.Printf("pointer: %p\n", &p)

    // 控制宽度和精度
    fmt.Printf("width1: |%6d|%6d|\n", 12, 345)
    fmt.Printf("width2: |%6.2f|%6.2f|\n", 1.2, 3.45)
    fmt.Printf("width3: |%-6.2f|%-6.2f|\n", 1.2, 3.45)
    fmt.Printf("width4: |%6s|%6s|\n", "foo", "b")
    fmt.Printf("width5: |%-6s|%-6s|\n", "foo", "b")

    // `Sprintf` 返回一个格式化后的字符串
    s := fmt.Sprintf("sprintf: a %s", "string")
    fmt.Println(s)

    // `Fprintf` 将格式化后的字符串写入 `io.Writer`
    fmt.Fprintf(os.Stderr, "io: an %s\n", "error")
}