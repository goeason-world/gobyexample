package main

import "fmt"

// `zeroval` 有一个 `int` 参数，所以参数将按值传递给它。
// `zeroval` 将得到 `ival` 的一个副本，与调用函数中的变量是不同的。
func zeroval(ival int) {
    ival = 0 // 只修改副本，不影响原始变量
}

// `zeroptr` 有一个 `*int` 参数，意味着它接收一个 `int` 指针。
// 函数体中的 `*iptr` 代码将指针从其内存地址解引用到该地址的当前值。
// 给解引用的指针赋值会改变引用地址处的值。
func zeroptr(iptr *int) {
    *iptr = 0 // 通过指针修改原始变量的值
}

func main() {
    i := 1
    fmt.Println("initial:", i)

    // 值传递：不会改变原始变量 `i`
    zeroval(i)
    fmt.Println("zeroval:", i)

    // `&i` 语法给出 `i` 的内存地址，即指向 `i` 的指针。
    // 指针传递：会改变原始变量 `i`
    zeroptr(&i)
    fmt.Println("zeroptr:", i)

    // 指针也可以被打印出来。
    fmt.Println("pointer:", &i)
}