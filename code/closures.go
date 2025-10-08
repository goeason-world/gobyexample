package main

import "fmt"

// `intSeq` 函数返回另一个函数，这个返回的函数在 `intSeq` 的函数体内匿名定义。
// 返回的函数对变量 `i` 形成闭包。
func intSeq() func() int {
    i := 0
    return func() int {
        i++
        return i
    }
}

func main() {
    // 调用 `intSeq`，将结果（一个函数）赋值给 `nextInt`。
    // 这个函数值捕获了它自己的 `i` 值，每次调用 `nextInt` 时都会更新这个值。
    nextInt := intSeq()

    // 通过多次调用 `nextInt` 来观察闭包的效果。
    fmt.Println(nextInt()) // 输出: 1
    fmt.Println(nextInt()) // 输出: 2
    fmt.Println(nextInt()) // 输出: 3

    // 为了确认状态对于特定函数是唯一的，创建并测试一个新的闭包。
    newInts := intSeq()
    fmt.Println(newInts()) // 输出: 1（重新开始计数）
}