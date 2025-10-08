package main

import "fmt"

// Go 没有内置的枚举类型，但可以使用语言习惯用法来实现枚举。
// 这里我们创建一个 `ServerState` 类型来表示服务器的状态。
type ServerState int

// 使用 `iota` 自动生成递增的常量值。
const (
    StateIdle ServerState = iota // 0
    StateConnected               // 1
    StateError                   // 2
    StateRetrying                // 3
)

// `stateName` 映射表提供了枚举值的字符串表示。
var stateName = map[ServerState]string{
    StateIdle:      "idle",
    StateConnected: "connected",
    StateError:     "error",
    StateRetrying:  "retrying",
}

// 通过实现 `fmt.Stringer` 接口，我们可以自定义枚举值的打印格式。
func (ss ServerState) String() string {
    return stateName[ss]
}

func main() {
    // 调用 `transition` 函数来模拟状态转换。
    ns := transition(StateIdle)
    fmt.Println(ns)

    ns2 := transition(ns)
    fmt.Println(ns2)
}

// `transition` 函数模拟服务器状态的变化。
func transition(s ServerState) ServerState {
    switch s {
    case StateIdle:
        return StateConnected
    case StateConnected, StateRetrying:
        return StateIdle
    case StateError:
        return StateError
    default:
        panic(fmt.Errorf("unknown state: %s", s))
    }
}