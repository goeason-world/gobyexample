package main

import (
    "fmt"
    "slices"
)

func main() {

    // 与数组不同，切片的类型仅由其包含的元素决定，而不包括元素的数量。
    // 一个未初始化的切片的值为 `nil`，其长度为 0。
    var s []string
    fmt.Println("uninit:", s, s == nil, len(s) == 0)

    // 使用内置的 `make` 函数创建一个非零长度的空切片。
    // 这里我们创建一个长度为3的字符串切片（初始值为零值 ""）。
    s = make([]string, 3)
    fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

    // 我们可以像数组一样设置和获取元素。
    s[0] = "a"
    s[1] = "b"
    s[2] = "c"
    fmt.Println("set:", s)
    fmt.Println("get:", s[2])

    // `len` 返回切片的长度。
    fmt.Println("len:", len(s))

    // `append` 是一个内置函数，用于向切片追加元素。
    // 它会返回一个包含新元素的、可能指向新底层数组的切片。
    s = append(s, "d")
    s = append(s, "e", "f")
    fmt.Println("apd:", s)

    // `copy` 函数可以将一个切片的内容复制到另一个切片。
    c := make([]string, len(s))
    copy(c, s)
    fmt.Println("cpy:", c)

    // 切片支持 "slice" 操作符，形式为 `slice[low:high]`。
    // 这会从索引 `low` 到 `high-1` 选取元素。
    l := s[2:5]
    fmt.Println("sl1:", l)

    // 这个切片从索引0到4 (5-1)
    l = s[:5]
    fmt.Println("sl2:", l)

    // 这个切片从索引2到结尾
    l = s[2:]
    fmt.Println("sl3:", l)

    // 使用切片字面量 (slice literal) 在一行内声明并初始化。
    t := []string{"g", "h", "i"}
    fmt.Println("dcl:", t)

    // `slices.Equal` 检查两个切片是否有相同的长度和对应位置上相等的元素。
    t2 := []string{"g", "h", "i"}
    if slices.Equal(t, t2) {
        fmt.Println("t == t2")
    }

    // 切片可以组合成多维数据结构。
    // 与多维数组不同，内部切片的长度可以变化。
    twoD := make([][]int, 3)
    for i := 0; i < 3; i++ {
        innerLen := i + 1
        twoD[i] = make([]int, innerLen)
        for j := 0; j < innerLen; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)
}