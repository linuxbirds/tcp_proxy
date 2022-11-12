package main

import (
	"fmt"

	"myProject/TcpProxy"
)

func main() {
	fmt.Println("This is a TCP Proxy program.")
	fmt.Println("You just need run as:Proxy.exe")
	// 开始调用代理程序
	TcpProxy.Proxy()

}
