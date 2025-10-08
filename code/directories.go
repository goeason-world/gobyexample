package main

import (
    "fmt"
    "io/fs"
    "os"
    "path/filepath"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    // `os.Mkdir` 用于创建一个新目录。
    err := os.Mkdir("subdir", 0755)
    check(err)

    // `os.RemoveAll` 会删除整个目录树。
    defer os.RemoveAll("subdir")

    // 辅助函数：创建一个空文件。
    createEmptyFile := func(name string) {
        d := []byte("")
        check(os.WriteFile(name, d, 0644))
    }

    createEmptyFile("subdir/file1")

    // `os.MkdirAll` 可以递归地创建多层目录。
    err = os.MkdirAll("subdir/parent/child", 0755)
    check(err)

    createEmptyFile("subdir/parent/file2")
    createEmptyFile("subdir/parent/file3")
    createEmptyFile("subdir/parent/child/file4")

    // `os.ReadDir` 用于读取一个目录的内容。
    c, err := os.ReadDir("subdir/parent")
    check(err)

    fmt.Println("Listing subdir/parent")
    for _, entry := range c {
        fmt.Println(" ", entry.Name(), entry.IsDir())
    }

    // `os.Chdir` 用于改变当前的G作目录。
    err = os.Chdir("subdir/parent/child")
    check(err)

    // 读取当前目录的内容。
    c, err = os.ReadDir(".")
    check(err)

    fmt.Println("Listing .")
    for _, entry := range c {
        fmt.Println(" ", entry.Name(), entry.IsDir())
    }

    // 切换回初始目录。
    err = os.Chdir("../../..")
    check(err)

    // `filepath.WalkDir` 用于递归地遍历一个目录树。
    fmt.Println("Visiting subdir")
    err = filepath.WalkDir("subdir", visit)
    check(err)
}

// `WalkDir` 的访问函数。
func visit(path string, d fs.DirEntry, err error) error {
    if err != nil {
        return err
    }
    fmt.Println(" ", path, d.IsDir())
    return nil
}
