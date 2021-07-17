package routers

import (
	gs "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"web_app/controllers"
	"web_app/logger"
	"web_app/middleware"

	_ "web_app/docs" // 千万不要忘了导入把你上一步生成的docs

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {

	r := gin.Default()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	r.Use(middleware.Cors())
	r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
	v1 := r.Group("/api/v1")

	//  注册
	v1.POST("/signup", controllers.SignUpHandler)
	//  登陆
	v1.POST("/login", controllers.LoginHandler)
	v1.GET("/refresh_token", controllers.RefreshTokenHandler)
	v1.Use(middleware.JWTAuthMiddleware())
	{

		v1.GET("/community", controllers.CommunityHandler)
		v1.GET("/community/:id", controllers.CommunityDetailHandler)

		v1.POST("/post", controllers.CreatePostHandler)
		v1.GET("/post/:id", controllers.PostDetailHandler)
		v1.GET("/posts", controllers.PostListHandler)
		v1.GET("/posts2", controllers.PostListHandler2) //新版帖子列表。已经将下方接口合并
		//v1.GET("/communityPosts", controllers.CommunityIDPostListHandler) //根据社区返回帖子列表,废弃。
		v1.POST("/post/vote", controllers.PostVoteHandler) //帖子投票

	}
	return r
}
