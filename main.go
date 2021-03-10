package main

import (
	"fmt"
	"github.com/crazy-me/os_tel/entity"
	"github.com/crazy-me/os_tel/service"
	"os"
)

func main() {
	tel := parseArgs()
	run := service.Run(tel)
	fmt.Println(run)
}

func parseArgs() *entity.Tel {
	commandArgs := os.Args
	if len(commandArgs) == 1 {
		toPrint(commandArgs[0])
		os.Exit(1)
	}

	argsValue := commandArgs[1:]
	if len(argsValue) == 1 {
		toPrint(commandArgs[0])
		os.Exit(1)
	}

	tel := entity.NewTel()
	tel.Mobile = argsValue[0]
	tel.Msg = argsValue[1]

	if len(argsValue) == 3 {
		tel.SerialPort = argsValue[2]
	}

	return tel
}

func toPrint(e string) {
	// 打印帮助信息
	fmt.Printf(`
请输入正确的运行参数!
 *exec: %s 13800138000 6d88606f6d4b8bd5 /dev/ttyUSB0
 	* 参数选项:
		$1: option for tel
		$2: option for Unicode message
		$3: option for Serial port name (optional)
\n`, e)
	os.Exit(1)
}
