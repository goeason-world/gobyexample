package main

import (
    "errors"
    "fmt"
)

// 在 Go 中，错误处理遵循“返回值 + `error`”的模式。
// `f` 函数在遇到特殊输入时返回一个错误。
func f(arg int) (int, error) {
    if arg == 42 {
        // `errors.New` 返回一个带固定消息的错误值。
        return -1, errors.New("can't work with 42")
    }
    // 返回 `nil` 表示没有错误。
    return arg + 3, nil
}

// 哨兵错误（sentinel error）是预声明的全局变量，用来描述某种固定的失败状态。
var ErrOutOfTea = fmt.Errorf("no more tea available")
var ErrPower = fmt.Errorf("can't boil water")

// `makeTea` 函数演示了如何返回哨兵错误和包装错误。
func makeTea(arg int) error {
    if arg == 2 {
        return ErrOutOfTea
    } else if arg == 4 {
        // `fmt.Errorf` 的 `%w` 动词可以将下层错误包装起来，形成错误链。
        return fmt.Errorf("making tea: %w", ErrPower)
    }
    return nil
}

func main() {
    // 使用 `if err != nil` 是 Go 中最常见的错误处理模式。
    for _, i := range []int{7, 42} {
        if r, err := f(i); err != nil {
            fmt.Println("f failed:", err)
        } else {
            fmt.Println("f worked:", r)
        }
    }

    // `errors.Is` 会沿着错误链查找目标哨兵错误。
    for i := range 5 {
        if err := makeTea(i); err != nil {
            if errors.Is(err, ErrOutOfTea) {
                fmt.Println("We should buy new tea!")
            } else if errors.Is(err, ErrPower) {
                fmt.Println("Now it is dark.")
            } else {
                fmt.Printf("unknown error: %s\n", err)
            }
            continue
        }
        fmt.Println("Tea is ready!")
    }
}