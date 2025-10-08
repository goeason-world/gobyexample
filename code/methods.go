package main

import "fmt"

// 定义一个矩形结构体
type rect struct {
    width, height int
}

// `area` 方法有一个指针接收者 `*rect`。
// 指针接收者允许方法修改接收者的值。
func (r *rect) area() int {
    return r.width * r.height
}

// `perim` 方法有一个值接收者 `rect`。
// 值接收者会复制结构体，不能修改原始值。
func (r rect) perim() int {
    return 2*r.width + 2*r.height
}

func main() {
    r := rect{width: 10, height: 5}

    // 调用为结构体定义的两个方法。
    fmt.Println("area: ", r.area())
    fmt.Println("perim:", r.perim())

    // Go 自动处理方法调用时值和指针之间的转换。
    // 你可能希望使用指针接收者类型来避免在方法调用时复制，
    // 或者允许方法修改接收结构体。
    rp := &r
    fmt.Println("area: ", rp.area())
    fmt.Println("perim:", rp.perim())
}