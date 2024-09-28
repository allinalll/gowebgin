package main

import (
	"926new/database"
	"926new/routers"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	database.Init()
	r := gin.Default()
	routers.SetupRouter(r)
	r.LoadHTMLGlob("templates/*")
	r.Run(":8080")

}
