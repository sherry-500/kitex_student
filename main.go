package main

import (
	demo "github.com/sherry-500/kitex_student/kitex_gen/demo/studentservice"
	"log"
)

func main() {
	svr := demo.NewServer(new(StudentServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
