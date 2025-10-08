package main

import (
    "flag"
    "fmt"
)

func main() {
    // `flag.String` 声明一个字符串标志，返回一个指向该值的指针。
    wordPtr := flag.String("word", "foo", "a string")

    // 声明整数和布尔标志。
    numbPtr := flag.Int("numb", 42, "an int")
    boolPtr := flag.Bool("fork", false, "a bool")

    // 你也可以将标志绑定到一个已有的变量上。
    var svar string
    flag.StringVar(&svar, "svar", "bar", "a string var")

    // 在声明完所有标志后，必须调用 `flag.Parse()` 来从命令行参数中解析它们。
    flag.Parse()

    // 访问标志的值。
    fmt.Println("word:", *wordPtr)
    fmt.Println("numb:", *numbPtr)
    fmt.Println("fork:", *boolPtr)
    fmt.Println("svar:", svar)

    // `flag.Args()` 返回一个包含所有非标志参数的字符串切片。
    fmt.Println("tail:", flag.Args())
}
