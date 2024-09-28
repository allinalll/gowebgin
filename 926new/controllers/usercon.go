package controllers

import (
	"926new/database"
	"926new/models"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

func PostUser(c *gin.Context) {
	//db, err := sql.Open("mysql", "root:1234@tcp(localhost:3306)/gomarket?charset=utf8&parseTime=true&loc=Local")
	//if err != nil {
	//	c.JSON(http.StatusInternalServerError, gin.H{"error": "数据库连接失败"})
	//	return
	//}
	name := c.PostForm("username")
	password := c.PostForm("password")
	var User models.User
	_ = database.Db.QueryRow("SELECT * FROM user WHERE Username = ?", name).Scan(&User.Id, &User.Username, &User.Password)
	if User.Password == password {
		products := LoadingAllproduct()
		totalProducts := len(products)
		pageSize := 15
		totalPages := (totalProducts + pageSize - 1) / pageSize // 向上取整

		// 获取当前页，默认为 1
		currentPage := 1

		// 计算当前页的商品索引
		start := (currentPage - 1) * pageSize
		end := start + pageSize
		if end > totalProducts {
			end = totalProducts
		}

		// 获取当前页商品
		currentProducts := products[start:end]

		// 生成页码数组
		pages := make([]int, totalPages)
		for i := 0; i < totalPages; i++ {
			pages[i] = i + 1
		}
		c.HTML(http.StatusOK, "home.html", gin.H{
			"title":       "Main website",
			"user":        name, // 将 user 数据传递给前端
			"Products":    currentProducts,
			"CurrentPage": currentPage,
			"TotalPages":  totalPages,
			"Pages":       pages,
		})
		//c.JSON(http.StatusOK, gin.H{
		//	"name":     name,
		//	"password": password,
		//	"check":    UserPassword,
		//})
	} else {
		fmt.Println(User)
		c.JSON(http.StatusOK, gin.H{
			"error": "error",
		})
	}
}

func LoginOut(c *gin.Context) {
	c.Redirect(http.StatusFound, "/home")
}

func Regist(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", gin.H{
		"meassage": "succeess",
	})
}

func PostRegist(c *gin.Context) {
	name := c.PostForm("username")
	password := c.PostForm("password")
	var user models.User
	_ = database.Db.QueryRow("SELECT Username FROM user WHERE Username = ?", name).Scan(&user.Username)
	fmt.Println(user.Username, password)
	if user.Username != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "用户名重复"})
	} else {
		_, err := database.Db.Exec("INSERT INTO user (Username,Password) VALUES (?,?)", name, password)
		if err != nil {
			log.Println("插入用户错误:", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "无法注册用户",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "注册成功",
		})
	}
}
