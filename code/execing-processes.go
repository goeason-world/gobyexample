package main

import (
    "os"
    "os/exec"
    "syscall"
)

func main() {
    // `exec.LookPath` 用于在系统的 `PATH` 环境变量中查找一个可执行文件的绝对路径。
    binary, lookErr := exec.LookPath("ls")
    if lookErr != nil {
        panic(lookErr)
    }

    // `syscall.Exec` 需要一个参数切片。
    args := []string{"ls", "-a", "-l", "-h"}

    // `syscall.Exec` 还需要一个环境变量的切片。
    env := os.Environ()

    // `syscall.Exec` 会将当前的 Go 进程完全替换为一个全新的进程。
    // 如果这个调用成功，它将永远不会返回。
    execErr := syscall.Exec(binary, args, env)
    if execErr != nil {
        panic(execErr)
    }
}
