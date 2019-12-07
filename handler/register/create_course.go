package register

import (
	"github.com/hly/ginny1/model"
	"github.com/gin-gonic/gin"
)

func Createcourse (c *gin.Context) {
	var data model.Course
	if err := c.BindJSON(&data); err != nil {
		c.JSON(400, gin.H{
			"message":"Bad Requset!",
		})
		return
	}
	if model.CheckCourse(data.Name) {
		c.JSON(401,gin.H{
			"message":"Course Already Existed!",
		})
		return 
	}
	model.CreateCourse(data.Name) 
	c.JSON(200,gin.H{
		"message":"Create Course Successful!",
	})
}

