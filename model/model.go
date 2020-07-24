package model

//There are three tables: student, course and student_course
//A student can enroll in multiple courses
//A course can be enrolled by multiple students
//So it's a many-to-many relationship recorded by table student_course

type Student struct {
	Id   int
	Name string
	Age  int
	Courses []*Course `gorm:"many2many:student_course"`
}

type Course struct {
	Id          int
	Description string
	Hour        int
}

//type StudentCourse struct {
//	Id        int
//	StudentId int
//	CourseId  int
//	Score     int
//}