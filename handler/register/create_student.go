package register

import (
	"github.com/gin-gonic/gin"
	"github.com/hly/ginny1/model"
)

func Createstudent(c *gin.Context) {
	var data model.Student
	if err := c.BindJSON(&data); err != nil {
		c.JSON(400, gin.H{
			"message": "Bad Requset!",
		})
		return
	}
	if model.CheckStudent(data.Name) {
		c.JSON(401, gin.H{
			"message": "Student Already Existed!",
		})
		return
	}
	model.CreateStudent(data.Name)
	c.JSON(200, gin.H{
		"message": "Create Student Successful!",
	})
}
