package main

import (
    "errors"
    "fmt"
)

// 自定义错误允许我们把额外的上下文信息绑定在错误值上。
// `argError` 结构体保存了触发错误的参数值和错误信息。
type argError struct {
    arg     int
    message string
}

// 只要类型实现了 `Error() string` 方法，就能被当作 `error` 使用。
func (e *argError) Error() string {
    return fmt.Sprintf("%d - %s", e.arg, e.message)
}

// `f` 函数在遇到特殊输入时返回一个自定义错误。
func f(arg int) (int, error) {
    if arg == 42 {
        // 返回一个指向 `argError` 结构体的指针。
        return -1, &argError{arg, "can't work with it"}
    }
    return arg + 3, nil
}

func main() {
    // 调用 `f` 函数并检查返回的错误。
    _, err := f(42)
    var ae *argError
    // `errors.As` 会沿着错误链查找第一个满足目标类型的错误，
    // 并把它解包到 `ae` 中，让我们方便地访问额外字段。
    if errors.As(err, &ae) {
        fmt.Println(ae.arg)
        fmt.Println(ae.message)
    } else {
        fmt.Println("err doesn't match argError")
    }
}