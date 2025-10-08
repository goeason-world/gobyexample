package main

import "fmt"

// 可变参数函数（Variadic functions）是一种可以接受任意数量的尾随参数的函数。
// 通过在参数类型前加上 `...` 来定义一个可变参数函数。
// 在函数内部，这个参数会变成一个相应类型的切片。
func sum(nums ...int) {
    fmt.Print(nums, " ") // nums 的类型是 []int
    total := 0
    for _, num := range nums {
        total += num
    }
    fmt.Println(total)
}

func main() {
    // 使用独立参数调用可变参数函数。
    sum(1, 2)
    sum(1, 2, 3)

    // 如果参数已经存在于一个切片中，可以通过在切片名后加上 `...` 的方式，
    // 将切片中的所有元素作为独立参数传递给函数。
    nums := []int{1, 2, 3, 4}
    sum(nums...)
}