package model

//There are three tables: student, course and student_course
//A student can enroll in multiple courses
//A course can be enrolled by multiple students
//So it's a many-to-many relationship recorded by table student_course

type Student struct {
	//column tag declare the name of column in sql table
	Id   int `gorm:"primary_key;column:id"`
	Name string `gorm:"column:name"`
	Age  int `gorm:"column:age"`
	//many2many tag declare the name of table which define many-to-many relationship
	Courses []*Course `gorm:"many2many:student_course;foreignkey:id;jointable_foreignkey:student_id;association_foreignkey:id;association_jointable_foreignkey:course_id"`
}

//if the name of golang struct is different from the name of sql table,
//then the table name must be explicitly declared through TableName() method
func (*Student) TableName() string {
	return "student"
}

type Course struct {
	Id          int `gorm:"primary_key;column:id"`
	Description string `gorm:"column:description"`
	Hour        int `gorm:"column:hour"`
}

func (*Course) TableName() string {
	return "course"
}

//type StudentCourse struct {
//	Id        int `gorm:"primary_key;column:id"`
//	StudentId int `gorm:"student_id"`
//	CourseId  int `gorm:"course_id"`
//	Score     int `gorm:"score"`
//	Course *Course `gorm:"foreignkey:id;association_foreignkey:course_id"`
//}
//
//func (*StudentCourse) TableName() string {
//	return "student_course"
//}