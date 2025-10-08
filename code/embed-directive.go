package main

import (
    "embed"
    "fmt"
)

// `//go:embed` 是一个编译器指令，它允许在编译时将文件内容直接嵌入到 Go 二进制文件中。

// 将文件内容嵌入到 `string`。
//go:embed folder/single_file.txt
var fileString string

// 将文件内容嵌入到 `[]byte`。
//go:embed folder/single_file.txt
var fileByte []byte

// 将多个文件嵌入到一个虚拟文件系统。
//go:embed folder/single_file.txt
//go:embed folder/*.hash
var folder embed.FS

func main() {
    // 直接访问嵌入的 `string` 和 `[]byte`。
    fmt.Print(fileString)
    fmt.Print(string(fileByte))

    // 从虚拟文件系统中读取文件。
    content1, _ := folder.ReadFile("folder/file1.hash")
    fmt.Print(string(content1))

    content2, _ := folder.ReadFile("folder/file2.hash")
    fmt.Print(string(content2))
}
