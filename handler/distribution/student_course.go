package distribution

import (
	"github.com/hly/ginny1/model"
	"github.com/gin-gonic/gin"
)

func DistributeStudentAndCourse (c *gin.Context) {
	var data model.Students_course
	if err := c.BindJSON(&data); err != nil {
		c.JSON(400, gin.H{
			"message":"Bad Requset!",
		})
		return
	}
	if model.CheckStudentAndCourse(data.Student_id,data.Course_id) {
		c.JSON(401,gin.H{
			"message":"Distribution Of Student And Course Already Existed!",
		})
		return
	}
	model.DistributeStudentAndCourse(data.Student_id,data.Course_id)
	c.JSON(200,gin.H{
		"message":"Distribute Student And Course Successful!",
	})
}
