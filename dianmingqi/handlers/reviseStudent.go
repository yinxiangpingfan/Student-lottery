package handlers

import (
	"dianmingqi/database"
	"dianmingqi/modles"

	"github.com/gin-gonic/gin"
)

func AddStudent(c *gin.Context) {
	var students modles.Student
	students.Name = c.PostForm("name")
	students.ID = c.PostForm("id")
	students.Sex = c.PostForm("sex")
	err := database.DB.Create(&students).Error
	if err != nil {
		c.JSON(500, gin.H{"message": "添加失败",
			"code": 3})
	} else {
		c.JSON(200, gin.H{"code": 0})
	}
}
func DeleteStudent(c *gin.Context) {
	tempID := c.PostForm("id")
	err := database.DB.Where("id = ?", tempID).Delete(&modles.Student{}).Error
	if err != nil {
		c.JSON(500, gin.H{"code": 2,
			"message": "删除失败"})
	} else {
		c.JSON(200, gin.H{"code": 0})
	}
}
func ShowAllStudents(c *gin.Context) {
	var students []modles.Student
	err := database.DB.Find(&students).Error
	if err != nil {
		c.JSON(500, gin.H{"message": "查询失败",
			"code": 1})
	} else {
		c.JSON(200, gin.H{"code": 0, "data": students})
	}
}
