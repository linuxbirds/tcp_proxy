package main

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/linuxbirds/lfshook"
	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

func NewLoggerWithRotate() *logrus.Logger {
	if Log != nil {
		return Log
	}

	path := "E:\\GoProject\\TcpServer\\test.log"
	writer, _ := rotatelogs.New(
		path+".%Y%m%d%H%M",
		rotatelogs.WithLinkName(path),               // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(60*time.Second),       // 文件最大保存时间
		rotatelogs.WithRotationTime(20*time.Second), // 日志切割时间间隔
	)

	pathMap := lfshook.WriterMap{
		logrus.InfoLevel:  writer,
		logrus.PanicLevel: writer,
	}

	Log = logrus.New()
	Log.Hooks.Add(lfshook.NewHook(
		pathMap,
		&logrus.JSONFormatter{},
	))

	return Log
}

func main() {
	Log := NewLoggerWithRotate()
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
			Log.Info("new connect:" + fromConn.RemoteAddr().String())
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

		//value := gjson.Get(string(buf[:n]), "类型")
		//fmt.Println(value)
	}

}
