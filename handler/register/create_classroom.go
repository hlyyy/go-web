package register

import (
	"github.com/gin-gonic/gin"
	"github.com/hly/ginny1/model"
)

func Createclassroom(c *gin.Context) {
	var data model.Classroom
	if err := c.BindJSON(&data); err != nil {
		c.JSON(400, gin.H{
			"message": "Bad Requset!",
		})
		return
	}
	if model.CheckClassroom(data.Name) {
		c.JSON(401, gin.H{
			"message": "Classroom Already Existed!",
		})
		return
	}
	model.CreateClassroom(data.Name)
	c.JSON(200, gin.H{
		"message": "Create Classroom Successful!",
	})
}
