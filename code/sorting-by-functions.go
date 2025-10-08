package main

import (
    "cmp"
    "fmt"
    "slices"
)

func main() {
    fruits := []string{"peach", "banana", "kiwi"}

    // `lenCmp` 是一个比较函数，它比较两个字符串的长度。
    // `cmp.Compare` 会返回 -1 (a<b), 0 (a==b), 或 +1 (a>b)。
    lenCmp := func(a, b string) int {
        return cmp.Compare(len(a), len(b))
    }

    // `slices.SortFunc` 使用我们自定义的比较函数进行排序。
    slices.SortFunc(fruits, lenCmp)
    fmt.Println(fruits)

    // --- 另一个例子：使用结构体切片 ---
    type Person struct {
        name string
        age  int
    }

    people := []Person{
        {"Alice", 30},
        {"Bob", 25},
        {"Charlie", 35},
    }

    // 按年龄排序。
    // 这里我们直接在 `SortFunc` 中使用一个匿名函数。
    slices.SortFunc(people, func(a, b Person) int {
        return cmp.Compare(a.age, b.age)
    })

    fmt.Println(people)
}