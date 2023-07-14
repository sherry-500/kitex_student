package main

import (
	"context"
	demo "github.com/sherry-500/kitex_student/kitex_gen/demo"
)

// StudentServiceImpl implements the last service interface defined in the IDL.
type StudentServiceImpl struct{}

// Register implements the StudentServiceImpl interface.
func (s *StudentServiceImpl) Register(ctx context.Context, student *demo.Student) (resp *demo.RegisterResp, err error) {
	// TODO: Your code here...
	return
}

// Query implements the StudentServiceImpl interface.
func (s *StudentServiceImpl) Query(ctx context.Context, req *demo.QueryReq) (resp *demo.Student, err error) {
	// TODO: Your code here...
	resp = &demo.Student{
		Id : 2,
		Name : "sherry",
		College: &demo.College{
			Name: "nju",
		},
	}
	return
}
