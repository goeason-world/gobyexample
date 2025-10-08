package main

import (
    "fmt"
    "math"
)

// `geometry` 是一个几何图形的基本接口。
// 接口是方法签名的命名集合。
type geometry interface {
    area() float64
    perim() float64
}

// `rect` 和 `circle` 是我们将要实现 `geometry` 接口的两个类型。
type rect struct {
    width, height float64
}
type circle struct {
    radius float64
}

// 要在 Go 中实现一个接口，我们只需要实现接口中的所有方法。
// 这里我们在 `rect` 上实现 `geometry` 接口。
func (r rect) area() float64 {
    return r.width * r.height
}
func (r rect) perim() float64 {
    return 2*r.width + 2*r.height
}

// `circle` 的实现。
func (c circle) area() float64 {
    return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
    return 2 * math.Pi * c.radius
}

// `measure` 函数接收一个 `geometry` 接口类型的参数，
// 它可以处理任何实现了 `geometry` 接口的类型。
func measure(g geometry) {
    fmt.Println(g)
    fmt.Println(g.area())
    fmt.Println(g.perim())
}

// `detectCircle` 函数使用类型断言来检查接口值的运行时类型。
func detectCircle(g geometry) {
    // `g.(circle)` 尝试将接口 `g` 断言为 `circle` 类型。
    // 如果断言成功，`ok` 为 `true`，`c` 为转换后的 `circle` 值。
    if c, ok := g.(circle); ok {
        fmt.Println("circle with radius:", c.radius)
    }
}

func main() {
    r := rect{width: 3, height: 4}
    c := circle{radius: 5}

    // `circle` 和 `rect` 结构体类型都实现了 `geometry` 接口，
    // 所以我们可以使用这些结构体的实例作为 `measure` 的参数。
    measure(r)
    measure(c)

    // 使用类型断言来检测具体的类型。
    detectCircle(r) // 不会输出任何内容，因为 r 不是 circle
    detectCircle(c) // 输出圆的半径信息
}