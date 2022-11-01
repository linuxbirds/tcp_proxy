package main

import (
	"fmt"
	"io"
	"log"
	"net"

	"github.com/fatih/color"
	"github.com/go-ini/ini"
)

var red = color.New(color.FgRed).SprintFunc()
var blue = color.New(color.FgBlue).SprintFunc()

func main() {

	// Parse the starting parameters
	//fromAddress := flag.String("from", "127.0.0.1:8888", blue("From which address to proxied"))
	//toAddress := flag.String("to", "127.0.0.1:8080", blue("To which address to proxied"))
	//
	//flag.Parse()

	cfgs, err := ini.Load("tcpproxy.ini")
	if err != nil {
		fmt.Println("Ini format err:", err)
		return
	}

	section := cfgs.Section("TcpProxy")
	fromAddress := section.Key("fromAddress").Value()
	toAddress := section.Key("toAddress").Value()

	fromListener, err := net.Listen("tcp", fromAddress)

	if err != nil {
		log.Fatal(red("Unable to listen on: %s, error : %s\n", fromAddress, err.Error()))
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
			log.Print(red("Unable to accept a request, error: %s\n", err.Error()))
		} else {
			fmt.Println("new connect:" + blue(fromConn.RemoteAddr().String()))
		}

		//这边最好也做个协程，防止阻塞
		toConn, err := net.Dial("tcp", toAddress)
		if err != nil {
			fmt.Print(red("can not connect to %s\n", toAddress))
			continue
		}

		go func() {
			_, err := io.Copy(fromConn, toConn)
			if err != nil {
				fmt.Println("from -> to failure:", err)
				err := fromConn.Close()
				if err != nil {
					fmt.Println("", err)
				}
				err = toConn.Close()
				if err != nil {
					fmt.Println("", err)
				}
			}
		}()
		go func() {
			_, err := io.Copy(toConn, fromConn)
			if err != nil {
				fmt.Println("to -> from failure:", err)
				err := toConn.Close()
				if err != nil {
					fmt.Println("", err)
				}
				err = fromConn.Close()
				if err != nil {
					fmt.Println("", err)
				}
			}
		}()

	}

}
