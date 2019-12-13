package query

import (
	"fmt"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/hly/ginny1/model"
)

func Querycourse(c *gin.Context) {
	var data model.Student
	var result []model.Course
	var l []string
	if err := c.BindJSON(&data); err != nil {
		c.JSON(400, gin.H{
			"message": "Bad Request!",
		})
		return
	}
	if !model.CheckStudent(data.Name) {
		c.JSON(401, gin.H{
			"message": "Student Not Existed!",
		})
		return
	}
	result = model.QueryCourseOfStudent(data.Name)
	fmt.Println(result)
	for _,u := range result {
		l=append(l,u.Name)
	}
	s,err1 := json.Marshal(result)
	fmt.Println(s)
	if err1 != nil {
		panic(err1)
	}
	c.JSON(200,gin.H{
		"message":l,
	})
}

