package route

import (
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

	auth.POST("/login", user.Login)
	auth.POST("/register", user.Register)
	// auth.POST("/logout", Logout)

	userG.POST("/addPost", user.DddPost)
	userG.GET("/GetPost", user.GetPost)
	userG.GET("/GetPostList", user.GetPostList)
	userG.POST("/updatePost", user.SetPostInfo)
	userG.POST("/deletePost", user.DeletePost)
	userG.POST("/addComment", user.AddPostComment)

	userG.Use(utils.JWTAuth())

	return Router
}
