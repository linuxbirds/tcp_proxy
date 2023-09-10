package main

import (
	"fmt"
	"github.com/getlantern/systray"
	"github.com/go-ini/ini"
	"os"
	"os/exec"
	"path/filepath"
	"sync"
)

var mutex sync.Mutex

func main() {
	// 1. 获取当前程序所在目录
	exePath, _ := os.Executable()
	currentDir := filepath.Dir(exePath)

	//currentDir = "H:\\果果冒险岛"
	fmt.Println("当前运行目录：" + currentDir)
	// 2. 进入子目录并修改ini文件
	subDir := filepath.Join(currentDir, "\\MapleStory")
	//fmt.Println("切换目录至：" + subDir)
	err := os.Chdir(subDir)
	if err != nil {
		fmt.Println("目录不存在，请确保游戏目录已经正常解压！")
		return
	}
	updateIniFile(subDir)

	// 3. 调用另一个exe程序
	runOtherExecutable(subDir)

	// 4. 创建带托盘菜单的Windows应用
	go createTrayMenu()

	// 保持程序运行
	select {}
}

func updateIniFile(subDir string) {
	// 在这里添加代码来修改ini文件
	//fmt.Println("ini文件修改成功！")
	iniFile := filepath.Join(subDir, "\\HShield\\ehsvc.ini")
	cfg, _ := ini.Load(iniFile)
	section := cfg.Section("Settings")
	section.Key("GamePath").SetValue(filepath.Join(subDir, "\\MapleStory.exe"))

	// 保存修改后的INI文件
	err := cfg.SaveTo(iniFile)
	if err != nil {
		fmt.Println("无法保存INI文件:", err)
		return
	}

}

func runOtherExecutable(subDir string) {
	// 保证只有一个实例在运行
	mutex.Lock()
	defer mutex.Unlock()
	fmt.Println("开始启动冒险岛主程序！")
	// 调用另一个exe程序并传递参数
	cmd := exec.Command("cmd", "/C", filepath.Join(subDir+"\\MapleStory.exe"), "svn.guoziweb.com", "3338")
	currentDir, _ := os.Getwd()
	cmd.Dir = currentDir
	//fmt.Println(cmd)
	// 将标准输出和标准错误输出重定向到空设备，以避免阻塞
	cmd.Stdout = nil
	cmd.Stderr = nil

	err := cmd.Start()
	if err != nil {
		fmt.Println("启动程序时发生错误,请将此错误报告给管理员", err)
		return
	} else {
		fmt.Println("程序已经启动，您可以关闭本窗口！！！\r\n如果程序长时间未启动，请尝试以管理员方式运行本文件！")
	}

	// 等待程序执行完成
	err = cmd.Wait()
	if err != nil {
		fmt.Println("程序执行时发生错误:", err)
	}
}

func createTrayMenu() {
	// 初始化托盘菜单
	systray.Run(onReady, onExit)
}

func onReady() {
	systray.SetTitle("Your App")
	systray.SetTooltip("Your App Tooltip")

	openMenu := systray.AddMenuItem("打开", "打开应用")
	logMenu := systray.AddMenuItem("日志", "查看日志")
	systray.AddSeparator()
	quitMenu := systray.AddMenuItem("退出", "退出应用")

	go func() {
		for {
			select {
			case <-openMenu.ClickedCh:
				// 处理打开菜单点击事件
			case <-logMenu.ClickedCh:
				// 处理日志菜单点击事件
			case <-quitMenu.ClickedCh:
				// 处理退出菜单点击事件
				systray.Quit()
				return
			}
		}
	}()
}

func onExit() {
	// 在退出应用时可以进行一些清理工作
}
