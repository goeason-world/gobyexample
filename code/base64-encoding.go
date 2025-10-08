package main

import (
    "encoding/base64"
    "fmt"
)

func main() {
    data := "hello world"

    // 标准 Base64 编码。
    sEnc := base64.StdEncoding.EncodeToString([]byte(data))
    fmt.Println(sEnc)

    // 标准 Base64 解码。
    sDec, _ := base64.StdEncoding.DecodeString(sEnc)
    fmt.Println(string(sDec))

    // URL 兼容的 Base64 编码。
    // 使用不同的字符集，使其在 URL 中安全。
    uEnc := base64.URLEncoding.EncodeToString([]byte(data))
    fmt.Println(uEnc)

    // URL 兼容的 Base64 解码。
    uDec, _ := base.URLEncoding.DecodeString(uEnc)
    fmt.Println(string(uDec))
}
