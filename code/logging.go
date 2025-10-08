package main

import (
    "bytes"
    "fmt"
    "log"
    "log/slog"
    "os"
)

func main() {
    // --- 使用传统的 `log` 包 ---
    log.Println("standard logger")

    // `log.SetFlags` 可以用来配置日志输出的格式。
    log.SetFlags(log.LstdFlags | log.Lmicroseconds)
    log.Println("with micro")

    log.SetFlags(log.LstdFlags | log.Lshortfile)
    log.Println("with file/line")

    // `log.New` 可以创建一个新的记录器实例。
    mylog := log.New(os.Stdout, "my:", log.LstdFlags)
    mylog.Println("from mylog")

    mylog.SetPrefix("ohmy:")
    mylog.Println("from mylog")

    // 你可以将日志输出重定向到一个内存中的缓冲区。
    var buf bytes.Buffer
    buflog := log.New(&buf, "buf:", log.LstdFlags)
    buflog.Println("hello")
    fmt.Print("from buflog:", buf.String())

    // --- 使用 Go 1.21+ 的 `log/slog` 包 ---
    // `slog` 提供了结构化日志的功能。
    jsonHandler := slog.NewJSONHandler(os.Stderr, nil)
    myslog := slog.New(jsonHandler)
    myslog.Info("hi there")

    myslog.Info("hello again", "key", "val", "age", 25)
}
