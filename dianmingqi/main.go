package main

import (
	"dianmingqi/database"
	"dianmingqi/modles"
	"dianmingqi/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	R := gin.Default()
	R.LoadHTMLGlob("templates/*")
	database.OpenDatabase()
	database.DB.AutoMigrate(&modles.Student{})
	routers.InitRouter(R)
	R.Run(":8900")
}
