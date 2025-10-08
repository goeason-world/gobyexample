package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
)

// 定义一个辅助函数来检查错误。
func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    // `os.ReadFile` 是最简单的文件读取方式，它将文件的全部内容一次性读入一个字节切片中。
    dat, err := os.ReadFile("/tmp/dat")
    check(err)
    fmt.Print(string(dat))

    // 对于更复杂的读取需求，可以先用 `os.Open` 打开文件，得到一个 `*os.File` 对象。
    f, err := os.Open("/tmp/dat")
    check(err)

    // 使用 `defer` 确保文件在函数结束时被关闭。
    defer func() {
        if err := f.Close(); err != nil {
            panic(err)
        }
    }()

    // `Read` 方法可以从文件中读取指定数量的字节到一个切片中。
    b1 := make([]byte, 5)
    n1, err := f.Read(b1)
    check(err)
    fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))

    // `Seek` 方法可以在文件内移动读取指针。
    o2, err := f.Seek(6, 0)
    check(err)
    b2 := make([]byte, 2)
    n2, err := f.Read(b2)
    check(err)
    fmt.Printf("%d bytes @ %d: ", n2, o2)
    fmt.Printf("%v\n", string(b2[:n2]))

    // `io.ReadAtLeast` 可以保证至少读取指定数量的字节。
    o3, err := f.Seek(6, 0)
    check(err)
    b3 := make([]byte, 2)
    n3, err := io.ReadAtLeast(f, b3, 2)
    check(err)
    fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

    // `bufio` 包提供了带缓冲的读取器。
    _, err = f.Seek(0, 0)
    check(err)
    r4 := bufio.NewReader(f)
    b4, err := r4.Peek(5)
    check(err)
    fmt.Printf("5 bytes: %s\n", string(b4))
}
