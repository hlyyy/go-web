package router

import(
	"github.com/hly/ginny1/handler/distribution"
	"github.com/hly/ginny1/handler/register"
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func InitRouter() {
	Router = gin.Default()
	Router.POST("/register/create_student",register.Createstudent)
	Router.POST("/register/create_course",register.Createcourse)
	Router.POST("/register/creste_classroom",register.Createclassroom)
	Router.POST("/distribution/student_course",distribution.DistributeStudentAndCourse)
	Router.POST("/distribution/course_classroom",distribution.DistributeCourseAndClassroom)
	return
}
