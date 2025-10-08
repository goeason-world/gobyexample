package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    // `bufio.NewScanner` 创建一个从标准输入读取的扫描器。
    scanner := bufio.NewScanner(os.Stdin)

    // `scanner.Scan()` 逐行扫描输入。
    for scanner.Scan() {
        // `scanner.Text()` 返回当前行的内容。
        ucl := strings.ToUpper(scanner.Text())
        // 将处理结果打印到标准输出。
        fmt.Println(ucl)
    }

    // 检查扫描过程中是否出错。
    if err := scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "error:", err)
        os.Exit(1)
    }
}
