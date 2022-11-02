package main

import (
	"fmt"
	"log"
	"net"

	"github.com/tidwall/gjson"
)

func main() {
	fmt.Println("hello")
	fromAddress := "0.0.0.0:3534"
	fromListener, err := net.Listen("tcp", fromAddress)

	if err != nil {
		log.Fatalf("Unable to listen on: %s, error : %s\n", fromAddress, err.Error())
		return
	}

	defer func(fromListener net.Listener) {
		err := fromListener.Close()
		if err != nil {

		}
	}(fromListener)

	for {
		fromConn, err := fromListener.Accept()
		if err != nil {
			log.Printf("Unable to accept a request, error: %s\n", err.Error())
		} else {
			fmt.Println("new connect:" + fromConn.RemoteAddr().String())
		}

		go acceptProcess(fromConn)
	}

}

func acceptProcess(fromConn net.Conn) {
	defer func(fromConn net.Conn) {
		err := fromConn.Close()
		if err != nil {

		}
	}(fromConn)

	var buf = make([]byte, 100000)

	for {
		n, err := fromConn.Read(buf)
		if err != nil {
			break
		}
		fmt.Println(buf[:n])

		value := gjson.Get(string(buf[:n]), "类型")
		fmt.Println(value)
	}

}
