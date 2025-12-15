package main

import (
	"class4_gin/dao"
	"class4_gin/models"
	"class4_gin/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	models.Logger = utils.InitLogger()
	// 初始化数据库
	models.Db = dao.GormMysql()

	// 自动迁移模型
	models.Db.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{})

	router := gin.Default()

	router.Run() // 默认监听 0.0.0.0:8080

}
