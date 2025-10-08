package main

import (
    "encoding/xml"
    "fmt"
)

// `Plant` 结构体将被映射到 XML。
// 字段标签决定了 XML 元素的名称和属性。
type Plant struct {
    XMLName xml.Name `xml:"plant"`
    Id      int      `xml:"id,attr"`
    Name    string   `xml:"name"`
    Origin  []string `xml:"origin"`
}

// `String` 方法实现了 `fmt.Stringer` 接口，用于自定义打印格式。
func (p Plant) String() string {
    return fmt.Sprintf("Plant id=%v, name=%v, origin=%v",
        p.Id, p.Name, p.Origin)
}

func main() {
    // --- 编码 (Marshal) ---
    coffee := &Plant{Id: 27, Name: "Coffee"}
    coffee.Origin = []string{"Ethiopia", "Brazil"}

    // `xml.MarshalIndent` 用于生成带缩进的、人类可读的 XML。
    out, _ := xml.MarshalIndent(coffee, " ", "  ")
    fmt.Println(string(out))

    // 添加通用的 XML 头部。
    fmt.Println(xml.Header + string(out))

    // --- 解码 (Unmarshal) ---
    var p Plant
    if err := xml.Unmarshal(out, &p); err != nil {
        panic(err)
    }
    fmt.Println(p)

    // --- 处理嵌套元素 ---
    tomato := &Plant{Id: 81, Name: "Tomato"}
    tomato.Origin = []string{"Mexico", "California"}

    type Nesting struct {
        XMLName xml.Name `xml:"nesting"`
        Plants  []*Plant `xml:"parent>child>plant"`
    }

    nesting := &Nesting{}
    nesting.Plants = []*Plant{coffee, tomato}

    out, _ = xml.MarshalIndent(nesting, " ", "  ")
    fmt.Println(string(out))
}