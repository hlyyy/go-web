package router

import (
	"github.com/gin-gonic/gin"
	"github.com/hly/ginny1/handler/distribution"
	"github.com/hly/ginny1/handler/register"
	"github.com/hly/ginny1/handler/query"
)

var Router *gin.Engine

func InitRouter() {
	Router = gin.Default()
	Router.POST("/register/create_student", register.Createstudent)
	Router.POST("/register/create_course", register.Createcourse)
	Router.POST("/register/create_classroom", register.Createclassroom)
	Router.POST("/distribution/student_course", distribution.DistributeStudentAndCourse)
	Router.POST("/distribution/course_classroom", distribution.DistributeCourseAndClassroom)
	Router.POST("/query/query_course",query.Querycourse)
	return
}
