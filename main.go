package main

import (
	"fmt"
	"shared_lib"
)

func main() {
	fmt.Println("Hello, World!")
	// 导入本地模块 需要在go.work中添加相对路径
	shared_lib.PrintMessage("Hello, World!")
}
