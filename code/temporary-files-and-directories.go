package main

import (
    "fmt"
    "os"
    "path/filepath"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    // `os.CreateTemp` 用于在系统默认的临时目录中创建一个新的临时文件。
    f, err := os.CreateTemp("", "sample")
    check(err)

    fmt.Println("Temp file name:", f.Name())

    // 使用 `defer` 和 `os.Remove` 来确保在程序退出时删除该文件。
    defer os.Remove(f.Name())

    // 向临时文件写入数据。
    _, err = f.Write([]byte{1, 2, 3, 4})
    check(err)

    // `os.MkdirTemp` 用于创建临时目录。
    dname, err := os.MkdirTemp("", "sampledir")
    check(err)
    fmt.Println("Temp dir name:", dname)

    // 使用 `defer` 和 `os.RemoveAll` 来递归删除目录及其所有内容。
    defer os.RemoveAll(dname)

    // 在临时目录中创建一个文件。
    fname := filepath.Join(dname, "file1")
    err = os.WriteFile(fname, []byte{1, 2}, 0666)
    check(err)
}
