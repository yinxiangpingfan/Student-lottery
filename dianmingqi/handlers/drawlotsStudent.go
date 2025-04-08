package handlers

import (
	"dianmingqi/database"
	"dianmingqi/modles"
	"math/rand"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func DrawlotsStudent(c *gin.Context) {
	// 抽签，从数据库中随机抽取一个学生
	var students modles.Student
	err := database.DB.Order("RANDOM()").Limit(1).Find(&students).Error
	if err != nil {
		c.JSON(500, gin.H{"message": "抽签失败",
			"code": 1})
		return
	} else {
		c.JSON(200, gin.H{"code": 0,
			"data": students})
	}
}

func RandomNumber(c *gin.Context) {
	var frm, to string
	frm = c.PostForm("frm")
	to = c.PostForm("to")
	frm1, err1 := strconv.Atoi(frm)
	to1, err2 := strconv.Atoi(to)
	if err1 != nil || err2 != nil {
		c.JSON(500, gin.H{"message": "输入错误",
			"code": 1})
	} else {
		// 从frm1到to1随机抽取一个数字
		rand.Seed(time.Now().UnixNano())
		randNum := rand.Intn(to1-frm1+1) + frm1
		c.JSON(200, gin.H{"code": 0,
			"data": randNum})
	}
}
