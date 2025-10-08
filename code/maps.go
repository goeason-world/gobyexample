package main

import (
    "fmt"
    "maps"
)

func main() {

    // Map (映射) 是 Go 语言内置的关联数据类型（在其他语言中称为哈希或字典）。
    // 使用 `make(map[key-type]val-type)` 创建一个空的 Map。
    m := make(map[string]int)

    // 使用 `map[key] = val` 语法设置键值对。
    m["k1"] = 7
    m["k2"] = 13
    fmt.Println("map:", m)

    // 使用 `map[key]` 获取值。
    v1 := m["k1"]
    fmt.Println("v1:", v1)

    // 如果获取一个不存在的键，会返回该值类型的零值。
    v3 := m["k3"]
    fmt.Println("v3:", v3) // int 的零值为 0

    // `len()` 函数获取 Map 中键值对的数量。
    fmt.Println("len:", len(m))

    // `delete` 函数从 Map 中删除一个键值对。
    delete(m, "k2")
    fmt.Println("map:", m)

    // 从 Map 中获取值时，可以接收一个可选的第二个返回值，
    // 该值是一个布尔类型，表示键是否存在。
    // 这对于区分“值为零值”和“键不存在”的场景非常有用。
    _, prs := m["k2"] // "k2" 已被删除
    fmt.Println("prs:", prs)

    // 可以在声明时使用 Map 字面量进行初始化。
    n := map[string]int{"foo": 1, "bar": 2}
    fmt.Println("dcl:", n)

    // `maps.Equal` 函数可以比较两个 Map 的内容是否相等。
    n2 := map[string]int{"foo": 1, "bar": 2}
    if maps.Equal(n, n2) {
        fmt.Println("n == n2")
    }
}