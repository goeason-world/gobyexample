package main

import (
    "fmt"
    "path/filepath"
    "strings"
)

func main() {
    // `filepath.Join` 用于构造路径，它会自动处理不同操作系统的路径分隔符。
    p := filepath.Join("dir1", "dir2", "filename")
    fmt.Println("p:", p)

    // `Join` 会自动处理多余的分隔符和 `..`。
    fmt.Println(filepath.Join("dir1//", "filename"))
    fmt.Println(filepath.Join("dir1/../dir1", "filename"))

    // `Dir` 和 `Base` 用于分解路径。
    fmt.Println("Dir(p):", filepath.Dir(p))
    fmt.Println("Base(p):", filepath.Base(p))

    // `IsAbs` 判断是否为绝对路径。
    fmt.Println(filepath.IsAbs("dir/file"))
    fmt.Println(filepath.IsAbs("/dir/file"))

    // `Ext` 获取文件扩展名。
    filename := "config.json"
    ext := filepath.Ext(filename)
    fmt.Println(ext)

    // 获取不带扩展名的文件名。
    fmt.Println(strings.TrimSuffix(filename, ext))

    // `Rel` 计算相对路径。
    rel, err := filepath.Rel("a/b", "a/b/t/file")
    if err != nil {
        panic(err)
    }
    fmt.Println(rel)

    rel, err = filepath.Rel("a/b", "a/c/t/file")
    if err != nil {
        panic(err)
    }
    fmt.Println(rel)
}
