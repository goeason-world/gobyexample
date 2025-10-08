package main

import (
    "fmt"
    "unicode/utf8"
)

func main() {
    // `s` 是一个字符串，被赋予一个字面值，表示泰语中的"hello"。
    // Go 字符串字面值是 UTF-8 编码的文本。
    const s = "สวัสดี"

    // 由于字符串等价于 `[]byte`，`len()` 将产生存储在其中的原始字节的长度。
    fmt.Println("Len:", len(s))

    // 对字符串进行索引会产生每个索引处的原始字节值。
    // 这个循环生成构成 `s` 中 rune 的所有字节的十六进制值。
    for i := 0; i < len(s); i++ {
        fmt.Printf("%x ", s[i])
    }
    fmt.Println()

    // 要计算字符串中有多少个 rune，我们可以使用 `unicode/utf8` 包。
    fmt.Println("Rune count:", utf8.RuneCountInString(s))

    // `range` 循环对字符串进行特殊处理，解码每个 rune 及其在字符串中的偏移量。
    for idx, runeValue := range s {
        fmt.Printf("%#U starts at %d\n", runeValue, idx)
    }

    // 我们可以通过显式使用 `utf8.DecodeRuneInString` 函数来实现相同的迭代。
    fmt.Println("\nUsing DecodeRuneInString")
    for i, w := 0, 0; i < len(s); i += w {
        runeValue, width := utf8.DecodeRuneInString(s[i:])
        fmt.Printf("%#U starts at %d\n", runeValue, i)
        w = width

        // 这是一个演示如何比较 rune 值的函数调用。
        examineRune(runeValue)
    }
}

func examineRune(r rune) {
    // 用单引号括起来的值是 rune 字面值。
    // 我们可以直接将 rune 值与 rune 字面值进行比较。
    if r == 't' {
        fmt.Println("found tee")
    } else if r == 'ส' {
        fmt.Println("found so sua")
    }
}