package main

import (
    "encoding/json"
    "fmt"
    "os"
)

// `Response1` 演示了基本类型的编码。
// 只有导出的字段（首字母大写）才会被编码。
type Response1 struct {
    Page   int
    Fruits []string
}

// `Response2` 演示了使用 struct tag 自定义 JSON 键名。
type Response2 struct {
    Page   int      `json:"page"`
    Fruits []string `json:"fruits"`
}

func main() {
    // --- 编码 (Marshal) ---

    // 基本数据类型到 JSON
    bolB, _ := json.Marshal(true)
    fmt.Println(string(bolB))

    intB, _ := json.Marshal(1)
    fmt.Println(string(intB))

    fltB, _ := json.Marshal(2.34)
    fmt.Println(string(fltB))

    strB, _ := json.Marshal("gopher")
    fmt.Println(string(strB))

    // 切片和 map 到 JSON
    slcD := []string{"apple", "peach", "pear"}
    slcB, _ := json.Marshal(slcD)
    fmt.Println(string(slcB))

    mapD := map[string]int{"apple": 5, "lettuce": 7}
    mapB, _ := json.Marshal(mapD)
    fmt.Println(string(mapB))

    // 自定义结构体到 JSON
    res1D := &Response1{
        Page:   1,
        Fruits: []string{"apple", "peach", "pear"}}
    res1B, _ := json.Marshal(res1D)
    fmt.Println(string(res1B))

    // 使用带 struct tag 的结构体
    res2D := &Response2{
        Page:   1,
        Fruits: []string{"apple", "peach", "pear"}}
    res2B, _ := json.Marshal(res2D)
    fmt.Println(string(res2B))

    // --- 解码 (Unmarshal) ---

    // 将 JSON 解码到通用接口 (map)
    byt := []byte(`{"num":6.13,"strs":["a","b"]}`)
    var dat map[string]interface{}
    if err := json.Unmarshal(byt, &dat); err != nil {
        panic(err)
    }
    fmt.Println(dat)

    // 为了访问 map 中的值，需要进行类型断言
    num := dat["num"].(float64)
    fmt.Println(num)

    // 访问嵌套数据
    strs := dat["strs"].([]interface{})
    str1 := strs[0].(string)
    fmt.Println(str1)

    // 将 JSON 解码到自定义结构体
    str := `{"page": 1, "fruits": ["apple", "peach"]}`
    res := Response2{}
    json.Unmarshal([]byte(str), &res)
    fmt.Println(res)
    fmt.Println(res.Fruits[0])

    // --- 流式编码 ---
    // 可以将 JSON 编码直接流式传输到 `os.Writer` (如 `os.Stdout`) 或 HTTP 响应体。
    enc := json.NewEncoder(os.Stdout)
    d := map[string]int{"apple": 5, "lettuce": 7}
    enc.Encode(d)
}