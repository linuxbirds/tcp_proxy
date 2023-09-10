package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	err := os.Chdir("H:\\果果冒险岛\\MapleStory")
	if err != nil {
		fmt.Println("切换程序目录时发生错误:", err)
		return
	}

	//fmt.Println("111111")
	cmd := exec.Command("H:\\果果冒险岛\\MapleStory\\MapleStory.exe", "192.168.3.88", "3338")
	//cmd := exec.Command("简易登录器.bat")
	// 将标准输出和标准错误输出重定向到空设备，以避免阻塞
	//cmd.Stdout = nil
	//cmd.Stderr = nil

	currentDir, _ := os.Getwd()
	fmt.Println("正在启动应用程序" + currentDir)
	fmt.Println(cmd)
	cmd.Dir = currentDir
	err = cmd.Start()
	if err != nil {
		fmt.Println("启动程序时发生错误,请将此错误报告给管理员", err)
		return
	}

	// 等待程序执行完成
	err = cmd.Wait()
	if err != nil {
		fmt.Println("程序执行时发生错误:", err)
	}

}
