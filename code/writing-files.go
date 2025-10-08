package main

import (
    "bufio"
    "fmt"
    "os"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    // `os.WriteFile` 是最简单的文件写入方式。
    d1 := []byte("hello\ngo\n")
    err := os.WriteFile("/tmp/dat1", d1, 0644)
    check(err)

    // 对于需要分步或更精细控制的写入操作，可以先用 `os.Create` 创建一个文件。
    f, err := os.Create("/tmp/dat2")
    check(err)

    // 使用 `defer` 确保文件在函数结束时关闭。
    defer f.Close()

    // `Write` 方法用于写入字节切片。
    d2 := []byte{115, 111, 109, 101, 10}
    n2, err := f.Write(d2)
    check(err)
    fmt.Printf("wrote %d bytes\n", n2)

    // `WriteString` 方法用于写入字符串。
    n3, err := f.WriteString("writes\n")
    check(err)
    fmt.Printf("wrote %d bytes\n", n3)

    // `Sync` 方法用于将缓冲区中的数据刷新到稳定的存储设备。
    f.Sync()

    // `bufio.NewWriter` 提供了一个带缓冲的写入器。
    w := bufio.NewWriter(f)
    n4, err := w.WriteString("buffered\n")
    check(err)
    fmt.Printf("wrote %d bytes\n", n4)

    // `Flush` 方法用于确保所有缓冲的数据都被写入到底层的文件中。
    w.Flush()
}
