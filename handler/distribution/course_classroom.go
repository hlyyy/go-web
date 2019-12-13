package distribution

import (
	"github.com/gin-gonic/gin"
	"github.com/hly/ginny1/model"
)

func DistributeCourseAndClassroom(c *gin.Context) {
	var data model.Courses_classroom
	if err := c.BindJSON(&data); err != nil {
		c.JSON(400, gin.H{
			"message": "Bad Requset!",
		})
		return
	}
	if model.CheckCourseAndClassroom(data.Course_id, data.Classroom_id) {
		c.JSON(401, gin.H{
			"message": "Distribution Of Course And Classroom Already Existed!",
		})
		return
	}
	model.DistributeCourseAndClassroom(data.Course_id, data.Classroom_id)
	c.JSON(200, gin.H{
		"message": "Distribute Course And Classroom Successful!",
	})
}
