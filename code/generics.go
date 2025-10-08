package main

import "fmt"

// `SlicesIndex` 是一个泛型函数，它可以在一个切片里查找某个元素的位置。
// `[S ~[]E, E comparable]` 定义了类型参数和约束。
// `S` 是一个切片类型，其元素类型是 `E`。
// `E` 必须是 `comparable` (可比较的)。
func SlicesIndex[S ~[]E, E comparable](s S, v E) int {
    for i, vs := range s {
        if v == vs {
            return i
        }
    }
    return -1
}

// `List` 是一个泛型类型，它可以存放任何类型的元素。
// `[T any]` 表示类型参数 `T` 可以是任何类型 (`any` 是 `interface{}` 的别名)。
type List[T any] struct {
    head, tail *element[T]
}

// `element` 是链表中的节点，也是一个泛型类型。
type element[T any] struct {
    next *element[T]
    val  T
}

// `Push` 方法在链表尾部添加一个元素。
func (lst *List[T]) Push(v T) {
    if lst.tail == nil {
        lst.head = &element[T]{val: v}
        lst.tail = lst.head
    } else {
        lst.tail.next = &element[T]{val: v}
        lst.tail = lst.tail.next
    }
}

// `GetAll` 方法返回链表中所有元素的切片。
func (lst *List[T]) GetAll() []T {
    var elems []T
    for e := lst.head; e != nil; e = e.next {
        elems = append(elems, e.val)
    }
    return elems
}

func main() {
    // --- 测试泛型函数 SlicesIndex ---
    xi := []int{10, 20, 15, -10}
    // 在整数切片中查找 15，它在索引 2 的位置。
    // Go 编译器会自动推断出类型参数。
    fmt.Println(SlicesIndex(xi, 15))

    xs := []string{"foo", "bar", "baz"}
    // 在字符串切片中查找 "hello"，它不存在，所以返回 -1。
    fmt.Println(SlicesIndex(xs, "hello"))

    // --- 测试泛型类型 List ---
    // 实例化一个存放 int 的链表。
    var lst List[int]
    lst.Push(10)
    lst.Push(20)
    lst.Push(30)
    // 获取并打印链表中的所有元素。
    fmt.Println("list:", lst.GetAll())
}