package routers

import (
	"926new/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	r.GET("/home/", controllers.GetHome)

	//user
	r.GET("/login", controllers.GetLogin)
	r.POST("/login", controllers.PostUser)
	r.GET("/logout", controllers.LoginOut)
	r.GET("/register", controllers.Regist)
	r.POST("/register", controllers.PostRegist)
	//
	r.GET("/product/:id", controllers.GetProduct)

}
