package main

import (
    "fmt"
    "slices"
)

func main() {
    // Go 的 `slices` 包提供了对切片进行排序的功能。
    // 对于常见的内置类型，Go 提供了开箱即用的排序函数。

    // 创建一个字符串切片。
    strs := []string{"c", "a", "b"}
    // `slices.Sort` 会就地对切片进行排序。
    slices.Sort(strs)
    fmt.Println("Strings:", strs)

    // 创建一个整数切片。
    ints := []int{7, 2, 4}
    // 同样使用 `slices.Sort` 进行排序。
    slices.Sort(ints)
    fmt.Println("Ints:   ", ints)

    // `slices.IsSorted` 可以检查一个切片是否已经排好序。
    s := slices.IsSorted(ints)
    fmt.Println("Sorted: ", s)
}