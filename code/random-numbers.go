package main

import (
    "fmt"
    "math/rand/v2"
)

func main() {
    // `rand.IntN` 返回一个在 `[0, n)` 区间内的随机整数。
    fmt.Print(rand.IntN(100), ",")
    fmt.Print(rand.IntN(100))
    fmt.Println()

    // `rand.Float64` 返回一个在 `[0.0, 1.0)` 区间内的随机 `float64` 值。
    fmt.Println(rand.Float64())

    // 你可以通过这个函数来生成其他范围内的浮点数。
    fmt.Print((rand.Float64()*5)+5, ",")
    fmt.Print((rand.Float64()*5)+5)
    fmt.Println()

    // 如果你需要一个可复现的随机数序列，可以提供一个种子。
    s1 := rand.NewSource(42)
    r1 := rand.New(s1)

    fmt.Print(r1.IntN(100), ",")
    fmt.Print(r1.IntN(100))
    fmt.Println()

    // 使用相同种子的源将产生相同的随机数序列。
    s2 := rand.NewSource(42)
    r2 := rand.New(s2)
    fmt.Print(r2.IntN(100), ",")
    fmt.Print(r2.IntN(100))
    fmt.Println()
}
