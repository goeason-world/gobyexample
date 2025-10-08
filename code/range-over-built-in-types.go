package main

import "fmt"

func main() {

    // `range` 用于遍历各种内置数据结构中的元素。

    // 使用 `range` 来计算切片中数字的和。
    // 数组的遍历方式与切片完全相同。
    nums := []int{2, 3, 4}
    sum := 0
    // `range` 返回索引和值，这里用 `_` 忽略索引。
    for _, num := range nums {
        sum += num
    }
    fmt.Println("sum:", sum)

    // `range` 在数组和切片上提供索引和值。
    // 如果需要索引，可以像这样使用。
    for i, num := range nums {
        if num == 3 {
            fmt.Println("index:", i)
        }
    }

    // `range` 在 map 上迭代键值对。
    kvs := map[string]string{"a": "apple", "b": "banana"}
    for k, v := range kvs {
        fmt.Printf("%s -> %s\n", k, v)
    }

    // `range` 也可以只遍历 map 的键。
    for k := range kvs {
        fmt.Println("key:", k)
    }

    // `range` 在字符串上遍历 Unicode rune。
    // 第一个值是 rune 的起始字节索引，第二个是 rune 本身。
    for i, c := range "go" {
        fmt.Println(i, c)
    }
}