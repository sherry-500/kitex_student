package main

import (
	"github.com/cloudwego/kitex/server"
	demo "github.com/sherry-500/kitex_student/kitex_gen/demo/studentservice"
	"log"
	"net"
	//"sever"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
)


func main() {
	//svr := demo.NewServer(new(StudentServiceImpl))
	r, err := etcd.NewEtcdRegistry([]string{"127.0.0.1:2379"})
	if err != nil {
		log.Fatal(err)
	}

	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:8889")

	studentService := &StudentServiceImpl{}
	studentService.InitDB()

	svr := demo.NewServer(studentService, server.WithRegistry(r), server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "studentservice"}), server.WithServiceAddr(addr))

    //  svr := demo.NewServer(new(StudentServiceImpl), server.WithServiceAddr(addr))

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
