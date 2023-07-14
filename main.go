package main

import (
	"github.com/cloudwego/kitex/server"
	demo "github.com/sherry-500/kitex_student/kitex_gen/demo/studentservice"
	"log"
	"net"
	//"sever"
)

func main() {
	//svr := demo.NewServer(new(StudentServiceImpl))

	 addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:8889")
     svr := demo.NewServer(new(StudentServiceImpl), server.WithServiceAddr(addr))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
