package main

import "fmt"

// `person` 结构体类型有 `name` 和 `age` 两个字段。
type person struct {
    name string
    age  int
}

// `newPerson` 是一个构造函数，它接收一个名字，并返回一个初始化好的 `person` 结构体的指针。
func newPerson(name string) *person {
    p := person{name: name}
    p.age = 42
    // Go 语言可以安全地返回一个指向局部变量的指针。
    return &p
}

func main() {
    // 这种语法创建一个新的结构体实例。
    fmt.Println(person{"Bob", 20})

    // 初始化结构体时可以命名字段。
    fmt.Println(person{name: "Alice", age: 30})

    // 省略的字段将被零值化。
    fmt.Println(person{name: "Fred"}) // age 字段为 0

    // `&` 前缀产生指向结构体的指针。
    fmt.Println(&person{name: "Ann", age: 40})

    // 使用构造函数来创建一个新的结构体实例。
    fmt.Println(newPerson("Jon"))

    // 使用点号访问结构体字段。
    s := person{name: "Sean", age: 50}
    fmt.Println(s.name)

    // 也可以对结构体指针使用点号 - 指针会被自动解引用。
    sp := &s
    fmt.Println(sp.age)

    // 结构体是可变的。
    sp.age = 51
    fmt.Println(sp.age)

    // 如果一个结构体类型只在单个地方使用，可以声明为匿名结构体。
    dog := struct {
        name   string
        isGood bool
    }{
        "Rex",
        true,
    }
    fmt.Println(dog)
}