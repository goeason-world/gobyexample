package main

import (
    "fmt"
    "io"
    "os/exec"
)

func main() {
    // `exec.Command` 用于创建一个代表外部命令的 `Cmd` 对象。
    dateCmd := exec.Command("date")

    // `Output` 方法会执行命令，并等待其完成，然后返回命令的标准输出。
    dateOut, err := dateCmd.Output()
    if err != nil {
        panic(err)
    }
    fmt.Println("> date")
    fmt.Println(string(dateOut))

    // 错误处理
    _, err = exec.Command("date", "-x").Output()
    if err != nil {
        switch e := err.(type) {
        case *exec.Error:
            fmt.Println("failed executing:", err)
        case *exec.ExitError:
            fmt.Println("command exit rc =", e.ExitCode())
        default:
            panic(err)
        }
    }

    // 通过管道与子进程交互。
    grepCmd := exec.Command("grep", "hello")

    grepIn, _ := grepCmd.StdinPipe()
    grepOut, _ := grepCmd.StdoutPipe()
    grepCmd.Start()
    grepIn.Write([]byte("hello grep\ngoodbye grep"))
    grepIn.Close()
    grepBytes, _ := io.ReadAll(grepOut)
    grepCmd.Wait()

    fmt.Println("> grep hello")
    fmt.Println(string(grepBytes))

    // 执行复杂的 Shell 命令。
    lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
    lsOut, err := lsCmd.Output()
    if err != nil {
        panic(err)
    }
    fmt.Println("> ls -a -l -h")
    fmt.Println(string(lsOut))
}
