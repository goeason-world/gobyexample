// 每个 Go 程序都由包构成。
// Go 程序从 `main` 包中的代码开始运行。
package main

// `import` 语句用于导入其他包。
// `fmt` 包是 Go 的标准库之一，提供了格式化 I/O 的功能（类似于 C 的 printf 和 scanf）。
import "fmt"

// `main` 函数是程序的入口点。
// 当程序执行时，会首先调用 `main` 函数。
func main() {
    // `fmt.Println` 是一个函数，用于在控制台打印文本。
    // 它会自动在文本末尾添加一个换行符。
    fmt.Println("Hello World")
}

/*
要运行这个程序，请执行：
$ go run hello-world.go

要将程序编译成二进制文件，请执行：
$ go build hello-world.go
然后可以运行这个二进制文件：
$ ./hello-world
*/