package cmd

import (
	"os"
	"os/exec"
)

func RunSwagInit() error {
	// 要执行的命令和参数
	cmd := exec.Command("swag","init")

	// 设置输出到标准输出和标准错误
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// 执行命令
	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}