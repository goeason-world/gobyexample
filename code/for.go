package main

import "fmt"

func main() {

    // `for` 是 Go 语言中唯一的循环结构。这里展示了 `for` 循环的三种基本形式。

    // 1. 单一条件循环（类似其他语言的 `while`）
    i := 1
    for i <= 3 {
        fmt.Println(i)
        i = i + 1
    }

    // 2. 经典的三部分 `for` 循环：初始化/条件/后续
    for j := 7; j <= 9; j++ {
        fmt.Println(j)
    }

    // 3. 无条件的无限循环，需要使用 `break` 或 `return` 来退出
    for {
        fmt.Println("loop")
        break
    }

    // 4. 使用 `continue` 跳过当前迭代，进入下一次循环
    // 这个循环打印出 0-5 中的奇数
    for n := 0; n <= 5; n++ {
        if n%2 == 0 {
            continue
        }
        fmt.Println(n)
    }
}