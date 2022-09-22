package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"

	"github.com/fatih/color"
)

var red = color.New(color.FgRed).SprintFunc()
var blue = color.New(color.FgBlue).SprintFunc()

func main() {

	// Parse the starting parameters
	fromAddress := flag.String("from_addr", "127.0.0.1:8888", blue("From which address to proxied"))
	toAddress := flag.String("to_addr", "127.0.0.1:8080", blue("To which address to proxied"))

	flag.Parse()

	fromListener, err := net.Listen("tcp", *fromAddress)

	if err != nil {
		log.Fatal(red("Unable to listen on: %s, error : %s\n", *fromAddress, err.Error()))
	}
	defer fromListener.Close()

	for {
		fromConn, err := fromListener.Accept()
		if err != nil {
			log.Print(red("Unable to accept a request, error: %s\n", err.Error()))
		} else {
			fmt.Println("new connect:" + blue(fromConn.RemoteAddr().String()))
		}

		//这边最好也做个协程，防止阻塞
		toConn, err := net.Dial("tcp", *toAddress)
		if err != nil {
			fmt.Print(red("can not connect to %s\n", *toAddress))
			continue
		}

		go io.Copy(fromConn, toConn)
		go io.Copy(toConn, fromConn)

	}

}
