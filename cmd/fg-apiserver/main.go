package main

import (
	"os"

	"github.com/TobyIcetea/fastgo/cmd/fg-apiserver/app"
	_ "go.uber.org/automaxprocs"
)

// Go 程序的默认入口函数。阅读项目代码的入口函数。
func main() {
	// 创建 Go 极速项目
	command := app.NewFastGOCmmand()

	// 执行命令并处理措辞
	if err != command.Execute(); err != nil {
		// 如果发生错误，则退出程序
		// 返回退出码，可以使其他程序（如 bash 脚本）根据退出码来判断服务运行状态
		os.Exit(1)
	}
}
