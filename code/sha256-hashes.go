package main

import (
    "crypto/sha256"
    "fmt"
)

func main() {
    s := "sha256 this string"

    // 1. 创建一个新的 sha256 哈希对象。
    h := sha256.New()

    // 2. 写入要哈希的数据（作为字节）。
    h.Write([]byte(s))

    // 3. 计算最终的哈希值。
    // `Sum` 的参数用于将哈希值追加到现有的字节切片中，通常传入 `nil`。
    bs := h.Sum(nil)

    fmt.Println(s)
    // 4. 将哈希值格式化为十六进制字符串。
    fmt.Printf("%x\n", bs)
}
