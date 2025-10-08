package main

import "fmt"

// `fact` 函数调用自身直到达到基础情况 `fact(0)`。
func fact(n int) int {
    if n == 0 {
        return 1 // 基础情况：0! = 1
    }
    // 递归情况：n! = n * (n-1)!
    return n * fact(n-1)
}

func main() {
    fmt.Println(fact(7)) // 计算 7! = 5040

    // 匿名函数也可以是递归的，但这需要在定义函数之前
    // 用 `var` 显式声明一个变量来存储函数。
    var fib func(n int) int

    fib = func(n int) int {
        if n < 2 {
            return n // 基础情况：fib(0)=0, fib(1)=1
        }
        // 由于 `fib` 之前在 main 中声明过，Go 知道这里调用的是哪个函数。
        return fib(n-1) + fib(n-2) // 递归情况：fib(n) = fib(n-1) + fib(n-2)
    }

    fmt.Println(fib(7)) // 计算斐波那契数列第7项 = 13
}