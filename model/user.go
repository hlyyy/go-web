package model

//定义各种结构体
type Classroom struct {
	Id   int    `json:"-"`
	Name string `json:"name";gorm:"column:name"`
}

type Course struct {
	Id   int    `json:"-"`
	Name string `json:"name";gorm:"column:name"`
}

type Student struct {
	Id   int    `json:"-"`
	Name string `json:"name";gorm:"column:name"`
}

type Courses_classroom struct {
	Id           int `json:"-"`
	Course_id    int `json:"course_id";gorm:"column:course_id"`
	Classroom_id int `json:"classroom_id";gorm:"column:classroom_id"`
}

type Students_course struct {
	Id         int `json:"-"`
	Student_id int `json:"student_id";gorm:"column:student_id"`
	Course_id  int `json:"course_id";gorm:"column:course_id"`
}

//新建学生
func CreateStudent(name string) {
	DB.Self.Model(&Student{}).Create(&Student{Name: name})
}

//检查学生是否存在
func CheckStudent(name string) bool {
	var l Student
	res := DB.Self.Table("students").Where("name = ?", name).Find(&l)
	if res.RecordNotFound() {
		return false
	} 
	return true
}

//新建课程
func CreateCourse(name string) {
	DB.Self.Model(&Course{}).Create(&Course{Name: name})
}

//检查课程是否存在
func CheckCourse(name string) bool {
	var l Course
	res := DB.Self.Model(&Course{}).Where(Course{Name: name}).First(&l)
	if res.RecordNotFound() {
		return false
	}
	return true
}

//新建教室
func CreateClassroom(name string) {
	DB.Self.Model(&Classroom{}).Create(&Classroom{Name: name})
}

//检查教室是否存在
func CheckClassroom(name string) bool {
	var l Classroom
	res := DB.Self.Model(&Classroom{}).Where(Classroom{Name: name}).First(&l)
	if res.RecordNotFound() {
		return false
	}
	return true
}

//分配学生和课程
func DistributeStudentAndCourse(student_id, course_id int) {
	DB.Self.Model(&Students_course{}).Create(&Students_course{Student_id: student_id, Course_id: course_id})
}

//检查学生和课程是否已经分配
func CheckStudentAndCourse(student_id, course_id int) bool {
	var l Students_course
	res := DB.Self.Model(&Students_course{}).Where(Students_course{Student_id: student_id, Course_id: course_id}).First(&l)
	if res.RecordNotFound() {
		return false
	}
	return true
}

//分配课程和教室
func DistributeCourseAndClassroom(course_id, classroom_id int) {
	DB.Self.Model(&Courses_classroom{}).Create(&Courses_classroom{Course_id: course_id, Classroom_id: classroom_id})
}

//检查课程和教室是否已经分配
func CheckCourseAndClassroom(course_id, classroom_id int) bool {
	var l Courses_classroom
	res := DB.Self.Model(&Courses_classroom{}).Where(Courses_classroom{Course_id: course_id, Classroom_id: classroom_id}).First(&l)
	if res.RecordNotFound() {
		return false
	}
	return true
}

//关联查询学生的课程
func QueryCourseOfStudent (name string) (d Course) {
	var l Student
	DB.Self.Model(&Student{}).Where(Student{Name:name}).First(&l)
	a := l.Id
	var b Students_course
	DB.Self.Model(&Students_course{}).Where(Students_course{Student_id:a}).First(&b)
	c := b.Course_id
	DB.Self.Model(&Course{}).Where(Course{Id:c}).Find(&d)
	return d
}















