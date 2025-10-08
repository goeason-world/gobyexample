package main

import (
    "os"
    "text/template"
)

func main() {
    // 创建一个简单的模板。
    t1 := template.New("t1")
    // `Parse` 方法用于解析模板字符串。
    t1, err := t1.Parse("Value is {{.}}\n")
    if err != nil {
        panic(err)
    }

    // `template.Must` 是一个辅助函数，它包装了一个返回 `(*Template, error)` 的函数调用，
    // 并在错误不为 `nil` 时 panic。这在初始化全局模板变量时非常有用。
    t1 = template.Must(t1.Parse("Value is {{.}}\n"))

    // `Execute` 方法用于执行模板，并将结果写入 `io.Writer`。
    // `{{.}}` 是一个动作，它会插入传递给模板的当前数据。
    t1.Execute(os.Stdout, "some string")
    t1.Execute(os.Stdout, 5)
    t1.Execute(os.Stdout, []string{
        "Go",
        "Rust",
        "C++",
        "C#",
    })

    // `Create` 是一个辅助函数，用于创建和解析模板。
    Create := func(name, t string) *template.Template {
        return template.Must(template.New(name).Parse(t))
    }

    // 演示 `if/else`。
    t2 := Create("t2",
        "{{if . -}} yes {{else -}} no {{end}}\n")
    t2.Execute(os.Stdout, "not empty")
    t2.Execute(os.Stdout, "")

    // 演示空白符控制。
    // `-` 用于去除动作任意一侧的空白符。
    t3 := Create("t3",
        "{{- if . -}} yes {{else -}} no {{end -}}\n")
    t3.Execute(os.Stdout, "not empty")
    t3.Execute(os.Stdout, "")

    // 演示 `range`。
    // `range` 用于遍历切片、数组、映射或通道。
    t4 := Create("t4",
        "Range: {{range .}}{{.}} {{end}}\n")
    t4.Execute(os.Stdout,
        []string{
            "Go",
            "Rust",
            "C++",
            "C#",
        })
}