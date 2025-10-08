package main

import "fmt"

// `base` 是一个基础结构体。
type base struct {
    num int
}

// `describe` 是 `base` 的一个方法。
func (b base) describe() string {
    return fmt.Sprintf("base with num=%v", b.num)
}

// `container` 结构体嵌入了 `base` 结构体。
// 嵌入时，我们只提供类型名，不提供字段名。
type container struct {
    base
    str string
}

func main() {
    // 初始化 `container` 时，我们可以像普通字段一样初始化嵌入的 `base` 结构体。
    co := container{
        base: base{
            num: 1,
        },
        str: "some name",
    }

    // 我们可以直接通过 `container` 实例访问 `base` 的字段 `num`。
    fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)

    // 也可以通过完整的路径访问嵌入的字段。
    fmt.Println("also num:", co.base.num)

    // `base` 的方法 `describe` 也被提升到了 `container`，
    // 因此我们可以直接在 `container` 实例上调用它。
    fmt.Println("describe:", co.describe())

    // 由于 `container` 拥有 `describe` 方法，它也隐式地实现了包含该方法的接口。
    type describer interface {
        describe() string
    }

    var d describer = co
    fmt.Println("describer:", d.describe())
}