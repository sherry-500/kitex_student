package main

import (
	"context"
	"errors"
	"log"
	//"time"
	"strings"

	demo "github.com/sherry-500/kitex_student/kitex_gen/demo"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// StudentServiceImpl implements the last service interface defined in the IDL.
type StudentServiceImpl struct{
	db *gorm.DB
}

//Student entity
type Student struct {
	Id int32
	Name string
	Email string
	CollegeName string
	CollegeAddress string
	//CreatedAt time.Time `gorm:"default:CURRENT_TMESTAMP`
}

// Register implements the StudentServiceImpl interface.
func (s *StudentServiceImpl) Register(ctx context.Context, student *demo.Student) (resp *demo.RegisterResp, err error) {
	// TODO: Your code here...
	result := s.db.Table("students").Create(student2Model(student))
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	return
}

// Query implements the StudentServiceImpl interface.
func (s *StudentServiceImpl) Query(ctx context.Context, req *demo.QueryReq) (resp *demo.Student, err error) {
	// TODO: Your code here...
	var stuRes Student

	result := s.db.Table("students").First(&stuRes, req.Id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		resp = &demo.Student{}
		panic("the result of query is null")
	}else{
		resp = &demo.Student{
			Id : req.Id,
			Name : stuRes.Name,
			College: &demo.College{
				Name: stuRes.CollegeName,
				Address: stuRes.CollegeAddress,
			},
			Email: strings.Split(stuRes.Email, ","),
		}
	}
	return
}

func (s *StudentServiceImpl) InitDB() {
    db, err := gorm.Open(sqlite.Open("foo.db"), &gorm.Config{})
    if err != nil {
       panic(err)
    }
    // drop table
    db.Migrator().DropTable(Student{})
    // create table
    err = db.Migrator().CreateTable(Student{})
    if err != nil {
       panic(err)
    }
    s.db = db
}

func student2Model(student *demo.Student) *Student{
	stu := &Student{
		Id : student.Id,
		Name: student.Name,
		CollegeName: student.College.Name,
		CollegeAddress: student.College.Address,
		Email : strings.Join(student.Email, ","),
	}
	return stu
}