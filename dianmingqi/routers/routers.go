package routers

import (
	handlers "dianmingqi/handlers"

	"github.com/gin-gonic/gin"
)

func InitRouter(R *gin.Engine) {
	R.GET("/all", handlers.ShowAllStudents)
	R.POST("/add", handlers.AddStudent)
	R.POST("/delete", handlers.DeleteStudent)
	R.GET("/drawlots", handlers.DrawlotsStudent)
	R.POST("/number", handlers.RandomNumber)
	R.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})
}
