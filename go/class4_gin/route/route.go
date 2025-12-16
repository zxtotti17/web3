package route

import (
	"class4_gin/models"
	"class4_gin/user"
	"class4_gin/utils"

	"github.com/gin-gonic/gin"
)

// 初始化总路由

func Routers() *gin.Engine {
	Router := gin.New()
	Router.Use(gin.Recovery())
	if gin.Mode() == gin.DebugMode {
		Router.Use(gin.Logger())
	}

	userG := Router.Group("/user")
	auth := Router.Group("/auth")
	userG.Use(utils.JWTAuth())

	auth.POST("/login", user.Login)
	auth.POST("/register", user.Register)
	// auth.POST("/logout", Logout)

	userG.POST("/addPost", user.DddPost)
	Router.GET("/getPost/:id", user.GetPost)
	Router.GET("/getPostList", user.GetPostList)
	userG.POST("/updatePost", user.SetPostInfo)
	userG.POST("/deletePost", user.DeletePost)
	userG.POST("/addComment", user.AddPostComment)

	Router.POST("/getComments", user.GetPostComments)

	models.Logger.Info("Router initialized successfully")
	return Router
}
