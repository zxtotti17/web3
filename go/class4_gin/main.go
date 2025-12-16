package main

import (
	"class4_gin/dao"
	"class4_gin/models"
	"class4_gin/route"
	"class4_gin/utils"
)

func main() {
	// gin.SetMode(gin.ReleaseMode)
	models.Logger = utils.InitLogger()
	// 初始化数据库
	models.Db = dao.GormMysql()
	if models.Db != nil {
		models.Logger.Info("mysql initialized successfully")
	} else {
		return
	}

	// 自动迁移模型
	models.Db.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{})

	router := route.Routers()

	router.Run() // 默认监听 0.0.0.0:8080

}
