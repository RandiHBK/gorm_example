package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/kr/pretty"
	"gorm_example/model"
)

func main() {
	user := ""
	password := ""
	host := ""
	port := ""
	schema := "school"

	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, schema))

	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.SingularTable(true)
	db.LogMode(true)

	//no preload
	student := new(model.Student)
	db.Take(student)
	pretty.Println(student)

	//preload student's courses
	student2 := new(model.Student)
	db.Preload("Courses").Take(student2)
	pretty.Println(student2)

	//if there are multiple layers of table dependency, load them all
	student3 := new(model.Student)
	db.Set("gorm:auto_preload", true).Take(student3)
	pretty.Println(student3)

	course := new(model.Course)
	db.Take(course)
	pretty.Println(course)

	students := new([]*model.Student)
	db.Preload("Courses").Find(students)
	pretty.Println(students)
}