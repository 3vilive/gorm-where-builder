package main

import (
	"net/http"

	builder "github.com/3vilive/gorm-where-builder"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var mysqlDb *gorm.DB

func init() {
	db, err := gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	mysqlDb = db
}

type Article struct {
	Id         int64  `gorm:"column:id"`
	Title      string `gorm:"column:title"`
	WebSiteId  int64  `gorm:"column:web_site_id"`
	CategoryId int64  `gorm:"column:category_id"`
	Status     string `gorm:"column:status"`
	Content    string `gorm:"column:content"`
	ExtContent string `gorm:"column:ext_content"`
}

type QueryFilter struct {
	Id         *int64 `json:"id" where:"id,eq"`
	Title      string `json:"title" where:"title,like"`
	WebSiteId  int64  `json:"web_site_id" where:"web_site_id,eq"`
	CategoryId int64  `json:"category_id" where:"category_id,eq"`
	Status     string `json:"status" where:"status,eq"`
}

func OnQuery(c *gin.Context) {
	var request QueryFilter
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var articles []Article
	where := builder.NewBuilderFromStruct(request)
	err := mysqlDb.Find(&articles, where.Build()...).Error
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, articles)
}

func main() {
	r := gin.Default()
	r.POST("/query", OnQuery)
	r.Run(":8080")
}
