package register

import (
	"github.com/gin-gonic/gin"
	"github.com/hly/ginny1/model"
)

func Createcourse(c *gin.Context) {
	var data model.Course
	if err := c.BindJSON(&data); err != nil {
		c.JSON(400, gin.H{
			"message": "Bad Requset!",
		})
		return
	}
	if model.CheckCourse(data.Name) {
		c.JSON(401, gin.H{
			"message": "Course Already Existed!",
		})
		return
	}
	model.CreateCourse(data.Name)
	c.JSON(200, gin.H{
		"message": "Create Course Successful!",
	})
}
