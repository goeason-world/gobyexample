package main

import "fmt"

func main() {

    // 在 Go 中，数组是具有特定长度的、编号的元素序列。
    // 声明一个包含5个整数的数组 `a`。
    // 默认情况下，数组是零值的，对于整数来说就是 0。
    var a [5]int
    fmt.Println("emp:", a)

    // 我们可以使用 `array[index] = value` 语法来设置数组指定索引的元素值，
    // 或者使用 `array[index]` 来获取值。
    a[4] = 100
    fmt.Println("set:", a)
    fmt.Println("get:", a[4])

    // 使用内置的 `len()` 函数可以获取数组的长度。
    fmt.Println("len:", len(a))

    // 使用数组字面量在声明时初始化数组。
    b := [5]int{1, 2, 3, 4, 5}
    fmt.Println("dcl:", b)

    // 使用 `...` 语法，可以让编译器自动计算数组的长度。
    b = [...]int{1, 2, 3, 4, 5}
    fmt.Println("dcl:", b)

    // 可以在初始化时指定特定索引的值。
    // 未被初始化的元素将被赋予其类型的零值。
    b = [...]int{100, 3: 400, 500} // 索引1和2的元素为0
    fmt.Println("idx:", b)

    // 数组也可以是多维的。
    var twoD [2][3]int
    for i := 0; i < 2; i++ {
        for j := 0; j < 3; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)

    // 也可以在声明时初始化多维数组。
    twoD = [2][3]int{
        {1, 2, 3},
        {1, 2, 3},
    }
    fmt.Println("2d: ", twoD)
}