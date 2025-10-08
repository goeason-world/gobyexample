package main

import (
    "fmt"
    "iter"
    "slices"
)

// `List` 是一个泛型单向链表。
type List[T any] struct {
    head, tail *element[T]
}

// `element` 是链表中的节点。
type element[T any] struct {
    next *element[T]
    val  T
}

// `Push` 方法向链表尾部添加元素。
func (lst *List[T]) Push(v T) {
    if lst.tail == nil {
        lst.head = &element[T]{val: v}
        lst.tail = lst.head
    } else {
        lst.tail.next = &element[T]{val: v}
        lst.tail = lst.tail.next
    }
}

// `All` 方法返回一个 `iter.Seq[T]` 类型的迭代器，用于遍历链表。
func (lst *List[T]) All() iter.Seq[T] {
    return func(yield func(T) bool) {
        for e := lst.head; e != nil; e = e.next {
            // 调用 `yield` 函数来“产生”值。
            // 如果 `yield` 返回 `false`，则停止迭代。
            if !yield(e.val) {
                return
            }
        }
    }
}

// `Fibonacci` 函数返回一个无限的斐波那契数列迭代器。
func Fibonacci() iter.Seq[int] {
    return func(yield func(int) bool) {
        a, b := 0, 1
        // 只要 `yield` 返回 `true`，就一直产生下一个数。
        for yield(a) {
            a, b = b, a+b
        }
    }
}

func main() {
    // 遍历自定义链表。
    var lst List[int]
    lst.Push(1)
    lst.Push(2)
    lst.Push(3)
    for v := range lst.All() {
        fmt.Println(v)
    }

    // 遍历斐波那契数列，直到值大于 50。
    for v := range Fibonacci() {
        if v > 50 {
            break
        }
        fmt.Println(v)
    }

    // 使用 `iter.Pull` 创建一个“拉”模式的迭代器。
    pull, stop := iter.Pull(Fibonacci())
    defer stop() // 确保资源被释放
    for i := 0; i < 10; i++ {
        v, ok := pull()
        if !ok {
            break
        }
        fmt.Println(v)
    }

    // `slices.Collect` 函数可以从一个迭代器中收集所有的值，并将它们存入一个切片。
    fib10 := slices.Collect(iter.Take(Fibonacci(), 10))
    fmt.Println(fib10)
}