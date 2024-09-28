package controllers

import (
	"926new/database"
	"926new/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func GetHome(c *gin.Context) {

	products := LoadingAllproduct()

	totalProducts := len(products)
	pageSize := 15
	totalPages := (totalProducts + pageSize - 1) / pageSize // 向上取整，计算总页数

	// 从 URL 获取当前页
	currentPageStr := c.Query("currentpage")
	currentPage := 1 // 默认第一页
	fmt.Println(currentPageStr, "sasdas")
	if currentPageStr != "" {
		pageInt, err := strconv.Atoi(currentPageStr)
		if err == nil && pageInt > 0 {
			currentPage = pageInt
		}

	}

	// 计算起始和结束索引
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
	c.HTML(200, "home.html", gin.H{
		"title": "Main website",
		//"Products": products,
		"Products":    currentProducts,
		"CurrentPage": currentPage,
		"TotalPages":  totalPages,
		"Pages":       pages,
	})
}

func LoadingAllproduct() []models.Product {

	var products []models.Product
	query := "SELECT id, name, description, price, stock, category FROM Product"
	rows, err := database.Db.Query(query)
	if err != nil {
		log.Println("查询商品失败:", err)
		return nil
	}
	defer rows.Close()
	for rows.Next() {
		var product models.Product
		if err := rows.Scan(&product.Id, &product.Name, &product.Description, &product.Price, &product.Stock, &product.Category); err != nil {
			log.Println("扫描商品失败:", err)
			return nil
		}
		products = append(products, product)
	}
	if err := rows.Err(); err != nil {
		log.Println("行扫描错误:", err)
		return nil
	}

	return products
}

func GetLogin(c *gin.Context) {
	c.HTML(200, "login.html", gin.H{
		"title": "login website",
	})
}

func GetProduct(c *gin.Context) {
	Id := c.Param("id")
	var product models.Product
	row, _ := database.Db.Query("SELECT * FROM Product WHERE Id = ?", Id)
	if row.Next() {
		err := row.Scan(&product.Id, &product.Name, &product.Description, &product.Price, &product.Stock, &product.Category)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	c.HTML(http.StatusOK, "product_detail.html", gin.H{
		"title":   "product_detail",
		"product": product,
	})

}
